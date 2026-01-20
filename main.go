package main

import (
	"context"
	"log"

	// "net/http"
	"io"
	"os"

	"resume-app/models"
	"resume-app/templates"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	supabase "github.com/supabase-community/supabase-go"
)

var supabaseClient *supabase.Client
var port string

func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

func init() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading configuration from environment variables")
	}

	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_KEY")
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if url == "" || key == "" {
		log.Fatal("SUPABASE_URL and SUPABASE_KEY environment variables must be set")
	}

	var err error
	supabaseClient, err = supabase.NewClient(url, key, nil)
	if err != nil {
		log.Fatal("Error creating Supabase client:", err)
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Static("/static", "static")

	e.GET("/", handleResume)

	log.Println("Server starting on port :" + port)
	e.Logger.Fatal(e.Start(":" + port))
}

func handleResume(c echo.Context) error {
	// ctx := context.Background()

	// Fetch resume data from Supabase
	// var resumes models.Resume
	// data, _, err := supabaseClient.From("resumes").Select("*", "", false).Execute()

	// if err != nil {
	// 	log.Printf("Error fetching resume: %v", err)
	// 	// Return sample data if database fails
	// 	resume := getSampleResume()
	// 	return renderTempl(c, templates.Resume(resume))
	// }

	// Unmarshal the data into resumes slice
	// if err := json.Unmarshal(data, &resumes); err != nil {
	// 	log.Printf("Error unmarshaling resume: %v", err)
	// 	// Return sample data if database fails
	// 	resume := getSampleResume()
	// 	return renderTempl(c, templates.Resume(resume))
	// }

	// if len(resumes) == 0 {
	// 	resume := getSampleResume()
	// 	return renderTempl(c, templates.Resume(resume))
	// }

	resume := getSampleResume()
	return renderTempl(c, templates.Resume(resume))
	// return renderTempl(c, templates.Resume(resumes[0]))
}

func renderTempl(c echo.Context, component templ.Component) error {
	return component.Render(context.Background(), c.Response().Writer)
}

