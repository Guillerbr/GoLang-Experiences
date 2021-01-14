package main

import (
	"fmt"
	"net/http"
	//"math/operations"
)

func main() {

	// resultado = math.Soma(2, 3)
	// fmt.Println("%v", resultado);

	// fmt.Println("Hello GO");
	var current string = 27321010

	fmt.Println(current);


	defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Error")
		}

	res, err := http.Get("https://viacep.com.br/ws/" + current + "/json/")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(res.header)

}
