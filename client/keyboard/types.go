package keyboard

type Button struct {
	Action ButtonText `json:"action"`
	Color  string     `json:"color"`
}

type ButtonText struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
	Label   string `json:"label"`
}

type Keyboard struct {
	Inline  bool       `json:"inline"`
	OneTime bool       `json:"one_time"`
	Buttons [][]Button `json:"buttons"`
}

type TypesKeyboard struct {
	HoroscopeKeyboard    [][]Button
	StartKeyboard        [][]Button
	RepeateKeyboard      [][]Button
	RepeatePhotoKeyboard [][]Button
}
