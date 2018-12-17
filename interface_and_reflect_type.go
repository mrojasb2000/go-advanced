package main

import (
	"log"
	"reflect"
)

type Gift struct {
	Sender    string
	Recipient string
	Number    uint
	Contents  string
}

func main() {
	g := Gift{
		Sender:    "Hank",
		Recipient: "Sue",
		Number:    1,
		Contents:  "Scarf",
	}

	t := reflect.TypeOf(g)

	if kind := t.Kind(); kind != reflect.Struct {
		log.Fatal("this program expects to work on a struct; we got a %v instead.", kind)
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		log.Printf("Field %03d: %-10.10s %v", i, f.Name, f.Type.Kind())
	}
}
