package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func StrictUnmarshal(body interface{}, c echo.Context) error {
	err := c.Bind(body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// if !errors.Is(err, io.EOF) {
	// 	return errors.New("error binding request")
	// }

	fmt.Printf("payload bind success: %v\n", body)
	return nil
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	responseData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling json: %s\n", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	if _, err := w.Write(responseData); err != nil {
		log.Printf("error sending json: %s\n", err)
	}
}
