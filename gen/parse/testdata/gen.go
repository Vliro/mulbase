package mulgen

// Code generated by mulgen. DO NOT EDIT (or feel free but it will be lost!).
import (
	"context"
	"mulbase"
)

//Populates the field of r.
func (r *Todo) GetTodoUser(count int, filter int) error {
	if r.UID() == "" {
		return mulbase.ErrUID
	}
	return mulbase.GetChild(r, GetField("User"), &r.User)
}

func (r *Todo) AddTodoUser(input *User) error {
	if input.UID() == "" {
		return mulbase.ErrUID
	}
	return nil
}

func (r *Todo) Fields() mulbase.FieldList {
	return TodoFields
}

type asyncTodoUser struct {
	Err   error
	Value User
}

func AsyncTodoUser(q mulbase.Query, input *User, txn *mulbase.Txn) error {
	ch := make(chan asyncTodoUser, 1)
	go func() {
		var obj User
		err := txn.RunQuery(context.Background(), q, &obj)
		ch <- asyncTodoUser{err, obj}
	}()
	return nil
}

func (r *User) Fields() mulbase.FieldList {
	return UserFields
}

//Populates the field of r.
func (r *Character) GetCharacterAppearsIn(count int, filter int) error {
	if r.UID() == "" {
		return mulbase.ErrUID
	}
	return mulbase.GetChild(r, GetField("AppearsIn"), &r.AppearsIn)
}

func (r *Character) AddCharacterAppearsIn(input *Episode) error {
	if input.UID() == "" {
		return mulbase.ErrUID
	}
	return nil
}

func (r *Character) Fields() mulbase.FieldList {
	return CharacterFields
}

type asyncCharacterAppearsIn struct {
	Err   error
	Value Episode
}

func AsyncCharacterAppearsIn(q mulbase.Query, input *Episode, txn *mulbase.Txn) error {
	ch := make(chan asyncCharacterAppearsIn, 1)
	go func() {
		var obj Episode
		err := txn.RunQuery(context.Background(), q, &obj)
		ch <- asyncCharacterAppearsIn{err, obj}
	}()
	return nil
}

func (r *Episode) Fields() mulbase.FieldList {
	return EpisodeFields
}

//Populates the field of r.
func (r *Query) GetQueryTodos(count int, filter int) error {
	if r.UID() == "" {
		return mulbase.ErrUID
	}
	return mulbase.GetChild(r, GetField("Todos"), &r.Todos)
}

func (r *Query) AddQueryTodos(input *Todo) error {
	if input.UID() == "" {
		return mulbase.ErrUID
	}
	return nil
}

func (r *Query) Fields() mulbase.FieldList {
	return QueryFields
}

type asyncQueryTodos struct {
	Err   error
	Value Todo
}

func AsyncQueryTodos(q mulbase.Query, input *Todo, txn *mulbase.Txn) error {
	ch := make(chan asyncQueryTodos, 1)
	go func() {
		var obj Todo
		err := txn.RunQuery(context.Background(), q, &obj)
		ch <- asyncQueryTodos{err, obj}
	}()
	return nil
}

var globalFields = make(map[string]mulbase.Field)

func GetField(name string) mulbase.Field {
	return globalFields[name]
}

func MakeField(name string) mulbase.Field {
	var field = mulbase.MakeField(name)
	if _, ok := globalFields[name]; !ok {
		globalFields[name] = field
	}
	return field
}
