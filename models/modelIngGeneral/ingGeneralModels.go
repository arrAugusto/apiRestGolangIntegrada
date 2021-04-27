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
		log.Fatal("Error al guardar el ingreso general")
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.RespSQL)
		if err != nil {
			log.Fatal("Error al guardar el ingreso general")
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
		log.Fatal("Error al guardar el ingreso general")
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
		log.Fatal("Error al guardar el ingreso general")
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
		log.Fatal("Error al guardar el ingreso general")
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
		log.Fatal("Error al guardar el ingreso general")
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
		log.Fatal("Error al guardar el ingreso general " + err.Error())
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
		log.Fatal("Error al guardar el ingreso general " + err.Error())
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
		log.Fatal("Error al guardar el ingreso general " + err.Error())
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
		log.Fatal("Error al guardar el ingreso general" + err.Error())
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.RespSQL)
		if err != nil {
			log.Fatal("Error al guardar el ingreso general")
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
		log.Fatal("Error al guardar el ingreso general" + err.Error())
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.RutaImg)
		if err != nil {
			log.Fatal("Error al guardar el ingreso general")
		}
		respSQLRuta = append(respSQLRuta, resp)
		//	names = append(names, id)
	}

	return respSQLRuta
}
