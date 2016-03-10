package xkcdcom

type Xkcd struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"Link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}
