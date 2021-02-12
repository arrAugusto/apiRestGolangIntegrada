package modelUser

import (
	"encoding/json"
	"fmt"
	"net/http"

	Conecta "../databaseSQL"
	//StructModUser "../structures/structUser"
)

func ConsultaUsuarios(w http.ResponseWriter, r *http.Request) {

	//declarando string para guardar los datos de la consulta
	Conecta.ConectionSQL()
	defer Conecta.ConectionSQL().Close()
	//structModelUs := StructModUser.DataConsultaUSer{}
	//consulta query
	names := make([]int, 0)

	tsql := fmt.Sprintf("EXEC [spMostrarUsuarios]")
	rows, err := Conecta.ConectionSQL().Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {

		var id int
		err := rows.Scan(&id)
		fmt.Println(err)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return
		}
		names = append(names, id)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rows)

}
