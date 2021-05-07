package structuresIngGeneral

///***********************************************************************************************************///
//								ESTRUCTURAS PARA GUARDAR DATA EN LA DB
//
///***********************************************************************************************************///

//Struct para guardar un nuevo ingreso de bodega general paso 1

type NewRetiroGeneral struct {
	TotalBultos int    `valid: IsNumeric`
	TokenReq    string `valid: string`
}

//Struct para guardar un nuevo ingreso de bodega general paso 1

type DetalleRetGen struct {
	IdRet       int    `valid: IsNumeric`
	IdDetalle   int    `valid: IsNumeric`
	TotalBultos int    `valid: "IsNumeric"`
	TokenReq    string `valid: string`
}

///***********************************************************************************************************///
//									ESTRUCTURAS PARA ANULAR DATA EN LA DB
//
///***********************************************************************************************************///

//Struct remove ingresos

type AnulacionFormas struct {
	Motivo   string `valid: "string"`
	TokenReq string `valid: "string"`
}
