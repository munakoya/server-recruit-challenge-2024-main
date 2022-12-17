package service

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

type AlbumSingerService interface {
	GetAlbumSingerListService(ctx context.Context) ([]*model.AlbumSinger, error)
	GetAlbumSingerService(ctx context.Context, albumID model.AlbumID) (*model.AlbumSinger, error)
	PostAlbumSingerService(ctx context.Context, album *model.Album) error
	DeleteAlbumSingerService(ctx context.Context, albumID model.AlbumID) error
}

type albumSingerService struct {
	albumSingerRepository repository.AlbumSingerRepository
}

var _ AlbumSingerService = (*albumSingerService)(nil)

func NewAlbumSingerService(albumSingerRepository repository.AlbumSingerRepository) *albumSingerService {
	return &albumSingerService{albumSingerRepository: albumSingerRepository}
}

func (s *albumSingerService) GetAlbumSingerListService(ctx context.Context) ([]*model.AlbumSinger, error) {
	Albums, err := s.albumSingerRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return Albums, nil
}

func (s *albumSingerService) GetAlbumSingerService(ctx context.Context, AlbumID model.AlbumID) (*model.AlbumSinger, error) {
	Album, err := s.albumSingerRepository.Get(ctx, AlbumID)
	if err != nil {
		return nil, err
	}
	return Album, nil
}

func (s *albumSingerService) PostAlbumSingerService(ctx context.Context, album *model.Album) error {
	if err := s.albumSingerRepository.Add(ctx, album); err != nil {
		return err
	}
	return nil
}

func (s *albumSingerService) DeleteAlbumSingerService(ctx context.Context, albumID model.AlbumID) error {
	if err := s.albumSingerRepository.Delete(ctx, albumID); err != nil {
		return err
	}
	return nil
}
