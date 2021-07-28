// TO DO: Implementer a classe api e a lib json, ao invés de slice de strings
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type api struct {
	copyright       string `json:"copyright"`
	date            string `json:"date"`
	explanation     string `json:"explanation"`
	hdurl           string `json:"hdurl"`
	media_type      string `json:"media_type"`
	service_version string `json:"service_version"`
	title           string `json:"title"`
	url             string `json:"url"`
}

func main() {
	resposta, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY")
	if err != nil {
		fmt.Println(err)
	}
	corpo, err := ioutil.ReadAll(resposta.Body)

	conteudo := api{}

	err = json.Unmarshal(corpo, &conteudo)

	corte1 := string(corpo)[strings.Index(string(corpo), `"url":`)+7:]

	corte2 := string(corpo)[strings.Index(string(corpo), `"url":`)+7 : strings.Index(string(corpo), `"url":`)+7+len(corte1)-3]

	fmt.Printf("\nURL: %v\n", corte2)

}
