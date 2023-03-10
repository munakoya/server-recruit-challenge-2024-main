package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/service"
)

type albumSingerController struct {
	service service.AlbumSingerService
}

func NewAlbumSingerController(s service.AlbumSingerService) *albumSingerController {
	return &albumSingerController{service: s}
}

// これ修正
// GET /albums のハンドラー
func (c *albumSingerController) GetAlbumSingerListHandler(w http.ResponseWriter, r *http.Request) {
	albums, err := c.service.GetAlbumSingerListService(r.Context())
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(albums)
}

// GET /albums/{id} のハンドラー
func (c *albumSingerController) GetAlbumSingerDetailHandler(w http.ResponseWriter, r *http.Request) {
	albumID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		err = fmt.Errorf("invalid path param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	album, err := c.service.GetAlbumSingerService(r.Context(), model.AlbumID(albumID))
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(album)
}
