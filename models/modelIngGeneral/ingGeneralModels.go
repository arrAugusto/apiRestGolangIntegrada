package modelIngGeneral

import (
	"fmt"
	"log"
	"time"

	StructDB "../../structures/structuresIngGeneral"
	Conecta "../databaseSQL"
)

//Guardando en la base de datos el nuevo ingreso
func MdlNuevoIngresoGeneral(IdUser int, IdBod int, IdNit int, CantBlts int, ValTotal float64) []StructDB.RespuestaInsertInGeneral {
	//SETEANDO LA DATA EN EL STRUCT
	respuesta := []StructDB.RespuestaInsertInGeneral{}
	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//Tomando la hora y fecha actual para la fecha de registro
	//instanciando el objeto
	var resp StructDB.RespuestaInsertInGeneral
	dt := time.Now()
	//	fmt.Println(dt)
	rows, err := Conecta.ConectionSQL().Query("EXEC spInsertIngGeneral ?, ?, ?, ?, ?, ?, ?, ?, ?", IdBod, IdUser, IdNit, CantBlts, ValTotal, CantBlts, ValTotal, dt, dt)
	if err != nil {
		log.Fatal("Error de lectura de procedimiento almacenado")
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.RespSQL)
		if err != nil {
			log.Fatal("Error de lectura de procedimiento almacenado")
		}
		respuesta = append(respuesta, resp)
		//	names = append(names, id)
	}
	fmt.Println(rows)
	return respuesta

}

//Guardando un nuevo producto en bodega general
//Guardando en la base de datos el nuevo ingreso
func MdlNewProductGeneral(IdIng int, IdUser int, IdProduct int, Bultos int, ValorUnitario float64) []StructDB.RespuestaInsertInGeneral {
	//SETEANDO LA DATA EN EL STRUCT
	respuesta := []StructDB.RespuestaInsertInGeneral{}
	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//Tomando la hora y fecha actual para la fecha de registro
	//instanciando el objeto
	var resp StructDB.RespuestaInsertInGeneral
	dt := time.Now()
	fmt.Println(dt)
	//	fmt.Println(dt)
	rows, err := Conecta.ConectionSQL().Query("EXEC spNewProducGeneral ?, ?, ?, ?, ?, ? ", IdIng, IdUser, IdProduct, Bultos, ValorUnitario, dt)
	if err != nil {
		log.Fatal("Error de lectura de procedimiento almacenado")
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.RespSQL)
		if err != nil {
			log.Fatal("Error al gardar el producto")
		}
		respuesta = append(respuesta, resp)
		//	names = append(names, id)
	}
	fmt.Println(rows)
	return respuesta

}

//Guardando un nuevo producto en bodega general
//Guardando en la base de datos el nuevo ingreso
func MdlRemoveEstadosActivos(idOperacion int, idUsuario int, motivo string, storeProduce string) []StructDB.RespuestaInsertInGeneral {
	//SETEANDO LA DATA EN EL STRUCT
	respuesta := []StructDB.RespuestaInsertInGeneral{}
	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//Tomando la hora y fecha actual para la fecha de registro
	//instanciando el objeto
	var resp StructDB.RespuestaInsertInGeneral
	dt := time.Now()
	fmt.Println(dt)
	//	fmt.Println(dt)
	rows, err := Conecta.ConectionSQL().Query("EXEC "+storeProduce+" ?, ?, ?, ?", idOperacion, idUsuario, motivo, dt)
	if err != nil {
		log.Fatal("Error de lectura de procedimiento almacenado")
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.RespSQL)
		if err != nil {
			log.Fatal("Error al gardar el producto")
		}
		respuesta = append(respuesta, resp)
		//	names = append(names, id)
	}
	fmt.Println(rows)
	return respuesta

}

