package dto

type InfoByGroupAndSong struct {
	Group string `query:"group"`
	Song  string `query:"song"`
}
