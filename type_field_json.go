package goparsedata

import "fmt"

type TypeFieldJson int

const (
	TcjNone         TypeFieldJson = 0
	TcjObject       TypeFieldJson = 1
	TcjList         TypeFieldJson = 2
	TcjObjectList   TypeFieldJson = 3
	TcjFieldElement TypeFieldJson = 4
	TcjAttribute    TypeFieldJson = 5
)

func GetTypeJson(value int) (TypeFieldJson, error) {
	switch value {
	case 0:
		return TcjNone, nil
	case 1:
		return TcjObject, nil
	case 2:
		return TcjList, nil
	case 3:
		return TcjObjectList, nil
	case 4:
		return TcjFieldElement, nil
	case 5:
		return TcjAttribute, nil
	default:
		return -1, fmt.Errorf("type not found")
	}
}