//Guardando un nuevo producto en bodega general
//Guardando en la base de datos el nuevo ingreso
func MdlRemoveEstdDetGeneral(IdIngGeneral int, idDetalle int, IdUserInt int, motivo string, storeProduce string) []StructDB.RespuestaInsertInGeneral {
	//SETEANDO LA DATA EN EL STRUCT
	respuesta := []StructDB.RespuestaInsertInGeneral{}
	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//Tomando la hora y fecha actual para la fecha de registro
	//instanciando el objeto
	var resp StructDB.RespuestaInsertInGeneral
	dt := time.Now()

	rows, err := Conecta.ConectionSQL().Query("EXEC "+storeProduce+" ?, ?, ?, ?, ?", IdIngGeneral, idDetalle, IdUserInt, motivo, dt)
	if err != nil {
		log.Fatal("Error de lectura de procedimiento almacenado")
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.RespSQL)
		if err != nil {
			log.Fatal("Error al gardar el producto")
		}
		respuesta = append(respuesta, resp)
		//	names = append(names, id)
	}
	fmt.Println(rows)
	return respuesta

}

//Guardando un nuevo producto en bodega general
func MdlNewProducto(Producto string, IdUser int) []StructDB.RespuestaInsertInGeneral {
	//SETEANDO LA DATA EN EL STRUCT
	respuesta := []StructDB.RespuestaInsertInGeneral{}
	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//Tomando la hora y fecha actual para la fecha de registro
	//instanciando el objeto
	var resp StructDB.RespuestaInsertInGeneral
	dt := time.Now()
	fmt.Println(dt)
	//	fmt.Println(dt)
	rows, err := Conecta.ConectionSQL().Query("EXEC spNewProducto ?, ?, ?", IdUser, Producto, dt)
	if err != nil {
		log.Fatal("Error de lectura de procedimiento almacenado")
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.RespSQL)
		if err != nil {
			log.Fatal("Error al gardar el producto")
		}
		respuesta = append(respuesta, resp)
		//	names = append(names, id)
	}
	fmt.Println(rows)
	return respuesta

}

//Guardando un nuevo producto en bodega general
func MdlIncidenciaDesGenerla(IdIng int, IdDetalle int, IdUser int, Descripcion string) []StructDB.RespuestaInsertInGeneral {
	//SETEANDO LA DATA EN EL STRUCT
	respuesta := []StructDB.RespuestaInsertInGeneral{}
	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//Tomando la hora y fecha actual para la fecha de registro
	//instanciando el objeto
	var resp StructDB.RespuestaInsertInGeneral
	dt := time.Now()
	fmt.Println(dt)
	//	fmt.Println(dt)
	rows, err := Conecta.ConectionSQL().Query("EXEC spNewDescRec ?, ?, ?, ?, ?", IdIng, IdDetalle, IdUser, Descripcion, dt)
	if err != nil {
		log.Fatal("Error de lectura de procedimiento almacenado " + err.Error())
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.RespSQL)
		if err != nil {
			log.Fatal("Error al gardar el producto")
		}
		respuesta = append(respuesta, resp)
		//	names = append(names, id)
	}
	fmt.Println(rows)
	return respuesta

}

//Guardando un nuevo producto en bodega general
func MdlNewMetrajeBodGeneral(IdIng int, IdDetalle int, IdAreaBod int, IdUserInt int, Metros float64, Posiciones int, PromedioTarima float64, MetrosStock float64, PosicionesStock int) []StructDB.RespuestaInsertInGeneral {
	//SETEANDO LA DATA EN EL STRUCT
	respuesta := []StructDB.RespuestaInsertInGeneral{}
	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//Tomando la hora y fecha actual para la fecha de registro
	//instanciando el objeto
	var resp StructDB.RespuestaInsertInGeneral
	dt := time.Now()
	fmt.Println(dt)
	//	fmt.Println(dt)
	rows, err := Conecta.ConectionSQL().Query("EXEC spNewMetrajeGeneral ?, ?, ?, ?, ?, ?, ?, ?, ?, ?", IdIng, IdDetalle, IdAreaBod, IdUserInt, Metros, Posiciones, PromedioTarima, MetrosStock, PosicionesStock, dt)
	if err != nil {
		log.Fatal("Error de lectura de procedimiento almacenado " + err.Error())
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.RespSQL)
		if err != nil {
			log.Fatal("Error al gardar el producto")
		}
		respuesta = append(respuesta, resp)
		//	names = append(names, id)
	}
	fmt.Println(rows)
	return respuesta

}

