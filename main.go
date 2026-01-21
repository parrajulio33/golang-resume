package main

import (
	"log"
	"net/http"

	"os"

	"resume-app/repositories"
	"resume-app/supabase"
	"resume-app/templates"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var port string

func init() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found in the current directory, reading configuration from environment variables")
	}

	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_ANON_KEY")
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if url == "" || key == "" {
		log.Fatal("SUPABASE_URL and SUPABASE_ANON_KEY environment variables must be set")
	}

}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.CORS())

	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		client, err := supabase.NewClient()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		resume, err := repositories.GetResume(client)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
		return templates.Resume(resume).Render(
			c.Request().Context(),
			c.Response(),
		)
	})

	e.Logger.Fatal(e.Start(":" + port))

}
