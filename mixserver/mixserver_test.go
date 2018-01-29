// Copyright © 2018 Privategrity Corporation
// All rights reserved.

package mixserver

import (
	"os"
	"testing"
	"time"

	pb "gitlab.com/privategrity/comms/mixmessages"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func TestMain(m *testing.M) {
	addr := "localhost:5555"
	go StartServer(addr)
	os.Exit(m.Run())
}

// Smoke test the NetworkError endpoint
func TestNetworkError(t *testing.T) {
	addr := "localhost:5555"
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithInsecure(),
		grpc.WithTimeout(time.Second))
	if err != nil {
		t.Errorf("NetworkError: Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMixMessageServiceClient(conn)

	// Send error, check that we get an ErrorAck back
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	r, err := c.NetworkError(ctx, &pb.ErrorMessage{Message: "Hello, world!"})
	if err != nil {
		t.Errorf("NetworkError: Error received: %s", err)
	}
	if r.MsgLen != 13 {
		t.Errorf("NetworkError: Expected len of %v, got %v", 13, r)
	}
	defer cancel()

}

// Smoke test the AskOnline endpoint
func TestAskOnline(t *testing.T) {
	addr := "localhost:5555"
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithInsecure(),
		grpc.WithTimeout(time.Second))
	if err != nil {
		t.Errorf("AskOnline: Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMixMessageServiceClient(conn)

	// Send AskOnline Request and check that we get an AskOnlineAck back
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	_, err = c.AskOnline(ctx, &pb.Ping{})
	if err != nil {
		t.Errorf("AskOnline: Error received: %s", err)
	}
	defer cancel()

}

func TestStartServer(t *testing.T) {
	addr := "localhost:5555"
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithInsecure(),
		grpc.WithTimeout(time.Second))
	if err != nil {
		t.Errorf("TestStartServer: Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMixMessageServiceClient(conn)

	// Say hello, check that we get the correct response
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	_, err = c.AskOnline(ctx, &pb.Ping{})
	if err != nil {
		t.Errorf("TestStartServer: Could not greet: %v", err)
	}
	defer cancel()

	time.Sleep(time.Millisecond * 600)

	// Send it again, this time expect a timeout error.
	ctx2, cancel2 := context.WithTimeout(context.Background(), 300*time.Millisecond)
	_, err = c.AskOnline(ctx2, &pb.Ping{})
	if err == nil {
		t.Errorf("TestStartServer: Somehow able to greet!")
	}
	defer cancel2()
}
