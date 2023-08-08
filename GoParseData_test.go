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

	json, err := ps.ToString()

	fmt.Print(json)
}

func TestParseDataMasterDetail2(t *testing.T) {
	connectStr := "oracle://nbsama:new@100.0.65.224:1521/fab"

	db, err := data.NewConnection(data.ORACLE, connectStr)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	ps := NewGoParseData(db, TypeJsonContent)
	Obj1 := ps.AddObject("").
		AddSql("select 1 as ID from dual")

	Obj1.HideField("ID")

	Obj2 := Obj1.AddObject("request")

	Obj2.AddSql("SELECT").
		AddSql("1 as ID").
		AddSql(",B.COD_AGENDA_NISSAN AS \"idAgendamento\"").
		AddSql(",B.STATUS as \"status\"").
		AddSql(",CASE").
		AddSql("WHEN B.STATUS <> 3 THEN").
		AddSql("TO_CHAR(C.DATA_AGENDADA,'RRRR-MM-DD')").
		AddSql("ELSE").
		AddSql("TO_CHAR(B.DATA_CANCELAMENTO,'RRRR-MM-DD')").
		AddSql("END AS \"data\"").
		AddSql(",CASE").
		AddSql("WHEN B.STATUS <> 3 THEN").
		AddSql("TO_CHAR(C.DATA_AGENDADA,'hh24:mi:ss')").
		AddSql("ELSE").
		AddSql("TO_CHAR(B.DATA_CANCELAMENTO,'hh24:mi:ss')").
		AddSql("END AS \"hora\"").
		AddSql("FROM FAB_MOV_EI_AGD_NISS_SBOK A").
		AddSql("LEFT JOIN FAB_EI_AGD_NISS_SBOK B ON B.ID = A.ID_EI_AGD_NISS_SBOK").
		AddSql("LEFT JOIN OS_AGENDA C ON C.COD_EMPRESA = B.COD_EMPRESA AND C.COD_OS_AGENDA = B.COD_OS_AGENDA").
		AddSql("LEFT JOIN OS_AGENDA_CANC D ON D.COD_EMPRESA = B.COD_EMPRESA AND D.COD_OS_AGENDA = B.COD_OS_AGENDA").
		AddSql("LEFT JOIN FAB_MOV_NISS_SBOK E ON E.ID = A.ID_MOV_NISS_SBOK").
		AddSql("LEFT JOIN FAB_CONFIG F ON F.ID_EMPRESA = E.ID_EMPRESA AND F.ID_OPERACAO = E.ID_OPERACAO").
		AddSql("WHERE").
		AddSql("A.ID_MOV_NISS_SBOK = 1405").
		//AddSql("select id, codigo, descricao, id_processo from fab_operacao").
		AddMasterSource(Obj1.DataSet).
		AddDetailFields("ID").
		AddMasterFields("ID")

	Obj2.HideField("ID")

	json, err := ps.ToString()

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

	content.(*JsonContent).EscapeStringLineBreak = "\n"
	content.(*JsonContent).QuotedFields = true

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

	json, err := ps.ToString()

	fmt.Print(json)
}
