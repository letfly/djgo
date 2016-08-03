package djgo

import (
	"fmt"
	"log"
	"net/http"
)

// Declare variables
var (
	DjApp    *App
	HttpAddr string
	HttpPort int
)

func init() {
	DjApp = NewApp()
	HttpAddr = ""
	HttpPort = 8000
}

type App struct {
	Handlers *ControllerRegistor
}

func NewApp() *App {
	cr := NewControllerRegistor()
	app := &App{Handlers: cr}
	return app
}

// Router interface
func (app *App) Router(path string, c ControllerInterface) *App {
	app.Handlers.Add(path, c)
	return app
}
func Router(path string, c ControllerInterface) *App {
	DjApp.Router(path, c)
	return DjApp
}

// Run interface
func (app *App) Run() {
	addr := fmt.Sprintf("%s:%d", HttpAddr, HttpPort)
	err := http.ListenAndServe(addr, app.Handlers)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func Run() {
	DjApp.Run()
}
