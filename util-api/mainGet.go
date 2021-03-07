package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	apiUtil "github.com/henryinfanteg/heyadivinen_backgo/util-api/util"
)

type Persona struct {
	UsuarioID string `bson:"usuarioId" json:"usuarioId"`
	Nombres   string `bson:"nombres" json:"nombres,omitempty"`
	Apellidos string `bson:"apellidos" json:"apellidos,omitempty"`
}

var wg sync.WaitGroup

func main() {

	interacciones := 1
	wg.Add(interacciones)

	fmt.Println("Running for loop…")
	fmt.Println("")

	// inicializamos los headers
	headers := apiUtil.GenerateHeaders("123", "golang", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2wiOjEsInBlcm1pc29zIjp7IkFERCI6dHJ1ZSwiREVMRVRFIjp0cnVlLCJSRUFEIjp0cnVlLCJVUERBVEUiOnRydWV9LCJhdWQiOiJBTEwiLCJpYXQiOjE1MzU5NTA2ODYsImlzcyI6IlRFU1QiLCJzdWIiOiJhZG1pbiJ9.Ewt7rbpqeUFwH1pDItgFeFNyOEFDJAc0iWEt0mFuhXM")
	endpoint := "http://localhost:3002/api/personas/5c4d5a7d9fcfda000190d2a5"

	for i := 1; i <= interacciones; i++ {
		// go servicio(i)
		go func(i int) {
			defer wg.Done()
			// servicio(i, headers)
			var obj Persona
			statusCode, err := apiUtil.GetResponse(apiUtil.Get, endpoint, 10000, headers, nil, &obj)
			fmt.Println("statusCode", statusCode)
			if err != nil {
				fmt.Println("errrrr", err)
			}
			fmt.Println(fmt.Sprintf("- %d%s%s", i, ")", obj))
		}(i)
	}

	wg.Wait()
	fmt.Println("")
	fmt.Println("Finished for loop")

}

func servicio(i int, headers http.Header) {
	// defer wg.Done()
	fmt.Println("Entro", i)

	url := "http://localhost:3001/api/seguridad/usuarios/5c4d5a7d6f1afa00014f4ae9"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header = headers
	// req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2wiOjEsInBlcm1pc29zIjp7IkFERCI6dHJ1ZSwiREVMRVRFIjp0cnVlLCJSRUFEIjp0cnVlLCJVUERBVEUiOnRydWV9LCJhdWQiOiJBTEwiLCJpYXQiOjE1MzU5NTA2ODYsImlzcyI6IlRFU1QiLCJzdWIiOiJhZG1pbiJ9.Ewt7rbpqeUFwH1pDItgFeFNyOEFDJAc0iWEt0mFuhXM")
	// req.Header.Add("Request-Id", fmt.Sprintf("%s%d", "123-", i))
	// req.Header.Add("App-Id", "postman")
	// req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	} else {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		fmt.Println(fmt.Sprintf("- %d%s%s", i, ")", string(body)))
		// fmt.Println(res)
		// fmt.Println(string(body))
	}
	fmt.Println("-------")
	// }

}
