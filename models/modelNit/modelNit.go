package modelNit

import (
	"log"

	StructDB "../../structures/structNit"
	Conecta "../databaseSQL"
)

func MdlConsultaNit(nitEmpresa string, storeProduce string) []StructDB.NitClienteResponse {
	//SETEANDO LA DATA EN EL STRUCT
	respuesta := []StructDB.NitClienteResponse{}
	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//Tomando la hora y fecha actual para la fecha de registro
	//instanciando el objeto
	var resp StructDB.NitClienteResponse
	rows, err := Conecta.ConectionSQL().Query("EXEC "+storeProduce+" ? ", nitEmpresa)
	if err != nil {
		log.Fatal("Error al guardar el ingreso general" + err.Error())
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&resp.IdNit, &resp.NumNit, &resp.NombreEmprea, &resp.DireccionEmpresa, &resp.ContactoEmpresa, &resp.Telefonoempresa, &resp.CorreoEmpresa, &resp.NombreEjecutivo, &resp.TelefonoEjecutivo)
		if err != nil {
			log.Fatal("Error en set objeto" + err.Error())
		}
		respuesta = append(respuesta, resp)
		//	names = append(names, id)
	}
	return respuesta

}
