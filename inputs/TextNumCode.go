package inputs

import "errors"

func TextNumberCode(params ...any) textNumCode {
	new := textNumCode{
		attributes: attributes{
			htmlName:   "tel",
			customName: "textNumCode",
			// Pattern: `^[A-Za-z0-9-_]{2,15}$`,
		},
		permitted: permitted{
			Letters:    true,
			Numbers:    true,
			Characters: []rune{'_', '-'},
			Minimum:    2,
			Maximum:    15,
		},
	}
	new.Set(&new.permitted, params)

	return new
}

// texto y numero para código ej: V234
type textNumCode struct {
	attributes
	permitted
}

func (t textNumCode) ValidateInput(value string) error {

	if len(value) >= 1 {
		var ok bool
		char := value[0]

		if valid_letters[rune(char)] {
			ok = true
		}

		if valid_number[rune(char)] {
			ok = true
		}

		if !ok {
			return errors.New("no se puede comenzar con " + string(char))
		}
	}

	return t.Validate(value)
}

func (t textNumCode) GoodTestData() (out []string) {

	out = []string{
		"et1",
		"12f",
		"GH03",
		"JJ10",
		"Wednesday",
		"los567",
		"677GH",
		"son_24_botellas",
	}

	return
}

func (t textNumCode) WrongTestData() (out []string) {

	out = []string{
		"los cuatro",
		"son 2 cuadros",
	}
	out = append(out, wrong_data...)

	return
}