//Guardando una nueva ubicación en bodega
func MdlNewUbicacionBodegaGeneral(IdIng int, IdDetalle int, IdAreaBod int, IdUserInt int, Pasillo int, Columna int, Comentarios string) []StructDB.RespuestaInsertInGeneral {
	//SETEANDO LA DATA EN EL STRUCT
	respuesta := []StructDB.RespuestaInsertInGeneral{}
	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//Tomando la hora y fecha actual para la fecha de registro
	//instanciando el objeto
	var resp StructDB.RespuestaInsertInGeneral
	dt := time.Now()
	fmt.Println(dt)
	//	fmt.Println(dt)
	rows, err := Conecta.ConectionSQL().Query("EXEC spNewUbicacion ?, ?, ?, ?, ?, ?, ?, ?", IdIng, IdDetalle, IdAreaBod, IdUserInt, Pasillo, Columna, Comentarios, dt)
	if err != nil {
		log.Fatal("Error de lectura de procedimiento almacenado " + err.Error())
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.RespSQL)
		if err != nil {
			log.Fatal("Error al gardar el producto")
		}
		respuesta = append(respuesta, resp)
		//	names = append(names, id)
	}
	fmt.Println(rows)
	return respuesta

}

//Guardando la ruta de la certificación contable o imagen contable
func MdlNewDocSistema(IdIngGeneral int, IdUserInt int, spExecute string, Ruta string) []StructDB.RespuestaInsertInGeneral {
	//SETEANDO LA DATA EN EL STRUCT
	respSQL := []StructDB.RespuestaInsertInGeneral{}
	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//Tomando la hora y fecha actual para la fecha de registro
	//instanciando el objeto
	var resp StructDB.RespuestaInsertInGeneral
	dt := time.Now()
	//	fmt.Println(dt)
	rows, err := Conecta.ConectionSQL().Query("EXEC "+spExecute+" ?, ?, ?, ?", IdIngGeneral, IdUserInt, Ruta, dt)
	if err != nil {
		log.Fatal("Error de lectura de procedimiento almacenado" + err.Error())
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.RespSQL)
		if err != nil {
			log.Fatal("Error de lectura de procedimiento almacenado")
		}
		respSQL = append(respSQL, resp)
		//	names = append(names, id)
	}

	return respSQL
}

func MdlRutaImagen(IngING int) []StructDB.ImagenesRuta {
	//SETEANDO LA DATA EN EL STRUCT
	respSQLRuta := []StructDB.ImagenesRuta{}
	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//Tomando la hora y fecha actual para la fecha de registro
	//instanciando el objeto
	var resp StructDB.ImagenesRuta

	//	fmt.Println(dt)
	rows, err := Conecta.ConectionSQL().Query("EXEC spGetRutaImgDocDesc ?", IngING)
	if err != nil {
		log.Fatal("Error de lectura de procedimiento almacenado" + err.Error())
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.RutaImg)
		if err != nil {
			log.Fatal("Error de lectura de procedimiento almacenado")
		}
		respSQLRuta = append(respSQLRuta, resp)
		//	names = append(names, id)
	}

	return respSQLRuta
}

//Guardando en la base de datos el nuevo ingreso
func MdlUpdateIngresoGeneral(idIng int, IdUserInt int, bultos int, ValorTotal float64, Motivo string) []StructDB.RespuestaInsertInGeneral {
	//SETEANDO LA DATA EN EL STRUCT
	respuesta := []StructDB.RespuestaInsertInGeneral{}
	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//Tomando la hora y fecha actual para la fecha de registro
	//instanciando el objeto
	var resp StructDB.RespuestaInsertInGeneral
	//Hora de la transacción
	dt := time.Now()
	//Ejecutando el query
	rows, err := Conecta.ConectionSQL().Query("EXEC spUpdateIngGeneral ?, ?, ?, ?, ?, ?, ?", idIng, IdUserInt, bultos, ValorTotal, dt, dt, Motivo)
	if err != nil {
		log.Fatal("Error de lectura de procedimiento almacenado")
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.RespSQL)
		if err != nil {
			log.Fatal("Error de lectura de procedimiento almacenado")
		}
		respuesta = append(respuesta, resp)
		//	names = append(names, id)
	}
	fmt.Println(rows)
	return respuesta

}

