package bot

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strings"
)

const (
	urlQuote     = "https://api.forismatic.com/api/1.0/?method=getQuote&format=json&jsonp=parseQuote"
	urlHoroscope = "https://ignio.com/r/export/utf/xml/daily/com.xml"
	urlPhoto     = "https://api.rand.by/image"
)

type Bot struct {
	LastAction func(string) (string, error)
	Parameters string
}

func (b *Bot) ExecuteLastAction() (string, error) {
	s, err := b.LastAction(b.Parameters)
	if err != nil {
		return "", err
	}

	return s, nil
}

func (b *Bot) GetPredictions(string) (string, error) {
	predictions := []string{"Бесспорно", "Предрешено", "Никаких сомнений", "Определённо да",
		"Можешь быть уверен в этом", "Мне кажется — да", "Вероятнее всего", "Хорошие перспективы",
		"Знаки говорят — да", "Да", "Пока не ясно, попробуй снова", "Спроси позже", "Лучше не рассказывать",
		"Сейчас нельзя предсказать", "Сконцентрируйся и спроси опять", "Даже не думай", "Мой ответ — нет",
		"По моим данным — нет", "Перспективы не очень хорошие", "Весьма сомнительно"}
	b.LastAction = b.GetPredictions
	b.Parameters = ""
	return predictions[rand.Intn(len(predictions))], nil
}

func (b *Bot) GetHoroscope(zodiac string) (string, error) {
	resp, err := getResponse(urlHoroscope)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	var response Horo
	err = xml.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	text := parseXml(response, zodiac)
	if len(text) == 0 {
		return "", errors.New("Что-то пошло не так")
	}

	return text, nil
}

func parseXml(resp Horo, zodiac string) string {
	metric := reflect.ValueOf(&resp).Elem().FieldByName(zodiac)
	parseText := fmt.Sprintf("%v", metric)
	if len(parseText) < 2 {
		return ""
	}

	return strings.Trim(parseText[1:len(parseText)-1], "\n")
}

func (b *Bot) GetQuote(string) (string, error) {
	b.LastAction = b.GetQuote
	b.Parameters = ""
	resp, err := getResponse(urlQuote)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	var response Quote
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s\n%s", response.QuoteText, response.QuoteAuthor), nil
}

func (b *Bot) GetPhoto(string) (str string, err error) {
	resp, err := getResponse(urlPhoto)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	var responsePhoto Photo
	err = json.NewDecoder(resp.Body).Decode(&responsePhoto)
	if err != nil {
		return
	}

	response, err := http.Get(responsePhoto.Urls.Regular)
	if err != nil {
		return
	}

	defer response.Body.Close()
	file, err := os.Create("photo.jpeg")
	if err != nil {
		return
	}

	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return
	}

	return
}

func getResponse(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
