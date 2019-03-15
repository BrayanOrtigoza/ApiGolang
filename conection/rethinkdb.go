package conection

import (
	"API/models"
	"fmt"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
	"time"
)

const (
	_ = iota
	STATUS_CREATE
	STATUS_ON_HOLD
	STATUS_ASSIGNED
	STATUS_
	STATUS_FINALIZED
	STATUS_CANCELLED
)

var (
	pluckPeople = [20]interface{}{
		"id",
		"name",
		"LastName",
		"Age",
	}
)


func NewPeople (m models.People) {

	item := NewTodoItem(m)

     fmt.Println(item)
	 _, err := r.Table("people").Insert(item).Run(session)

	if err != nil {
	   fmt.Println(err)
	}
}


func NewTodoItem(m models.People) *models.People {
	return &models.People{
		Name:   name,
		LastName: "active",
		Age:    age,
	}
}

