package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/", handler)
	app.Get("/home", handlerHome)

	test()

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))

}

func handler(c fiber.Ctx) error {
	// Send a string response to the client
	//return c.SendString("Hello, World " + c.Query("q"))

	return c.Redirect().To(c.Query("q"))
}

func handlerHome(c fiber.Ctx) error {
	// Send a string response to the client
	return c.SendString("Hello, you are home!")
}

type someStruct struct {
	Key  string  `json:key`
	Key2 float64 `json:key2`
}

func test() {
	var s string
	s = `{"key":"hello","key2":5.4}`

	obj := someStruct{}
	json.Unmarshal([]byte(s), &obj)
	fmt.Println(obj)
}
