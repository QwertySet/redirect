package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

type users struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/", handler)
	app.Get("/home", handlerHome)
	app.Post("/", handlerPost)

	test()

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))

}
func handlerPost(c fiber.Ctx) error {
	var objUsers users
	err := json.Unmarshal(c.Body(), &objUsers)
	if err != nil {
		return err
	}
	return c.SendString(objUsers.Name)
	//return c.SendString("its work!" + string(c.Body()))
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
