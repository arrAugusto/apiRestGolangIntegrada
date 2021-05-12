package structuresIngGeneral

///***********************************************************************************************************///
//								ESTRUCTURAS PARA GUARDAR DATA EN LA DB
//
///***********************************************************************************************************///

//Struct para guardar un nuevo ingreso de bodega general paso 1

type IngresoGeneralDat struct {
	IdBod    int     `valid: "IsNumeric"`
	IdNit    int     `valid: "IsNumeric"`
	CantBlts int     `valid: "IsNumeric"`
	ValTotal float64 `valid: "IsFloat"`
	TokenReq string  `valid: "string"`
}

//Struct para guardar un detalle de producto posiblemente del struct anterior

type NewProductGeneral struct {
	IdIng         int     `valid: "IsNumeric"`
	IdUser        int     `valid: "IsNumeric"`
	IdProduct     int     `valid: "IsNumeric"`
	Bultos        int     `valid: "IsNumeric"`
	ValorUnitario float64 `valid: "IsFloat"`
	TokenReq      string  `valid: "string"`
}

//Struct para guardar un detalle de producto posiblemente del struct anterior

type IncidenciaDescarga struct {
	IdIngReq       int    `valid: "IsNumeric"`
	IdDetalleReq   int    `valid: "IsNumeric"`
	DescripcionReq string `valid: "string"`
	TokenReq       string `valid: "string"`
}

//Struct para guardar el metraje de bodega
type NewMetraje struct {
	IdIngReq           int     `valid: "IsNumeric"`
	IdDetalleReq       int     `valid: "IsNumeric"`
	IdAreaBodReq       int     `valid: "IsNumeric"`
	MetrosReq          float64 `valid: "IsFloat"`
	PosicionesReq      int     `valid: "IsNumeric"`
	PromedioTarimaReq  float64 `valid: "IsFloat"`
	MetrosStockReq     float64 `valid: "IsFloat"`
	PosicionesStockReq int     `valid: "IsNumeric"`
	TokenReq           string  `valid: "string"`
}

//Struct para guardar la ubicación de mercaderia
type NewUbicaciones struct {
	IdIngReq     int    `valid: "IsNumeric"`
	IdDetalleReq int    `valid: "IsNumeric"`
	IdAreaBodReq int    `valid: "IsNumeric"`
	Pasillo      int    `valid: "IsNumeric"`
	Columna      int    `valid: "IsNumeric"`
	Comentarios  string `valid: "string"`
	TokenReq     string `valid: "string"`
}

///***********************************************************************************************************///
//								ESTRUCTURAS PARA GUARDAR DATA EN LA DB
//
///***********************************************************************************************************///

//Struct para almacenar datos que modifcaran el ingreso
type UpdateIng struct {
	IdIngReq    int     `valid: "IsNumeric"`
	BultosTotal int     `valid: "IsNumeric"`
	ValorTotal  float64 `valid: "IsFloat"`
	Motivo      string  `valid: "string"`
	TokenReq    string  `valid: "string"`
}

//Struct para almacenar datos que modifcaran los detalles de mercaderia
type UpdateDetalle struct {
	IdIngReq  int     `valid: "IsNumeric"`
	IdProduct int     `valid: "IsNumeric"`
	Bultos    int     `valid: "IsNumeric"`
	PUnitario float64 `valid: "IsFloat"`
	Motivo    string  `valid: "string"`
	TokenReq  string  `valid: "string"`
}

//Struct para almacenar datos que modifcaran los detalles de mercaderia
type UpdateIncidencia struct {
	Motivo      string `valid: string`
	Descripcion string `valid: "string"`
	TokenReq    string `valid: "string"`
}

//Struct para almacenar datos que modifcaran los detalles de mercaderia
type UpdateMetraje struct {
	IdAreaBodReq int     `valid: "IsNumeric"`
	Metros       float64 `valid: "float"`
	Posiciones   float64 `valid: "float"`
	Promedio     float64 `valid: "float"`
	Motivo       string  `valid: string`
	TokenReq     string  `valid: "string"`
}

//Struct para almacenar datos que modifcaran los detalles de mercaderia
type UpdateUbicacion struct {
	IdAreaBodReq int    `valid: "IsNumeric"`
	Pasillo      int    `valid: "isNumeric"`
	Columna      int    `valid: "isNumeric"`
	Comentario   string `valid: "string"`
	Motivo       string `valid: string`
	TokenReq     string `valid: "string"`
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

//Struct Anulacion de detalles de mercaderia
type AnulacionDetallesGeneral struct {
	IdDetalle int    `valid: "isNumeric"`
	Motivo    string `valid: "string"`
	TokenReq  string `valid: "string"`
}

//Struct para guardar un detalle de producto posiblemente del struct anterior

type NewProducto struct {
	Producto string `valid: "string"`
	TokenReq string `valid: "string"`
}

///***********************************************************************************************************///
//						ESTRUCTURAS PARA CONSULTAS RELACIONADAS A INGRESOS DATA EN LA DB
//
///***********************************************************************************************************///
type ConsultaProducto struct {
	Id   string `valid: "int"`
	Name string `valid: "string"`
}

///***********************************************************************************************************///
//								ESTRUCTURAS PARA AUTENTICACION Y RESPUESTAS GENERICAS
//
///***********************************************************************************************************///

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
