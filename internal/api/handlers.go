package api

import (
	"errors"
	"music_info/internal/dto"
	"music_info/internal/model"
	"music_info/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	CreateSong(c echo.Context) error
	UpdateSong(c echo.Context) error
	GetInfoBy(e echo.Context) error
}

func New(service service.Service) Handler {
	return &handler{service: service}
}

type handler struct {
	service service.Service
}

func (h *handler) CreateSong(c echo.Context) error {
	var song dto.Song

	// parse request
	if err := c.Bind(&song); err != nil {
		return errors.Join(model.ErrBadData, err)
	}

	// create song
	err := h.service.CreateSong(c.Request().Context(), song)
	if err != nil {
		return err
	}

	// ok
	return c.String(http.StatusCreated, "Success")
}

func (h *handler) UpdateSong(c echo.Context) error {
	var song dto.UpdateSong

	// parse request
	if err := c.Bind(&song); err != nil {
		return errors.Join(model.ErrBadData, err)
	}

	// update song
	err := h.service.UpdateSong(c.Request().Context(), song)
	if err != nil {
		return err
	}

	// ok
	return c.String(http.StatusOK, "Updated")
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
	var filter dto.Song

	// parse request
	if err := c.Bind(&filter); err != nil {
		return errors.Join(model.ErrBadData, err)
	}

	// get song
	song, err := h.service.GetInfoByGroupAndSong(c.Request().Context(), filter)
	if err != nil {
		return err
	}

	// ok
	return c.JSON(http.StatusOK, song)
}
