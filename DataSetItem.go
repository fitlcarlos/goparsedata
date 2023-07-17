package goparsedata

import (
	"github.com/fitlcarlos/godata"
	"strings"
)

type DataSetItem struct {
	Owner       any
	Connection  *godata.Conn
	Caption     string
	BeginText   string
	EndText     string
	Separator   string
	FieldType   TypeFieldJson
	FieldsHide  []string
	FieldsConf  []*Field
	DataSet     *godata.DataSet
	SubQueries  *DataSetCollection
	RootNode    bool
	Level       int
	Indentation int
	Index       int
}

func NewDataSetItem(collection any) *DataSetItem {
	dsi := &DataSetItem{
		Owner:      collection,
		FieldsHide: []string{},
		FieldsConf: []*Field{},
	}

	dsc, ok := collection.(*DataSetCollection)
	if ok {
		dsi.DataSet = godata.NewDataSet(dsc.getConnection())
	}

	dsi.FieldType = TcjFieldElement
	dsi.SubQueries = NewDataSetCollection(dsi)
	dsi.SubQueries.RootNode = dsi.RootNode
	dsi.RootNode = collection.(*DataSetCollection).RootNode

	if dsi.RootNode {
		dsi.Level = 1
		dsi.Indentation = 0
	} else {
		dsi.Level = collection.(*DataSetCollection).Owner.(*DataSetItem).Level + 1
		dsi.Indentation = collection.(*DataSetCollection).Owner.(*DataSetItem).Indentation + 2
	}

	return dsi
}
func (dsi *DataSetItem) getConnection() *godata.Conn {
	var conn *godata.Conn

	dsc, ok := dsi.Owner.(*DataSetCollection)
	if ok {
		conn = dsc.getConnection()
	} else {
		conn = nil
	}
	return conn
}
func (dsi *DataSetItem) And() *DataSetItem {
	dsc, ok := dsi.Owner.(*DataSetCollection)
	if ok {
		item, ok := dsc.Owner.(*DataSetItem)
		if ok {
			return item
		} else {
			return nil
		}
	} else {
		return nil
	}
}
func (dsi *DataSetItem) AddObject(caption string) *DataSetItem {
	return dsi.SubQueries.AddObject(caption)
}
func (dsi *DataSetItem) AddList(caption string) *DataSetItem {
	return dsi.SubQueries.AddList(caption)
}
func (dsi *DataSetItem) AddObjectList(caption string) *DataSetItem {
	return dsi.SubQueries.AddObjectList(caption)
}

func (dsi *DataSetItem) AddSql(sql string) *DataSetItem {
	dsi.DataSet.AddSql(sql)
	return dsi
}

func (dsi *DataSetItem) Prepare(sql string) *DataSetItem {
	dsi.DataSet.Prepare()
	return dsi
}

func (dsi *DataSetItem) AddMasterSource(dataSource *godata.DataSet) *DataSetItem {
	dsi.DataSet.AddMasterSource(dataSource)
	return dsi
}
func (dsi *DataSetItem) AddDetailFields(fields ...string) *DataSetItem {
	dsi.DataSet.AddDetailFields(fields...)
	return dsi
}

func (dsi *DataSetItem) AddMasterFields(fields ...string) *DataSetItem {
	dsi.DataSet.AddMasterFields(fields...)
	return dsi
}

func (dsi *DataSetItem) ConfigField(fieldName string, caption string, boolValue bool, trueValue string, falseValue string, acceptNull bool) *DataSetItem {
	var field *Field
	field = dsi.fieldByName(fieldName)
	if field == nil {
		field := &Field{
			Name:       strings.ToUpper(fieldName),
			Caption:    caption,
			BoolValue:  boolValue,
			TrueValue:  trueValue,
			FalseValue: falseValue,
			AcceptNull: acceptNull,
		}
		dsi.FieldsConf = append(dsi.FieldsConf, field)
	}
	return dsi
}

func (dsi *DataSetItem) HideField(fieldName string) *DataSetItem {
	dsi.FieldsHide = append(dsi.FieldsHide, fieldName)
	return dsi
}

func (dsi *DataSetItem) HideFields(fieldsName ...string) *DataSetItem {
	for i := 0; i < len(fieldsName); i++ {
		dsi.HideField(fieldsName[i])
	}
	return dsi
}

func (dsi *DataSetItem) SetInputParam(paramName string, paramValue any) *DataSetItem {
	dsi.DataSet.SetInputParam(paramName, paramValue)
	return dsi
}

func (dsi *DataSetItem) SetOutputParam(paramName string, paramType any) *DataSetItem {
	dsi.DataSet.SetOutputParam(paramName, paramType)
	return dsi
}

func (dsi *DataSetItem) SetMacro(macroName string, macroValue any) *DataSetItem {
	dsi.DataSet.SetMacro(macroName, macroValue)
	return dsi
}

func (dsi *DataSetItem) fieldByName(fieldName string) *Field {
	for i := 0; i < len(dsi.FieldsConf); i++ {
		if strings.ToUpper(dsi.FieldsConf[i].Name) == strings.ToUpper(fieldName) {
			return dsi.FieldsConf[i]
		}
	}
	return nil
}
