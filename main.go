package main

import (
	"embed"
	"os"
	"todo/internal/repository"
	"todo/pkg/datastore/sqlite"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	db, err := sqlite.Open()
	if err != nil {
		os.Exit(1)
	}
	if err := sqlite.Init(db); err != nil {
		os.Exit(1)
	}
	repo := repository.New(db)
	// Create an instance of the app structure
	app := NewApp(repo)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "MyToDo",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
