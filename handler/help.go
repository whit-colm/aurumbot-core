package cmd

import (
	"github.com/aurumbot/dat/data"
	f "github.com/aurumbot/dat/foundation"
	"github.com/aurumbot/flags"
	dsg "github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

func init() {
	Cmd["help"] = &f.Command{
		Name: "Command Help Page Search",
		Help: `Info  : The built-in helper to get information about all of the bots commands
Flags:
-c --command <command>	: get help for the specific <command>
-ls --list		: get a list of all available commands
Usage : ` + f.MyBot.Prefs.Prefix + `help -c <command>
	` + f.MyBot.Prefs.Prefix + `help -ls`,
		Perms:   -1,
		Version: "v2.0.1β",
		Action:  help,
	}
}

/* # Get bot help
* Overcomplecated for little good reason
*
* Note that this function handles responding instead of returning a value to
* its parent to be sent out.
*
* Flags:
* -d  | Sends the result via dm (not working). //TODO: Figure this out
* -ls | Gets a list of available commands given the users perm level
* -c  | Gets the result for a specific command
 */
func help(session *dsg.Session, message *dsg.MessageCreate) {
	msg := strings.Split(message.Content, " ")
	if len(msg) <= 1 {
		h := "Help Page Found:\n```" + Cmd["help"].Name + "\n" + Cmd["help"].Help + "```"
		session.ChannelMessageSend(message.ChannelID, h)
		return
	}

	flagsParsed := flags.Parse(msg)

	// These are some cop-out variables so I don't nest to eternity.
	if len(flagsParsed) == 0 {
		h := "Help Page Found:\n```" + Cmd["help"].Name + "\n" + Cmd["help"].Help + "```"
		session.ChannelMessageSend(message.ChannelID, h)
		return
	}

	for i := range flagsParsed {
		if flagsParsed[i].Type == flags.Dash && flagsParsed[i].Name == "c" {
			session.ChannelMessageSend(message.ChannelID, search(flagsParsed[i]))
		} else if flagsParsed[i].Type == flags.Dash && flagsParsed[i].Name == "ls" {
			session.ChannelMessageSend(message.ChannelID, list(session, message))
		} else if flagsParsed[i].Type == flags.DoubleDash && flagsParsed[i].Name == "command" {
			session.ChannelMessageSend(message.ChannelID, search(flagsParsed[i]))
		}
	}
}

// TODO: Stop making repeat calls via HasPermissions
func list(session *dsg.Session, message *dsg.MessageCreate) string {
	t1 := time.Now()
	msg := "**Available Commands:**"
	for command, action := range Cmd {
		u, err := f.HasPermissions(session, message.Message, message.Author.ID, action.Perms)
		if err != nil {
			dat.Log.Println(err)
			dat.AlertDiscord(session, message, err)
			return ""
		}
		if u {
			msg += "\n" + f.MyBot.Prefs.Prefix + command + " : " + action.Name
		}
	}
	t2 := time.Since(t1)
	tStr := t2.String()
	msg += "\nUse `" + f.MyBot.Prefs.Prefix + "help -c <command>` to get more info on a command (" + tStr + ")"
	return msg
}

func search(cmd *flags.Flag) string {
	for command, action := range Cmd {
		if cmd.Value == command {
			help := "Help Page Found:\n**" + action.Name + "**\n```" + action.Help + "\nVersion: " + action.Version + "```"
			return help
		}
	}
	return "A help page for that command couldn't be found."
}
