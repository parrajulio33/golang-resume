package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"resume-app/custom_middleware"
	"resume-app/repositories"
	"resume-app/session"
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
	session.Init(os.Getenv("SESSION_SECRET"))
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
	// e.Pre(middleware.RemoveTrailingSlash())
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

	adminGroup.GET("/", func(c echo.Context) error {
		if session.GetUserID(c) != "" {
			return c.Redirect(http.StatusFound, "/admin/dashboard")
		}
		c.Response().Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		c.Response().Header().Set("Pragma", "no-cache")
		c.Response().Header().Set("Expires", "0")

		return templates.Login("").Render(c.Request().Context(), c.Response())
	})

	adminGroup.POST("/login/", func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		client, err := supabase.NewClient()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		user, err := repositories.Login(client, email, password)
		if err != nil || user.ID.String() == "" {
			return templates.Login("Invalid admin credentials").Render(c.Request().Context(), c.Response())
		}

		if err := session.SetUser(c, user.ID.String(), user.Email, user.DisplayName); err != nil {
			return c.String(http.StatusInternalServerError, "Failed to create session")
		}

		// without re-calling Login every time
		return c.Redirect(http.StatusFound, "/admin/dashboard/")
	})

	adminGroup.POST("/logout/", func(c echo.Context) error {
		if err := session.Clear(c); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.Redirect(http.StatusFound, "/admin/")
	})

	adminGroup.GET("/dashboard/", func(c echo.Context) error {
		return templates.Dashboard(session.GetDisplayName(c)).Render(c.Request().Context(), c.Response())
	}, custom_middleware.RequireAdmin)

	// for _, r := range e.Routes() {
	// 	fmt.Println(r.Method, r.Path)
	// }

	log.Printf("Starting server on port %s", port)

	e.Logger.Fatal(e.Start(":" + port))

}
