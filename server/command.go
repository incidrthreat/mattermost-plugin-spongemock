package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

// CommandHelp displays command info
const CommandHelp = "## **_Mattermost SpongeMock Plugin - cOmMaNd HeLp_**\n" +
	"#### Basic Usage:\n" +
	"* |/spongemock <phrase>| - Takes in a phrase/input and returns a Spongebob mocked phrase.\n" +
	"    * |/spongemock this is a test| will return |tHiS Is a tEsT|"

// InitCommand ...
func getCommand() *model.Command {
	return &model.Command{
		Trigger:          "spongemock",
		DisplayName:      "SpongeMock Bot",
		Description:      "SpongeMock bot takes in a phrase and returns a Spongebob mocking phrase. i.e `tHiS sHoUlD bE a PlUgIn`",
		AutoComplete:     true,
		AutoCompleteDesc: "Available command(s): help",
		AutoCompleteHint: "[command]",
	}
}

func (p *Plugin) postCommandResponse(args *model.CommandArgs, text string) {
	post := &model.Post{
		UserId:    p.BotUserID,
		ChannelId: args.ChannelId,
		Message:   text,
	}
	_ = p.API.SendEphemeralPost(args.UserId, post)
}

// ExecuteCommand ...
func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	// remove slash command syntax
	phrase := strings.TrimLeft(args.Command, "/spongemock")
	// remove leading and trailing white space
	phrase = strings.TrimSpace(phrase)

	// Displays the help menu
	if phrase == "help" {
		text := strings.Replace(CommandHelp, "|", "`", -1)
		p.postCommandResponse(args, text)
		return &model.CommandResponse{}, nil
	}

	spongemock := []string{}
	for i, c := range phrase {
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
