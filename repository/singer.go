package repository

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/model"
)

// データ操作4つのメソッド
type SingerRepository interface {
	// singersで一覧表示 			ここは戻り値の型？ → 課題4はGetAllとGetををいじればできそう
	GetAll(ctx context.Context) ([]*model.Singer, error)
	// singers/id でidに紐づいたデータを表示		ポインタ
	Get(ctx context.Context, id model.SingerID) (*model.Singer, error)
	Add(ctx context.Context, singer *model.Singer) error
	Delete(ctx context.Context, id model.SingerID) error
}
