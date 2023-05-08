package service

import (
	"Vk-internship/pkg/config"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	_ "github.com/SevereCloud/vksdk/v2/object"
	"log"
)

const (
	DefaultPath = "pkg/keyboards/keyboards.json"
	Button1     = "mad_lads"
	Button2     = "solcasino"
	Button3     = "lily"
	Button4     = "okay_bears"
)

var m = params.NewMessagesSendBuilder()

type Bot struct {
	messages config.Responses
	vk       *api.VK
}

func NewBot(messages config.Responses, vk *api.VK) *Bot {
	return &Bot{messages: messages, vk: vk}
}

func (b *Bot) HandleButtons(obj events.MessageNewObject) {

	path := fmt.Sprintf("pkg/keyboards/%s.json", obj.Message.Text)

	if obj.Message.Text == "Начать" || obj.Message.Text == "начать" {
		sender(b.vk, b.messages.Choice, obj, DefaultPath)
	}

	// Switch construction for different messages/buttons
	switch {
	// First button
	case obj.Message.Text == Button1:
		sender(b.vk, b.messages.MadLads, obj, path)
	case obj.Message.Text == "Прайс Mad Lads":
		b.PrintFloor(Button1, obj)
		sender(b.vk, b.messages.Choice, obj, DefaultPath)
	case obj.Message.Text == "Объем Mad Lads":
		b.PrintVolume(Button1, obj)
		sender(b.vk, b.messages.Choice, obj, DefaultPath)

	// Second button
	case obj.Message.Text == Button2:
		sender(b.vk, b.messages.SolCasino, obj, path)
	case obj.Message.Text == "Прайс solcasino":
		b.PrintFloor(Button2, obj)
		sender(b.vk, b.messages.Choice, obj, DefaultPath)
	case obj.Message.Text == "Объем solcasino":
		b.PrintVolume(Button2, obj)
		sender(b.vk, b.messages.Choice, obj, DefaultPath)

	// Third button
	case obj.Message.Text == Button3:
		sender(b.vk, b.messages.Lily, obj, path)
	case obj.Message.Text == "Прайс lily":
		b.PrintFloor(Button3, obj)
		sender(b.vk, b.messages.Choice, obj, DefaultPath)
	case obj.Message.Text == "Объем lily":
		b.PrintVolume(Button3, obj)
		sender(b.vk, b.messages.Choice, obj, DefaultPath)

	// Fourth button
	case obj.Message.Text == Button4:
		sender(b.vk, b.messages.OkayBears, obj, path)
	case obj.Message.Text == "Прайс okay_bears":
		b.PrintFloor(Button4, obj)
		sender(b.vk, b.messages.Choice, obj, DefaultPath)
	case obj.Message.Text == "Объем okay_bears":
		b.PrintVolume(Button4, obj)
		sender(b.vk, b.messages.Choice, obj, DefaultPath)

	}

}

// sender - creates message -> creates keyboard -> send it all
func sender(vk *api.VK, msg string, obj events.MessageNewObject, pathToButtons string) {
	m := params.NewMessagesSendBuilder()
	m.PeerID(obj.Message.PeerID)
	m.RandomID(0)
	// Keyboard creation
	buttons, err := GetKeyboardFromFile(pathToButtons)
	if err != nil {
		log.Fatal(err)
	}
	m.Keyboard(buttons)
	m.Message(msg)
	_, err = vk.MessagesSend(m.Params)
	if err != nil {
		log.Fatal(err)
	}
}
