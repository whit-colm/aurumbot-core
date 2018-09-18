package cmd

import (
	"errors"
	"github.com/aurumbot/dat/data"
	f "github.com/aurumbot/dat/foundation"
	dsg "github.com/bwmarrin/discordgo"
	"strings"
)

/* # MessageCreate
* The world's bigest switch statment
*
* This is a very big switch statment run commands. It reads all the messages in
* all the servers its in, determines which ones are commands, and then sees
* what in all the commands mean and then takes the appropriate action.
*
* Parameters:
* - s (type *discordgo.Session) | The current running discord session,
*     (discordgo needs that always apparently)
* - m (type *discordgo.Message) | The message thats to be acted upon.
*
* TODO: See if it can be made so it doesn't have to read every single message
*       ever.
*
* TODO: Break this one function up to smaller functions that only run if a user
*       has a certain role
*
* NOTE: Please delegate what the command actually does to a function. This
*       method should only be used to determine what the user is acutally
*       trying to do.
 */
func MessageCreate(session *dsg.Session, message *dsg.MessageCreate) {
	s := session
	m := message.Message
	// The message is checked to see if its a command and can be run
	canRunCommand, err := canTriggerBot(s, m)
	if err != nil {
		dat.Log.Println(err.Error())
		dat.AlertDiscord(s, m, err)
		return
	}
	if !canRunCommand {
		return
	}

	// Removing case sensitivity:
	messageSanatized := strings.ToLower(m.Content)

	// The prefix is cut off the message so the commands can be more easily handled.
	var msg []string
	if strings.HasPrefix(m.Content, f.MyBot.Auth.Prefix) {
		msg = strings.SplitAfterN(messageSanatized, f.MyBot.Auth.Prefix, 2)
		m.Content = msg[1]
		//TODO: Check if there is a way to use a mention() method of discordgo rather than
		//this string frankenstein
	} else if strings.HasPrefix(m.Content, "<@!"+f.MyBot.Auth.ClientID+">") {
		msg = strings.SplitAfterN(messageSanatized, "<@!"+f.MyBot.Auth.ClientID+">", 2)
		m.Content = strings.TrimSpace(msg[1])
	} else {
		err := errors.New("Message passed 'can run' checks but does not start with prefix:\n" + m.Content)
		dat.Log.Println(err.Error())
		dat.AlertDiscord(s, m, err)
		return
	}

	msgSplit := strings.Split(m.Content, " ")

	// Now the message is run to see if its a valid command and acted upon.
	for command, action := range Cmd {
		if msgSplit[0] == command {
			if action.Perms != -1 {
				perm, err := f.HasPermissions(s, m, m.Author.ID, action.Perms)
				if err != nil {
					dat.Log.Println(err)
					dat.AlertDiscord(s, m, err)
					return
				}
				if !perm {
					s.ChannelMessageSend(m.ChannelID, "Sorry, you do not have permission to use this command.")
					return
				}
			}
			action.Action(s, m)
			return
		}
	}

	if strings.Contains(m.Content, "@") {
		s.ChannelMessageSend(m.ChannelID, "Sorry <@"+m.Author.ID+">, but I don't understand.")
	} else {
		s.ChannelMessageSend(m.ChannelID, "Sorry <@"+m.Author.ID+">, but I don't understand what you mean by \"`"+m.Content+"`\".")
	}

}
