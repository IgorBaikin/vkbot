package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"vkbot/client/bot"
	"vkbot/client/keyboard"
	"vkbot/client/vk"
)

type Handler struct {
	Vk  *vk.VKApi
	Bot *bot.Bot
}

func InitHandler() *Handler {
	vk := vk.NewVKApi()
	b := &bot.Bot{}
	return &Handler{
		Vk:  vk,
		Bot: b,
	}
}

func (h *Handler) handleMessage(msg MessageResponse) {
	payload := make(map[string]interface{})
	keyboard := keyboard.Keyboard{}

	if len(msg.Message.Payload) != 0 {
		err := json.Unmarshal([]byte(msg.Message.Payload), &payload)
		if err != nil {
			log.Println(err)
			return
		}

		if value, ok := payload["button"]; ok {
			switch value {
			case "1":
				keyboard.Buttons = h.Vk.TypesKeyboard.RepeateKeyboard
				text, err := h.Bot.GetQuote("")
				if err != nil {
					log.Println(err)
					return
				}

				err = h.Vk.SendMessage(text, msg.Message.PeerID, keyboard)
				if err != nil {
					log.Println(err)
					return
				}

				return
			case "2":
				keyboard.Buttons = h.Vk.TypesKeyboard.RepeateKeyboard
				text, _ := h.Bot.GetPredictions("")
				err = h.Vk.SendMessage(text, msg.Message.PeerID, keyboard)
				if err != nil {
					log.Println(err)
					return
				}

				return
			case "3":
				keyboard.Buttons = h.Vk.TypesKeyboard.HoroscopeKeyboard
				err = h.Vk.SendMessage("Выберите знак зодиака:", msg.Message.PeerID, keyboard)
				if err != nil {
					log.Println(err)
					return
				}

				return
			case "4", "40":
				keyboard.Buttons = h.Vk.TypesKeyboard.RepeatePhotoKeyboard
				_, err := h.Bot.GetPhoto("")
				if err != nil {
					log.Println(err)
					return
				}

				h.handlePhotoMessage(msg.Message.PeerID, keyboard)

				return
			case "30":
				keyboard.Buttons = h.Vk.TypesKeyboard.RepeateKeyboard
				var text string
				if h.Bot.LastAction != nil {
					text, err = h.Bot.ExecuteLastAction()
					if err != nil {
						log.Println(err)
						return
					}
				}

				err = h.Vk.SendMessage(text, msg.Message.PeerID, keyboard)
				if err != nil {
					log.Println(err)
					return
				}

				return
			case "31":
				keyboard.Buttons = h.Vk.TypesKeyboard.StartKeyboard
				err = h.Vk.SendMessage("Возвращаюсь", msg.Message.PeerID, keyboard)
				if err != nil {
					log.Println(err)
					return
				}

				return
			}
		}

		if value, ok := payload["horoscopeButton"]; ok {
			keyboard.Buttons = h.Vk.TypesKeyboard.HoroscopeKeyboard
			text, err := h.Bot.GetHoroscope(fmt.Sprintf("%v", value))
			if err != nil {
				log.Println(err)
				return
			}

			err = h.Vk.SendMessage(text, msg.Message.PeerID, keyboard)
			if err != nil {
				log.Println(err)
				return
			}

			return
		}
	}

	keyboard.Buttons = h.Vk.TypesKeyboard.StartKeyboard
	err := h.Vk.SendMessage("Привет! Используй клавиатуру", msg.Message.PeerID, keyboard)
	if err != nil {
		log.Println(err)
		return
	}

}

func (h *Handler) handlePhotoMessage(peerID int, keyboard keyboard.Keyboard) {
	handler, err := h.Vk.UploadPhoto()
	if err != nil {
		log.Println(err)
		return
	}

	response, err := h.Vk.SaveMessagePhoto(handler)
	if err != nil {
		log.Println(err)
		return
	}

	err = h.Vk.SendMessagePhoto(peerID, keyboard, response.Response[0])
	if err != nil {
		log.Println(err)
		return
	}
}
