package mulgen

// Code generated by mulgen. DO NOT EDIT (or feel free but it will be lost!).
import (
	"mulbase"
)

//Populates the field of r.
func (r *Level) GetLevelOwner(filter int, txn *mulbase.Txn) error {
	if r.UID() == "" {
		return mulbase.ErrUID
	}
	return mulbase.GetChild(r, "Level.owner", UserFields, -1, txn, &r.Owner)
}

func (r *Level) AddLevelOwner(input *User) error {
	if input.UID() == "" {
		return mulbase.ErrUID
	}
	return nil
}

func (r *Level) Fields() mulbase.FieldList {
	return LevelFields
}

type asyncLevelOwner struct {
	Err   error
	Value *User
}

//Populates the field of r.
func (r *Level) GetLevelOwnerAsync(filter int, txn *mulbase.Txn) chan asyncLevelOwner {
	if r.UID() == "" {
		return nil
	}
	var ch = make(chan asyncLevelOwner, 1)
	go func() {
		var output = new(User)
		var err = mulbase.GetChild(r, "Level.owner", UserFields, -1, txn, output)
		var result asyncLevelOwner
		result.Err = err
		result.Value = output
		ch <- result
	}()
	return ch
}

/*
   func (r *Level) AddLevelOwner(input *User) error {
       if input.UID() == "" {
           return mulbase.ErrUID
       }
       return nil
   }*/

//Populates the field of r.
func (r *User) GetUserLevels(count int, filter int, txn *mulbase.Txn) error {
	if r.UID() == "" {
		return mulbase.ErrUID
	}
	return mulbase.GetChild(r, "User.levels", LevelFields, count, txn, &r.Levels)
}

func (r *User) AddUserLevels(input *Level) error {
	if input.UID() == "" {
		return mulbase.ErrUID
	}
	return nil
}

func (r *User) Fields() mulbase.FieldList {
	return UserFields
}

type asyncUserLevels struct {
	Err   error
	Value *Level
}

//Populates the field of r.
func (r *User) GetUserLevelsAsync(count int, filter int, txn *mulbase.Txn) chan asyncUserLevels {
	if r.UID() == "" {
		return nil
	}
	var ch = make(chan asyncUserLevels, 1)
	go func() {
		var output = new(Level)
		var err = mulbase.GetChild(r, "User.levels", LevelFields, count, txn, output)
		var result asyncUserLevels
		result.Err = err
		result.Value = output
		ch <- result
	}()
	return ch
}

/*
   func (r *User) AddUserLevels(input *Level) error {
       if input.UID() == "" {
           return mulbase.ErrUID
       }
       return nil
   }*/

//Beginning of field.template. General functions.
var globalFields = make(map[string]mulbase.Field)

func GetField(name string) mulbase.Field {
	return globalFields[name]
}

func MakeField(name string, flags mulbase.FieldMeta) mulbase.Field {
	var field = mulbase.Field{Name: name, Meta: flags}
	if _, ok := globalFields[name]; !ok {
		globalFields[name] = field
	}
	return field
}

func GetGlobalFields() map[string]mulbase.Field {
	return globalFields
}
