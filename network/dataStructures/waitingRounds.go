///////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////
package dataStructures

import (
	"github.com/elliotchance/orderedmap"
	"github.com/pkg/errors"
	jww "github.com/spf13/jwalterweatherman"
	pb "gitlab.com/elixxir/comms/mixmessages"
	"gitlab.com/elixxir/primitives/excludedRounds"
	"gitlab.com/elixxir/primitives/states"
	"gitlab.com/xx_network/primitives/netTime"
	"sync"
	"sync/atomic"
	"time"
)

var timeOutError = errors.New("Timed out getting round furthest in the future.")

// maxGetClosestTries is the maximum amount of rounds pulled by
// WaitingRounds.GetUpcomingRealtime. Exceeding this amount causes
// WaitingRounds.GetUpcomingRealtime to switch from using
// WaitingRounds.getClosest to using WaitingRounds.getFurthest
const maxGetClosestTries = 2

// WaitingRounds contains a list of all queued rounds ordered by which occurs
// furthest in the future with the furthest in the the back.
type WaitingRounds struct {
	readRounds  *atomic.Value
	writeRounds *orderedmap.OrderedMap
	mux         sync.Mutex
	signal 	    chan struct{}
}

// NewWaitingRounds generates a new WaitingRounds with an empty round list.
func NewWaitingRounds() *WaitingRounds {
	wr := WaitingRounds{
		writeRounds: orderedmap.NewOrderedMap(),
		readRounds: &atomic.Value{},
		// this is intentionally unbuffered,
		// do not change
		signal: make(chan struct{}),
	}


	return &wr
}

// Len returns the number of rounds in the list.
func (wr *WaitingRounds) Len() int {
	wr.mux.Lock()
	defer wr.mux.Unlock()
	return wr.writeRounds.Len()
}

// Insert inserts a queued round into the list in order of its timestamp, from
// smallest to greatest. If the new round is not in a QUEUED state, then it is
// not inserted. If the new round already exists in the list but is no longer
// queued, then it is removed.
func (wr *WaitingRounds) Insert(added, removed []*Round) {
	wr.mux.Lock()
	defer wr.mux.Unlock()

	//add any round which should be added
	for i:= range added{
		toAdd := added[i]
		if !time.Now().After(time.Unix(0,int64(toAdd.info.Timestamps[states.QUEUED]))){
			wr.writeRounds.Set(toAdd.info.ID,toAdd)
		}

	}

	//remove any round which should be removed
	for i := range removed {
		toRemove := removed[i]
		wr.writeRounds.Delete(toRemove.info.ID)
	}

	// If changes occured, update the atomic
	if len(removed)>0||len(added)>0{
		wr.storeReadRounds()
	}

	// If inserts occured, signal to any waiting threads
	// only do this on inserts because only inserts will change the
	// evaluation by callers of GetUpcomingRealtime
	if len(added)>0 {
		go func(){
			// this will loop for as many people are waiting on the
			// channel which is why it is in a seperate function
			for{
				select{
				case wr.signal<- struct{}{}:
				default:
					//exit when there are no listeners
					return
				}
			}
		}()
	}
}

func(wr *WaitingRounds)storeReadRounds(){
	roundsList := make([]*Round,0,wr.writeRounds.Len())

	toDelete  := make([]*Round,0)

	for e := wr.writeRounds.Front(); e != nil; e = e.Next() {
		rnd := e.Value.(*Round)
		if time.Since(time.Unix(0,int64(rnd.info.Timestamps[states.QUEUED])))<time.Hour{
			roundsList = append(roundsList,rnd)
		}else{
			toDelete = append(toDelete,rnd)
		}

	}
	wr.readRounds.Store(roundsList)

	if len(toDelete)>0{
		for i := range toDelete {
			toRemove := toDelete[i]
			wr.writeRounds.Delete(toRemove.info.ID)
		}
	}
}

// getTime returns the timestamp for the round's realtime.
func getTime(round *Round) uint64 {
	return round.info.Timestamps[states.QUEUED]
}

