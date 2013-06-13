package example

import (
	"fmt"
)

// Person represents an object you can query for its name.
type Person interface {
	Name() string
}

// User represents a person with an ID number.
type User struct {
	id   int
	name string
}

func NewUser(id int, name string) *User {
	return &User{id, name}
}

func (u User) Name() string {
	return u.name
}

// Names demonstrates structs and interfaces.
func Names() {
	var someone Person
	someoneInParticular := NewUser(5, "John Jacobs")
	someone = someoneInParticular
	fmt.Println(someone.Name())
}
