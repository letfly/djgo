package djgo

import (
	"fmt"
	"net/http"
)

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
}

func NewApp() *App {
	app := &App{}
	return app
}

func (app *App) Run() {
	addr := fmt.Sprintf("%s:%d", HttpAddr, HttpPort)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("err=%s", err)
	}
}
