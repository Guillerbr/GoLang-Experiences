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
	res, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal(err.Error());
	}

    fmt.Println(res.header);
	
	}


