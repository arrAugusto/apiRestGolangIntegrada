package modelIngGeneral

import (
	"fmt"
	"log"
	"time"

	StructDB "../../structures/structuresIngGeneral"
	Conecta "../databaseSQL"
)

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
