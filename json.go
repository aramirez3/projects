package main

import (
	"fmt"

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
