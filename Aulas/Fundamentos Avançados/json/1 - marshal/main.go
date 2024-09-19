package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
}

func main() {
	c := cachorro{"Pituxo", "Pinscher"}
	cachorroEmJSON, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(cachorroEmJSON))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}
	cachorro2EmJSON, err := json.Marshal(c2)
	fmt.Println(string(cachorro2EmJSON))

}
