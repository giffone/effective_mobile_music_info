package dto

type InfoByGroupAndSong struct {
	Group string `query:"group"`
	Song  string `query:"song"`
}

func (i InfoByGroupAndSong) Empty() bool {
	return i.Group == "" || i.Song == ""
}
