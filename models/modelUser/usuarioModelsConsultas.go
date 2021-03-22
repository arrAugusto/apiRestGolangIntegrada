package modelUser

import (
	"fmt"

	StructUser "../../structures/structuresUser"
	Conecta "../databaseSQL"
)

/**
*	METODO UTILIZADO PARA LA CONSULTA DEL USUARIO DEL QUE SE CONECTARA
**/
func MdlConsultaUsuarios(UserIng int, PwIng string) ([]StructUser.Usuario, error) {
	//declaro objeto a retornar
	usuarios := []StructUser.Usuario{}
	//userJWT := []StructUser.JWTAutentication{}

	//instanciando la conexión
	Conecta.ConectionSQL()
	//cerrar la conexión al final de script
	defer Conecta.ConectionSQL().Close()
	//consulta query

	rows, err := Conecta.ConectionSQL().Query("EXEC spMostrarUsuario ?", UserIng)
	//instanciando el objeto
	var user StructUser.Usuario
	if err != nil {
		fmt.Println("No se encontro el usuario: " + err.Error())
		return nil, err
	}
	//Destruir los rows que se almacenan en memoria dinamica al final del script
	defer rows.Close()
	for rows.Next() {
		//Leyendo cada una de las rows
		err := rows.Scan(&user.IdUser, &user.Usuario, &user.Password, &user.Nombre, &user.Apellidos, &user.Fecha_creacion, &user.Celular, &user.Email, &user.Niveles, &user.Dependencia, &user.IdDeBodega, &user.Foto, &user.Estado, &user.Departamentos)
		if err != nil {
			fmt.Println("No se encontro el usuario: " + err.Error())
			return nil, err
		}
		usuarios = append(usuarios, user)
		//	names = append(names, id)
	}

	if len(usuarios) == 1 && PwIng == usuarios[0].Password {
		PwIng = ""
		usuarios[0].Password = ""
		if err != nil {
			fmt.Println("No se encontro el usuario: " + err.Error())
			return nil, err

		}
		return usuarios, nil
	}
	return nil, err
}
