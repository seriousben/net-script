package executor

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/seriousben/net-script/internal/types"
)

var defaultTransport = &http.Transport{
	// Remove "Accept-Encoding: gzip"
	DisableCompression: true,
}

// ExecuteMany executes multiple commands.
func ExecuteMany(cmds []types.Command) ([]types.CommandOutput, error) {
	outputs := []types.CommandOutput{}
	for _, cmd := range cmds {
		out, err := Execute(cmd)
		if err != nil {
			return nil, err
		}
		outputs = append(outputs, out)
	}
	return outputs, nil
}

// Execute one command.
func Execute(cmd types.Command) (types.CommandOutput, error) {
	if cmd.Method == "WS" {
		return execWebsocket(cmd)
	} else if cmd.Method == "TCP" {
		return execTCP(cmd)
	}
	return execHTTP(cmd)
}

func execTCP(cmd types.Command) (types.CommandOutput, error) {
	conn, err := net.Dial("tcp", cmd.URL)
	if err != nil {
		return types.CommandOutput{}, err
	}
	defer conn.Close()

	_, err = conn.Write(cmd.Body)
	if err != nil {
		return types.CommandOutput{}, err
	}

	body := make([]byte, 256)
	_, err = conn.Read(body)
	if err != nil {
		return types.CommandOutput{}, err
	}

	return types.CommandOutput{
		Body:     bytes.Trim(body, "\x00"),
		Response: nil,
	}, nil
}

func execHTTP(cmd types.Command) (types.CommandOutput, error) {
	client := &http.Client{
		Transport: defaultTransport,
	}

	req, err := http.NewRequest(cmd.Method, cmd.URL, bytes.NewBuffer(cmd.Body))
	if err != nil {
		return types.CommandOutput{}, err
	}

	req.Header = cmd.Headers

	uaHack := false
	// Don't send default User-Agent
	if _, ok := req.Header["User-Agent"]; !ok {
		req.Header.Add("User-Agent", "")
		uaHack = true
	}

	resp, err := client.Do(req)
	if err != nil {
		return types.CommandOutput{}, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return types.CommandOutput{}, err
	}

	if uaHack {
		// Mutating back to expected
		req.Header.Del("User-Agent")
	}

	return types.CommandOutput{
		Body:     b,
		Response: resp,
	}, nil
}

func execWebsocket(cmd types.Command) (types.CommandOutput, error) {
	// return errors.New("WS transport not implemented")
	return types.CommandOutput{}, nil
}
