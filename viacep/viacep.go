package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func checkCep(cep []string) map[string]bool {

	ceps := make(map[string]bool)

	for i := 0; i < len(cep); i++ {

		var current string = cep[i]

		resp, err := http.Get("https://viacep.com.br/ws/" + current + "/json/")

		if err != nil {
			fmt.Println("Error")
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Error")
		}

		var value string = string(body)

		var exists bool = strings.Contains(value, "erro")

		if exists == false {
			ceps[current] = true
		} else {
			ceps[current] = false
		}

	}

	return ceps

}

func main() {

	var ceps = []string{"26116700", "26116800", "26116801"}
	fmt.Println(checkCep(ceps))



	//var ceps = []string ioutil.ReadFile("./ceps.txt")
	//ioutil.ReadFile("./ceps.txt")
	//var ceps = ioutil.ReadFile("./ceps.txt")
	
}
