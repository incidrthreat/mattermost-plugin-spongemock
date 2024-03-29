package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "com.mattermost.plugin-spongemock",
  "name": "Plugin Starter Template",
  "description": "This plugin serves as a starting point for writing a Mattermost plugin.",
  "version": "1.0.4",
  "min_server_version": "5.12.0",
  "server": {
    "executables": {
      "linux-amd64": "server/dist/plugin-linux-amd64",
      "darwin-amd64": "server/dist/plugin-darwin-amd64",
      "windows-amd64": "server/dist/plugin-windows-amd64.exe"
    },
    "executable": ""
  },
}
`

func init() {
	manifest = model.ManifestFromJson(strings.NewReader(manifestStr))
}
