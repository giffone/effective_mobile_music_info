package api

import (
	"log"
	"music_info/internal/dto"
	"music_info/internal/service"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	GetInfoBy(e echo.Context) error
}

func New(service service.Service) Handler {
	return &handler{service: service}
}

type handler struct {
	service service.Service
}

// @Summary Get song info
// @Description Get detailed information about a song based on group and song name.
// @Tags music
// @Accept json
// @Produce json
// @Param group query string true "Group name"
// @Param song query string true "Song name"
// @Success 200 {object} view.SongDetail
// @Failure 400 {object} echo.HTTPError "Bad Request"
// @Failure 404 {object} echo.HTTPError "Not found"
// @Failure 500 {object} echo.HTTPError "Internal Server Error"
// @Router /info [get]
func (h *handler) GetInfoBy(c echo.Context) error {
	var filter dto.InfoByGroupAndSong

	// parse request
	if err := c.Bind(&filter); err != nil || filter.Empty() {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	// get song
	song, err := h.service.GetInfoByGroupAndSong(c.Request().Context(), filter)
	if err != nil {
		if err == pgx.ErrNoRows {
			return c.String(http.StatusNotFound, "Song not found")
		}
		log.Println("err is: %w", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// ok
	return c.JSON(http.StatusOK, song)
}
