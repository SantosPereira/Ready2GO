package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	arquivo, err := ioutil.ReadFile("texto.txt")

	if err != nil {
		fmt.Println(err)
	}

	texto := string(arquivo)

	fmt.Print(texto)
}