func getSampleResume() models.Resume {
	return models.Resume{
		ID:       "1",
		Name:     "Julio Parra Sanchez",
		Title:    "Profesional Software Engineer",
		Email:    "parrajulio33@gmail.com",
		Phone:    "+1 (720) 277-8846",
		Location: "Denver, CO",
		LinkedIn: "https://www.linkedin.com/in/julio-parra-sanchez/",
		GitHub:   "https://github.com/parrajulio33",
		About:    "I started my passion in software developing after studying electrical engineering and realizing that it was the software, not the hardware that really spoke to me. To me, nothing is more involved than creating something new to bring into the world. I've been in both, building and managing phases of an organization and it's both the challenges and opportunities that come during hyper-growth that I'm most passionate about.",
		Skills: []models.Skill{
			{Name: "HTML", Icon: "https://cdn.jsdelivr.net/npm/simple-icons@3.13.0/icons/html5.svg"}, {Name: "CSS", Icon: "https://cdn.jsdelivr.net/npm/simple-icons@3.13.0/icons/css3.svg"}, {Name: "Go", Icon: "https://cdn.jsdelivr.net/npm/simple-icons@3.13.0/icons/go.svg"}, {Name: "JavaScript", Icon: "https://cdn.jsdelivr.net/npm/simple-icons@3.13.0/icons/javascript.svg"}, {Name: "React", Icon: "https://cdn.jsdelivr.net/npm/simple-icons@3.13.0/icons/react.svg"}, {Name: "React Native", Icon: "https://cdn.jsdelivr.net/npm/simple-icons@3.13.0/icons/react.svg"}, {Name: "Docker", Icon: "https://cdn.jsdelivr.net/npm/simple-icons@3.13.0/icons/docker.svg"}, {Name: "Expo", Icon: "https://cdn.jsdelivr.net/npm/simple-icons@3.13.0/icons/expo.svg"},
			{Name: "PostgreSQL", Icon: "https://cdn.jsdelivr.net/npm/simple-icons@3.13.0/icons/postgresql.svg"}, {Name: "MongoDB", Icon: "https://cdn.jsdelivr.net/npm/simple-icons@3.13.0/icons/mongodb.svg"}, {Name: "Python", Icon: "https://cdn.jsdelivr.net/npm/simple-icons@3.13.0/icons/python.svg"}, {Name: "GraphQL", Icon: "https://cdn.jsdelivr.net/npm/simple-icons@3.13.0/icons/graphql.svg"}, {Name: "REST APIs", Icon: "https://cdn.jsdelivr.net/npm/simple-icons@3.13.0/icons/postman.svg"}, {Name: "SQL", Icon: "https://cdn.jsdelivr.net/npm/simple-icons@3.13.0/icons/microsoftsqlserver.svg"}, {Name: "Angular", Icon: "https://cdn.jsdelivr.net/npm/simple-icons@3.13.0/icons/angular.svg"}, {Name: "Firebase", Icon: "https://cdn.jsdelivr.net/npm/simple-icons@3.13.0/icons/firebase.svg"},
			{Name: "Django", Icon: "https://cdn.jsdelivr.net/npm/simple-icons@3.13.0/icons/django.svg"},
		},
		Experience: []models.Experience{
			{
				Company:     "Wausau Coated Products Inc.",
				Position:    "Programmer Analyst IT",
				WebPage:     "https://www.wausaucoated.com/",
				StartDate:   "March 2022",
				EndDate:     "December 2025",
				Description: "Work with a senior IT programmer/ analyst(s), users and management to assets software needs and system requirements; Mentor Entry level IT programmer/analyst(s) to assist them with their tasks and projects; Design, develop, test, document, implement and support new cus-tom-written software, as well as modify and support existing soft-ware; Monitor existing software and increase performance, as needed; Provide 24/7 support on a rotating schedule with other members of the IT Department.",
			},
			{
				Company:     "Alfi Inc",
				Position:    "Front End Engineer",
				WebPage:     "https://www.getalfi.com/",
				StartDate:   "April 2021",
				EndDate:     "March 2022",
				Description: "Create reusable components using MaterialUI framework; Design, develop, test, document, implement, modify and support existing software ReactJS and React Native. Troubleshoot, debug and upgrade existing software; CI/CD pipeline using GitHub Actions.",
			},
			{
				Company:     "Techtonic Group",
				Position:    "Apprentice & Software Developer",
				WebPage:     "https://techtonic.com/",
				StartDate:   "March 2019",
				EndDate:     "December 2020",
				Description: "AWS cognito, S3 in a project using Django & Python. Also, I worked on a project using React & Node.js; Development React Native project: Using framework EXPO and GCP; Install, config and implement Docker and Kubernetes for orchestration containers; Learn new technologies as required.",
			},
			{
				Company:     "PDVSA Servicios Electricos Tamare",
				WebPage:     "https://www.pdvsa.com/",
				Position:    "Electrical Engineer",
				StartDate:   "May 2002",
				EndDate:     "December 2005",
				Description: "Worked management plans on land distribution projects; Basic operation on different electrical PDVSA's facilities; create reports and technical analysis of electrical systems; Involved in electrical maintenance works in the field; Ensure compliance with safety regulations in all electrical works. Calculation used electrical software such as ETAP and AutoCAD.",
			},
		},
		Projects: []models.Projects{
			{
				Name:        "PPO performance optimization for Employee",
				Description: "A Web application to optimize performance for employees using ReactJS.",
				Link:        "https://www.henselphelps.com/",
			},
			{
				Name:        "Liga Latina Softball Denver App",
				Description: "Build a website (in construction), and a mobile app for the league (in construction) using React Native and Expo.",
				Link:        "https://github.com/parrajulio33/ligaLatinaDenver",
			},
			{
				Name:        "Job Application Tracker",
				Description: "Mobile app for tracking job applications and interviews using React Native and Expo and EAS. Supabase as backend.",
				Link:        "https://github.com/parrajulio33/jobs",
			},
		},
		Education: []models.Education{
			{
				Institution: "University Rafael Urdaneta",
				WebPage:     "https://uru.edu/",
				Degree:      "Bachelor Degree of Electrical Engineering",
				Year:        "June 1998 - December 2003",
			},
		},
	}
}
