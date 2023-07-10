package forms

type errors map[string][]string

// Add adds a error message for a given field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns first error message for field
func (e errors) Get(field string) string {
	error_stings := e[field]

	if len(error_stings) == 0 {
		return ""
	}

	return error_stings[0]
}
