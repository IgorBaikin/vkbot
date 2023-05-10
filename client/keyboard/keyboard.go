package keyboard

func startKeyboard() [][]Button {
	var buttons = [][]Button{
		{
			{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"button":"1"}`,
					Label:   "Случайная цитата",
				},
				Color: "primary",
			},
			{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"button":"2"}`,
					Label:   "8ball",
				},
				Color: "primary",
			},
		},
		{
			Button{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"button":"3"}`,
					Label:   "Гороскоп",
				},
				Color: "primary",
			},
			Button{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"button":"4"}`,
					Label:   "Фото дня",
				},
				Color: "primary",
			},
		},
	}

	return buttons
}

func horoscopeKeyboard() [][]Button {
	var buttons = [][]Button{
		{
			{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"horoscopeButton":"Aries"}`,
					Label:   "♈ Овен",
				},
				Color: "primary",
			},
			{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"horoscopeButton":"Taurus"}`,
					Label:   "♉ Телец",
				},
				Color: "primary",
			},
			{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"horoscopeButton":"Gemini"}`,
					Label:   "♊ Близнецы",
				},
				Color: "primary",
			},
		},
		{
			{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"horoscopeButton":"Leo"}`,
					Label:   "♌ Лев",
				},
				Color: "primary",
			},
			{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"horoscopeButton":"Virgo"}`,
					Label:   "♍	Дева",
				},
				Color: "primary",
			},
			{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"horoscopeButton":"Libra"}`,
					Label:   "♎ Весы",
				},
				Color: "primary",
			},
		},
		{
			{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"horoscopeButton":"Sagittarius"}`,
					Label:   "♐ Стрелец",
				},
				Color: "primary",
			},
			{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"horoscopeButton":"Capricorn"}`,
					Label:   "♑ Козерог",
				},
				Color: "primary",
			},
			{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"horoscopeButton":"Aquarius"}`,
					Label:   "♒ Водолей",
				},
				Color: "primary",
			},
		},
		{
			{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"horoscopeButton":"Cancer"}`,
					Label:   "♋ Рак",
				},
				Color: "primary",
			},
			{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"horoscopeButton":"Pisces"}`,
					Label:   "♓ Рыбы",
				},
				Color: "primary",
			},
			{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"horoscopeButton":"Scorpio"}`,
					Label:   "♏ Скорпион",
				},
				Color: "primary",
			},
		},
		{
			{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"button":"31"}`,
					Label:   "Назад",
				},
				Color: "primary",
			},
		},
	}

	return buttons
}

func repeateKeyboard() [][]Button {
	var buttons = [][]Button{
		{
			{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"button":"30"}`,
					Label:   "Давай еще раз",
				},
				Color: "primary",
			},
			{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"button":"31"}`,
					Label:   "Назад",
				},
				Color: "primary",
			},
		},
	}

	return buttons
}

func repeatePhotoKeyboard() [][]Button {
	var buttons = [][]Button{
		{
			{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"button":"40"}`,
					Label:   "Еще фото?",
				},
				Color: "primary",
			},
			{
				Action: ButtonText{
					Type:    "text",
					Payload: `{"button":"31"}`,
					Label:   "Назад",
				},
				Color: "primary",
			},
		},
	}

	return buttons
}

func (t *TypesKeyboard) InitKeyboards() TypesKeyboard {
	return TypesKeyboard{
		HoroscopeKeyboard:    horoscopeKeyboard(),
		StartKeyboard:        startKeyboard(),
		RepeateKeyboard:      repeateKeyboard(),
		RepeatePhotoKeyboard: repeatePhotoKeyboard(),
	}
}
