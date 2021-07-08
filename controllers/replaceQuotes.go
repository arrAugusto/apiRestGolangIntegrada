package controllerIngGeneral

func ReplaceString(str string) string {

	if str[0] == '"' {
		str = str[1:]
	}
	if i := len(str) - 1; str[i] == '"' {
		str = str[:i]
	}

	return str
}
