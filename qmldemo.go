package main

import (
	"fmt"
	"gopkg.in/qml.v1"
	"os"
)

type Login struct {
	Id       string
	Password string
	IsValid  bool

	Engine     *qml.Engine
	Ctx        *qml.Context
	Window     *qml.Window
	WindowRoot qml.Object
}

func (self *Login) HandleOk() {
	self.Id = self.WindowRoot.ObjectByName("idField").String("text")
	self.Password = self.WindowRoot.ObjectByName("passwordField").String("text")

	fmt.Printf("Id: %s | Password: %s\n", self.Id, self.Password)

	self.Window.Destroy()
}

func (self *Login) HandleCancel() {
	fmt.Println("Login cancelled.")

	self.Window.Destroy()
}

func main() {
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	login := Login{}
	engine := qml.NewEngine()

	ctx := engine.Context()
	ctx.SetVar("loginObj", &login)

	// Use main.qml file.
	// app, err := engine.LoadFile("main.qml")

	// Use included qml file.
	app, err := engine.LoadFile("qrc:///ui/ui.qml")
	if err != nil {
		return err
	}

	window := app.CreateWindow(nil)

	login.IsValid = false
	login.Engine = engine
	login.Ctx = ctx
	login.Window = window
	login.WindowRoot = window.Root()

	window.Show()
	window.Wait()

	return nil
}
