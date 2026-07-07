package models

type PageData struct {
	Artists []Artist

	Query string
	Order string
	Field string

	CreationMin string
	CreationMax string

	AlbumMin string
	AlbumMax string

	SelectedMembers []int
}