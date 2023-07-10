package goparsedata

import (
	"fmt"
	"github.com/fitlcarlos/godata"
	"reflect"
	"strings"
)

type JsonContent struct {
	QuotedFields bool
	CustomContent
}

func NewJsonContent() Content {
	json := &JsonContent{}
	json.QuotedFields = true
	json.CustomContent.EscapeStringLineBreak = ""
	json.CustomContent.EscapeCharacters = false

	return json
}

func (jc *JsonContent) GetContent() CustomContent {
	return jc.CustomContent
}

func (jc *JsonContent) ReadTree(rsc *DataSetCollection, list *godata.Strings) error {

	if rsc.count() > 0 {
		for i := 0; i < rsc.count(); i++ {
			item := rsc.Items[i]

			if item.DataSet != nil {
				err := item.DataSet.Open()
				if err != nil {
					return err
				}

				switch item.FieldType {
				case TcjObject:
					return jc.loadObject(item, list, TcjObject, item.Indentation)
				case TcjList:
					return jc.loadList(item, list, TcjList, item.Indentation)
				case TcjObjectList:
					return jc.loadObjectList(item, list, TcjObjectList, item.Indentation)
				}
			}

		}
	}
	return nil
}

func (jc *JsonContent) loadObject(dsi *DataSetItem, list *godata.Strings, fieldType TypeFieldJson,
	indentation int) error {
	var key string
	var value string
	fmt.Println(key, value)

	dsi.DataSet.First()
	for !dsi.DataSet.Eof() {
		jc.openObject(dsi, list, fieldType, indentation)
		for i := 0; i < len(dsi.DataSet.Fields.List); i++ {
			field := dsi.DataSet.Fields.List[i]
			if field.Visible {
				key = jc.getQuotedField(field.Caption)

				if field.AsValue() == nil {
					if field.AcceptNull {
						value = "null"
					} else {
						switch field.DataType.ScanType().Kind() {
						case reflect.String:
							value = "\"\""
						default:
							value = "null"
						}
					}
				} else {
					switch dsi.DataSet.Connection.Dialect {
					case godata.FIREBIRD:
					case godata.INTERBASE:
					case godata.MYSQL:
					case godata.ORACLE:
						switch field.DataType.DatabaseTypeName() {
						case "NUMBER", "LONG", "UINT":
							_, size, _ := field.DataType.DecimalSize()
							if size == 0 {
								value = field.AsString()
							} else if size > 0 {
								value = fmt.Sprintf("%.2f", field.AsFloat64())
							}
						default:
							if field.BoolValue {
								if field.AsString() == field.TrueValue {
									value = "true"
								} else {
									value = "false"
								}
							} else {
								value = "\"" + jc.removeEscape(field.AsString()) + "\""
							}
						}
					case godata.POSTGRESQL:
					case godata.SQLSERVER:
					case godata.SQLITE:

					}
				}

			}

			if field.Index == len(dsi.DataSet.Fields.List)-1 {
				if dsi.SubQueries.count() == 0 {
					list.Add(jc.getSpace(indentation+2) + key + " : " + value)
				} else {
					list.Add(jc.getSpace(indentation+2) + key + " : " + value + ",")
				}
			} else {
				list.Add(jc.getSpace(indentation+2) + key + " : " + value + ",")
			}
		}

		if dsi.SubQueries != nil {
			if dsi.SubQueries.count() > 0 {
				jc.ReadTree(dsi.SubQueries, list)
			}
		}

		jc.closeObject(dsi, list, fieldType, indentation)

		dsi.DataSet.Next()
	}

	return nil
}

func (jc *JsonContent) loadObjectList(dsi *DataSetItem, list *godata.Strings, fieldType TypeFieldJson, indentation int) error {
	jc.openObject(dsi, list, fieldType, indentation)

	err := jc.loadObject(dsi, list, TcjObject, indentation+2)

	if err != nil {
		return err
	}

	jc.closeObject(dsi, list, fieldType, indentation)

	return nil
}

