package main

import (
	"github.com/canonical/pebble/client"
	"github.com/jessevdk/go-flags"
)

type cmdStopExtend struct {
	clientMixin
	Positional struct {
		Services []string `positional-arg-name:"<service>"`
	} `positional-args:"yes" required:"yes"`
}

func init() {
	addCommand("stop-extend", "", "", func() flags.Commander {
		return &cmdStopExtend{}
	}, nil, nil)
}

func (cmd *cmdStopExtend) Execute(args []string) error {
	if len(args) > 1 {
		return ErrExtraArgs
	}

	servopts := client.ServiceOptions{
		Names: cmd.Positional.Services,
	}
	_, err := cmd.client.StopExtend(&servopts)
	if err != nil {
		return err
	}

	return nil
}
