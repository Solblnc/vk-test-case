package service

import (
	"Vk-internship/internal/config"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	_ "github.com/SevereCloud/vksdk/v2/object"
	"log"
	"strings"
)

const (
	Start       = "начать"
	DefaultPath = "internal/keyboards/keyboards.json"
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

// NewBot - returns a new bot
func NewBot(messages config.Responses, vk *api.VK) *Bot {
	return &Bot{messages: messages, vk: vk}
}

func (b *Bot) HandleButtons(obj events.MessageNewObject) {

	path := fmt.Sprintf("internal/keyboards/%s.json", obj.Message.Text)

	if strings.EqualFold(obj.Message.Text, Start) {
		b.sender(b.messages.Start, obj, DefaultPath)
	}

	switch obj.Message.Text {
	// First button
	case Button1:
		b.sender(b.messages.MadLads, obj, path)
	case "Прайс Mad Lads":
		b.PrintFloor(Button1, obj)
		b.sender(b.messages.Choice, obj, DefaultPath)
	case "Объем Mad Lads":
		b.PrintVolume(Button1, obj)
		b.sender(b.messages.Choice, obj, DefaultPath)

	// Second button
	case Button2:
		b.sender(b.messages.SolCasino, obj, path)
	case "Прайс solcasino":
		b.PrintFloor(Button2, obj)
		b.sender(b.messages.Choice, obj, DefaultPath)
	case "Объем solcasino":
		b.PrintVolume(Button2, obj)
		b.sender(b.messages.Choice, obj, DefaultPath)

	// Third button
	case Button3:
		b.sender(b.messages.Lily, obj, path)
	case "Прайс lily":
		b.PrintFloor(Button3, obj)
		b.sender(b.messages.Choice, obj, DefaultPath)
	case "Объем lily":
		b.PrintVolume(Button3, obj)
		b.sender(b.messages.Choice, obj, DefaultPath)

	// Fourth button
	case Button4:
		b.sender(b.messages.OkayBears, obj, path)
	case "Прайс okay_bears":
		b.PrintFloor(Button4, obj)
		b.sender(b.messages.Choice, obj, DefaultPath)
	case "Объем okay_bears":
		b.PrintVolume(Button4, obj)
		b.sender(b.messages.Choice, obj, DefaultPath)
	}

}

// sender - creates message -> creates keyboard -> send it all
func (b *Bot) sender(msg string, obj events.MessageNewObject, pathToButtons string) {
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
	_, err = b.vk.MessagesSend(m.Params)
	if err != nil {
		log.Fatal(err)
	}
}
