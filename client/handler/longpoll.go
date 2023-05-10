package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"vkbot/client/vk"
)

func (handler *Handler) LongPollServer(group_id int) (LongPoolServerResponse, error) {
	var response LongPoolServerResponse
	client := &http.Client{}

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s%s?group_id=%d&v=%s",
		vk.VkApiUrl, "groups.getLongPollServer", group_id, handler.Vk.Version), nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", handler.Vk.Token))

	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (handler *Handler) Start() error {
	client := &http.Client{}

	groupId, _ := strconv.Atoi(os.Getenv("vk_group_id"))
	lpserver, err := handler.LongPollServer(groupId)
	if err != nil {
		return err
	}

	handler.Vk.TypesKeyboard = handler.Vk.TypesKeyboard.InitKeyboards()

	for {
		req, _ := http.NewRequest("GET", fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=25",
			lpserver.Response.Server, lpserver.Response.Key, lpserver.Response.Ts), nil)

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", handler.Vk.Token))

		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		var response Response
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return err
		}

		switch response.Failed {
		case 0:
			lpserver.Response.Ts = response.Ts
		case 1:
			lpserver.Response.Ts = response.Ts
		case 2:
			lpserver, err = handler.LongPollServer(groupId)
			if err != nil {
				return err
			}
		case 3:
			lpserver, err = handler.LongPollServer(groupId)
			if err != nil {
				return err
			}
		default:
		}

		go handler.handleUpdates(response)
	}
}

func (handler *Handler) handleUpdates(response Response) {
	updates := response.Updates
	for _, value := range updates {
		if value.Type == "message_new" {
			var message MessageResponse
			err := json.Unmarshal(value.Object, &message)
			if err != nil {
				log.Println(err)
			}

			go handler.handleMessage(message)

		}
	}
}
