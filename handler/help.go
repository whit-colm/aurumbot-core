package cmd

import (
	"github.com/aurumbot/flags"
	"github.com/aurumbot/lib/dat"
	f "github.com/aurumbot/lib/foundation"
	dsg "github.com/bwmarrin/discordgo"
	"sort"
	"time"
)

func init() {
	Cmd["help"] = &f.Command{
		Name: "Command Help Page Search",
		Help: `# Info: 
The built-in helper to get information about all of the bots commands
# Flags:
<none> <command> : get help for the specific <command>
<-ls|--list>     : get a list of all available commands
# Usage : 
` + f.Config.Prefix + `help reloadplugins
` + f.Config.Prefix + `help -ls
---
*Powered by Aurum at https://github.com/aurumbot/core*`,
		Perms:   -1,
		Version: "v2.1.0β",
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
func help(session *dsg.Session, message *dsg.Message) {
	flagsParsed := flags.Parse(message.Content)

	for i := range flagsParsed {
		if flagsParsed[i].Name == "--unflagged" && flagsParsed[i].Value != "" {
			session.ChannelMessageSend(message.ChannelID, search(flagsParsed[i]))
			return
		} else if flagsParsed[i].Name == "-ls" {
			session.ChannelMessageSend(message.ChannelID, list(session, message))
			return
		} else if flagsParsed[i].Name == "--list" {
			session.ChannelMessageSend(message.ChannelID, list(session, message))
			return
		}
	}
	h := "Help Page Found:\n```" + Cmd["help"].Name + "\n" + Cmd["help"].Help + "```"
	session.ChannelMessageSend(message.ChannelID, h)
}

func list(session *dsg.Session, message *dsg.Message) string {
	t1 := time.Now()
	// Due to the fact that to verify permissions, a call to HasPermissions is
	// required, which goes to discord each time, I'm committing a classic sin
	// of making my code less dry to improve speed

	// Gets the guild
	guild, err := f.GetGuild(session, message)
	if err != nil {
		dat.Log.Println(err)
		dat.AlertDiscord(session, message, err)
		return ""
	}
	// Gets the message author as a guild member
	member, err := session.GuildMember(guild.ID, message.Author.ID)
	if err != nil {
		dat.Log.Println(err)
		dat.AlertDiscord(session, message, err)
		return ""
	}
	// Grabs all of the roles of the guild
	roles, err := session.GuildRoles(guild.ID)
	if err != nil {
		dat.Log.Println(err)
		dat.AlertDiscord(session, message, err)
		return ""
	}
	// msg is the final string tht will be sent to discord
	msg := "**Available Commands:**"
	// This slice will store all the commands that the user *can* run.
	// It is a slice instead of a string because it will be sorted
	// alphabetically later.
	var availableCommands []string
	for command, action := range Cmd {
		// This is here unstead of with the rest because if a user has no
		// roles, they aren't checked even if the perm is open to everyone.
		if action.Perms == -1 {
			availableCommands = append(availableCommands, "\n"+f.Config.Prefix+command+" : "+action.Name)
			continue
		}
		for _, role := range roles {
			// This sorts through the users roles, if they have
			// its permissions are checked, otherwise it moves on
			// to the next role
			if !f.Contains(member.Roles, role.ID) {
				continue
			}
			// checks permissions of the role, now that we know
			// the user has it. This also checks if they have an
			// "administrator" role as defined in the bot's config
			// docs.
			// This is repetitive, yes, but its broken up to
			// prevent 1 ajsdillion character long lines.
			if role.Permissions&action.Perms != 0 {
				availableCommands = append(availableCommands, "\n"+f.Config.Prefix+command+" : "+action.Name)
				break
			} else if role.Permissions&dsg.PermissionAdministrator != 0 {
				availableCommands = append(availableCommands, "\n"+f.Config.Prefix+command+" : "+action.Name)
				break
			} else if f.Contains(f.Config.Admins, role.ID) {
				availableCommands = append(availableCommands, "\n"+f.Config.Prefix+command+" : "+action.Name)
				break
			}
		}
	}
	// Now availableCommands is sorted and written to msg
	sort.Strings(availableCommands)
	for _, c := range availableCommands {
		msg += c
	}

	tStr := time.Since(t1).String()
	msg += "\nUse `" + f.Config.Prefix + "help -c <command>` to get more info on a command (" + tStr + ")"
	return msg
}

// NOTE: This search function is really inefficent as it makes checks to
// discord each time. I don't know where to store this so it will be like this
// for now. oof.
func search(cmd *flags.Flag) string {
	for command, action := range Cmd {
		if cmd.Value == command {
			help := "**__" + action.Name + "__**\n```markdown\n" + action.Help + "\nVersion: " + action.Version + "```"
			return help
		}
	}
	return "A help page for that command couldn't be found."
}
