package main

import (
	"Vk-internship/internal/config"
	"Vk-internship/internal/service"
	"github.com/SevereCloud/vksdk/v2/api"
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

	group, err := vk.GroupsGetByID(nil)
	if err != nil {
		log.Fatal(err)
	}
	bot := service.NewBot(cfg.Responses, vk, group)

	if err = bot.Start(); err != nil {
		log.Fatal(err)
	}
}
