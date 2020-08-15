package internal

import (
	"fmt"
	"github.com/Floor-Gang/utilpkg"
	"github.com/Floor-Gang/utilpkg/botutil"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"strings"
)

func (bot *Bot) onReady(_ *discordgo.Session, ready *discordgo.Ready) {
	println("inside OnReady Func")
	//the function for when the bot turns on/logs on
	err := bot.client.UpdateStatus(0, "Mega cool awesome sauce")
	if err != nil {
		println(err)
	}
	fmt.Printf("Running bot %s", bot.client.State.User.Username)
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func (bot *Bot) messageCreate(_ *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == bot.client.State.User.ID {
		return
	}

	if m.Author.Bot || !strings.HasPrefix(m.Content, bot.botPrefix) {
		return
	}

	args := strings.Fields(m.Content)

	if len(args) < 2 {
		return
	}

	args = utilpkg.RemoveFromSlice(bot.botPrefix, args)
	fmt.Println(args)
	switch args[0] {
	case "ping":
		bot.sendMessage(m.ChannelID, "Pong!🏓")
		break
	case "pong":
		bot.sendMessage(m.ChannelID, "Ping!🏓")
		break
	case "say":
		bot.sendMessage(m.ChannelID, strings.Join(args, " "))
		break
	case "mute":
		bot.voiceMute(true, m)
		break
	case "unmute":
		bot.voiceMute(false, m)
		break

	case "write":
		bot.writeToFile(m.Content, m)
		break

	case "read":
		bot.readFromFile(m.Content, m)
		break
	}
}

func (bot Bot) voiceMute(muted bool, m *discordgo.MessageCreate) {
	args := strings.Fields(m.Content)
	args = utilpkg.RemoveFromSlice(args[0], args)

	member, err := bot.client.GuildMember(m.GuildID, botutil.FilterTag(args[1]))
	if err != nil {
		bot.sendMessage(m.ChannelID, "Failed cmd is used like this\n`!bot (un)mute <@user>`")
		return
	}

	err = botutil.ServerMute(bot.client, m.GuildID, member.User.ID, muted)

	if muted == false {
		a := "<@" + member.User.ID + "> has been unServer muted"
		bot.sendMessage(m.ChannelID, a)
		return
	}
	if muted == true {
		a := "<@" + member.User.ID + "> has been Server muted"
		bot.sendMessage(m.ChannelID, a)
		return
	}

	if err != nil {
		bot.sendMessage(m.ChannelID, "User is not in voice channel")
		fmt.Println(err)
	}
}

func (bot Bot) writeToFile(content string, m *discordgo.MessageCreate) {
	d1 := []byte(content)
	err := ioutil.WriteFile("discgobottest.txt", d1, 0644)
	if err != nil {
		println(err)
	}
	bot.sendMessage(m.ChannelID, "Wrote **"+content+"** to file")

}
func (bot Bot) readFromFile(content string, m *discordgo.MessageCreate) {
	f, err := ioutil.ReadFile("discgobottest.txt")
	if err != nil {
		println(err)
	}

	str := string(f)
	println((str))
	bot.sendMessage(m.ChannelID, "File says: \n"+"**"+str+"**")

}

func (bot Bot) sendMessage(channelID string, content string) {
	_, err := bot.client.ChannelMessageSend(channelID, content)

	if err != nil {
		fmt.Println(err)
	}
}
