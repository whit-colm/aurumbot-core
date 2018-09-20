package cmd

import (
	"github.com/aurumbot/dat/data"
	f "github.com/aurumbot/dat/foundation"
	dsg "github.com/bwmarrin/discordgo"
	"strings"
)

/* # Check if user can run command
* This switch statment makes sure the bot runs when its triggered and the user has the perms to trigger it.
* Prevents:
* - Bot posted something that would trigger itself, possibly creating an infinite loop
* - Message posted doesn't have the bot's prefix
* - Command was posted in a channel where the bot shouldn't respond to commands
* - Bot whitelists channels and the command was run in a channel not on the whitelist.
* - Users with a blacklisted role from running the bot
*
* NOTE: Users who have "admin" roles (according to the bot's json data) or
*       permissions will have the ability to run commands regardless of any
*       other rules
*
* NOTE: IF THESE CONDITIONS ARE MET THEN NO ERROR WILL BE SENT TO EITHER DISCORD OR LOGGED.
* THIS IS BY DESIGN. DON'T CHANGE IT THINKING I WAS JUST LAZY.
 */
func canTriggerBot(s *dsg.Session, m *dsg.Message) (bool, error) {
	if m.Author.Bot {
		return false, nil
	}

	admin, err := f.HasPermissions(s, m, m.Author.ID, dsg.PermissionAdministrator)
	if err != nil {
		dat.Log.Println(err)
		return false, err
	}

	switch true {
	case m.Author.ID == s.State.User.ID:
		return false, nil
	//TODO: look at this stupid line. that seems like it shouldn't work.
	//NOTE: Well it doesn't!
	case !strings.HasPrefix(m.Content, f.MyBot.Prefix) && !strings.HasPrefix(m.Content, "<@!"+f.MyBot.ClientID+">"):
		return false, nil
	case admin:
		return true, nil
	case f.Contains(f.MyBot.BlacklistedChannels, m.ChannelID) == true:
		return false, nil
	}
	for _, b := range f.MyBot.BlacklistedRoles {
		guild, err := f.GetGuild(s, m)
		if err != nil {
			return false, err
		}
		member, err := s.GuildMember(guild.ID, m.Author.ID)
		if err != nil {
			return false, err
		}
		blacklisted := f.Contains(member.Roles, b)
		if blacklisted {
			return false, nil
		}
	}
	return true, nil
}
