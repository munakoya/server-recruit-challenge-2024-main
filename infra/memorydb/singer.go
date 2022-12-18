package memorydb

import (
	"context"
	"errors"
	"sync"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

// 排他制御とsingerデータが入るフィールドを持つ構造体
type singerRepository struct {
	sync.RWMutex                                  // 共通データの排他制御 → ex) RLock() 読み取り専用
	singerMap    map[model.SingerID]*model.Singer // キーが SingerID、値が model.Singer のマップ
}

// repository/singer.goで定義したインターフェース → オーバーライドして実装？
var _ repository.SingerRepository = (*singerRepository)(nil)

// model.singerをインスタンス化 初期化
func NewSingerRepository() *singerRepository {
	var initMap = map[model.SingerID]*model.Singer{
		1: {ID: 1, Name: "Alice"},
		2: {ID: 2, Name: "Bella"},
		3: {ID: 3, Name: "Chris"},
		4: {ID: 4, Name: "Daisy"},
		5: {ID: 5, Name: "Ellen"},
	}

	return &singerRepository{
		singerMap: initMap,
	}
}

// 1. GetAll → インスタンス化したsingerデータを一覧表示
func (r *singerRepository) GetAll(ctx context.Context) ([]*model.Singer, error) {
	r.RLock()         //  読み取り専用
	defer r.RUnlock() // 処理完了 → ロック解除

	// model.Singer型のスライス作成
	singers := make([]*model.Singer, 0, len(r.singerMap))
	// singerRepositoryのsingerMapから要素を取り出してappend
	for _, s := range r.singerMap {
		singers = append(singers, s)
	}
	return singers, nil // 戻り値は[]*model.Singer
}

// 2. Get → 引数で指定されたidに該当するsingerデータを取り出す
func (r *singerRepository) Get(ctx context.Context, id model.SingerID) (*model.Singer, error) {
	r.RLock()
	defer r.RUnlock()

	// 指定されたid要素の取り出し
	singer, ok := r.singerMap[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return singer, nil
}

// 3. Add → 引数に指定したsingerデータを追加する
func (r *singerRepository) Add(ctx context.Context, singer *model.Singer) error {
	r.Lock() // 書き込みロック → このメソッドが実行されるとロック → 処理が終わると解除
	r.singerMap[singer.ID] = singer
	r.Unlock()
	return nil
}

// 4. Delete → 引数に指定されたidのsingerデータを削除する
func (r *singerRepository) Delete(ctx context.Context, id model.SingerID) error {
	r.Lock()
	delete(r.singerMap, id)
	r.Unlock()
	return nil
}
