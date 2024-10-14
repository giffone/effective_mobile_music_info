package dto

import "log"

type Song struct {
	Group string `query:"group"`
	Song  string `query:"song"`
}

func (f Song) Empty() bool {
	return f.Group == "" || f.Song == ""
}

type UpdateSong struct {
	GroupID     string `query:"group_id"`
	Song        string `query:"song"`
	ReleaseDate string `query:"release_date"`
	Text        string `query:"text"`
	Link        string `query:"link"`
}

func (s UpdateSong) Empty() bool {
	log.Println("id -", s.GroupID)

	log.Println(s)
	return s.GroupID == "" || s.Song == "" || (s.ReleaseDate+s.Text+s.Link == "")
}
