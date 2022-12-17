package repository

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/model"
)

type AlbumSingerRepository interface {
	GetAll(ctx context.Context) ([]*model.AlbumSinger, error)
	Get(ctx context.Context, id model.AlbumID) (*model.AlbumSinger, error)
	Add(ctx context.Context, album *model.Album) error
	Delete(ctx context.Context, id model.AlbumID) error
}
