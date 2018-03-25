package executor

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
	"github.com/seriousben/net-script/internal/types"
)

var ignoredHeaders = []string{"Content-Length"}

func removeIgnoredHeaders(h http.Header) {
	for _, hk := range ignoredHeaders {
		h.Del(hk)
	}
}

type echoHTTPPayload struct {
	Path    string
	Method  string
	Headers http.Header
	Body    []byte
}

func startEchoHTTPServer(t *testing.T) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rb, err := ioutil.ReadAll(r.Body)
		assert.NilError(t, err)

		payload := echoHTTPPayload{
			Path:    r.URL.Path,
			Method:  r.Method,
			Headers: r.Header,
			Body:    rb,
		}
		b, err := json.Marshal(payload)
		assert.NilError(t, err)
		_, err = w.Write(b)
		assert.NilError(t, err)
	}))
}

func TestExecHTTP(t *testing.T) {
	srv := startEchoHTTPServer(t)
	defer srv.Close()

	tests := []struct {
		name    string
		command types.Command
	}{
		{
			"simple GET",
			types.Command{
				Method:  "GET",
				URL:     srv.URL,
				Headers: http.Header{"Auth": []string{"authorization data"}},
				Body:    []byte("LINE 1\nLINE 2\nLINE 3\n"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out, err := Execute(test.command)
			assert.NilError(t, err)

			resp := out.Response.(*http.Response)
			assert.Equal(t, resp.StatusCode, 200)

			payload := new(echoHTTPPayload)
			err = json.Unmarshal(out.Body, &payload)
			assert.NilError(t, err)

			removeIgnoredHeaders(payload.Headers)

			assert.DeepEqual(t, payload.Headers, test.command.Headers)
			assert.DeepEqual(t, payload.Body, test.command.Body)
		})
	}
}

func startTCPServer(t *testing.T) net.Listener {
	t.Helper()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		assert.NilError(t, err)
	}

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			assert.NilError(t, err)
		}

		body := make([]byte, 256)
		_, err = conn.Read(body)
		if err != nil {
			assert.NilError(t, err)
		}

		_, err = conn.Write(bytes.Trim(body, "\x00"))
		if err != nil {
			assert.NilError(t, err)
		}
	}()

	return ln
}

func TestExecuteTCP(t *testing.T) {
	srv := startTCPServer(t)
	defer srv.Close()

	tests := []struct {
		name    string
		command types.Command
	}{
		{
			"simple TCP",
			types.Command{
				Method: "TCP",
				URL:    srv.Addr().String(),
				Body:   []byte("LINE 1\nLINE 2\nLINE 3\n"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out, err := Execute(test.command)
			assert.NilError(t, err)
			assert.DeepEqual(t, out.Body, test.command.Body)
		})
	}
}

func TestExecuteWS(t *testing.T) {
	t.Skip()
	srv := startTCPServer(t)
	defer srv.Close()

	tests := []struct {
		name    string
		command types.Command
	}{
		{
			"simple WS",
			types.Command{
				Method: "WS",
				URL:    srv.Addr().String(),
				Body:   []byte("LINE 1\nLINE 2\nLINE 3\n"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out, err := Execute(test.command)
			assert.NilError(t, err)
			assert.DeepEqual(t, out.Body, test.command.Body)
		})
	}
}