// getFurthest returns the round that will occur furthest in the future. If the
// list is empty, then nil is returned. If the round is on the exclusion list,
// then the next round is checked.
// this is assumed to be called on an operation already under the cond's lock
func (wr *WaitingRounds) getFurthest(exclude excludedRounds.ExcludedRounds, cutoffDelta time.Duration) *Round {
	earliestStart := netTime.Now().Add(cutoffDelta)


	roundsList := wr.readRounds.Load().([]*Round)

	// Return the last non-excluded round in the list
	for i:=len(roundsList)-1;i>=0;i-- {
		r := roundsList[i]
		// Cannot guarantee that the round object's pointers will be exact match
		// of value in set
		RoundStartTime := time.Unix(0, int64(r.info.Timestamps[states.QUEUED]))
		if RoundStartTime.After(earliestStart) && !isExcluded(exclude, r.info) {
			return r
		}
	}

	// If all the rounds in the list are excluded, then return nil
	return nil
}

// getClosest returns the round that will occur soonest in the future. If the
// list is empty, then nil is returned. If the round is on the exclusion list,
// then the next round is checked.
// this is assumed to be called on an operation already under the cond's lock
func (wr *WaitingRounds) getClosest(exclude excludedRounds.ExcludedRounds, minRoundAge time.Duration) *Round {
	earliestStart := netTime.Now().Add(minRoundAge)

	roundsList := wr.readRounds.Load().([]*Round)

	// Return the first non-excluded round in the list
	for i:=0;i<len(roundsList);i++ {
		r := roundsList[i]
		// Cannot guarantee that the round object's pointers will be exact match
		// of value in set
		RoundStartTime := time.Unix(0, int64(r.info.Timestamps[states.QUEUED]))
		if RoundStartTime.After(earliestStart) && !isExcluded(exclude, r.info) {
			return r
		}
	}

	// If all the rounds in the list are excluded, then return nil
	return nil
}

func isExcluded(exclude excludedRounds.ExcludedRounds, r *pb.RoundInfo) bool {
	if exclude == nil {
		return false
	}

	return exclude.Has(r.GetRoundId())
}

// GetSlice returns a slice of all round infos in the list that have yet to
// occur.
func (wr *WaitingRounds) GetSlice() []*pb.RoundInfo {

	roundsList := wr.readRounds.Load().([]*Round)

	now := uint64(netTime.Now().Nanosecond())
	var roundInfos []*pb.RoundInfo
	for i:=0;i<len(roundsList);i++ {
		if getTime(roundsList[i]) > now {
			roundInfos = append(roundInfos, roundsList[i].info)
		}
	}

	// Return the last round in the list, which is the furthest in the future
	return roundInfos
}

// GetUpcomingRealtime returns the round that will occur furthest in the future.
// If the list is empty, then it waits waits for a round to be added for the
// specified duration. If no round is added, then an error is returned.
//
// The length of the excluded set indicates how many times the client has
// called GetUpcomingRealtime trying to retrieve a round to send on.
// GetUpcomingRealtime defaults to retrieving the closest non-excluded round
// from WaitingRounds. If the length of the excluded set exceeds the maximum
// attempts at pulling the closest round, GetUpcomingRealtime will retrieve
// the furthest non-excluded round from WaitingRounds.
func (wr *WaitingRounds) GetUpcomingRealtime(timeout time.Duration,
	exclude excludedRounds.ExcludedRounds, minRoundAge time.Duration) (*pb.RoundInfo, error) {

	// Start timeout timer
	timer := time.NewTimer(timeout)


	// Start seeing if an acceptable round exists
	round := wr.get(exclude,minRoundAge)
	if round!=nil{
		return round, nil
	}

	jww.INFO.Printf("Could not find round to send on, waiting for update")
	// If the no round exists, wait for an update to the list.
	for {
		select {
		case <-timer.C:
			return nil, timeOutError
		case <-wr.signal:
			round = wr.get(exclude,minRoundAge)
			if round!=nil{
				return round, nil
			}
		}
	}
}

func (wr *WaitingRounds) get(exclude excludedRounds.ExcludedRounds, minRoundAge time.Duration)*pb.RoundInfo{
	if exclude.Len() < maxGetClosestTries {
		// Use getClosest when excluded set's length is small
		round := wr.getClosest(exclude, minRoundAge)
		if round != nil {
			return round.Get()
		}
	} else {
		// Use getFurthest when excluded set's length exceeds maxGetClosestTries
		round := wr.getFurthest(exclude, minRoundAge)
		if round != nil {
			return round.Get()
		}
	}
	return nil

}
