package memorydb

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

type albumRepository struct {
	sync.RWMutex
	albumMap map[model.AlbumID]*model.Album // キーが AlbumID 値がmodel.Album のマップ
	// model.albumが構造体 → 実際にデータが入っているのがalbumRepository

}

type albumSingerRepository struct {
	sync.RWMutex
	albumSingerMap map[model.AlbumID]*model.AlbumSinger
}

var _ repository.AlbumRepository = (*albumRepository)(nil)

func NewAlbumSingerRepository() *albumSingerRepository {
	var initMap = map[model.AlbumID]*model.AlbumSinger{}
	return &albumSingerRepository{
		albumSingerMap: initMap,
	}
}

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

// SingerIDをもとにsinger:{~~~~}ってやりたい
// レシーバー → albumRepositoryで登録したメソッドが使用できる
func (r *albumRepository) GetAll(ctx context.Context) ([]*model.Album, error) {
	r.RLock()
	defer r.RUnlock()

	// SingerID → からsingerのデータを取得したい → albumSinger構造体作成
	// make([]Tスライスの要素の型, スライスの長さ, スライスの容量)
	albums := make([]*model.Album, 0, len(r.albumMap))
	albumSinger := make([]model.AlbumSinger, 0, len(r.albumMap))
	singers := make([]model.Singer, 0, len(NewSingerRepository().singerMap))

	var albumSingers = []model.AlbumSinger{}
	// singersにsingersデータを入れる
	for _, value := range NewSingerRepository().singerMap {
		fmt.Println("singerValue : ", value) // singerのvalue表示 value.ID value.Name
		singers = append(singers, *value)    // *valueで値が追加
	}
	fmt.Println("singers : ", singers)

	// albumsにalbumデータを入れる
	for _, s := range r.albumMap {
		// appendでalbumsにsを追加していってる s = &{1 Alice's 1st Album 1}

		for _, singersValue := range NewSingerRepository().singerMap {

			// idが同じであれば追加
			if int(s.SingerID) == int(singersValue.ID) {
				albumSingers = append(albumSingers, model.AlbumSinger{ID: s.ID, Title: s.Title, Singer: model.Singer{ID: singersValue.ID, Name: singersValue.Name}})
			}
		}
		albums = append(albums, s)

		// albumSinger = append(albumSinger, model.AlbumSinger{albums[index].ID, albums[index].Title,)
	}
	fmt.Println("albums : ", albums)
	// fmt.Println("albumSinger : ", albumSinger)
	fmt.Println("albumSingers : ", albumSingers)

	// albumSingerにalbumsとsingersのデータを追加したい
	for index, albumSingerValue := range NewAlbumSingerRepository().albumSingerMap {
		albumSingerValue.ID = albums[index].ID
		albumSingerValue.Title = albums[index].Title
		albumSingerValue.Singer = singers[index]
	}
	// シンプルにalbumSingerっていう配列に追加する
	// albumSinger = append(albumSinger, {albumSin*NewAlbumSingerRepository().albumSingerMap[ID]})
	fmt.Println()
	fmt.Println(albumSinger) // ただのからの配列 → 追加したい
	fmt.Println(albums[1].ID)
	// return albums, nil
	return albums, nil
}

func (r *albumRepository) Get(ctx context.Context, id model.AlbumID) (*model.Album, error) {
	r.RLock()
	defer r.RUnlock()

	album, ok := r.albumMap[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return album, nil
}

func (r *albumRepository) Add(ctx context.Context, album *model.Album) error {
	r.Lock()
	r.albumMap[album.ID] = album
	r.Unlock()
	return nil
}

func (r *albumRepository) Delete(ctx context.Context, id model.AlbumID) error {
	r.Lock()
	delete(r.albumMap, id)
	r.Unlock()
	return nil
}
