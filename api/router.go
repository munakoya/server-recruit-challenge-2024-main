package api

// 初期化とルーティング？
import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pulse227/server-recruit-challenge-sample/api/middleware"
	"github.com/pulse227/server-recruit-challenge-sample/controller"
	"github.com/pulse227/server-recruit-challenge-sample/infra/memorydb"
	"github.com/pulse227/server-recruit-challenge-sample/service"
)

func NewRouter() *mux.Router {
	// repository service controllerの3つ
	// model.singerをインスタンス化 初期化
	singerRepo := memorydb.NewSingerRepository()
	// 初期化処理？
	singerService := service.NewSingerService(singerRepo)
	// 初期化
	singerController := controller.NewSingerController(singerService)

	// アルバムデータの初期化？
	albumRepo := memorydb.NewAlbumRepository()
	albumService := service.NewAlbumService(albumRepo)
	albumController := controller.NewAlbumController(albumService)

	// アルバムデータの初期化？
	albumSingerRepo := memorydb.NewAlbumSingerRepository()
	albumSingerService := service.NewAlbumSingerService(albumSingerRepo)
	albumSingerController := controller.NewAlbumSingerController(albumSingerService)

	r := mux.NewRouter()

	// controllo/singersに定義された関数の実行
	r.HandleFunc("/singers", singerController.GetSingerListHandler).Methods(http.MethodGet)
	r.HandleFunc("/singers/{id:[0-9]+}", singerController.GetSingerDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/singers", singerController.PostSingerHandler).Methods(http.MethodPost)
	r.HandleFunc("/singers/{id:[0-9]+}", singerController.DeleteSingerHandler).Methods(http.MethodDelete)

	// ここにalbumsのルーティング追加されるはず
	// r.HandleFunc("/albums", albumController.GetAlbumListHandler).Methods(http.MethodGet)
	// r.HandleFunc("/albums/{id:[0-9]+}", albumController.GetAlbumDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/albums", albumController.PostAlbumHandler).Methods(http.MethodPost)
	r.HandleFunc("/albums/{id:[0-9]+}", albumController.DeleteAlbumHandler).Methods(http.MethodDelete)

	// albumsでalbumSinger~~~メソッド実行 もとのルーティングはコメントアウト
	r.HandleFunc("/albums", albumSingerController.GetAlbumSingerListHandler).Methods(http.MethodGet)
	r.HandleFunc("/albums/{id:[0-9]+}", albumSingerController.GetAlbumSingerDetailHandler).Methods(http.MethodGet)

	r.Use(middleware.LoggingMiddleware)

	return r
}
