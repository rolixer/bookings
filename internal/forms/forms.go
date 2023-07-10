package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form creates custom form struct
type Form struct {
	url.Values
	Errors errors
}

// New initializes a Form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string) bool {
	x := f.Get(field)

	return x != ""

}

// Required validates that fields are not empty
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannnot be empty")
		}
	}
}

// MinLength validates that field is at least `length` characters long
func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)

	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}

	return true
}

// Valid checks if form is valid
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// IsEmail checks for valid email
func (f *Form) IsEmail(field string) bool {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email")
		return false
	}
	return true
}
