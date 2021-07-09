package main

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

const (
	preferenceCategory = "opendyslexic"
	preferenceName     = "enabled"

	commandTrigger = "opendyslexic"
	commandDesc    = "Configure the OpenDyslexic plugin."
)

type Subcommand struct {
	Trigger      string
	AutoComplete bool
	HelpText     string
	Hint         string
	Execute      func(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError)
}

func (p *Plugin) getSubcommands() []*Subcommand {
	return []*Subcommand{
		{
			Trigger:      "enable",
			AutoComplete: true,
			HelpText:     "Enable the OpenDyslexic font.",
			Execute:      p.enableFontCommand,
		},
		{
			Trigger:      "disable",
			AutoComplete: true,
			HelpText:     "Enable the OpenDyslexic font.",
			Execute:      p.disableFontCommand,
		},
		{
			Trigger:      "help",
			AutoComplete: false,
			HelpText:     "Show this help text.",
			Execute:      p.printHelpCommand,
		},
	}
}

func (p *Plugin) registerCommands() error {
	if err := p.API.RegisterCommand(&model.Command{
		Trigger:          commandTrigger,
		AutoComplete:     true,
		AutoCompleteHint: "",
		AutoCompleteDesc: commandDesc,
		AutocompleteData: p.getAutocompleteData(),
	}); err != nil {
		return errors.Wrapf(err, "failed to register %s command", "opendyslexic")
	}
	// for _, command := range p.getCommands() {
	// 	if err := p.API.RegisterCommand(command.Command); err != nil {
	// 		return errors.Wrapf(err, "failed to register %s command", command.Command.Trigger)
	// 	}

	// 	fmt.Println("Registered", command.Command.Trigger)
	// }

	return nil
}

func (p *Plugin) getAutocompleteData() *model.AutocompleteData {
	autocomplete := model.NewAutocompleteData(commandTrigger, "", commandDesc)

	var items []model.AutocompleteListItem
	for _, command := range p.getSubcommands() {
		items = append(items, model.AutocompleteListItem{
			Item:     command.Trigger,
			Hint:     command.Hint,
			HelpText: command.HelpText,
		})
	}
	autocomplete.AddStaticListArgument("", true, items)

	return autocomplete
}

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	subcommandTrigger := ""
	if len(args.Command) > 1 {
		subcommandTrigger = strings.TrimPrefix(strings.Fields(args.Command)[1], "/")
	}

	for _, command := range p.getSubcommands() {
		if subcommandTrigger == command.Trigger {
			return command.Execute(c, args)
		}
	}

	return &model.CommandResponse{
		ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
		Text:         fmt.Sprintf("Unknown command: " + args.Command),
	}, nil
}

func (p *Plugin) enableFontCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	return &model.CommandResponse{}, p.API.UpdatePreferencesForUser(args.UserId, []model.Preference{
		{
			UserId:   args.UserId,
			Category: preferenceCategory,
			Name:     preferenceName,
			Value:    "true",
		},
	})
}

func (p *Plugin) disableFontCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	return &model.CommandResponse{}, p.API.UpdatePreferencesForUser(args.UserId, []model.Preference{
		{
			UserId:   args.UserId,
			Category: preferenceCategory,
			Name:     preferenceName,
			Value:    "false",
		},
	})
}

func (p *Plugin) printHelpCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	var lines []string
	for _, command := range p.getSubcommands() {
		lines = append(lines, fmt.Sprintf("- `/opendyslexic %s` - %s", command.Trigger, command.HelpText))
	}

	return &model.CommandResponse{
		ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
		Text:         "###### OpenDyslexic Plugin Help\n" + strings.Join(lines, "\n"),
	}, nil
}
