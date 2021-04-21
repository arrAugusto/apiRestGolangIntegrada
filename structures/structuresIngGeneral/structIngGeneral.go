package structuresIngGeneral

type IngresoGeneralDat struct {
	IdUser   int     `valid: "IsNumeric"`
	IdBod    int     `valid: "IsNumeric"`
	IdNit    int     `valid: "IsNumeric"`
	CantBlts int     `valid: "IsNumeric"`
	ValTotal float64 `valid: "IsFloat"`
}

type RespuestaInsertInGeneral struct {
	RespSQL string `json: respSQL`
}
