package memorydb

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

// 排他制御とalbumSingerデータが入るフィールドを持つ構造体
type albumSingerRepository struct {
	sync.RWMutex
	albumSingerMap map[model.AlbumID]*model.AlbumSinger
}

// repository/albumSinger.goで定義したインターフェース → オーバーライド
var _ repository.AlbumSingerRepository = (*albumSingerRepository)(nil)

// 初期化
func NewAlbumSingerRepository() *albumSingerRepository {
	var initMap = map[model.AlbumID]*model.AlbumSinger{}
	return &albumSingerRepository{
		albumSingerMap: initMap,
	}
}

// 1. GetAll → インスタンス化したalbumSingerデータを一覧表示
func (r *albumSingerRepository) GetAll(ctx context.Context) ([]*model.AlbumSinger, error) {
	// 制御系
	r.RLock()
	defer r.RUnlock()

	// make([]Tスライスの要素の型, スライスの長さ, スライスの容量)
	albumSinger := make([]*model.AlbumSinger, 0, len(NewAlbumRepository().albumMap))
	albumSinger_test := make([]model.AlbumSinger, 0, len(NewAlbumRepository().albumMap)) // 確認用

	// albumSingerにalbumとsingerデータを入れる
	for _, s := range NewAlbumRepository().albumMap { // 一覧表示できない理由は初期化したマップを回しているから → 別の方法でできれば追加後も表示できる ただし、album, singerのadd必須かな？ → albumSingerにaddとdelはいらない
		for _, singersValue := range NewSingerRepository().singerMap {
			// idが同じであれば追加
			if int(s.SingerID) == int(singersValue.ID) {
				albumSinger = append(albumSinger, &model.AlbumSinger{ID: s.ID, Title: s.Title, Singer: model.Singer{ID: singersValue.ID, Name: singersValue.Name}})
				albumSinger_test = append(albumSinger_test, model.AlbumSinger{ID: s.ID, Title: s.Title, Singer: model.Singer{ID: singersValue.ID, Name: singersValue.Name}})
			}
		}
	}
	fmt.Println("albumSinger : ", albumSinger)
	fmt.Println("albumSinger_test : ", albumSinger_test)
	return albumSinger, nil
}

// 2. Get → 引数で指定されたidに該当するsingerデータを取り出す
func (r *albumSingerRepository) Get(ctx context.Context, id model.AlbumID) (*model.AlbumSinger, error) {
	// 制御系
	r.RLock()
	defer r.RUnlock()
	album, ok := NewAlbumRepository().albumMap[id] // 初期化マップ → 追加後のマップにアクセスするように
	var albumSinger *model.AlbumSinger             // 指定したidの要素のみ → スライスでない
	if !ok {
		return nil, errors.New("not found")
	}
	for _, singerValue := range NewSingerRepository().singerMap { // 上と同様
		if singerValue.ID == album.SingerID {
			// idが等しいsingerデータを取り出し
			albumSinger = &model.AlbumSinger{ID: album.ID, Title: album.Title, Singer: *singerValue}
		}
	}
	return albumSinger, nil
}

// TODO
/*
addされたもの、deleteされたものを一覧表示に反映させる
→ range mapの部分 → 初期化してるものではなく、追加されているmapを指定
albumSingerは、albumsとsingersの両方のaddで表示 → どちらかだけの追加で表示できないっていうのは考慮しなくていい
*/
