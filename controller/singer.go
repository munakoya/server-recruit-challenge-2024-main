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

type singerController struct {
	service service.SingerService
}

// 初期化
func NewSingerController(s service.SingerService) *singerController {
	return &singerController{service: s}
}

// GET /singers のハンドラー
func (c *singerController) GetSingerListHandler(w http.ResponseWriter, r *http.Request) {
	// GetSingerListServiceの呼び出し → Getの呼び出し → 3段階にする必要性
	singers, err := c.service.GetSingerListService(r.Context())
	if err != nil {
		// エラー発生時に実行される
		errorHandler(w, r, 500, err.Error())
		return
	}
	// リクエストヘッダーにjsonを指定
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(singers)
}

// GET /singers/{id} のハンドラー
func (c *singerController) GetSingerDetailHandler(w http.ResponseWriter, r *http.Request) {
	// strconv.Atoi()引数に指定した文字列を数値型に変換 文字列のid → 数値id
	singerID, err := strconv.Atoi(mux.Vars(r)["id"])
	// エラー処理
	if err != nil {
		err = fmt.Errorf("invalid path param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	// service.GetsingerService → Get
	singer, err := c.service.GetSingerService(r.Context(), model.SingerID(singerID))
	// エラー処理
	if err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(singer)
}

// POST /singers のハンドラー
func (c *singerController) PostSingerHandler(w http.ResponseWriter, r *http.Request) {
	var singer *model.Singer
	if err := json.NewDecoder(r.Body).Decode(&singer); err != nil {
		err = fmt.Errorf("invalid body param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	if err := c.service.PostSingerService(r.Context(), singer); err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(singer)
}

// DELETE /singers/{id} のハンドラー
func (c *singerController) DeleteSingerHandler(w http.ResponseWriter, r *http.Request) {
	singerID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		err = fmt.Errorf("invalid path param: %w", err)
		errorHandler(w, r, 400, err.Error())
		return
	}

	if err := c.service.DeleteSingerService(r.Context(), model.SingerID(singerID)); err != nil {
		errorHandler(w, r, 500, err.Error())
		return
	}
	w.WriteHeader(204)
}
