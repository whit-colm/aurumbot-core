package cmd

import (
	"github.com/aurumbot/dat/data"
	f "github.com/aurumbot/dat/foundation"
	"io/ioutil"
	"plugin"
)

var Cmd = map[string]*f.Command{}

/* FOR THE PERSON RUNNING THIS BOT: Adding packages to the command list
* As of now, the bot has no commands set to it so while it may boot up, it
* won't actually do anything. You will need to add the maps of the command
* modules you have imported or made into the main Cmd map. To do this, add
* each of the command's public map[string]*f.Command type into the following
* init statment. 2 commands, `info` and `ping` have already been added to help
* show what you need to do:
 */

func init() {
	Cmd["reloadplugins"] = &f.Command{
		Name:    "Reload bot plugins",
		Help:    "Reloads plugins from the ./plugin directory.",
		Perms:   dsg.PermissionAdministrator,
		Version: "1.0.0α",
		Action: func(session *dsg.Session, message *dsg.MessageCreate) {
			err := reloadPlugins()
			if err != nil {
				dat.AlertDiscord(s, m, err)
			} else {
				session.ChannelMessageSend(message.Message.ChannelID, "Successfully reloaded plugins")
			}
		},
	}
	_ = reloadPlugins()
}

func reloadPlugins() error {
	files, err := ioutil.ReadDir("./plugins/")
	if err != nil {
		dat.Log.Println(err)
		return
	}
	for _, module := range files {
		p, err := plugin.Open("./plugins/" + module)
		if err != nil {
			dat.Log.Println(err)
			return
		}
		s, err := p.Lookup("Commands")
		if err != nil {
			dat.Log.Println(err)
			return
		}
		for key, value := range s {
			Cmd[key] = value
		}
	}
}
