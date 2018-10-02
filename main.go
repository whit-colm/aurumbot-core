package main

import (
	"github.com/aurumbot/core/handler"
	"github.com/aurumbot/lib/dat"
	f "github.com/aurumbot/lib/foundation"
	dsg "github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

var bot *f.Bot

func main() {
	runBot()
	// And the main should end here. so what to do about this?
}

func runBot() {
	err := dat.Load("aurum/preferences.json", &bot)
	dat.Log.Println("Reading bot prefs file...")
	if err != nil {
		dat.Log.Fatalln(err.Error())
	} else {
		dat.Log.Println("Bot prefrences recived.")
	}
	dat.Log.Println("Creating bot session")
	dg, err := dsg.New("Bot " + bot.Token)
	if err != nil {
		dat.Log.Fatalln(err)
	} else {
		dat.Log.Println("Session successfully created.")
	}

	dg.AddHandler(cmd.MessageCreate)

	dat.Log.Println("Opening websocket to Discord")
	err = dg.Open()
	if err != nil {
		dat.Log.Fatalln(err.Error())
	} else {
		dat.Log.Println("Socket successfully opened.")
	}
	f.Session = dg
	f.Config = *bot
	dat.Log.Println("Loading Plugins")
	err = cmd.ReloadPlugins()
	if err != nil {
		dat.Log.Fatalln(err)
	} else {
		dat.Log.Println("Successfully loaded plugins")
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
	dat.Log.Println("Escape for bot called. The system is now closing cleanly")
}
