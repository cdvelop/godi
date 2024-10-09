package inputs

import "fmt"

// ej: options=m:male,f:female
func Radio(params ...any) radio {
	new := radio{
		attributes: attributes{
			htmlName: "radio",
			Onchange: `onchange="RadioChange(this);"`,
		},
	}
	new.Set(params)

	return new
}

// ej: {"f": "Femenino", "m": "Masculino"}.
func RadioGender(params ...any) radio {
	options := append(params, "name=genre", `options=f:Femenino,m:Masculino`)
	return Radio(options...)
}

type radio struct {
	attributes
	dataSource
}

// validación con datos de entrada
func (r radio) ValidateInput(value string) error {
	return r.checkOptionKeys(value)
}

func (r radio) BuildHtmlInput(id string) string {
	var id3 string

	var tags string

	for i, opt := range r.options {

		for value, span := range opt {
			id3 = fmt.Sprintf("%v.%v", id, i)

			tags += `<label for="` + id3 + `" class="block-label">`

			r.Value = `value="` + value + `"`

			tags += r.buildHtml(id3)

			tags += `<span>` + span + `</span>`
			tags += `</label>`
		}
	}

	return tags
}
