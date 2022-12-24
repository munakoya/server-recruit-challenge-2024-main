package memorydb

import (
	"context"
	"errors"
	"sync"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

type albumRepository struct {
	sync.RWMutex                                // 排他制御
	albumMap     map[model.AlbumID]*model.Album // キーが AlbumID 値がmodel.Album のマップ
}

var globalAlbum map[model.AlbumID]*model.Album

// repository/album.goで定義したインターフェース
var _ repository.AlbumRepository = (*albumRepository)(nil)

// 初期化？
func NewAlbumRepository() *albumRepository {
	var initMap = map[model.AlbumID]*model.Album{
		1: {ID: 1, Title: "Alice's 1st Album", SingerID: 1}, // SingerIDをもとに参照したい
		2: {ID: 2, Title: "Alice's 2nd Album", SingerID: 1},
		3: {ID: 3, Title: "Bella's 1st Album", SingerID: 2},
	}
	return &albumRepository{
		albumMap: initMap,
	}
}

// 1. GetAll → インスタンス化したalbumデータを一覧表示
func (r *albumRepository) GetAll(ctx context.Context) ([]*model.Album, error) {
	// 制御系
	r.RLock()
	defer r.RUnlock()

	// make([]Tスライスの要素の型, スライスの長さ, スライスの容量)
	albums := make([]*model.Album, 0, len(NewAlbumRepository().albumMap))

	// albumsにalbumデータを入れる
	for _, s := range r.albumMap {
		// appendでalbumsにsを追加していってる s = &{1 Alice's 1st Album 1}
		albums = append(albums, s)
	}
	return albums, nil
}

// 2. Get → 引数で指定されたidに該当するalbumデータを取り出す
func (r *albumRepository) Get(ctx context.Context, id model.AlbumID) (*model.Album, error) {
	r.RLock()
	defer r.RUnlock()
	// 指定されたid要素の取り出し
	album, ok := r.albumMap[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return album, nil
}

// 3. Add → 引数に指定したalbumデータを追加する
func (r *albumRepository) Add(ctx context.Context, album *model.Album) error {
	r.Lock()
	r.albumMap[album.ID] = album
	// 更新
	globalAlbum = r.albumMap
	r.Unlock()
	return nil
}

// 4. Delete → 引数に指定されたidのalbumデータを削除する
func (r *albumRepository) Delete(ctx context.Context, id model.AlbumID) error {
	r.Lock()
	delete(r.albumMap, id)
	// 更新
	globalAlbum = r.albumMap
	r.Unlock()
	return nil
}
