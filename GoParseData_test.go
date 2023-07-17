package goparsedata

import (
	"fmt"
	data "github.com/fitlcarlos/godata"
	"testing"
)

func TestParseDataMasterDetail(t *testing.T) {
	connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"

	db, err := data.NewConnection(data.ORACLE, connectStr)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	ps := NewGoParseData(db, TypeJsonContent)
	Obj1 := ps.AddObject("teste").
		AddSql("select id, descricao from fab_processo").
		AddSql("where id = 41")

	Obj1.AddObjectList("lista").
		AddSql("select id, codigo, descricao, id_processo from fab_operacao").
		AddMasterSource(Obj1.DataSet).
		AddDetailFields("id_processo").
		AddMasterFields("id")

	json, err := ps.toString()

	fmt.Print(json)
}

func TestParseData(t *testing.T) {
	connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"

	db, err := data.NewConnectionOracle(connectStr)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	ps := NewGoParseData(db, TypeJsonContent)
	content := ps.GetContent()
	content.EscapeStringLineBreak = "\n"
	Obj1 := ps.AddObject("process1").
		AddSql("select id, descricao from fab_processo").
		AddSql("where id = 41")

	Obj2 := Obj1.AddObjectList("lista")
	Obj2.AddSql("select id, codigo, descricao, id_processo, ativo from fab_operacao").
		AddMasterSource(Obj1.DataSet).
		AddDetailFields("id_processo").
		AddMasterFields("id").
		HideFields("codigo", "id_processo").
		ConfigField("ativo", "Ativo", true, "S", "N", false)

	_ = ps.AddObject("process2").
		AddSql("select id, descricao from fab_processo").
		AddSql("where id = 42")

	json, err := ps.toString()

	fmt.Print(json)
}
