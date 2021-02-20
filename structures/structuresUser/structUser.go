package structuresUser

/**
*	ESTRUCTURA `Usuario` DE DATOS SE UTILIZA PARA TOMAR LOS VALORES DEL SP EL CUAL FORMARA PARTE DEL PAYLOAD DE JWT
*	¡Tomar nota que esta estructura sirve unicamente para leer la data de la consulta
**/
type Usuario struct {
	IdUser         int    `json: id`
	Usuario        int    `json: usuario`
	Password       string `json: password`
	Nombre         string `json: nombre`
	Apellidos      string `json: apellidos`
	Fecha_creacion string `json: fechaCreacion`
	Celular        string `json: telefono`
	Email          string `json: email`
	Niveles        string `json: niveles`
	Dependencia    int    `json: dependencia`
	IdDeBodega     int    `json: idBodega`
	Foto           string `json: foto`
	Estado         int    `json: estado`
	Departamentos  string `json: departamentos`
}
