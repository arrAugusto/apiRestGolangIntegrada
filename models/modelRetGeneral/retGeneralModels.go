package modelIngGeneral

import (
	"fmt"
	"log"
	"time"

	StructDB "../../structures/structuresIngGeneral"
	Conecta "../databaseSQL"
)

//Guardando en la base de datos el nuevo ingreso
func MdlNewRetiroGeneral(TotalBultos int, IdUserInt int) []StructDB.RespuestaInsertInGeneral {
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

	rows, err := Conecta.ConectionSQL().Query("EXEC spNewRetGeneral   ?, ?, ?, ? ", TotalBultos, IdUserInt, dt, dt)
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
		respuesta = append(respuesta, resp)
		//	names = append(names, id)
	}
	fmt.Println(rows)
	return respuesta

}

//Guardando en la base de datos el nuevo ingreso
func MdlNewDetRetGeneral(IdRet int, IdDetalle int, IdUserInt int, TotalBultos int) []StructDB.RespuestaInsertInGeneral {
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
	fmt.Println(IdRet)
	fmt.Println(IdDetalle)
	fmt.Println(IdUserInt)
	fmt.Println(TotalBultos)
	fmt.Println(dt)

	rows, err := Conecta.ConectionSQL().Query("EXEC spNewRebajaDetGen  ?, ?, ?, ?, ? ", IdRet, IdDetalle, IdUserInt, TotalBultos, dt)
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
		respuesta = append(respuesta, resp)
		//	names = append(names, id)
	}
	fmt.Println(rows)
	return respuesta

}

//Guardando la ruta de la certificación contable o imagen contable
func MdlMGRetiroGeneral(idRetGeneral int, IdUserInt int, spExecute string, Ruta string) []StructDB.RespuestaInsertInGeneral {
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
	rows, err := Conecta.ConectionSQL().Query("EXEC "+spExecute+" ?, ?, ?, ?", idRetGeneral, IdUserInt, Ruta, dt)
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

//Guardando un nuevo producto en bodega general
//Guardando en la base de datos el nuevo ingreso
func MdlAnulaTransaccionesRet(idOperacion int, idUsuario int, motivo string, storeProduce string) []StructDB.RespuestaInsertInGeneral {
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