//Guardando en la base de datos el nuevo ingreso
func MdlUpdateDetallesGeneral(idDetalle int, IdIngReq int, IdUserInt int, IdProduct int, Bultos int, PUnitario float64, Motivo string) []StructDB.RespuestaInsertInGeneral {
	//SETEANDO LA DATA EN EL STRUCT
	respuesta := []StructDB.RespuestaInsertInGeneral{}
	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//Tomando la hora y fecha actual para la fecha de registro
	//instanciando el objeto
	var resp StructDB.RespuestaInsertInGeneral
	//Hora de la transacción
	dt := time.Now()
	//Ejecutando el query
	rows, err := Conecta.ConectionSQL().Query("EXEC spUpdateDetGeneral ?, ?, ?, ?, ?, ?, ?, ?", idDetalle, IdIngReq, IdProduct, IdUserInt, Bultos, PUnitario, Motivo, dt)
	if err != nil {
		log.Fatal("Error de lectura de procedimiento almacenado")
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.RespSQL)
		if err != nil {
			log.Fatal("Error de lectura de procedimiento almacenado")
		}
		respuesta = append(respuesta, resp)
		//	names = append(names, id)
	}
	fmt.Println(rows)
	return respuesta

}

//Guardando en la base de datos el nuevo ingreso
func MdlDetallesGeneralRemove(idIncidencia int, IdUserInt int, Motivo string) []StructDB.RespuestaInsertInGeneral {
	//SETEANDO LA DATA EN EL STRUCT
	respuesta := []StructDB.RespuestaInsertInGeneral{}
	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//Tomando la hora y fecha actual para la fecha de registro
	//instanciando el objeto
	var resp StructDB.RespuestaInsertInGeneral
	//Hora de la transacción
	dt := time.Now()
	//Ejecutando el query
	rows, err := Conecta.ConectionSQL().Query("EXEC spCancelVisual ?, ?, ?, ? ", idIncidencia, IdUserInt, Motivo, dt)
	if err != nil {
		log.Fatal("Error de lectura de procedimiento almacenado")
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.RespSQL)
		if err != nil {
			log.Fatal("Error de lectura de procedimiento almacenado")
		}
		respuesta = append(respuesta, resp)
		//	names = append(names, id)
	}
	return respuesta

}

//Guardando en la base de datos el nuevo ingreso
func MdlIncDescGeneralUpdate(idIncidencia int, IdUserInt int, Descripcion string, Motivo string) []StructDB.RespuestaInsertInGeneral {
	//SETEANDO LA DATA EN EL STRUCT
	respuesta := []StructDB.RespuestaInsertInGeneral{}
	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//Tomando la hora y fecha actual para la fecha de registro
	//instanciando el objeto
	var resp StructDB.RespuestaInsertInGeneral
	//Hora de la transacción
	dt := time.Now()
	//Ejecutando el query
	rows, err := Conecta.ConectionSQL().Query("EXEC spEditProduc ?, ?, ?, ?, ?  ", idIncidencia, IdUserInt, Descripcion, Motivo, dt)
	if err != nil {
		log.Fatal("Error de lectura de procedimiento almacenado")
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.RespSQL)
		if err != nil {
			log.Fatal("Error de lectura de procedimiento almacenado")
		}
		respuesta = append(respuesta, resp)
		//	names = append(names, id)
	}
	fmt.Println(rows)
	return respuesta

}

