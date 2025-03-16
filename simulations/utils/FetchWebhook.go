package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchAPI(method string, duration int) {
	baseURL := "http://localhost:8080/events/" + method

	if duration != 0 && method == "duration" {
		baseURL = fmt.Sprintf("%s/%d", baseURL, duration)
	}

	req, err := http.NewRequest("POST", baseURL, nil)
	if err != nil {
		fmt.Println("Error al crear la solicitud:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al hacer la petición:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return
	}

	fmt.Println("Respuesta de la API:", string(body))
}