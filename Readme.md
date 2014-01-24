# gq

GQ is a Golang Query package for Postgres. It is inspired by Ruby's Sequel.

**This project is still under develoment and breaking changes may occur**

### Usage

Using gq directly

```go
package main

import (
	"github.com/daneharrigan/gq"
	"fmt"
)

type Person struct {
	Id string
	FirstName string
	Age int
}

func main() {
	// connect to the database
	gq.Connect("postgres://localhost/example")

	// build the statement
	statement := gq.From("people")
	statement = statement.Select("id", "first_name", "age")
	statement = statement.Where(gq.Equal("id", id))

	// what is it doing?
	fmt.Println(statement.SQL())
	// => SELECT id, first_name, age FROM people WHERE (id = $1)

	// get the first result
	result := statement.First()

	person := new(Person)
	err := result.Apply(&person.Id, &person.FirstName, &person.Age)

	fmt.Printf("fn=Apply error=%q\n", err)
	fmt.Printf("FirstName=%q\n", person.FirstName)
}
```

Using gq in a `Find` method

```go
package main

import (
	"github.com/daneharrigan/gq"
	"fmt"
)

type Person struct {
	Id string
	FirstName string
	Age int
}

func (p *Person) Find(id string) error {
	statement := gq.From("people")
	statement = statement.Select("id", "first_name", "age")
	statement = statement.Where(gq.Equal("id", id))

	// what is it doing?
	fmt.Println(statement.SQL())
	// => SELECT id, first_name, age FROM people WHERE (id = $1)

	result := statement.First() // talk to the database
	return result.Apply(&p.Id, &p.FirstName, &p.Age)
}

func init() {
	gq.Connect("postgres://localhost/example")
}

func main() {
	person := new(Person)
	person.Find("unique-identifier-here")
	fmt.Println(person.FirstName)
}
```
