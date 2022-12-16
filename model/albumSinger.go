package model

// type AlbumID int

type AlbumSinger struct {
	ID     AlbumID `json:"id"`
	Title  string  `json:"title"`
	Singer Singer  `json:"singer"` // モデル Singer の ID と紐づきます
}
