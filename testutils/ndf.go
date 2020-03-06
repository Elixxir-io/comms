package testutils

import (
	"encoding/base64"
	"gitlab.com/elixxir/primitives/ndf"
)

var (
	ExampleJSON = `{
	"Timestamp": "2019-06-04T20:48:48-07:00",
	"gateways": [
		{
			"Address": "2.5.3.122",
			"Tls_certificate": "-----BEGIN CERTIFICATE-----\nMIIDgTCCAmmgAwIBAgIJAKLdZ8UigIAeMA0GCSqGSIb3DQEBBQUAMG8xCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQx\nGzAZBgNVBAoMElByaXZhdGVncml0eSBDb3JwLjEaMBgGA1UEAwwRZ2F0ZXdheSou\nY21peC5yaXAwHhcNMTkwMzA1MTgzNTU0WhcNMjkwMzAyMTgzNTU0WjBvMQswCQYD\nVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJQ2xhcmVtb250\nMRswGQYDVQQKDBJQcml2YXRlZ3JpdHkgQ29ycC4xGjAYBgNVBAMMEWdhdGV3YXkq\nLmNtaXgucmlwMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA9+AaxwDP\nxHbhLmn4HoZu0oUM48Qufc6T5XEZTrpMrqJAouXk+61Jc0EFH96/sbj7VyvnXPRo\ngIENbk2Y84BkB9SkRMIXya/gh9dOEDSgnvj/yg24l3bdKFqBMKiFg00PYB30fU+A\nbe3OI/le0I+v++RwH2AV0BMq+T6PcAGjCC1Q1ZB0wP9/VqNMWq5lbK9wD46IQiSi\n+SgIQeE7HoiAZXrGO0Y7l9P3+VRoXjRQbqfn3ETNL9ZvQuarwAYC9Ix5MxUrS5ag\nOmfjc8bfkpYDFAXRXmdKNISJmtCebX2kDrpP8Bdasx7Fzsx59cEUHCl2aJOWXc7R\n5m3juOVL1HUxjQIDAQABoyAwHjAcBgNVHREEFTATghFnYXRld2F5Ki5jbWl4LnJp\ncDANBgkqhkiG9w0BAQUFAAOCAQEAMu3xoc2LW2UExAAIYYWEETggLNrlGonxteSu\njuJjOR+ik5SVLn0lEu22+z+FCA7gSk9FkWu+v9qnfOfm2Am+WKYWv3dJ5RypW/hD\nNXkOYxVJNYFxeShnHohNqq4eDKpdqSxEcuErFXJdLbZP1uNs4WIOKnThgzhkpuy7\ntZRosvOF1X5uL1frVJzHN5jASEDAa7hJNmQ24kh+ds/Ge39fGD8pK31CWhnIXeDo\nvKD7wivi/gSOBtcRWWLvU8SizZkS3hgTw0lSOf5geuzvasCEYlqrKFssj6cTzbCB\nxy3ra3WazRTNTW4TmkHlCUC9I3oWTTxw5iQxF/I2kQQnwR7L3w==\n-----END CERTIFICATE-----"
		},
		{
			"Address": "2.2.58.38",
			"Tls_certificate": "-----BEGIN CERTIFICATE-----\nMIIDgTCCAmmgAwIBAgIJAKLdZ8UigIAeMA0GCSqGSIb3DQEBBQUAMG8xCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQx\nGzAZBgNVBAoMElByaXZhdGVncml0eSBDb3JwLjEaMBgGA1UEAwwRZ2F0ZXdheSou\nY21peC5yaXAwHhcNMTkwMzA1MTgzNTU0WhcNMjkwMzAyMTgzNTU0WjBvMQswCQYD\nVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJQ2xhcmVtb250\nMRswGQYDVQQKDBJQcml2YXRlZ3JpdHkgQ29ycC4xGjAYBgNVBAMMEWdhdGV3YXkq\nLmNtaXgucmlwMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA9+AaxwDP\nxHbhLmn4HoZu0oUM48Qufc6T5XEZTrpMrqJAouXk+61Jc0EFH96/sbj7VyvnXPRo\ngIENbk2Y84BkB9SkRMIXya/gh9dOEDSgnvj/yg24l3bdKFqBMKiFg00PYB30fU+A\nbe3OI/le0I+v++RwH2AV0BMq+T6PcAGjCC1Q1ZB0wP9/VqNMWq5lbK9wD46IQiSi\n+SgIQeE7HoiAZXrGO0Y7l9P3+VRoXjRQbqfn3ETNL9ZvQuarwAYC9Ix5MxUrS5ag\nOmfjc8bfkpYDFAXRXmdKNISJmtCebX2kDrpP8Bdasx7Fzsx59cEUHCl2aJOWXc7R\n5m3juOVL1HUxjQIDAQABoyAwHjAcBgNVHREEFTATghFnYXRld2F5Ki5jbWl4LnJp\ncDANBgkqhkiG9w0BAQUFAAOCAQEAMu3xoc2LW2UExAAIYYWEETggLNrlGonxteSu\njuJjOR+ik5SVLn0lEu22+z+FCA7gSk9FkWu+v9qnfOfm2Am+WKYWv3dJ5RypW/hD\nNXkOYxVJNYFxeShnHohNqq4eDKpdqSxEcuErFXJdLbZP1uNs4WIOKnThgzhkpuy7\ntZRosvOF1X5uL1frVJzHN5jASEDAa7hJNmQ24kh+ds/Ge39fGD8pK31CWhnIXeDo\nvKD7wivi/gSOBtcRWWLvU8SizZkS3hgTw0lSOf5geuzvasCEYlqrKFssj6cTzbCB\nxy3ra3WazRTNTW4TmkHlCUC9I3oWTTxw5iQxF/I2kQQnwR7L3w==\n-----END CERTIFICATE-----"
		},
		{
			"Address": "52.41.80.104",
			"Tls_certificate": "-----BEGIN CERTIFICATE-----\nMIIDgTCCAmmgAwIBAgIJAKLdZ8UigIAeMA0GCSqGSIb3DQEBBQUAMG8xCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQx\nGzAZBgNVBAoMElByaXZhdGVncml0eSBDb3JwLjEaMBgGA1UEAwwRZ2F0ZXdheSou\nY21peC5yaXAwHhcNMTkwMzA1MTgzNTU0WhcNMjkwMzAyMTgzNTU0WjBvMQswCQYD\nVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJQ2xhcmVtb250\nMRswGQYDVQQKDBJQcml2YXRlZ3JpdHkgQ29ycC4xGjAYBgNVBAMMEWdhdGV3YXkq\nLmNtaXgucmlwMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA9+AaxwDP\nxHbhLmn4HoZu0oUM48Qufc6T5XEZTrpMrqJAouXk+61Jc0EFH96/sbj7VyvnXPRo\ngIENbk2Y84BkB9SkRMIXya/gh9dOEDSgnvj/yg24l3bdKFqBMKiFg00PYB30fU+A\nbe3OI/le0I+v++RwH2AV0BMq+T6PcAGjCC1Q1ZB0wP9/VqNMWq5lbK9wD46IQiSi\n+SgIQeE7HoiAZXrGO0Y7l9P3+VRoXjRQbqfn3ETNL9ZvQuarwAYC9Ix5MxUrS5ag\nOmfjc8bfkpYDFAXRXmdKNISJmtCebX2kDrpP8Bdasx7Fzsx59cEUHCl2aJOWXc7R\n5m3juOVL1HUxjQIDAQABoyAwHjAcBgNVHREEFTATghFnYXRld2F5Ki5jbWl4LnJp\ncDANBgkqhkiG9w0BAQUFAAOCAQEAMu3xoc2LW2UExAAIYYWEETggLNrlGonxteSu\njuJjOR+ik5SVLn0lEu22+z+FCA7gSk9FkWu+v9qnfOfm2Am+WKYWv3dJ5RypW/hD\nNXkOYxVJNYFxeShnHohNqq4eDKpdqSxEcuErFXJdLbZP1uNs4WIOKnThgzhkpuy7\ntZRosvOF1X5uL1frVJzHN5jASEDAa7hJNmQ24kh+ds/Ge39fGD8pK31CWhnIXeDo\nvKD7wivi/gSOBtcRWWLvU8SizZkS3hgTw0lSOf5geuzvasCEYlqrKFssj6cTzbCB\nxy3ra3WazRTNTW4TmkHlCUC9I3oWTTxw5iQxF/I2kQQnwR7L3w==\n-----END CERTIFICATE-----"
		}
	],
	"nodes": [
		{
			"Id": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
			"Address": "18.237.147.105",
			"Tls_certificate": "-----BEGIN CERTIFICATE-----\nMIIDbDCCAlSgAwIBAgIJAOUNtZneIYECMA0GCSqGSIb3DQEBBQUAMGgxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQx\nGzAZBgNVBAoMElByaXZhdGVncml0eSBDb3JwLjETMBEGA1UEAwwKKi5jbWl4LnJp\ncDAeFw0xOTAzMDUxODM1NDNaFw0yOTAzMDIxODM1NDNaMGgxCzAJBgNVBAYTAlVT\nMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQxGzAZBgNV\nBAoMElByaXZhdGVncml0eSBDb3JwLjETMBEGA1UEAwwKKi5jbWl4LnJpcDCCASIw\nDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAPP0WyVkfZA/CEd2DgKpcudn0oDh\nDwsjmx8LBDWsUgQzyLrFiVigfUmUefknUH3dTJjmiJtGqLsayCnWdqWLHPJYvFfs\nWYW0IGF93UG/4N5UAWO4okC3CYgKSi4ekpfw2zgZq0gmbzTnXcHF9gfmQ7jJUKSE\ntJPSNzXq+PZeJTC9zJAb4Lj8QzH18rDM8DaL2y1ns0Y2Hu0edBFn/OqavBJKb/uA\nm3AEjqeOhC7EQUjVamWlTBPt40+B/6aFJX5BYm2JFkRsGBIyBVL46MvC02MgzTT9\nbJIJfwqmBaTruwemNgzGu7Jk03hqqS1TUEvSI6/x8bVoba3orcKkf9HsDjECAwEA\nAaMZMBcwFQYDVR0RBA4wDIIKKi5jbWl4LnJpcDANBgkqhkiG9w0BAQUFAAOCAQEA\nneUocN4AbcQAC1+b3To8u5UGdaGxhcGyZBlAoenRVdjXK3lTjsMdMWb4QctgNfIf\nU/zuUn2mxTmF/ekP0gCCgtleZr9+DYKU5hlXk8K10uKxGD6EvoiXZzlfeUuotgp2\nqvI3ysOm/hvCfyEkqhfHtbxjV7j7v7eQFPbvNaXbLa0yr4C4vMK/Z09Ui9JrZ/Z4\ncyIkxfC6/rOqAirSdIp09EGiw7GM8guHyggE4IiZrDslT8V3xIl985cbCxSxeW1R\ntgH4rdEXuVe9+31oJhmXOE9ux2jCop9tEJMgWg7HStrJ5plPbb+HmjoX3nBO04E5\n6m52PyzMNV+2N21IPppKwA==\n-----END CERTIFICATE-----"
		},
		{
			"Id": [1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
			"Address": "52.11.136.238",
			"Tls_certificate": "-----BEGIN CERTIFICATE-----\nMIIDbDCCAlSgAwIBAgIJAOUNtZneIYECMA0GCSqGSIb3DQEBBQUAMGgxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQx\nGzAZBgNVBAoMElByaXZhdGVncml0eSBDb3JwLjETMBEGA1UEAwwKKi5jbWl4LnJp\ncDAeFw0xOTAzMDUxODM1NDNaFw0yOTAzMDIxODM1NDNaMGgxCzAJBgNVBAYTAlVT\nMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQxGzAZBgNV\nBAoMElByaXZhdGVncml0eSBDb3JwLjETMBEGA1UEAwwKKi5jbWl4LnJpcDCCASIw\nDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAPP0WyVkfZA/CEd2DgKpcudn0oDh\nDwsjmx8LBDWsUgQzyLrFiVigfUmUefknUH3dTJjmiJtGqLsayCnWdqWLHPJYvFfs\nWYW0IGF93UG/4N5UAWO4okC3CYgKSi4ekpfw2zgZq0gmbzTnXcHF9gfmQ7jJUKSE\ntJPSNzXq+PZeJTC9zJAb4Lj8QzH18rDM8DaL2y1ns0Y2Hu0edBFn/OqavBJKb/uA\nm3AEjqeOhC7EQUjVamWlTBPt40+B/6aFJX5BYm2JFkRsGBIyBVL46MvC02MgzTT9\nbJIJfwqmBaTruwemNgzGu7Jk03hqqS1TUEvSI6/x8bVoba3orcKkf9HsDjECAwEA\nAaMZMBcwFQYDVR0RBA4wDIIKKi5jbWl4LnJpcDANBgkqhkiG9w0BAQUFAAOCAQEA\nneUocN4AbcQAC1+b3To8u5UGdaGxhcGyZBlAoenRVdjXK3lTjsMdMWb4QctgNfIf\nU/zuUn2mxTmF/ekP0gCCgtleZr9+DYKU5hlXk8K10uKxGD6EvoiXZzlfeUuotgp2\nqvI3ysOm/hvCfyEkqhfHtbxjV7j7v7eQFPbvNaXbLa0yr4C4vMK/Z09Ui9JrZ/Z4\ncyIkxfC6/rOqAirSdIp09EGiw7GM8guHyggE4IiZrDslT8V3xIl985cbCxSxeW1R\ntgH4rdEXuVe9+31oJhmXOE9ux2jCop9tEJMgWg7HStrJ5plPbb+HmjoX3nBO04E5\n6m52PyzMNV+2N21IPppKwA==\n-----END CERTIFICATE-----"
		},
		{
			"Id": [2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
			"Address": "34.213.79.31",
			"Tls_certificate": "-----BEGIN CERTIFICATE-----\nMIIDbDCCAlSgAwIBAgIJAOUNtZneIYECMA0GCSqGSIb3DQEBBQUAMGgxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQx\nGzAZBgNVBAoMElByaXZhdGVncml0eSBDb3JwLjETMBEGA1UEAwwKKi5jbWl4LnJp\ncDAeFw0xOTAzMDUxODM1NDNaFw0yOTAzMDIxODM1NDNaMGgxCzAJBgNVBAYTAlVT\nMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQxGzAZBgNV\nBAoMElByaXZhdGVncml0eSBDb3JwLjETMBEGA1UEAwwKKi5jbWl4LnJpcDCCASIw\nDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAPP0WyVkfZA/CEd2DgKpcudn0oDh\nDwsjmx8LBDWsUgQzyLrFiVigfUmUefknUH3dTJjmiJtGqLsayCnWdqWLHPJYvFfs\nWYW0IGF93UG/4N5UAWO4okC3CYgKSi4ekpfw2zgZq0gmbzTnXcHF9gfmQ7jJUKSE\ntJPSNzXq+PZeJTC9zJAb4Lj8QzH18rDM8DaL2y1ns0Y2Hu0edBFn/OqavBJKb/uA\nm3AEjqeOhC7EQUjVamWlTBPt40+B/6aFJX5BYm2JFkRsGBIyBVL46MvC02MgzTT9\nbJIJfwqmBaTruwemNgzGu7Jk03hqqS1TUEvSI6/x8bVoba3orcKkf9HsDjECAwEA\nAaMZMBcwFQYDVR0RBA4wDIIKKi5jbWl4LnJpcDANBgkqhkiG9w0BAQUFAAOCAQEA\nneUocN4AbcQAC1+b3To8u5UGdaGxhcGyZBlAoenRVdjXK3lTjsMdMWb4QctgNfIf\nU/zuUn2mxTmF/ekP0gCCgtleZr9+DYKU5hlXk8K10uKxGD6EvoiXZzlfeUuotgp2\nqvI3ysOm/hvCfyEkqhfHtbxjV7j7v7eQFPbvNaXbLa0yr4C4vMK/Z09Ui9JrZ/Z4\ncyIkxfC6/rOqAirSdIp09EGiw7GM8guHyggE4IiZrDslT8V3xIl985cbCxSxeW1R\ntgH4rdEXuVe9+31oJhmXOE9ux2jCop9tEJMgWg7HStrJ5plPbb+HmjoX3nBO04E5\n6m52PyzMNV+2N21IPppKwA==\n-----END CERTIFICATE-----"
		}
	],
	"registration": {
		"Address": "92.42.125.61",
		"Tls_certificate": "-----BEGIN CERTIFICATE-----\nMIIDkDCCAnigAwIBAgIJAJnjosuSsP7gMA0GCSqGSIb3DQEBBQUAMHQxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQx\nGzAZBgNVBAoMElByaXZhdGVncml0eSBDb3JwLjEfMB0GA1UEAwwWcmVnaXN0cmF0\naW9uKi5jbWl4LnJpcDAeFw0xOTAzMDUyMTQ5NTZaFw0yOTAzMDIyMTQ5NTZaMHQx\nCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFy\nZW1vbnQxGzAZBgNVBAoMElByaXZhdGVncml0eSBDb3JwLjEfMB0GA1UEAwwWcmVn\naXN0cmF0aW9uKi5jbWl4LnJpcDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC\nggEBAOQKvqjdh35o+MECBhCwopJzPlQNmq2iPbewRNtI02bUNK3kLQUbFlYdzNGZ\nS4GYXGc5O+jdi8Slx82r1kdjz5PPCNFBARIsOP/L8r3DGeW+yeJdgBZjm1s3ylka\nmt4Ajiq/bNjysS6L/WSOp+sVumDxtBEzO/UTU1O6QRnzUphLaiWENmErGvsH0CZV\nq38Ia58k/QjCAzpUcYi4j2l1fb07xqFcQD8H6SmUM297UyQosDrp8ukdIo31Koxr\n4XDnnNNsYStC26tzHMeKuJ2Wl+3YzsSyflfM2YEcKE31sqB9DS36UkJ8J84eLsHN\nImGg3WodFAviDB67+jXDbB30NkMCAwEAAaMlMCMwIQYDVR0RBBowGIIWcmVnaXN0\ncmF0aW9uKi5jbWl4LnJpcDANBgkqhkiG9w0BAQUFAAOCAQEAF9mNzk+g+o626Rll\nt3f3/1qIyYQrYJ0BjSWCKYEFMCgZ4JibAJjAvIajhVYERtltffM+YKcdE2kTpdzJ\n0YJuUnRfuv6sVnXlVVugUUnd4IOigmjbCdM32k170CYMm0aiwGxl4FrNa8ei7AIa\nx/s1n+sqWq3HeW5LXjnoVb+s3HeCWIuLfcgrurfye8FnNhy14HFzxVYYefIKm0XL\n+DPlcGGGm/PPYt3u4a2+rP3xaihc65dTa0u5tf/XPXtPxTDPFj2JeQDFxo7QRREb\nPD89CtYnwuP937CrkvCKrL0GkW1FViXKqZY9F5uhxrvLIpzhbNrs/EbtweY35XGL\nDCCMkg==\n-----END CERTIFICATE-----"
	},
	"notification": {
		"Address": "notification.default.cmix.rip",
		"Tls_certificate": "-----BEGIN CERTIFICATE-----\nMIIDkDCCAnigAwIBAgIJAJnjosuSsP7gMA0GCSqGSIb3DQEBBQUAMHQxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFyZW1vbnQx\nGzAZBgNVBAoMElByaXZhdGVncml0eSBDb3JwLjEfMB0GA1UEAwwWcmVnaXN0cmF0\naW9uKi5jbWl4LnJpcDAeFw0xOTAzMDUyMTQ5NTZaFw0yOTAzMDIyMTQ5NTZaMHQx\nCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlDbGFy\nZW1vbnQxGzAZBgNVBAoMElByaXZhdGVncml0eSBDb3JwLjEfMB0GA1UEAwwWcmVn\naXN0cmF0aW9uKi5jbWl4LnJpcDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC\nggEBAOQKvqjdh35o+MECBhCwopJzPlQNmq2iPbewRNtI02bUNK3kLQUbFlYdzNGZ\nS4GYXGc5O+jdi8Slx82r1kdjz5PPCNFBARIsOP/L8r3DGeW+yeJdgBZjm1s3ylka\nmt4Ajiq/bNjysS6L/WSOp+sVumDxtBEzO/UTU1O6QRnzUphLaiWENmErGvsH0CZV\nq38Ia58k/QjCAzpUcYi4j2l1fb07xqFcQD8H6SmUM297UyQosDrp8ukdIo31Koxr\n4XDnnNNsYStC26tzHMeKuJ2Wl+3YzsSyflfM2YEcKE31sqB9DS36UkJ8J84eLsHN\nImGg3WodFAviDB67+jXDbB30NkMCAwEAAaMlMCMwIQYDVR0RBBowGIIWcmVnaXN0\ncmF0aW9uKi5jbWl4LnJpcDANBgkqhkiG9w0BAQUFAAOCAQEAF9mNzk+g+o626Rll\nt3f3/1qIyYQrYJ0BjSWCKYEFMCgZ4JibAJjAvIajhVYERtltffM+YKcdE2kTpdzJ\n0YJuUnRfuv6sVnXlVVugUUnd4IOigmjbCdM32k170CYMm0aiwGxl4FrNa8ei7AIa\nx/s1n+sqWq3HeW5LXjnoVb+s3HeCWIuLfcgrurfye8FnNhy14HFzxVYYefIKm0XL\n+DPlcGGGm/PPYt3u4a2+rP3xaihc65dTa0u5tf/XPXtPxTDPFj2JeQDFxo7QRREb\nPD89CtYnwuP937CrkvCKrL0GkW1FViXKqZY9F5uhxrvLIpzhbNrs/EbtweY35XGL\nDCCMkg==\n-----END CERTIFICATE-----"
	},
	"udb": {
		"Id": [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3]
	},
	"E2e": {
		"Prime": "FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE45B3DC2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F83655D23DCA3AD961C62F356208552BB9ED529077096966D670C354E4ABC9804F1746C08CA18217C32905E462E36CE3BE39E772C180E86039B2783A2EC07A28FB5C55DF06F4C52C9DE2BCBF6955817183995497CEA956AE515D2261898FA051015728E5A8AACAA68FFFFFFFFFFFFFFFF",
		"Small_prime": "7FFFFFFFFFFFFFFFE487ED5110B4611A62633145C06E0E68948127044533E63A0105DF531D89CD9128A5043CC71A026EF7CA8CD9E69D218D98158536F92F8A1BA7F09AB6B6A8E122F242DABB312F3F637A262174D31BF6B585FFAE5B7A035BF6F71C35FDAD44CFD2D74F9208BE258FF324943328F6722D9EE1003E5C50B1DF82CC6D241B0E2AE9CD348B1FD47E9267AFC1B2AE91EE51D6CB0E3179AB1042A95DCF6A9483B84B4B36B3861AA7255E4C0278BA3604650C10BE19482F23171B671DF1CF3B960C074301CD93C1D17603D147DAE2AEF837A62964EF15E5FB4AAC0B8C1CCAA4BE754AB5728AE9130C4C7D02880AB9472D455655347FFFFFFFFFFFFFFF",
		"Generator": "02"
	},
	"Cmix": {
		"Prime": "FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE45B3DC2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F83655D23DCA3AD961C62F356208552BB9ED529077096966D670C354E4ABC9804F1746C08CA18217C32905E462E36CE3BE39E772C180E86039B2783A2EC07A28FB5C55DF06F4C52C9DE2BCBF6955817183995497CEA956AE515D2261898FA051015728E5A8AACAA68FFFFFFFFFFFFFFFF",
		"Small_prime": "7FFFFFFFFFFFFFFFE487ED5110B4611A62633145C06E0E68948127044533E63A0105DF531D89CD9128A5043CC71A026EF7CA8CD9E69D218D98158536F92F8A1BA7F09AB6B6A8E122F242DABB312F3F637A262174D31BF6B585FFAE5B7A035BF6F71C35FDAD44CFD2D74F9208BE258FF324943328F6722D9EE1003E5C50B1DF82CC6D241B0E2AE9CD348B1FD47E9267AFC1B2AE91EE51D6CB0E3179AB1042A95DCF6A9483B84B4B36B3861AA7255E4C0278BA3604650C10BE19482F23171B671DF1CF3B960C074301CD93C1D17603D147DAE2AEF837A62964EF15E5FB4AAC0B8C1CCAA4BE754AB5728AE9130C4C7D02880AB9472D455655347FFFFFFFFFFFFFFF",
		"Generator": "02"
	}
}`
	ExampleSignature  = `gkh98J10rQiuVsEXd6xe8IeCINplnD93CFpXZFNjT1CgNMxgsHumiC5HsctjnF0xTxDPq3hn3/J0s+eblSVyGMMszTIoWNINVSS1fkm0EGkKafC1vKTZMmc9ivsWL7oY`
	ExampleNDF        = ExampleJSON + "\n" + ExampleSignature
	NDF, _, _         = ndf.DecodeNDF(ExampleNDF)
	EmptyNdf, _, _    = ndf.DecodeNDF("")
	JsonBytes, _      = base64.StdEncoding.DecodeString("AQAAAA7UiTKgAAAAAP5cNTIuMjUuMTM1LjUyLS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURnVENDQW1tZ0F3SUJBZ0lKQUtMZFo4VWlnSUFlTUEwR0NTcUdTSWIzRFFFQkJRVUFNRzh4Q3pBSkJnTlYKQkFZVEFsVlRNUk13RVFZRFZRUUlEQXBEWVd4cFptOXlibWxoTVJJd0VBWURWUVFIREFsRGJHRnlaVzF2Ym5ReApHekFaQmdOVkJBb01FbEJ5YVhaaGRHVm5jbWwwZVNCRGIzSndMakVhTUJnR0ExVUVBd3dSWjJGMFpYZGhlU291ClkyMXBlQzV5YVhBd0hoY05NVGt3TXpBMU1UZ3pOVFUwV2hjTk1qa3dNekF5TVRnek5UVTBXakJ2TVFzd0NRWUQKVlFRR0V3SlZVekVUTUJFR0ExVUVDQXdLUTJGc2FXWnZjbTVwWVRFU01CQUdBMVVFQnd3SlEyeGhjbVZ0YjI1MApNUnN3R1FZRFZRUUtEQkpRY21sMllYUmxaM0pwZEhrZ1EyOXljQzR4R2pBWUJnTlZCQU1NRVdkaGRHVjNZWGtxCkxtTnRhWGd1Y21sd01JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBOStBYXh3RFAKeEhiaExtbjRIb1p1MG9VTTQ4UXVmYzZUNVhFWlRycE1ycUpBb3VYays2MUpjMEVGSDk2L3NiajdWeXZuWFBSbwpnSUVOYmsyWTg0QmtCOVNrUk1JWHlhL2doOWRPRURTZ252ai95ZzI0bDNiZEtGcUJNS2lGZzAwUFlCMzBmVStBCmJlM09JL2xlMEkrdisrUndIMkFWMEJNcStUNlBjQUdqQ0MxUTFaQjB3UDkvVnFOTVdxNWxiSzl3RDQ2SVFpU2kKK1NnSVFlRTdIb2lBWlhyR08wWTdsOVAzK1ZSb1hqUlFicWZuM0VUTkw5WnZRdWFyd0FZQzlJeDVNeFVyUzVhZwpPbWZqYzhiZmtwWURGQVhSWG1kS05JU0ptdENlYlgya0RycFA4QmRhc3g3RnpzeDU5Y0VVSENsMmFKT1dYYzdSCjVtM2p1T1ZMMUhVeGpRSURBUUFCb3lBd0hqQWNCZ05WSFJFRUZUQVRnaEZuWVhSbGQyRjVLaTVqYldsNExuSnAKY0RBTkJna3Foa2lHOXcwQkFRVUZBQU9DQVFFQU11M3hvYzJMVzJVRXhBQUlZWVdFRVRnZ0xOcmxHb254dGVTdQpqdUpqT1IraWs1U1ZMbjBsRXUyMit6K0ZDQTdnU2s5RmtXdSt2OXFuZk9mbTJBbStXS1lXdjNkSjVSeXBXL2hECk5Ya09ZeFZKTllGeGVTaG5Ib2hOcXE0ZURLcGRxU3hFY3VFckZYSmRMYlpQMXVOczRXSU9LblRoZ3poa3B1eTcKdFpSb3N2T0YxWDV1TDFmclZKekhONWpBU0VEQWE3aEpObVEyNGtoK2RzL0dlMzlmR0Q4cEszMUNXaG5JWGVEbwp2S0Q3d2l2aS9nU09CdGNSV1dMdlU4U2l6WmtTM2hnVHcwbFNPZjVnZXV6dmFzQ0VZbHFyS0Zzc2o2Y1R6YkNCCnh5M3JhM1dhelJUTlRXNFRta0hsQ1VDOUkzb1dUVHh3NWlReEYvSTJrUVFud1I3TDN3PT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLTUyLjI1LjIxOS4zOC0tLS0tQkVHSU4gQ0VSVElGSUNBVEUtLS0tLQpNSUlEZ1RDQ0FtbWdBd0lCQWdJSkFLTGRaOFVpZ0lBZU1BMEdDU3FHU0liM0RRRUJCUVVBTUc4eEN6QUpCZ05WCkJBWVRBbFZUTVJNd0VRWURWUVFJREFwRFlXeHBabTl5Ym1saE1SSXdFQVlEVlFRSERBbERiR0Z5WlcxdmJuUXgKR3pBWkJnTlZCQW9NRWxCeWFYWmhkR1ZuY21sMGVTQkRiM0p3TGpFYU1CZ0dBMVVFQXd3UloyRjBaWGRoZVNvdQpZMjFwZUM1eWFYQXdIaGNOTVRrd016QTFNVGd6TlRVMFdoY05Namt3TXpBeU1UZ3pOVFUwV2pCdk1Rc3dDUVlEClZRUUdFd0pWVXpFVE1CRUdBMVVFQ0F3S1EyRnNhV1p2Y201cFlURVNNQkFHQTFVRUJ3d0pRMnhoY21WdGIyNTAKTVJzd0dRWURWUVFLREJKUWNtbDJZWFJsWjNKcGRIa2dRMjl5Y0M0eEdqQVlCZ05WQkFNTUVXZGhkR1YzWVhrcQpMbU50YVhndWNtbHdNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQTkrQWF4d0RQCnhIYmhMbW40SG9adTBvVU00OFF1ZmM2VDVYRVpUcnBNcnFKQW91WGsrNjFKYzBFRkg5Ni9zYmo3Vnl2blhQUm8KZ0lFTmJrMlk4NEJrQjlTa1JNSVh5YS9naDlkT0VEU2dudmoveWcyNGwzYmRLRnFCTUtpRmcwMFBZQjMwZlUrQQpiZTNPSS9sZTBJK3YrK1J3SDJBVjBCTXErVDZQY0FHakNDMVExWkIwd1A5L1ZxTk1XcTVsYks5d0Q0NklRaVNpCitTZ0lRZUU3SG9pQVpYckdPMFk3bDlQMytWUm9YalJRYnFmbjNFVE5MOVp2UXVhcndBWUM5SXg1TXhVclM1YWcKT21mamM4YmZrcFlERkFYUlhtZEtOSVNKbXRDZWJYMmtEcnBQOEJkYXN4N0Z6c3g1OWNFVUhDbDJhSk9XWGM3Ugo1bTNqdU9WTDFIVXhqUUlEQVFBQm95QXdIakFjQmdOVkhSRUVGVEFUZ2hGbllYUmxkMkY1S2k1amJXbDRMbkpwCmNEQU5CZ2txaGtpRzl3MEJBUVVGQUFPQ0FRRUFNdTN4b2MyTFcyVUV4QUFJWVlXRUVUZ2dMTnJsR29ueHRlU3UKanVKak9SK2lrNVNWTG4wbEV1MjIreitGQ0E3Z1NrOUZrV3UrdjlxbmZPZm0yQW0rV0tZV3YzZEo1UnlwVy9oRApOWGtPWXhWSk5ZRnhlU2huSG9oTnFxNGVES3BkcVN4RWN1RXJGWEpkTGJaUDF1TnM0V0lPS25UaGd6aGtwdXk3CnRaUm9zdk9GMVg1dUwxZnJWSnpITjVqQVNFREFhN2hKTm1RMjRraCtkcy9HZTM5ZkdEOHBLMzFDV2huSVhlRG8KdktEN3dpdmkvZ1NPQnRjUldXTHZVOFNpelprUzNoZ1R3MGxTT2Y1Z2V1enZhc0NFWWxxcktGc3NqNmNUemJDQgp4eTNyYTNXYXpSVE5UVzRUbWtIbENVQzlJM29XVFR4dzVpUXhGL0kya1FRbndSN0wzdz09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS01Mi40MS44MC4xMDQtLS0tLUJFR0lOIENFUlRJRklDQVRFLS0tLS0KTUlJRGdUQ0NBbW1nQXdJQkFnSUpBS0xkWjhVaWdJQWVNQTBHQ1NxR1NJYjNEUUVCQlFVQU1HOHhDekFKQmdOVgpCQVlUQWxWVE1STXdFUVlEVlFRSURBcERZV3hwWm05eWJtbGhNUkl3RUFZRFZRUUhEQWxEYkdGeVpXMXZiblF4Ckd6QVpCZ05WQkFvTUVsQnlhWFpoZEdWbmNtbDBlU0JEYjNKd0xqRWFNQmdHQTFVRUF3d1JaMkYwWlhkaGVTb3UKWTIxcGVDNXlhWEF3SGhjTk1Ua3dNekExTVRnek5UVTBXaGNOTWprd016QXlNVGd6TlRVMFdqQnZNUXN3Q1FZRApWUVFHRXdKVlV6RVRNQkVHQTFVRUNBd0tRMkZzYVdadmNtNXBZVEVTTUJBR0ExVUVCd3dKUTJ4aGNtVnRiMjUwCk1Sc3dHUVlEVlFRS0RCSlFjbWwyWVhSbFozSnBkSGtnUTI5eWNDNHhHakFZQmdOVkJBTU1FV2RoZEdWM1lYa3EKTG1OdGFYZ3VjbWx3TUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUE5K0FheHdEUAp4SGJoTG1uNEhvWnUwb1VNNDhRdWZjNlQ1WEVaVHJwTXJxSkFvdVhrKzYxSmMwRUZIOTYvc2JqN1Z5dm5YUFJvCmdJRU5iazJZODRCa0I5U2tSTUlYeWEvZ2g5ZE9FRFNnbnZqL3lnMjRsM2JkS0ZxQk1LaUZnMDBQWUIzMGZVK0EKYmUzT0kvbGUwSSt2KytSd0gyQVYwQk1xK1Q2UGNBR2pDQzFRMVpCMHdQOS9WcU5NV3E1bGJLOXdENDZJUWlTaQorU2dJUWVFN0hvaUFaWHJHTzBZN2w5UDMrVlJvWGpSUWJxZm4zRVROTDladlF1YXJ3QVlDOUl4NU14VXJTNWFnCk9tZmpjOGJma3BZREZBWFJYbWRLTklTSm10Q2ViWDJrRHJwUDhCZGFzeDdGenN4NTljRVVIQ2wyYUpPV1hjN1IKNW0zanVPVkwxSFV4alFJREFRQUJveUF3SGpBY0JnTlZIUkVFRlRBVGdoRm5ZWFJsZDJGNUtpNWpiV2w0TG5KcApjREFOQmdrcWhraUc5dzBCQVFVRkFBT0NBUUVBTXUzeG9jMkxXMlVFeEFBSVlZV0VFVGdnTE5ybEdvbnh0ZVN1Cmp1SmpPUitpazVTVkxuMGxFdTIyK3orRkNBN2dTazlGa1d1K3Y5cW5mT2ZtMkFtK1dLWVd2M2RKNVJ5cFcvaEQKTlhrT1l4VkpOWUZ4ZVNobkhvaE5xcTRlREtwZHFTeEVjdUVyRlhKZExiWlAxdU5zNFdJT0tuVGhnemhrcHV5Nwp0WlJvc3ZPRjFYNXVMMWZyVkp6SE41akFTRURBYTdoSk5tUTI0a2grZHMvR2UzOWZHRDhwSzMxQ1dobklYZURvCnZLRDd3aXZpL2dTT0J0Y1JXV0x2VThTaXpaa1MzaGdUdzBsU09mNWdldXp2YXNDRVlscXJLRnNzajZjVHpiQ0IKeHkzcmEzV2F6UlROVFc0VG1rSGxDVUM5STNvV1RUeHc1aVF4Ri9JMmtRUW53UjdMM3c9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAxOC4yMzcuMTQ3LjEwNS0tLS0tQkVHSU4gQ0VSVElGSUNBVEUtLS0tLQpNSUlEYkRDQ0FsU2dBd0lCQWdJSkFPVU50Wm5lSVlFQ01BMEdDU3FHU0liM0RRRUJCUVVBTUdneEN6QUpCZ05WCkJBWVRBbFZUTVJNd0VRWURWUVFJREFwRFlXeHBabTl5Ym1saE1SSXdFQVlEVlFRSERBbERiR0Z5WlcxdmJuUXgKR3pBWkJnTlZCQW9NRWxCeWFYWmhkR1ZuY21sMGVTQkRiM0p3TGpFVE1CRUdBMVVFQXd3S0tpNWpiV2w0TG5KcApjREFlRncweE9UQXpNRFV4T0RNMU5ETmFGdzB5T1RBek1ESXhPRE0xTkROYU1HZ3hDekFKQmdOVkJBWVRBbFZUCk1STXdFUVlEVlFRSURBcERZV3hwWm05eWJtbGhNUkl3RUFZRFZRUUhEQWxEYkdGeVpXMXZiblF4R3pBWkJnTlYKQkFvTUVsQnlhWFpoZEdWbmNtbDBlU0JEYjNKd0xqRVRNQkVHQTFVRUF3d0tLaTVqYldsNExuSnBjRENDQVNJdwpEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBUFAwV3lWa2ZaQS9DRWQyRGdLcGN1ZG4wb0RoCkR3c2pteDhMQkRXc1VnUXp5THJGaVZpZ2ZVbVVlZmtuVUgzZFRKam1pSnRHcUxzYXlDbldkcVdMSFBKWXZGZnMKV1lXMElHRjkzVUcvNE41VUFXTzRva0MzQ1lnS1NpNGVrcGZ3MnpnWnEwZ21ielRuWGNIRjlnZm1RN2pKVUtTRQp0SlBTTnpYcStQWmVKVEM5ekpBYjRMajhRekgxOHJETThEYUwyeTFuczBZMkh1MGVkQkZuL09xYXZCSktiL3VBCm0zQUVqcWVPaEM3RVFValZhbVdsVEJQdDQwK0IvNmFGSlg1QlltMkpGa1JzR0JJeUJWTDQ2TXZDMDJNZ3pUVDkKYkpJSmZ3cW1CYVRydXdlbU5nekd1N0prMDNocXFTMVRVRXZTSTYveDhiVm9iYTNvcmNLa2Y5SHNEakVDQXdFQQpBYU1aTUJjd0ZRWURWUjBSQkE0d0RJSUtLaTVqYldsNExuSnBjREFOQmdrcWhraUc5dzBCQVFVRkFBT0NBUUVBCm5lVW9jTjRBYmNRQUMxK2IzVG84dTVVR2RhR3hoY0d5WkJsQW9lblJWZGpYSzNsVGpzTWRNV2I0UWN0Z05mSWYKVS96dVVuMm14VG1GL2VrUDBnQ0NndGxlWnI5K0RZS1U1aGxYazhLMTB1S3hHRDZFdm9pWFp6bGZlVXVvdGdwMgpxdkkzeXNPbS9odkNmeUVrcWhmSHRieGpWN2o3djdlUUZQYnZOYVhiTGEweXI0QzR2TUsvWjA5VWk5SnJaL1o0CmN5SWt4ZkM2L3JPcUFpclNkSXAwOUVHaXc3R004Z3VIeWdnRTRJaVpyRHNsVDhWM3hJbDk4NWNiQ3hTeGVXMVIKdGdINHJkRVh1VmU5KzMxb0pobVhPRTl1eDJqQ29wOXRFSk1nV2c3SFN0cko1cGxQYmIrSG1qb1gzbkJPMDRFNQo2bTUyUHl6TU5WKzJOMjFJUHBwS3dBPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAANTIuMTEuMTM2LjIzOC0tLS0tQkVHSU4gQ0VSVElGSUNBVEUtLS0tLQpNSUlEYkRDQ0FsU2dBd0lCQWdJSkFPVU50Wm5lSVlFQ01BMEdDU3FHU0liM0RRRUJCUVVBTUdneEN6QUpCZ05WCkJBWVRBbFZUTVJNd0VRWURWUVFJREFwRFlXeHBabTl5Ym1saE1SSXdFQVlEVlFRSERBbERiR0Z5WlcxdmJuUXgKR3pBWkJnTlZCQW9NRWxCeWFYWmhkR1ZuY21sMGVTQkRiM0p3TGpFVE1CRUdBMVVFQXd3S0tpNWpiV2w0TG5KcApjREFlRncweE9UQXpNRFV4T0RNMU5ETmFGdzB5T1RBek1ESXhPRE0xTkROYU1HZ3hDekFKQmdOVkJBWVRBbFZUCk1STXdFUVlEVlFRSURBcERZV3hwWm05eWJtbGhNUkl3RUFZRFZRUUhEQWxEYkdGeVpXMXZiblF4R3pBWkJnTlYKQkFvTUVsQnlhWFpoZEdWbmNtbDBlU0JEYjNKd0xqRVRNQkVHQTFVRUF3d0tLaTVqYldsNExuSnBjRENDQVNJdwpEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBUFAwV3lWa2ZaQS9DRWQyRGdLcGN1ZG4wb0RoCkR3c2pteDhMQkRXc1VnUXp5THJGaVZpZ2ZVbVVlZmtuVUgzZFRKam1pSnRHcUxzYXlDbldkcVdMSFBKWXZGZnMKV1lXMElHRjkzVUcvNE41VUFXTzRva0MzQ1lnS1NpNGVrcGZ3MnpnWnEwZ21ielRuWGNIRjlnZm1RN2pKVUtTRQp0SlBTTnpYcStQWmVKVEM5ekpBYjRMajhRekgxOHJETThEYUwyeTFuczBZMkh1MGVkQkZuL09xYXZCSktiL3VBCm0zQUVqcWVPaEM3RVFValZhbVdsVEJQdDQwK0IvNmFGSlg1QlltMkpGa1JzR0JJeUJWTDQ2TXZDMDJNZ3pUVDkKYkpJSmZ3cW1CYVRydXdlbU5nekd1N0prMDNocXFTMVRVRXZTSTYveDhiVm9iYTNvcmNLa2Y5SHNEakVDQXdFQQpBYU1aTUJjd0ZRWURWUjBSQkE0d0RJSUtLaTVqYldsNExuSnBjREFOQmdrcWhraUc5dzBCQVFVRkFBT0NBUUVBCm5lVW9jTjRBYmNRQUMxK2IzVG84dTVVR2RhR3hoY0d5WkJsQW9lblJWZGpYSzNsVGpzTWRNV2I0UWN0Z05mSWYKVS96dVVuMm14VG1GL2VrUDBnQ0NndGxlWnI5K0RZS1U1aGxYazhLMTB1S3hHRDZFdm9pWFp6bGZlVXVvdGdwMgpxdkkzeXNPbS9odkNmeUVrcWhmSHRieGpWN2o3djdlUUZQYnZOYVhiTGEweXI0QzR2TUsvWjA5VWk5SnJaL1o0CmN5SWt4ZkM2L3JPcUFpclNkSXAwOUVHaXc3R004Z3VIeWdnRTRJaVpyRHNsVDhWM3hJbDk4NWNiQ3hTeGVXMVIKdGdINHJkRVh1VmU5KzMxb0pobVhPRTl1eDJqQ29wOXRFSk1nV2c3SFN0cko1cGxQYmIrSG1qb1gzbkJPMDRFNQo2bTUyUHl6TU5WKzJOMjFJUHBwS3dBPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMzQuMjEzLjc5LjMxLS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURiRENDQWxTZ0F3SUJBZ0lKQU9VTnRabmVJWUVDTUEwR0NTcUdTSWIzRFFFQkJRVUFNR2d4Q3pBSkJnTlYKQkFZVEFsVlRNUk13RVFZRFZRUUlEQXBEWVd4cFptOXlibWxoTVJJd0VBWURWUVFIREFsRGJHRnlaVzF2Ym5ReApHekFaQmdOVkJBb01FbEJ5YVhaaGRHVm5jbWwwZVNCRGIzSndMakVUTUJFR0ExVUVBd3dLS2k1amJXbDRMbkpwCmNEQWVGdzB4T1RBek1EVXhPRE0xTkROYUZ3MHlPVEF6TURJeE9ETTFORE5hTUdneEN6QUpCZ05WQkFZVEFsVlQKTVJNd0VRWURWUVFJREFwRFlXeHBabTl5Ym1saE1SSXdFQVlEVlFRSERBbERiR0Z5WlcxdmJuUXhHekFaQmdOVgpCQW9NRWxCeWFYWmhkR1ZuY21sMGVTQkRiM0p3TGpFVE1CRUdBMVVFQXd3S0tpNWpiV2w0TG5KcGNEQ0NBU0l3CkRRWUpLb1pJaHZjTkFRRUJCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFQUDBXeVZrZlpBL0NFZDJEZ0twY3VkbjBvRGgKRHdzam14OExCRFdzVWdRenlMckZpVmlnZlVtVWVma25VSDNkVEpqbWlKdEdxTHNheUNuV2RxV0xIUEpZdkZmcwpXWVcwSUdGOTNVRy80TjVVQVdPNG9rQzNDWWdLU2k0ZWtwZncyemdacTBnbWJ6VG5YY0hGOWdmbVE3akpVS1NFCnRKUFNOelhxK1BaZUpUQzl6SkFiNExqOFF6SDE4ckRNOERhTDJ5MW5zMFkySHUwZWRCRm4vT3FhdkJKS2IvdUEKbTNBRWpxZU9oQzdFUVVqVmFtV2xUQlB0NDArQi82YUZKWDVCWW0ySkZrUnNHQkl5QlZMNDZNdkMwMk1nelRUOQpiSklKZndxbUJhVHJ1d2VtTmd6R3U3SmswM2hxcVMxVFVFdlNJNi94OGJWb2JhM29yY0trZjlIc0RqRUNBd0VBCkFhTVpNQmN3RlFZRFZSMFJCQTR3RElJS0tpNWpiV2w0TG5KcGNEQU5CZ2txaGtpRzl3MEJBUVVGQUFPQ0FRRUEKbmVVb2NONEFiY1FBQzErYjNUbzh1NVVHZGFHeGhjR3laQmxBb2VuUlZkalhLM2xUanNNZE1XYjRRY3RnTmZJZgpVL3p1VW4ybXhUbUYvZWtQMGdDQ2d0bGVacjkrRFlLVTVobFhrOEsxMHVLeEdENkV2b2lYWnpsZmVVdW90Z3AyCnF2STN5c09tL2h2Q2Z5RWtxaGZIdGJ4alY3ajd2N2VRRlBidk5hWGJMYTB5cjRDNHZNSy9aMDlVaTlKclovWjQKY3lJa3hmQzYvck9xQWlyU2RJcDA5RUdpdzdHTThndUh5Z2dFNElpWnJEc2xUOFYzeElsOTg1Y2JDeFN4ZVcxUgp0Z0g0cmRFWHVWZTkrMzFvSmhtWE9FOXV4MmpDb3A5dEVKTWdXZzdIU3RySjVwbFBiYitIbWpvWDNuQk8wNEU1CjZtNTJQeXpNTlYrMk4yMUlQcHBLd0E9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tcmVnaXN0cmF0aW9uLmRlZmF1bHQuY21peC5yaXAtLS0tLUJFR0lOIENFUlRJRklDQVRFLS0tLS0KTUlJRGtEQ0NBbmlnQXdJQkFnSUpBSm5qb3N1U3NQN2dNQTBHQ1NxR1NJYjNEUUVCQlFVQU1IUXhDekFKQmdOVgpCQVlUQWxWVE1STXdFUVlEVlFRSURBcERZV3hwWm05eWJtbGhNUkl3RUFZRFZRUUhEQWxEYkdGeVpXMXZiblF4Ckd6QVpCZ05WQkFvTUVsQnlhWFpoZEdWbmNtbDBlU0JEYjNKd0xqRWZNQjBHQTFVRUF3d1djbVZuYVhOMGNtRjAKYVc5dUtpNWpiV2w0TG5KcGNEQWVGdzB4T1RBek1EVXlNVFE1TlRaYUZ3MHlPVEF6TURJeU1UUTVOVFphTUhReApDekFKQmdOVkJBWVRBbFZUTVJNd0VRWURWUVFJREFwRFlXeHBabTl5Ym1saE1SSXdFQVlEVlFRSERBbERiR0Z5ClpXMXZiblF4R3pBWkJnTlZCQW9NRWxCeWFYWmhkR1ZuY21sMGVTQkRiM0p3TGpFZk1CMEdBMVVFQXd3V2NtVm4KYVhOMGNtRjBhVzl1S2k1amJXbDRMbkpwY0RDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQwpnZ0VCQU9RS3ZxamRoMzVvK01FQ0JoQ3dvcEp6UGxRTm1xMmlQYmV3Uk50STAyYlVOSzNrTFFVYkZsWWR6TkdaClM0R1lYR2M1TytqZGk4U2x4ODJyMWtkano1UFBDTkZCQVJJc09QL0w4cjNER2VXK3llSmRnQlpqbTFzM3lsa2EKbXQ0QWppcS9iTmp5c1M2TC9XU09wK3NWdW1EeHRCRXpPL1VUVTFPNlFSbnpVcGhMYWlXRU5tRXJHdnNIMENaVgpxMzhJYTU4ay9RakNBenBVY1lpNGoybDFmYjA3eHFGY1FEOEg2U21VTTI5N1V5UW9zRHJwOHVrZElvMzFLb3hyCjRYRG5uTk5zWVN0QzI2dHpITWVLdUoyV2wrM1l6c1N5ZmxmTTJZRWNLRTMxc3FCOURTMzZVa0o4Sjg0ZUxzSE4KSW1HZzNXb2RGQXZpREI2NytqWERiQjMwTmtNQ0F3RUFBYU1sTUNNd0lRWURWUjBSQkJvd0dJSVdjbVZuYVhOMApjbUYwYVc5dUtpNWpiV2w0TG5KcGNEQU5CZ2txaGtpRzl3MEJBUVVGQUFPQ0FRRUFGOW1OemsrZytvNjI2UmxsCnQzZjMvMXFJeVlRcllKMEJqU1dDS1lFRk1DZ1o0SmliQUpqQXZJYWpoVllFUnRsdGZmTStZS2NkRTJrVHBkekoKMFlKdVVuUmZ1djZzVm5YbFZWdWdVVW5kNElPaWdtamJDZE0zMmsxNzBDWU1tMGFpd0d4bDRGck5hOGVpN0FJYQp4L3MxbitzcVdxM0hlVzVMWGpub1ZiK3MzSGVDV0l1TGZjZ3J1cmZ5ZThGbk5oeTE0SEZ6eFZZWWVmSUttMFhMCitEUGxjR0dHbS9QUFl0M3U0YTIrclAzeGFpaGM2NWRUYTB1NXRmL1hQWHRQeFREUEZqMkplUURGeG83UVJSRWIKUEQ4OUN0WW53dVA5MzdDcmt2Q0tyTDBHa1cxRlZpWEtxWlk5RjV1aHhydkxJcHpoYk5ycy9FYnR3ZVkzNVhHTApEQ0NNa2c9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAANGRkZGRkZGRkZGRkZGRkZGQzkwRkRBQTIyMTY4QzIzNEM0QzY2MjhCODBEQzFDRDEyOTAyNEUwODhBNjdDQzc0MDIwQkJFQTYzQjEzOUIyMjUxNEEwODc5OEUzNDA0RERFRjk1MTlCM0NEM0E0MzFCMzAyQjBBNkRGMjVGMTQzNzRGRTEzNTZENkQ1MUMyNDVFNDg1QjU3NjYyNUU3RUM2RjQ0QzQyRTlBNjM3RUQ2QjBCRkY1Q0I2RjQwNkI3RURFRTM4NkJGQjVBODk5RkE1QUU5RjI0MTE3QzRCMUZFNjQ5Mjg2NjUxRUNFNDVCM0RDMjAwN0NCOEExNjNCRjA1OThEQTQ4MzYxQzU1RDM5QTY5MTYzRkE4RkQyNENGNUY4MzY1NUQyM0RDQTNBRDk2MUM2MkYzNTYyMDg1NTJCQjlFRDUyOTA3NzA5Njk2NkQ2NzBDMzU0RTRBQkM5ODA0RjE3NDZDMDhDQTE4MjE3QzMyOTA1RTQ2MkUzNkNFM0JFMzlFNzcyQzE4MEU4NjAzOUIyNzgzQTJFQzA3QTI4RkI1QzU1REYwNkY0QzUyQzlERTJCQ0JGNjk1NTgxNzE4Mzk5NTQ5N0NFQTk1NkFFNTE1RDIyNjE4OThGQTA1MTAxNTcyOEU1QThBQUNBQTY4RkZGRkZGRkZGRkZGRkZGRjAyN0ZGRkZGRkZGRkZGRkZGRkU0ODdFRDUxMTBCNDYxMUE2MjYzMzE0NUMwNkUwRTY4OTQ4MTI3MDQ0NTMzRTYzQTAxMDVERjUzMUQ4OUNEOTEyOEE1MDQzQ0M3MUEwMjZFRjdDQThDRDlFNjlEMjE4RDk4MTU4NTM2RjkyRjhBMUJBN0YwOUFCNkI2QThFMTIyRjI0MkRBQkIzMTJGM0Y2MzdBMjYyMTc0RDMxQkY2QjU4NUZGQUU1QjdBMDM1QkY2RjcxQzM1RkRBRDQ0Q0ZEMkQ3NEY5MjA4QkUyNThGRjMyNDk0MzMyOEY2NzIyRDlFRTEwMDNFNUM1MEIxREY4MkNDNkQyNDFCMEUyQUU5Q0QzNDhCMUZENDdFOTI2N0FGQzFCMkFFOTFFRTUxRDZDQjBFMzE3OUFCMTA0MkE5NURDRjZBOTQ4M0I4NEI0QjM2QjM4NjFBQTcyNTVFNEMwMjc4QkEzNjA0NjUwQzEwQkUxOTQ4MkYyMzE3MUI2NzFERjFDRjNCOTYwQzA3NDMwMUNEOTNDMUQxNzYwM0QxNDdEQUUyQUVGODM3QTYyOTY0RUYxNUU1RkI0QUFDMEI4QzFDQ0FBNEJFNzU0QUI1NzI4QUU5MTMwQzRDN0QwMjg4MEFCOTQ3MkQ0NTU2NTUzNDdGRkZGRkZGRkZGRkZGRkZGRkZGRkZGRkZGRkZGRkZGQzkwRkRBQTIyMTY4QzIzNEM0QzY2MjhCODBEQzFDRDEyOTAyNEUwODhBNjdDQzc0MDIwQkJFQTYzQjEzOUIyMjUxNEEwODc5OEUzNDA0RERFRjk1MTlCM0NEM0E0MzFCMzAyQjBBNkRGMjVGMTQzNzRGRTEzNTZENkQ1MUMyNDVFNDg1QjU3NjYyNUU3RUM2RjQ0QzQyRTlBNjM3RUQ2QjBCRkY1Q0I2RjQwNkI3RURFRTM4NkJGQjVBODk5RkE1QUU5RjI0MTE3QzRCMUZFNjQ5Mjg2NjUxRUNFNDVCM0RDMjAwN0NCOEExNjNCRjA1OThEQTQ4MzYxQzU1RDM5QTY5MTYzRkE4RkQyNENGNUY4MzY1NUQyM0RDQTNBRDk2MUM2MkYzNTYyMDg1NTJCQjlFRDUyOTA3NzA5Njk2NkQ2NzBDMzU0RTRBQkM5ODA0RjE3NDZDMDhDQTE4MjE3QzMyOTA1RTQ2MkUzNkNFM0JFMzlFNzcyQzE4MEU4NjAzOUIyNzgzQTJFQzA3QTI4RkI1QzU1REYwNkY0QzUyQzlERTJCQ0JGNjk1NTgxNzE4Mzk5NTQ5N0NFQTk1NkFFNTE1RDIyNjE4OThGQTA1MTAxNTcyOEU1QThBQUNBQTY4RkZGRkZGRkZGRkZGRkZGRjAyN0ZGRkZGRkZGRkZGRkZGRkU0ODdFRDUxMTBCNDYxMUE2MjYzMzE0NUMwNkUwRTY4OTQ4MTI3MDQ0NTMzRTYzQTAxMDVERjUzMUQ4OUNEOTEyOEE1MDQzQ0M3MUEwMjZFRjdDQThDRDlFNjlEMjE4RDk4MTU4NTM2RjkyRjhBMUJBN0YwOUFCNkI2QThFMTIyRjI0MkRBQkIzMTJGM0Y2MzdBMjYyMTc0RDMxQkY2QjU4NUZGQUU1QjdBMDM1QkY2RjcxQzM1RkRBRDQ0Q0ZEMkQ3NEY5MjA4QkUyNThGRjMyNDk0MzMyOEY2NzIyRDlFRTEwMDNFNUM1MEIxREY4MkNDNkQyNDFCMEUyQUU5Q0QzNDhCMUZENDdFOTI2N0FGQzFCMkFFOTFFRTUxRDZDQjBFMzE3OUFCMTA0MkE5NURDRjZBOTQ4M0I4NEI0QjM2QjM4NjFBQTcyNTVFNEMwMjc4QkEzNjA0NjUwQzEwQkUxOTQ4MkYyMzE3MUI2NzFERjFDRjNCOTYwQzA3NDMwMUNEOTNDMUQxNzYwM0QxNDdEQUUyQUVGODM3QTYyOTY0RUYxNUU1RkI0QUFDMEI4QzFDQ0FBNEJFNzU0QUI1NzI4QUU5MTMwQzRDN0QwMjg4MEFCOTQ3MkQ0NTU2NTUzNDdGRkZGRkZGRkZGRkZGRkY=")
	SignatureBytes, _ = base64.StdEncoding.DecodeString(ExampleSignature)
	E2eGrp            = NDF.E2E.String()
	CmixGrp           = NDF.CMIX.String()
)
