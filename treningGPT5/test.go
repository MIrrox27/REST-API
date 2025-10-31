package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	if len(os.Args) < 2 { // проверяем что передан путь к файлу через аргументы командной строки, файл передается как первый аргумент после имени программы пример: go run jsonstats.go data.json
		fmt.Println("usage: jsonstats <path>") // выводим инструкцию по использованию
		os.Exit(1)                             //
	}
	path := os.Args[1]
	fmt.Println("Processing file:", path)
	// TODO: open file, decode JSON array into []Person, compute count and average age, print result
	// сам все сделаю не подсказывай
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	if file != nil {
		fileDecoder := json.NewDecoder(file)
		fmt.Println(fileDecoder)
	}

}
