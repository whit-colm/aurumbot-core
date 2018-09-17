package cmd

import (
	"errors"
	"github.com/aurumbot/dat/data"
	f "github.com/aurumbot/dat/foundation"
	dsg "github.com/bwmarrin/discordgo"
	"io/ioutil"
	"plugin"
	"strings"
)

var Cmd = make(map[string]*f.Command)

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
		Action: func(session *dsg.Session, message *dsg.Message) {
			err := reloadPlugins()
			if err != nil {
				dat.AlertDiscord(session, message, err)
			} else {
				session.ChannelMessageSend(message.ChannelID, "Successfully reloaded plugins")
			}
		},
	}
	err := reloadPlugins()
	if err != nil {
		dat.Log.Println(err)
	}
}

func reloadPlugins() error {
	var files []string
	// A painfully complex `ls`
	filesUnchecked, err := ioutil.ReadDir("./plugins/")
	if err != nil {
		dat.Log.Println(err)
		return err
	}
	// Filters out non-.so files (i.e. .DS_Store)
	for _, file := range filesUnchecked {
		if strings.HasSuffix(file.Name(), ".so") && !file.IsDir() {
			files = append(files, file.Name())
		}
	}

	for _, module := range files {
		p, err := plugin.Open("./plugins/" + module)
		if err != nil {
			dat.Log.Println(err)
			return err
		}
		s, err := p.Lookup("Commands")
		if err != nil {
			dat.Log.Println(err)
			return err
		}
		cmds, ok := s.(*map[string]*f.Command)
		if !ok {
			err := errors.New("Unexpected type from module symbol")
			dat.Log.Println(err)
			return err
		}
		for key, value := range *cmds {
			Cmd[key] = value
		}
	}
	return nil
}
