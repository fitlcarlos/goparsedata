package goparsedata

import (
	"fmt"
	"github.com/fitlcarlos/godata"
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

func (dsc *GoParseData) AddObject(caption string) *DataSetItem {
	return dsc.DataSets.AddObject(caption)
}

func (dsc *GoParseData) AddObjectList(caption string) *DataSetItem {
	return dsc.DataSets.AddObjectList(caption)
}

func (gpd *GoParseData) toString() (string, error) {
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
	//cc := gpd.content.GetContent()
	//return cc.toString(gpd.DataSets)
}

func (gpd *GoParseData) toStream() ([]byte, error) {

	str, err := gpd.toString()

	if err != nil {
		return nil, err
	}

	output := []byte(str)

	return output, nil
}

func (gpd *GoParseData) saveToStream() ([]byte, error) {
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

//func (gpd *GoParseData) findNoCollectionByName(name string) (*DataSetCollection, error) {
//	return gpd.findNoCollection(gpd.DataSets, name)
//}

//func (gpd *GoParseData) findNoCollection(collection *DataSetCollection, name string) (*DataSetCollection, error) {
//	var rsc *DataSetCollection
//
//	if collection.Name == name {
//		rsc = collection
//	} else {
//		for i := 0; i < collection.count(); i++ {
//			gpd.findNoCollection(collection.getItem(i).SubQueries, name)
//		}
//	}
//
//	return rsc, nil
//}

//func (gpd *GoParseData) findNoItem(name string) (*DataSetItem, error) {
//	return nil, nil
//}

//func (gpd *GoParseData) findNoItem(rsc *DataSetCollection, name string) (*DataSetItem, error) {
//	return nil, nil
//}