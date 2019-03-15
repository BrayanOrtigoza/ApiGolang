package controler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func NewPeoplePost(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//var people models.People

	//error := json.NewDecoder(r.Body)
	time.Sleep(20 * time.Second)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))

	/*if error != nil {
		panic(error)
	}*/
/*
    conection.NewPeople(people)
	w.Header().Set("Content-Type", "application/json")
	j, error := json.Marshal(note)

	if error != nil{
		panic(error)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(j)*/

}


