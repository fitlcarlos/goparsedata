package goparsedata

func NewContent(typeContent TypeContent) Content {
	var content Content

	switch typeContent {
	case TypeCsvContent:
		content = NewCsvContent()
	case TypeJsonContent:
		content = NewJsonContent()
	}
	return content
}
