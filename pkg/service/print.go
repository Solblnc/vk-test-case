package service

import (
	"github.com/SevereCloud/vksdk/v2/events"
	"log"
	"strconv"
)

// PrintFloor - converting float value of floor to string and put it in VK message
func (b *Bot) PrintFloor(symbol string, obj events.MessageNewObject) {
	// getting a floor value and setting it like a message
	msg := GetFloor(symbol)
	m.RandomID(0)
	m.PeerID(obj.Message.PeerID)
	// Creating a message
	m.Message("Цена коллекции: " + strconv.FormatFloat(msg, 'f', 2, 64) + " sol")

	// Sending a message
	_, err := b.vk.MessagesSend(m.Params)
	if err != nil {
		log.Fatal(err)
	}
}

// PrintVolume - converting float value of volume to string and put it in VK message
func (b *Bot) PrintVolume(symbol string, obj events.MessageNewObject) {
	// getting a volume value
	msg := GetVolume(symbol)
	m.RandomID(0)
	m.PeerID(obj.Message.PeerID)
	// Creating a message
	m.Message("Объем коллекции за все время: " + strconv.FormatFloat(msg, 'f', 2, 64) + " sol")

	// Sending a message
	_, err := b.vk.MessagesSend(m.Params)
	if err != nil {
		log.Fatal(err)
	}
}
