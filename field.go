package monogo

import (
	"reflect"

	"github.com/cdvelop/monogo/inputs"
)

type field struct {
	Index      uint32 // index of the field
	Name       string // e.g.: id_user, name_user, phone
	Legend     string // e.g.: ID, Name, Phone
	Type       FieldType
	Unique     bool // unique and unalterable field in db
	NotNull    bool
	PrimaryKey bool    // primary key of the table
	ForeignKey *entity // foreign key of the table
	Input      input   //for representation in the form and validation
	Parent     *entity // pointer to the entity parent
}

func (f *field) isPrimaryKey() bool {
	_, isPrimary := fieldType(f.Parent.TableName, f.Name)
	return isPrimary
}

func setUnique(rf *reflect.StructField) bool {
	if unique := rf.Tag.Get("Unique"); unique != "" {
		return true
	}
	return false
}

func (f *field) setDataBaseParams() {

	// check if the field is a primary key
	f.PrimaryKey = f.isPrimaryKey()
	if f.PrimaryKey {
		f.Parent.IsTable = true
		f.Unique = true

		// f.Name = prefixNameID + f.Parent.TableName

		// check input is not set
		if f.Input == nil {
			f.Input = inputs.ID()
		}
	}
}
