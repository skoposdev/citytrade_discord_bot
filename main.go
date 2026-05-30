package main

import (
	"TradeBot/commands/flux"
	"TradeBot/commands/fun"
	"TradeBot/commands/info"
	"TradeBot/commands/moderation"
	"log"
	"os"

	"TradeBot/bot"
	"TradeBot/commands"
)

var (
	token = os.Getenv("TOKEN")
)

func main() {
	if token == "" {
		log.Fatal("missing TOKEN")
	}

	dg, err := bot.Start(token)
	if err != nil {
		log.Fatal(err)
	}

	commands.Register("echo", commands.Command{Exec: fun.Echo})
	commands.Register("user_info", commands.Command{Exec: info.UserInfo})
	commands.Register("rank", commands.Command{Exec: flux.StaffRank})
	commands.Register("unrank", commands.Command{Exec: flux.StaffUnrank})
	commands.Register("switch", commands.Command{Exec: flux.RankSwitch})
	commands.Register("mute", commands.Command{Exec: moderation.Mute})

	bot.GuildsHandler(dg)

	select {}
}
