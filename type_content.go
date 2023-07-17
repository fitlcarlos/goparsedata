package goparsedata

import "fmt"

type TypeContent int

const (
	TypeCsvContent  TypeContent = 0
	TypeJsonContent TypeContent = 1
)

func GetTypeContent(value int) (TypeContent, error) {
	switch value {
	case 0:
		return TypeCsvContent, nil
	case 1:
		return TypeJsonContent, nil
	default:
		return -1, fmt.Errorf("type not found")
	}
}
