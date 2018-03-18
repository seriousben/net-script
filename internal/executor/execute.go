package executor

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/seriousben/net-script/internal/types"
)

func ExecuteMany(cmds []types.Command) error {
	for _, cmd := range cmds {
		err := Execute(cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func Execute(cmd types.Command) error {
	fmt.Printf("Executing: %s\n", cmd)
	if cmd.Method == "WS" {
		return execWebsocket(cmd)
	} else if cmd.Method == "TCP" {
		return execTCP(cmd)
	}
	return execHTTP(cmd)
}

func execTCP(cmd types.Command) error {
	return errors.New("TCP transport not implemented")
}

func execHTTP(cmd types.Command) error {
	client := http.DefaultClient

	req, err := http.NewRequest(cmd.Method, cmd.URL, nil)
	if err != nil {
		return err
	}
	req.Header = cmd.Headers
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}

func execWebsocket(cmd types.Command) error {
	// return errors.New("WS transport not implemented")
	return nil
}
