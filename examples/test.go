package main

import (
	"fmt"

	"github.com/zekroTJA/configoration"
)

func main() {
	c, err := configoration.NewBuilder().
		SetBasePath("./data").
		AddJsonFile("test2.json", true).
		AddJsonFile("test1.json", true).
		AddYamlFile("test3.yaml", true).
		Build()

	if err != nil {
		panic(err)
	}

	fmt.Println(c.GetValue("y:e"))
}
