package cmd

import (
	"github.com/aurumbot/flags"
	"github.com/aurumbot/lib/dat"
	f "github.com/aurumbot/lib/foundation"
	dsg "github.com/bwmarrin/discordgo"
	"sort"
	"strings"
	"time"
)

func init() {
	Cmd["botcfg"] = &f.Command{
		Name: "Bot configuration tool",
		Help: `Info  : The built-in helper to get information about all of the bots commands
Options:
prefix <prefix>			: set the default prefix for the bot
admins <-a|-r|-l> [role ID]	: add or remove a botadmin role, which gives users with the role authorization to all bot abilities.
blchans <-a|-r|-l> [#channel]	: add or remove a channel to the blacklist. Blacklisted channels will never have the bot respond to commands.
blroles <-a|-r|-l> [roleID]	: add or remove a role to the blacklist. Users with blacklisted roles will never have the bot respond to their commands. (overwritten by admin roles)
Usage : ` + f.Config.Prefix + `botcfg <flag> <value>
	` + f.Config.Prefix + `botcfg -ls
Powered by Aurum at https://github.com/aurumbot/core`,
		Perms:   dsg.PermissionAdministrator,
		Version: "v1.0.0β",
		Action:  botcfg,
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
func botcfg(session *dsg.Session, message *dsg.Message) {
	msg := strings.Split(message.Content, " ")
	if len(msg) <= 1 {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("You need to provide a valid operator. Please use `%vhelp botcfg` for info.", f.Config.Prefix))
		return
	}

	flagsParsed := flags.Parse(msg)

	if len(flagsParsed) == 0 {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("You need to provide a valid operator. Please use `%vhelp botcfg` for info.", f.Config.Prefix))
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
			help := "Help Page Found:\n**" + action.Name + "**\n```" + action.Help + "\nVersion: " + action.Version + "```"
			return help
		}
	}
	return "A help page for that command couldn't be found."
}
