package storeReview

type Author struct {
	Name string `xml:"name"`
	URI  string `xml:"uri"`
}

type Link struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

type Comment struct {
	Type string `xml:"type,attr"`
	Text string `xml:",chardata"`
}

type XML struct {
	Reviews []struct {
		Id      string    `xml:"id"`
		Updated string    `xml:"updated"`
		Title   string    `xml:"title"`
		Comment []Comment `xml:"content"`
		Rating  int       `xml:"rating"`
		Version string    `xml:"version"`
		Author  Author    `xml:"author"`
		Link    Link      `xml:"link"`
	} `xml:"entry"`
}
