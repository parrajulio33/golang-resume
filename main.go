package main

import (
	"fmt"
	"log"
	"net/http"

	"os"

	"resume-app/models"
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

	// e.Use(middleware.Logger())
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Pre(middleware.AddTrailingSlash())
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

		// print the resume object to the console for debugging
		fmt.Printf("StartDate: '%s'\n", resume.Experience[0].Start)

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
		return templates.Resume(resume).Render(
			c.Request().Context(),
			c.Response(),
		)
	})

	// Admin routes
	adminGroup := e.Group("/admin")

	// adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	return username == os.Getenv("ADMIN_USERNAME") && password == os.Getenv("ADMIN_PASSWORD"), nil
	// }))

	adminGroup.GET("/", func(c echo.Context) error {
		client, err := supabase.NewClient()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		user, err := repositories.Login(client, os.Getenv("ADMIN_USERNAME"), os.Getenv("ADMIN_PASSWORD"))

		// if user == (models.User{}) {
		// 	return c.String(http.StatusUnauthorized, "Invalid admin credentials")
		// }
		if user == (models.User{}) {
			return c.String(http.StatusUnauthorized, "Invalid admin credentials")
		}

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.String(http.StatusOK, "Admin Dashboard")
	})

	adminGroup.GET("/edit/", func(c echo.Context) error {
		// client, err := supabase.NewClient()
		// if err != nil {
		// 	return c.String(http.StatusInternalServerError, err.Error())
		// }

		// resume, err := repositories.GetResume(client)
		// if err != nil {
		// 	return c.String(http.StatusInternalServerError, err.Error())
		// }

		// c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
		// return templates.EditResume(resume).Render(
		// 	c.Request().Context(),
		// 	c.Response(),
		// )
		return c.String(http.StatusOK, "Edit Resume Page - Not Implemented")
	})

	// adminGroup.POST("/edit", func(c echo.Context) error {
	// 	// client, err := supabase.NewClient()
	// 	// if err != nil {
	// 	// 	return c.String(http.StatusInternalServerError, err.Error())
	// 	// }

	// 	// resume, err := repositories.GetResume(client)
	// 	// if err != nil {
	// 	// 	return c.String(http.StatusInternalServerError, err.Error())
	// 	// }

	// 	// if err := c.Bind(&resume); err != nil {
	// 	// 	return c.String(http.StatusBadRequest, err.Error())
	// 	// }

	// 	// if err := repositories.UpdateResume(client, resume); err != nil {
	// 	// 	return c.String(http.StatusInternalServerError, err.Error())
	// 	// }

	// 	// return c.Redirect(http.StatusSeeOther, "/admin/edit")
	// })

	for _, r := range e.Routes() {
		fmt.Println(r.Method, r.Path)
	}

	log.Printf("Starting server on port %s", port)

	e.Logger.Fatal(e.Start(":" + port))

}