//Guardando en la base de datos el nuevo ingreso
func MdlMetrajeBodegaGeneralUpdate(idMetraje int, IdAreaBodReq int, IdUserInt int, Metros float64, Posiciones float64, Promedio float64, Motivo string) []StructDB.RespuestaInsertInGeneral {
	//SETEANDO LA DATA EN EL STRUCT
	respuesta := []StructDB.RespuestaInsertInGeneral{}
	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//Tomando la hora y fecha actual para la fecha de registro
	//instanciando el objeto
	var resp StructDB.RespuestaInsertInGeneral
	//Hora de la transacción
	dt := time.Now()
	//Ejecutando el query
	rows, err := Conecta.ConectionSQL().Query("EXEC spModMetrajeGeneral ?, ?, ?, ?, ?, ?, ?, ?  ", idMetraje, IdAreaBodReq, IdUserInt, Metros, Posiciones, Promedio, Motivo, dt)
	if err != nil {
		log.Fatal("Error de lectura de procedimiento almacenado" + err.Error())
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.RespSQL)
		if err != nil {
			log.Fatal("Error de lectura de procedimiento almacenado")
		}
		respuesta = append(respuesta, resp)
		//	names = append(names, id)
	}
	fmt.Println(rows)
	return respuesta

}

//Guardando en la base de datos el nuevo ingreso
func MdlUbicacionesMercaGUpdate(idUbica int, IdAreaBodReq int, IdUserInt int, Pasillo int, Columna int, Comentario string, Motivo string) []StructDB.RespuestaInsertInGeneral {
	//SETEANDO LA DATA EN EL STRUCT
	respuesta := []StructDB.RespuestaInsertInGeneral{}
	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//Tomando la hora y fecha actual para la fecha de registro
	//instanciando el objeto
	var resp StructDB.RespuestaInsertInGeneral
	//Hora de la transacción
	dt := time.Now()
	//Ejecutando el query
	rows, err := Conecta.ConectionSQL().Query("EXEC spModificaUbicacion ?, ?, ?, ?, ?, ?, ?, ?  ", idUbica, IdAreaBodReq, IdUserInt, Pasillo, Columna, Comentario, Motivo, dt)
	if err != nil {
		log.Fatal("Error de lectura de procedimiento almacenado" + err.Error())
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.RespSQL)
		if err != nil {
			log.Fatal("Error de lectura de procedimiento almacenado")
		}
		respuesta = append(respuesta, resp)
		//	names = append(names, id)
	}
	fmt.Println(rows)
	return respuesta

}

//Guardando en la base de datos el nuevo ingreso
func MdlConsultaProducto(Producto string) []StructDB.ConsultaProducto {
	//SETEANDO LA DATA EN EL STRUCT
	respuesta := []StructDB.ConsultaProducto{}
	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//Tomando la hora y fecha actual para la fecha de registro
	//instanciando el objeto
	var resp StructDB.ConsultaProducto
	//Ejecutando el query
	rows, err := Conecta.ConectionSQL().Query("EXEC spConsultProducto ? ", Producto)
	if err != nil {
		log.Fatal("Error de lectura de procedimiento almacenado")
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.Id, &resp.Name)
		if err != nil {
			log.Fatal("Error de lectura de procedimiento almacenado")
		}
		respuesta = append(respuesta, resp)
		//	names = append(names, id)
	}
	return respuesta

}

//buscando todos los productos de bodega general
func MdlConsultaProductoAll() []StructDB.ConsultaProducto {
	//SETEANDO LA DATA EN EL STRUCT
	respuesta := []StructDB.ConsultaProducto{}
	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//Tomando la hora y fecha actual para la fecha de registro
	//instanciando el objeto
	var resp StructDB.ConsultaProducto
	//Ejecutando el query
	rows, err := Conecta.ConectionSQL().Query("EXEC spConsultProductoAll")
	if err != nil {
		log.Fatal("Error de lectura de procedimiento almacenado")
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.Id, &resp.Name)
		if err != nil {
			log.Fatal("Error de lectura de procedimiento almacenado")
		}
		respuesta = append(respuesta, resp)
		//	names = append(names, id)
	}
	return respuesta

}
