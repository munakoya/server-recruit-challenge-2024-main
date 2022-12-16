package api

// ルーティング

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pulse227/server-recruit-challenge-sample/api/middleware"
	"github.com/pulse227/server-recruit-challenge-sample/controller"
	"github.com/pulse227/server-recruit-challenge-sample/infra/memorydb"
	"github.com/pulse227/server-recruit-challenge-sample/service"
)

func NewRouter() *mux.Router {
	// この3つ
	// repository → ?
	singerRepo := memorydb.NewSingerRepository() // singerデータを代入
	// service → ?
	singerService := service.NewSingerService(singerRepo) // singerServiceにsingerデータがセットされる
	// controller → ?
	singerController := controller.NewSingerController(singerService)

	// アルバムデータの初期化？
	albumRepo := memorydb.NewAlbumRepository()
	albumService := service.NewAlbumService(albumRepo)
	albumController := controller.NewAlbumController(albumService)

	r := mux.NewRouter()

	r.HandleFunc("/singers", singerController.GetSingerListHandler).Methods(http.MethodGet)
	r.HandleFunc("/singers/{id:[0-9]+}", singerController.GetSingerDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/singers", singerController.PostSingerHandler).Methods(http.MethodPost)
	r.HandleFunc("/singers/{id:[0-9]+}", singerController.DeleteSingerHandler).Methods(http.MethodDelete)

	// ここにalbumsのルーティング追加されるはず
	r.HandleFunc("/albums", albumController.GetAlbumListHandler).Methods(http.MethodGet)
	r.HandleFunc("/albums/{id:[0-9]+}", albumController.GetAlbumDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/albums", albumController.PostAlbumHandler).Methods(http.MethodPost)
	r.HandleFunc("/albums/{id:[0-9]+}", albumController.DeleteAlbumHandler).Methods(http.MethodDelete)

	r.Use(middleware.LoggingMiddleware)

	return r
}
