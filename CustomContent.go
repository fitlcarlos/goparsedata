package goparsedata

import (
	"fmt"
	"github.com/fitlcarlos/godata"
)

type Content interface {
	ReadTree(rsc *DataSetCollection, list *godata.Strings) error
	GetContent() CustomContent
}
type CustomContent struct {
	EscapeCharacters      bool
	EscapeStringLineBreak string
	Content
}

func (cc *CustomContent) toStrings(rsc *DataSetCollection) (*godata.Strings, error) {
	var list *godata.Strings

	if rsc != nil {
		if len(rsc.Items) > 0 {
			err := cc.Content.ReadTree(rsc, list)
			if err != nil {
				return nil, fmt.Errorf("error listing strings %w\n", err)
			}
		}
	}
	return list, nil
}

func (cc *CustomContent) toString(rsc *DataSetCollection) (string, error) {
	list, err := cc.toStrings(rsc)

	if err != nil {
		return "", err
	}
	return list.Text(), nil
}
