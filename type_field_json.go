package goparsedata

type TypeFieldJson int

const (
	TcjNone         TypeFieldJson = 0
	TcjObject       TypeFieldJson = 1
	TcjList         TypeFieldJson = 2
	TcjObjectList   TypeFieldJson = 3
	TcjFieldElement TypeFieldJson = 4
	TcjAttribute    TypeFieldJson = 5
)
