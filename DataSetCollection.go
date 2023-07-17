package goparsedata

import (
	"github.com/fitlcarlos/godata"
	"reflect"
)

type DataSetCollection struct {
	Owner     any
	Caption   string
	Name      string
	RootNode  bool
	ItemClass reflect.Type
	Items     []*DataSetItem
}

func NewDataSetCollection(owner any) *DataSetCollection {
	dsc := &DataSetCollection{
		Owner: owner,
		Items: []*DataSetItem{},
	}

	_, ok := owner.(*GoParseData)

	if ok {
		dsc.Name = "Root"
		dsc.RootNode = true
	} else {
		dsi, ok := owner.(*DataSetItem)
		if ok {
			dsc.Name = dsi.Name
			dsc.RootNode = false
		}
	}

	return dsc
}

func (dsc *DataSetCollection) count() int {
	return len(dsc.Items)
}

func (dsc *DataSetCollection) getItem(index int) *DataSetItem {
	return dsc.Items[index]
}

func (dsc *DataSetCollection) AddObject(caption string) *DataSetItem {
	dsi := NewDataSetItem(dsc)
	dsi.FieldType = TcjObject
	dsi.Caption = caption
	dsc.Items = append(dsc.Items, dsi)
	dsi.Index = len(dsc.Items) - 1

	return dsi
}

func (dsc *DataSetCollection) AddList(caption string) *DataSetItem {
	dsi := NewDataSetItem(dsc)
	dsi.FieldType = TcjList
	dsi.Caption = caption
	dsc.Items = append(dsc.Items, dsi)
	dsi.Index = len(dsc.Items) - 1

	return dsi
}

func (dsc *DataSetCollection) AddObjectList(caption string) *DataSetItem {
	dsi := NewDataSetItem(dsc)
	dsi.FieldType = TcjObjectList
	dsi.Caption = caption
	dsc.Items = append(dsc.Items, dsi)
	dsi.Index = len(dsc.Items) - 1

	return dsi
}

func (dsc *DataSetCollection) Add(name string, caption string, fieldType TypeFieldJson) *DataSetItem {
	dsi := NewDataSetItem(dsc)
	dsi.FieldType = fieldType
	dsi.Name = name
	dsi.Caption = caption
	dsc.Items = append(dsc.Items, dsi)
	dsi.Index = len(dsc.Items) - 1

	return dsi
}

func (dsc *DataSetCollection) addItem(dsi *DataSetItem) int {
	dsc.Items = append(dsc.Items, dsi)
	dsi.Index = len(dsc.Items) - 1

	return dsi.Index
}

func (dsc *DataSetCollection) getConnection() *godata.Conn {
	var conn *godata.Conn

	gpd, ok := dsc.Owner.(*GoParseData)

	if ok {
		conn = gpd.Connection
	} else {
		rsi, ok := dsc.Owner.(*DataSetItem)
		if ok {
			conn = rsi.getConnection()
		}
	}
	return conn
}
