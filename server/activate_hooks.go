package main

import (
	"github.com/pkg/errors"
)

func (p *Plugin) OnActivate() error {
	if err := p.registerCommands(); err != nil {
		return errors.Wrap(err, "failed to register commands")
	}

	return nil
}

func (p *Plugin) OnDeactivate() error {
	return nil
}
