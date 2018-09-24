package cmd

import (
	"fmt"
	"github.com/aurumbot/flags"
	"github.com/aurumbot/lib/dat"
	f "github.com/aurumbot/lib/foundation"
	dsg "github.com/bwmarrin/discordgo"
	"strings"
)

func init() {
	Cmd["botcfg"] = &f.Command{
		Name: "Bot configuration tool",
		Help: `Info: Allows administrators to set lower level configuration values for the bot.
Options:
**prefix  <-s> <prefix>** : set the default prefix for the bot
**admins  <-a|-r|-l> [role ID...]** : add or remove a botadmin role, which gives users with the role authorization to all bot abilities. Multiple items can be modified, separated by spaces.
**blchans <-a|-r|-l> [channelID]** : add or remove a channel to the blacklist. Blacklisted channels will never have the bot respond to commands (overwritted by admin permissions). Multiple items can be modified, separated by spaces.
**blroles <-a|-r|-l> [roleID]** : add or remove a role to the blacklist. Users with blacklisted roles will never have the bot respond to their commands. (overwritten by admin permissions) Multiple items can be modified, separated by spaces.
**Usage : ` + f.Config.Prefix + `botcfg <flag> <value [args...]>
	` + f.Config.Prefix + `botcfg admins -a 452901410065874954 485528736276414505
Powered by Aurum at https://github.com/aurumbot/core`,
		Perms:   dsg.PermissionAdministrator,
		Version: "v1.0.0β",
		Action:  botcfg,
	}
}

func botcfg(session *dsg.Session, message *dsg.Message) {
	if len(strings.Split(message.Content, " ")) <= 1 {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("You need to provide a valid operator. Please use `%vhelp botcfg` for info.", f.Config.Prefix))
		return
	}

	flagsParsed := flags.Parse(message.Content)

	for i := range flagsParsed {
		if flagsParsed[i].Name == "--unflagged" {
			switch flagsParsed[i].Value {
			case "prefix":
				session.ChannelMessageSend(message.ChannelID, prefix(flagsParsed))
			case "admins":
				session.ChannelMessageSend(message.ChannelID, admins(flagsParsed))
			case "blchans":
				session.ChannelMessageSend(message.ChannelID, blchans(flagsParsed))
			case "blroles":
				session.ChannelMessageSend(message.ChannelID, blroles(flagsParsed))
			default:
				session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("You need to provide a valid operator. Please use `%vhelp botcfg` for info.", f.Config.Prefix))
			}
		}
	}
}

func prefix(flgs []*flags.Flag) string {
	for i := range flgs {
		switch true {
		case flgs[i].Name == "-s":
			f.Config.Prefix = flgs[i].Value
			if err := dat.Save("aurum/preferences.json", f.Config); err != nil {
				dat.Log.Println(err)
				return fmt.Sprintf("Encountered error trying to save changes:\n```%v```", err)
			}
			return fmt.Sprintf("Prefix successfully changed to **%v**.", f.Config.Prefix)
		}
	}
	return fmt.Sprintf("Unable to change prefix. Are you using correct syntax?")
}

func admins(flgs []*flags.Flag) string {
	for i := range flgs {
		switch true {
		case flgs[i].Name == "-a":
			for _, role := range strings.Split(flgs[i].Value, " ") {
				f.Config.Admins = append(f.Config.Admins, role)
			}
			if err := dat.Save("aurum/preferences.json", f.Config); err != nil {
				dat.Log.Println(err)
				return fmt.Sprintf("Encountered error trying to save changes:\n```%v```", err)
			}
			return fmt.Sprintf("Successfully added admin roles.")
		case flgs[i].Name == "-r":
			for k := range f.Config.Admins {
				for _, role := range strings.Split(flgs[i].Value, " ") {
					if f.Config.Admins[k] == role {
						f.Config.Admins[k] = f.Config.Admins[len(f.Config.Admins)-1]
						f.Config.Admins[len(f.Config.Admins)-1] = ""
						f.Config.Admins = f.Config.Admins[:len(f.Config.Admins)-1]
						break
					}
				}
			}
			if err := dat.Save("aurum/preferences.json", f.Config); err != nil {
				dat.Log.Println(err)
				return fmt.Sprintf("Encountered error trying to save changes:\n```%v```", err)
			}
			return fmt.Sprintf("Successfully removed given admin roles.")
		case flgs[i].Name == "-l":
			msg := "**Administrator Role IDs:\n"
			for _, role := range f.Config.Admins {
				msg += fmt.Sprintf("\n- %v", role)
			}
			return msg
		}
	}
	return fmt.Sprintf("Unable to complete task. Are you using correct syntax?")
}

func blroles(flgs []*flags.Flag) string {
	return flgs[0].Name
}

func blchans(flgs []*flags.Flag) string {
	return flgs[0].Name
}
