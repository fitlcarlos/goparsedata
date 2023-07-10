package goparsedata

import (
	"fmt"
	"github.com/fitlcarlos/godata"
	"reflect"
	"strings"
)

type CsvContent struct {
	Separator string
	CustomContent
}

func NewCsvContent() Content {
	csv := &CsvContent{}
	csv.CustomContent.EscapeStringLineBreak = ""
	csv.CustomContent.EscapeCharacters = false

	return csv
}
func (csv *CsvContent) ReadTree(dsi *DataSetCollection, list *godata.Strings) error {
	if dsi.count() > 0 {
		if dsi.Items[0].DataSet != nil {

			err := dsi.Items[0].DataSet.Open()
			if err != nil {
				return err
			}

			err = csv.ReadLine(dsi.Items[0], list)

			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (csv *CsvContent) ReadLine(dsi *DataSetItem, list *godata.Strings) error {
	var value string
	var line string

	dsi.DataSet.First()
	for !dsi.DataSet.Eof() {
		for _, field := range dsi.DataSet.Fields.List {
			if field.Visible {
				if field.AsValue() != nil {
					switch field.DataType.ScanType().Kind() {

					case reflect.Float32, reflect.Float64:
						value = fmt.Sprintf("%.2f", field.AsFloat64())
					default:
						value = csv.removeEscape(field.AsString())
					}
				}
				if dsi.DataSet.Recno == dsi.DataSet.Count() {
					line = line + value
				} else {
					line = line + csv.Separator + value
				}
			}
		}
		list.Add(line)
		dsi.DataSet.Next()
	}

	return nil
}

func (csv *CsvContent) removeEscape(value string) string {
	if csv.EscapeCharacters {
		strings.ReplaceAll(value, "\r\n", csv.EscapeStringLineBreak)
		strings.ReplaceAll(value, "\r", csv.EscapeStringLineBreak)
		strings.ReplaceAll(value, "\n", csv.EscapeStringLineBreak)
	}
	return value
}
