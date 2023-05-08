package main

import (
	"Vk-internship/pkg/config"
	"Vk-internship/pkg/service"
	"context"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"log"
)

func main() {
	// AccessToken for vk bot
	token := config.FromEnv("token")

	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	vk := api.NewVK(token)

	bot := service.NewBot(cfg.Responses, vk)

	group, err := vk.GroupsGetByID(nil)
	if err != nil {
		log.Fatal(err)
	}
	// Starting long pool
	lp, err := longpoll.NewLongPoll(vk, group[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	// listening for a message
	lp.MessageNew(func(_ context.Context, obj events.MessageNewObject) {

		log.Printf("%d: %s", obj.Message.PeerID, obj.Message.Text)

		bot.HandleButtons(obj)
	})

	log.Println("Start Long Poll")
	if err := lp.Run(); err != nil {
		log.Fatal(err)
	}
}
