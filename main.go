package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

type Song struct {
	Path string `json:"path"`
	Filename string `json:"filename"`
}

func main() {
	txtRegEx, e := regexp.Compile("^.+\\.(txt)$")

	var currentPath string
	var songs []Song

	e = filepath.Walk("./songs", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			currentPath = filepath.ToSlash(path)
		}
		if err == nil && txtRegEx.MatchString(info.Name()) {
			song := Song{currentPath, info.Name()}
			songs = append(songs, song)
		}
		return nil
	})
	if e != nil {
		log.Fatal(e)
	}

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Use(cors.New())

	app.Static("/songs", "songs", fiber.Static{
		ByteRange: true,
		MaxAge: 24 * 60 * 60,
		CacheDuration: 10 * time.Minute,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(songs)
	})

	e = app.Listen(":3000")
	if e != nil {
		log.Fatal(e)
	}
}
