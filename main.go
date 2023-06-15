package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/calculator", calculator)
	http.HandleFunc("/doCulc", doCulc)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl/home.html") // Parse template file.
	t.Execute(w, nil)
}

type dataForCulc struct {
	Answer        int
	Error         string
	IsAnswerExist bool
	Num1          int
	Num2          int
	Operation     string
}

func calculator(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl/calculator.html") // Parse template file.
	err := t.Execute(w, dataForCulc{Num1: 6, Num2: 10, Operation: "add"}) //Добавляем вместо NIL - dataForCulc (с дефолтными значениями)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
}

func doCulc(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl/calculator.html") // Parse template file.

	r.ParseForm()

	var err error

	data := dataForCulc{}

	if len(r.Form["number1"]) > 0 { //Проверяем, что массив не пустой
		data.Num1, err = strconv.Atoi(r.Form["number1"][0]) //Конвертируем string --> INT + присваиваем значение полю структуры
		if err != nil {
			fmt.Printf("error %s\n", err)
		}
	}
	if len(r.Form["number2"]) > 0 {
		data.Num2, err = strconv.Atoi(r.Form["number2"][0])
		if err != nil {
			fmt.Printf("error %s\n", err)
		}
	}
	if len(r.Form["operation"]) > 0 { //Проверяем, что массив не пустой
		data.Operation = r.Form["operation"][0]
	}
	fmt.Printf("num1 %v num2 %v %v\n", data.Num1, data.Num2, data.Operation)

	switch data.Operation {
	case "add":
		data.Answer = data.Num1 + data.Num2
		data.IsAnswerExist = true
	case "subtract":
		data.Answer = data.Num1 - data.Num2
		data.IsAnswerExist = true
	case "multiply":
		data.Answer = data.Num1 * data.Num2
		data.IsAnswerExist = true
	case "divide":
		if data.Num2 != 0 {
			data.Answer = data.Num1 / data.Num2
			data.IsAnswerExist = true
		} else {
			data.Error = "делить на ноль нельзя"
		}
	}

	err = t.Execute(w, data)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
}
