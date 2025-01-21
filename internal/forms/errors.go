package forms

type errors map[string][]string

func (e errors) AddErr(field, message string) {
	e[field] = append(e[field], message)
}

func (e errors) Err(field string) string {
	errMessages := e[field]
	if len(errMessages) == 0 {
		return ""
	}

	return errMessages[0]
}
