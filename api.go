package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resposta, err := http.Get("https://swapi.dev/api/people/2")
	if err != nil {
		fmt.Println(err)
	}

	resposta, err = ioutil.ReadAll(resposta.Body)
	fmt.Println(resposta)
}
