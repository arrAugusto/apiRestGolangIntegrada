package structNit

type NitClienteResponse struct {
	Contacto         string `json: contacto`
	DireccionEmpresa string `json: direccionEmpresa`
	Ejecutivo        string `json: ejecutivo`
	EmailConta       string `json: emailConta`
	EmailEje         string `json: emailEje`
	Id               string `json: id`
	NitEmpresa       string `json: nitEmpresa`
	NombreEmpresa    string `json: nombreEmpresa`
	TelEjecutivo     string `json: telEjecutivo`
	TelEmpresa       string `json: telEmpresa`
}
