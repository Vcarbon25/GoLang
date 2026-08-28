package main

import (
	"fmt"

	"log"
	"os"
)

type Dados struct {
	Nm    string
	Info1 float32
	Info2 string
}

func CreateJSON() {
	Obj := Dados{"Test", 2026, "Created Json File"}
	fmt.Println(Obj.Info2)
	json, Wer1 := json.Marshal(Obj)
	if Wer1 != nil {
		log.Fatal(Wer1)
	} else {
		Wrfile, Wer2 := os.Create("Data.json")
		if Wer2 != nil {
			log.Fatal(Wer2)
		} else {
			defer Wrfile.Close()

			fmt.Fprintf(Wrfile, string(json))
			fmt.Println("File Created")
		}

	}

}

func ReadJSON() {
	RdFile, Rder1 := os.ReadFile("Data.json")
	if Rder1 != nil {
		log.Fatal(Rder1)
	} else {

		fmt.Println("Full File: %s", string(RdFile))
	}
	var Rcv Dados
	Rder2 := json.Unmarshal(RdFile, &Rcv)
	if Rder2 != nil {
		log.Fatal(Rder2)
	}
	fmt.Println("Acessing member: %s", Rcv.Info2)
}

func main() {
	CreateJSON()

	ReadJSON()
}
