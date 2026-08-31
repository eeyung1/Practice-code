package models

type FilterOptions struct {
	Query	string
	Order	string
	Field	string

	CreationMin	int
	CreationMax	int

	AlbumMin	int
	AlbumMax	int

	Members		[]int

	Countries	[]string
}