package service

import (
	"encoding/json"
	"github.com/SevereCloud/vksdk/object"
	"os"
)

// GetKeyboardFromFile - reads a provided a file and returns keyboard format for vk
func GetKeyboardFromFile(fileName string) (*object.MessagesKeyboard, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var keyboard object.MessagesKeyboard
	if err := json.Unmarshal(b, &keyboard); err != nil {
		return nil, err
	}
	return &keyboard, nil
}
