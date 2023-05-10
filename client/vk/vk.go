package vk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
	"vkbot/client/keyboard"
)

const VkApiUrl = "https://api.vk.com/method/"

type VKApi struct {
	Token         string
	Version       string
	TypesKeyboard keyboard.TypesKeyboard
}

func NewVKApi() *VKApi {
	return &VKApi{Token: os.Getenv("VK_TOKEN"), Version: os.Getenv("API_VERSION")}
}

func (v *VKApi) SetUrlValues(text string, peer_id int, keyboard keyboard.Keyboard, attach Attachments) url.Values {
	values := url.Values{}
	values.Set("user_id", strconv.Itoa(peer_id))
	values.Set("message", text)
	values.Set("access_token", v.Token)
	values.Set("v", v.Version)
	if len(text) == 0 {
		values.Set("attachment", fmt.Sprintf("photo%s_%s", strconv.Itoa(attach.OwnerID), strconv.Itoa(attach.ID)))
	}

	rand.Seed(time.Now().Unix())
	values.Set("random_id", strconv.Itoa(rand.Int()))
	if keyboard.Buttons != nil {
		keyboardJSON, err := json.Marshal(keyboard)
		if err != nil {
			log.Fatal()
		}

		values.Set("keyboard", string(keyboardJSON))
	}

	return values
}

func (v *VKApi) SendMessage(text string, peer_id int, keyboard keyboard.Keyboard) error {
	values := v.SetUrlValues(text, peer_id, keyboard, Attachments{})
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s?%s",
		VkApiUrl, "messages.send", values.Encode()), nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}

func (v *VKApi) SendMessagePhoto(peer_id int, keyboard keyboard.Keyboard, attach Attachments) error {
	values := v.SetUrlValues("", peer_id, keyboard, attach)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s?%s",
		VkApiUrl, "messages.send", values.Encode()), nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}

func (v *VKApi) UploadPhoto() (handler PhotosMessageUploadResponse, err error) {
	values := url.Values{}
	values.Set("group_id", os.Getenv("VK_GROUP_ID"))
	values.Set("access_token", v.Token)
	values.Set("v", v.Version)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s?%s",
		VkApiUrl, "photos.getMessagesUploadServer", values.Encode()), nil)
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	var response PhotosGetMessagesUploadServerResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return
	}

	file, err := os.Open("photo.jpeg")
	bodyContent, err := v.UploadFile(response.Response.UploadURL, file, "photo", "photo.jpeg")

	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		return
	}

	return
}

func (v *VKApi) SaveMessagePhoto(handler PhotosMessageUploadResponse) (response AttachmentsRespone, err error) {
	values := url.Values{}
	values.Add("access_token", v.Token)
	values.Add("photo", handler.Photo)
	values.Add("server", strconv.Itoa(handler.Server))
	values.Add("hash", handler.Hash)
	values.Add("v", v.Version)

	resp, err := http.PostForm(fmt.Sprintf("%s%s", VkApiUrl, "photos.saveMessagesPhoto"), values)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return
	}

	return
}

func (v *VKApi) UploadFile(url string, file io.Reader, fieldname, filename string) (bodyContent []byte, err error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(fieldname, filename)
	if err != nil {
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return
	}

	contentType := writer.FormDataContentType()
	_ = writer.Close()

	client := http.Client{}
	resp, err := client.Post(url, contentType, body)
	if err != nil {
		return
	}

	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyContent, err = io.ReadAll(resp.Body)

	return
}