func (jc *JsonContent) loadList(dsi *DataSetItem, list *godata.Strings, fieldType TypeFieldJson, indentation int) error {
	var value = ""

	jc.openObject(dsi, list, fieldType, indentation)

	dsi.DataSet.First()
	for !dsi.DataSet.Eof() {
		for _, column := range dsi.DataSet.Fields.List {
			field := dsi.DataSet.FieldByName(column.Name)
			if field.Visible {
				if field.AsValue() != nil {
					switch field.DataType.ScanType().Kind() {
					case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64,
						reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
						reflect.Uint64:
						value = field.AsString()
					case reflect.Float32, reflect.Float64:
						value = fmt.Sprintf("%.2f", field.AsFloat64())
					default:
						value = "\"" + jc.removeEscape(field.AsString()) + "\""
					}
				}

				if dsi.DataSet.Recno == dsi.DataSet.Count() {
					list.Add(jc.getSpace(indentation + 2))
				} else {
					list.Add(jc.getSpace(indentation+2) + value + ",")
				}
			}

			if dsi.SubQueries != nil {
				if dsi.SubQueries.count() > 0 {
					err := jc.ReadTree(dsi.SubQueries, list)

					if err != nil {
						return err
					}
				}
			}

			dsi.DataSet.Next()
		}
		jc.closeObject(dsi, list, fieldType, indentation)
	}

	return nil
}

func (jc *JsonContent) openObject(dsi *DataSetItem, list *godata.Strings, fieldType TypeFieldJson, indentation int) {
	space := jc.getSpace(indentation)
	var title = strings.TrimSpace(dsi.Caption)

	if title != "" {
		title = jc.getQuotedField(dsi.Caption) + ":"
	}

	switch fieldType {
	case TcjObject:
		if dsi.FieldType == TcjObject {
			list.Add(space + title + "{")
		} else {
			list.Add(space + "{")
		}
		break
	case TcjObjectList:
		list.Add(space + title + "[")
		break
	case TcjList:
		list.Add(space + title + "[")
		break
	}
}

func (jc *JsonContent) closeObject(dsi *DataSetItem, list *godata.Strings, fieldType TypeFieldJson, indentation int) {
	var space string
	var lastLine = false

	space = jc.getSpace(indentation)

	var line = list.Items[list.Count()-1]

	lineLength := len(line)

	var comma = line[lineLength-1 : lineLength]

	if comma == "," {
		list.Add(line[0 : len(line)-1])
	}

	switch fieldType {
	case TcjObject:
		if dsi.FieldType == TcjObject {
			lastLine = dsi.Index == dsi.Owner.(*DataSetCollection).count()-1
		} else {
			lastLine = dsi.DataSet.Recno == dsi.DataSet.Count()
		}

		if lastLine {
			if dsi.FieldType == TcjObject {
				list.Add(space + "}")
			} else {
				list.Add(space + "}")
			}
		} else {
			if dsi.FieldType == TcjObject {
				list.Add(space + "},")
			} else {
				list.Add(space + "},")
			}
		}
	case TcjObjectList:
		list.Add(space + "]")
	case TcjList:
		list.Add(space + "]")
	}
}

func (jc *JsonContent) removeEscape(value string) string {

	var character = value

	if jc.EscapeCharacters {
		strings.ReplaceAll(character, "\\", "\\\\")
		strings.ReplaceAll(character, "\"", "\\")
		strings.ReplaceAll(character, "\r\n", jc.EscapeStringLineBreak)
		strings.ReplaceAll(character, "\r", jc.EscapeStringLineBreak)
		strings.ReplaceAll(character, "\n", jc.EscapeStringLineBreak)

	}
	return character
}

func (jc *JsonContent) getQuotedField(value string) string {
	if jc.QuotedFields {
		return "\"" + value + "\""
	} else {
		return "" + value + ""
	}
}

func (jc *JsonContent) getSpace(quantity int) string {
	var space string
	var qt = quantity

	for qt > 0 {
		space = space + " "
		qt--
	}
	return space
}
