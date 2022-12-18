package service

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

type SingerService interface {
	GetSingerListService(ctx context.Context) ([]*model.Singer, error)
	GetSingerService(ctx context.Context, singerID model.SingerID) (*model.Singer, error)
	PostSingerService(ctx context.Context, singer *model.Singer) error
	DeleteSingerService(ctx context.Context, singerID model.SingerID) error
}

type singerService struct {
	singerRepository repository.SingerRepository
}

// 上で定義したインターフェース → オーバーライドして実装？
var _ SingerService = (*singerService)(nil)

// 初期化？
func NewSingerService(singerRepository repository.SingerRepository) *singerService {
	return &singerService{singerRepository: singerRepository}
}

// オーバーライド
func (s *singerService) GetSingerListService(ctx context.Context) ([]*model.Singer, error) {
	// GetAllの実行 → singersにsingerデータが入る  → 挟む必要性 戻り値も同じ
	singers, err := s.singerRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return singers, nil
}

// 引数で指定されたsingerIDに該当するsingerデータを出力
func (s *singerService) GetSingerService(ctx context.Context, singerID model.SingerID) (*model.Singer, error) {
	singer, err := s.singerRepository.Get(ctx, singerID)
	if err != nil {
		return nil, err
	}
	return singer, nil
}

// 追加
func (s *singerService) PostSingerService(ctx context.Context, singer *model.Singer) error {
	if err := s.singerRepository.Add(ctx, singer); err != nil {
		return err
	}
	return nil
}

// 削除
func (s *singerService) DeleteSingerService(ctx context.Context, singerID model.SingerID) error {
	if err := s.singerRepository.Delete(ctx, singerID); err != nil {
		return err
	}
	return nil
}
