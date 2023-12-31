package goparsedata

import (
	"fmt"
	"github.com/fitlcarlos/godata"
	"strings"
)

type GoParseData struct {
	Connection  *godata.Conn
	DataSets    *DataSetCollection
	Items       []any
	content     Content
	typeContent TypeContent
}

func NewGoParseData(conn *godata.Conn, typeContent TypeContent) *GoParseData {
	gpd := &GoParseData{
		Connection: conn,
	}

	gpd.DataSets = NewDataSetCollection(gpd)
	gpd.Items = []any{}
	gpd.content = NewContent(typeContent)

	return gpd
}

func (dsc *GoParseData) GetContent() Content {
	return dsc.content
}

func (dsc *GoParseData) AddObject(caption string) *DataSetItem {
	return dsc.DataSets.AddObject(caption)
}

func (dsc *GoParseData) AddObjectList(caption string) *DataSetItem {
	return dsc.DataSets.AddObjectList(caption)
}

func (gpd *GoParseData) ToString() (string, error) {
	var list godata.Strings

	if gpd.DataSets != nil {
		if len(gpd.DataSets.Items) > 0 {
			err := gpd.content.ReadTree(gpd.DataSets, &list)
			if err != nil {
				return "", fmt.Errorf("error listing strings %w\n", err)
			}
		}
	}
	return list.Text(), nil
}

func (gpd *GoParseData) ToStream() ([]byte, error) {

	str, err := gpd.ToString()

	if err != nil {
		return nil, err
	}

	output := []byte(str)

	return output, nil
}

func (gpd *GoParseData) SaveToStream() ([]byte, error) {
	var list *godata.Strings

	if gpd.DataSets != nil {
		if len(gpd.DataSets.Items) > 0 {
			err := gpd.content.ReadTree(gpd.DataSets, list)
			if err != nil {
				return nil, fmt.Errorf("error listing strings %w\n", err)
			}
		}
	}

	output := []byte(list.Text())

	return output, nil
}

func (gpd *GoParseData) FindNoCollectionByName(name string) (*DataSetCollection, error) {
	return gpd.FindNoCollection(gpd.DataSets, name)
}

func (gpd *GoParseData) FindNoCollection(collection *DataSetCollection, name string) (*DataSetCollection, error) {
	var rsc *DataSetCollection

	if collection.Name == name {
		rsc = collection
	} else {
		for i := 0; i < collection.count(); i++ {
			gpd.FindNoCollection(collection.getItem(i).SubQueries, name)
		}
	}

	return rsc, nil
}

func (gpd *GoParseData) FindNoItem(name string) *DataSetItem {
	return gpd.FindNoItemByCollection(gpd.DataSets, name)
}

func (gpd *GoParseData) FindNoItemByCollection(rsc *DataSetCollection, name string) *DataSetItem {
	var rsi *DataSetItem

	for i := 0; i < rsc.count(); i++ {
		if strings.ToUpper(rsc.getItem(i).Name) == strings.ToUpper(name) {
			rsi = rsc.getItem(i)
			break
		} else {
			rsi = gpd.FindNoItemByCollection(rsc.getItem(i).SubQueries, name)
		}
	}

	return rsi
}
