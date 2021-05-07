package structNit

type NitClienteResponse struct {
	IdNit             int    `json: nit`
	NumNit            string `json: numNit`
	NombreEmprea      string `json: nombreEmpresa`
	DireccionEmpresa  string `json: direccionEmpresa`
	ContactoEmpresa   string `json: contactoEmpresa`
	Telefonoempresa   string `json: telefonoEmpresa`
	CorreoEmpresa     string `json: correoEmpresa`
	NombreEjecutivo   string `json: nombreEjecutivo`
	TelefonoEjecutivo string `json: telefonoEjecutivo`
	CorreoEjecutivo   string `json: correoEjecutivo`
}
