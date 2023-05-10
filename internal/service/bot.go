package service

import (
	"Vk-internship/internal/config"
	"context"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"log"
)

type Bot struct {
	messages config.Responses
	vk       *api.VK
	group    api.GroupsGetByIDResponse
}

// NewBot - returns a new bot
func NewBot(messages config.Responses, vk *api.VK, group api.GroupsGetByIDResponse) *Bot {
	return &Bot{messages: messages, vk: vk, group: group}
}

func (b *Bot) Start() error {
	lp, err := longpoll.NewLongPoll(b.vk, b.group[0].ID)
	if err != nil {
		return err
	}
	// listening for a message
	lp.MessageNew(func(_ context.Context, obj events.MessageNewObject) {

		log.Printf("%d: %s", obj.Message.PeerID, obj.Message.Text)

		b.HandleButtons(obj)
	})

	log.Println("Start Long Poll")
	if err := lp.Run(); err != nil {
		return err
	}

	return nil
}
