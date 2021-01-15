package main

import (
	"fmt"
	"log"
	"net/http"
	//"math/operations"
)

type CepValidade struct {
	current string
}

func main() {

	// resultado = math.Soma(2, 3)
	// fmt.Println("%v", resultado);

	// fmt.Println("Hello GO");
	var current = 27321010

	fmt.Println(current)

	//defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)

	// if err != nil {
	// 	fmt.Println("Error")
	// }

	res, err := http.Get("https://viacep.com.br/ws/" + current + "/json/")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(res.Body)

}
