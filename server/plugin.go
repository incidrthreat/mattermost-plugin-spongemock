package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin
}

// See https://developers.mattermost.com/extend/plugins/server/reference/

// OnActivate checks if the configurations is valid and ensures the bot account exists
func (p *Plugin) OnActivate() error {
	return p.API.RegisterCommand(&model.Command{
		Trigger:          "spongemock",
		AutoComplete:     true,
		AutoCompleteDesc: "Gimmie a phrase to mock",
	})
}

// CommandHelp displays command info
const CommandHelp = "## **_Mattermost SpongeMock Plugin - cOmMaNd HeLp_**\n" +
	"#### Basic Usage:\n" +
	"* |/spongemock <input>| - Takes in a input and returns a Spongebob mocked output.\n" +
	"    * |/spongemock this is a test| will return |tHiS Is a tEsT|"

// ExecuteCommand ...
func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	// remove leading and trailing white space and removes slash command syntax from input
	input := strings.TrimSpace(strings.TrimPrefix(args.Command, "/spongemock"))

	// Displays the help menu
	if input == "help" || input == "" {
		text := strings.Replace(CommandHelp, "|", "`", -1)
		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text:         text,
		}, nil
	}

	spongemock := []string{}
	for i, c := range input {
		if i%2 != 0 {
			spongemock = append(spongemock, strings.ToUpper(string(c)))
		} else {
			spongemock = append(spongemock, strings.ToLower(string(c)))
		}
	}

	//p.postBotResponse(args, strings.Join(spongemock, ""))
	return &model.CommandResponse{
		ResponseType: model.COMMAND_RESPONSE_TYPE_IN_CHANNEL,
		Text:         strings.Join(spongemock, ""),
	}, nil
}

func main() {
	plugin.ClientMain(&Plugin{})
}
