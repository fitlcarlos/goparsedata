package goparsedata

type Field struct {
	Name       string
	Caption    string
	BoolValue  bool
	TrueValue  string
	FalseValue string
	AcceptNull bool
	Visible    bool
}
