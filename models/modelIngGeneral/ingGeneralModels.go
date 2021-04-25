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

//Guardando la ruta de la certificación contable o imagen contable
func MdlNewDocSistema(Ruta string, IngIMG int) []StructDB.RespuestaInsertInGeneral {
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
	rows, err := Conecta.ConectionSQL().Query("EXEC spNewDocDescarga ?, ?, ?", IngIMG, Ruta, dt)
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
