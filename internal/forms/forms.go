package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

func (f *Form) IsValid() bool {
	return len(f.Errors) == 0
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.AddErr(field, "Это поле не может быть пустым")
		}
	}
}

func (f *Form) CheckLength(field string, length int) bool {
	if len([]rune(f.Get(field))) < length {
		f.Errors.AddErr(
			field,
			fmt.Sprintf("Это поле должно быть длиной не менее %d символов", length),
		)
		return false
	}
	return true
}

func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.AddErr(field, "Неверный адрес эл.почты")
	}
}
