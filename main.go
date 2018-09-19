package main

import (
	"github.com/aurumbot/core/handler"
	"github.com/aurumbot/dat/data"
	f "github.com/aurumbot/dat/foundation"
	dsg "github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	runBot()
	// And the main should end here. so what to do about this?
}

func runBot() {
	bot, err := dat.GetBotInfo()
	dat.Log.Println("Reading bot prefs file...")
	if err != nil {
		dat.Log.Fatalln(err.Error())
	} else {
		dat.Log.Println("Bot prefrences recived.")
	}
	dat.Log.Println("Creating bot session")
	dg, err := dsg.New("Bot " + bot.Auth.Token)
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
	f.DG = dg

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dat.Log.Println("Escape for bot called. The system is now closing cleanly")
}
