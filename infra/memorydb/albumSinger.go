package memorydb

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

type albumSingerRepository struct {
	sync.RWMutex
	albumSingerMap map[model.AlbumID]*model.AlbumSinger
}

// var _ repository.AlbumRepository = (*albumRepository)(nil)

var _ repository.AlbumSingerRepository = (*albumSingerRepository)(nil)

func NewAlbumSingerRepository() *albumSingerRepository {
	var initMap = map[model.AlbumID]*model.AlbumSinger{}
	return &albumSingerRepository{
		albumSingerMap: initMap,
	}
}

// SingerIDをもとにsinger:{~~~~}ってやりたい
// レシーバー → albumRepositoryで登録したメソッドが使用できる
func (r *albumSingerRepository) GetAll(ctx context.Context) ([]*model.AlbumSinger, error) {
	r.RLock()
	defer r.RUnlock()

	// SingerID → からsingerのデータを取得したい → albumSinger構造体作成
	// make([]Tスライスの要素の型, スライスの長さ, スライスの容量)
	albums := make([]*model.Album, 0, len(NewAlbumRepository().albumMap))
	albumSinger := make([]*model.AlbumSinger, 0, len(NewAlbumRepository().albumMap)) // 本当は*
	singers := make([]model.Singer, 0, len(NewSingerRepository().singerMap))

	var albumSingers = []model.AlbumSinger{}
	// singersにsingersデータを入れる
	for _, value := range NewSingerRepository().singerMap {
		fmt.Println("singerValue : ", value) // singerのvalue表示 value.ID value.Name
		singers = append(singers, *value)    // *valueで値が追加
	}
	fmt.Println("singers : ", singers)

	// albumsにalbumデータを入れる
	for _, s := range NewAlbumRepository().albumMap {
		// appendでalbumsにsを追加していってる s = &{1 Alice's 1st Album 1}

		for _, singersValue := range NewSingerRepository().singerMap {

			// idが同じであれば追加
			if int(s.SingerID) == int(singersValue.ID) {
				albumSingers = append(albumSingers, model.AlbumSinger{ID: s.ID, Title: s.Title, Singer: model.Singer{ID: singersValue.ID, Name: singersValue.Name}})
				albumSinger = append(albumSinger, &model.AlbumSinger{ID: s.ID, Title: s.Title, Singer: model.Singer{ID: singersValue.ID, Name: singersValue.Name}})
			}
		}
		albums = append(albums, s)
		// albumSinger = append(albumSinger, model.AlbumSinger{albums[index].ID, albums[index].Title,)
	}
	fmt.Println("albums : ", albums)
	// fmt.Println("albumSinger : ", albumSinger)
	fmt.Println("albumSingers : ", albumSingers)
	fmt.Println("albumSinger : ", albumSinger)

	// albumSingerにalbumsとsingersのデータを追加したい
	// for index, albumSingerValue := range NewAlbumSingerRepository().albumSingerMap {
	// 	albumSingerValue.ID = albums[index].ID
	// 	albumSingerValue.Title = albums[index].Title
	// 	albumSingerValue.Singer = singers[index]
	// }
	// シンプルにalbumSingerっていう配列に追加する
	// albumSinger = append(albumSinger, {albumSin*NewAlbumSingerRepository().albumSingerMap[ID]})
	fmt.Println(albums[1].ID)
	// return albums, nil
	return albumSinger, nil
}

func (r *albumSingerRepository) Get(ctx context.Context, id model.AlbumID) (*model.AlbumSinger, error) {
	r.RLock()
	defer r.RUnlock()
	album, ok := NewAlbumRepository().albumMap[id]
	// var singer *model.Singer
	// var albumSinger = []*model.AlbumSinger{}
	var albumSinger *model.AlbumSinger

	if !ok {
		return nil, errors.New("not found")
	}
	for _, singerValue := range NewSingerRepository().singerMap {
		if singerValue.ID == album.SingerID {
			// idが等しいsingerデータを取り出し
			// singer = &model.Singer{ID: singerValue.ID, Name: singerValue.Name}
			albumSinger = &model.AlbumSinger{ID: album.ID, Title: album.Title, Singer: *singerValue}
		}
	}
	return albumSinger, nil
}

func (r *albumSingerRepository) Add(ctx context.Context, album *model.Album) error {
	r.Lock()
	NewAlbumRepository().albumMap[album.ID] = album
	r.Unlock()
	return nil
}

func (r *albumSingerRepository) Delete(ctx context.Context, id model.AlbumID) error {
	r.Lock()
	delete(NewAlbumRepository().albumMap, id)
	r.Unlock()
	return nil
}
