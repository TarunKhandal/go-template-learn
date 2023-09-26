package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	log := log.Default()
	templVar := template.New("simple")
	templates, err := templVar.ParseFiles("E:/projects/src/go-template-learn/WhatIsTemplate/simple.tmpl")
	if err != nil {

		log.Println("logged err" + err.Error())
	}

	data := map[string]interface{}{
		"hi": "Hello",
		"greet": map[string]interface{}{
			"gophers": "Gophers, Welcome!!",
		},
	}

	if err := templates.Execute(os.Stdout, data); err != nil {
		log.Println("logged err" + err.Error())
	}
}
