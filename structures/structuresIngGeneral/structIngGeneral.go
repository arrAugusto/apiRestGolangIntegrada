package structuresIngGeneral

//Struct para guardar un nuevo ingreso de bodega general paso 1
type IngresoGeneralDat struct {
	IdUser   int     `valid: "IsNumeric"`
	IdBod    int     `valid: "IsNumeric"`
	IdNit    int     `valid: "IsNumeric"`
	CantBlts int     `valid: "IsNumeric"`
	ValTotal float64 `valid: "IsFloat"`
}

//Struct para guardar un detalle de producto posiblemente del struct anterior

type NewProductGeneral struct {
	IdIng         int     `valid: "IsNumeric"`
	IdUser        int     `valid: "IsNumeric"`
	IdProduct     int     `valid: "IsNumeric"`
	Bultos        int     `valid: "IsNumeric"`
	ValorUnitario float64 `valid: "IsFloat"`
}

//Jwt Read
type JwtRead struct {
	Token string `json: "token"`
}

//Respuesta SQL
type RespuestaInsertInGeneral struct {
	RespSQL string `json: respSQL`
}

type ObjIdIng struct {
	IdIngReq int `valid: "IsNumeric"`
}

type ImagenesRuta struct {
	RutaImg string `json: "rutaImg"`
}
