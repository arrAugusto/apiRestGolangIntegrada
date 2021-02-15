package modelUser

import (
	"fmt"

	StructUser "../../structures/structuresUser"
	Conecta "../databaseSQL"
	//StructModUser "../structures/structUser"
)

func ConsultaUsuarios() ([]StructUser.Usuario, error) {

	usuarios := []StructUser.Usuario{}
	//declarando string para guardar los datos de la consulta
	Conecta.ConectionSQL()
	defer Conecta.ConectionSQL().Close()
	//structModelUs := StructModUser.DataConsultaUSer{}
	//consulta query
	//	names := make([]int, 0)
	tsql := fmt.Sprintf("EXEC [spMostrarUsuarios]")
	rows, err := Conecta.ConectionSQL().Query(tsql)
	var user StructUser.Usuario
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {

		//	var id int
		err := rows.Scan(&user.Id)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return nil, err
		}
		usuarios = append(usuarios, user)
		//	names = append(names, id)
	}
	return usuarios, nil

}
