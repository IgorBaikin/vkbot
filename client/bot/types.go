package bot

type Quote struct {
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
}

type Horo struct {
	Aries struct {
		Today string `xml:"today"`
	} `xml:"aries"`
	Taurus struct {
		Today string `xml:"today"`
	} `xml:"taurus"`
	Gemini struct {
		Today string `xml:"today"`
	} `xml:"gemini"`
	Cancer struct {
		Today string `xml:"today"`
	} `xml:"cancer"`
	Leo struct {
		Today string `xml:"today"`
	} `xml:"leo"`
	Virgo struct {
		Today string `xml:"today"`
	} `xml:"virgo"`
	Libra struct {
		Today string `xml:"today"`
	} `xml:"libra"`
	Scorpio struct {
		Today string `xml:"today"`
	} `xml:"scorpio"`
	Sagittarius struct {
		Today string `xml:"today"`
	} `xml:"sagittarius"`
	Capricorn struct {
		Today string `xml:"today"`
	} `xml:"capricorn"`
	Aquarius struct {
		Today string `xml:"today"`
	} `xml:"aquarius"`
	Pisces struct {
		Today string `xml:"today"`
	} `xml:"pisces"`
}

type Photo struct {
	Urls struct {
		Regular string `json:"regular"`
	} `json:"urls"`
}
