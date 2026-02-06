# 📄 Go Resume

A modern, responsive resume application built with Go, featuring UI and seamless database integration.

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-green)
![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)

## ✨ Features

- 🎨 **UI** - Modern, professional design with TailwindCSS
- 📱 **Fully Responsive** - Perfect on desktop, tablet, and mobile
- 🖨️ **Print Optimized** - Print-ready with preserved colors and styling
- ⌨️ **Keyboard Shortcuts** - Quick print with `Ctrl+P` or `Cmd+P`
- 🗄️ **Supabase Integration** - Optional cloud database support
- ⚡ **Fast & Lightweight** - Built with Go for optimal performance
- 🔒 **Type-Safe Templates** - Using Templ for compile-time safety
- 🎯 **Zero Config** - Works out of the box with sample data

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- Templ CLI (for template compilation)

### Installation

1. **Clone the repository**

```bash
git clone https://github.com/parrajulio33/golang-resume.git
cd golang-resume
```

2. **Install dependencies**

```bash
go mod download
```

3. **Install Templ CLI**

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

4. **Generate template code**

```bash
templ generate
```

5. **Run the application**

```bash
go run main.go
```

6. **Open your browser**

```
http://localhost:8080
```

That's it! 🎉 the resume is now running with sample data.

## 🏗️ Project Structure

```
resume-app/
├── main.go                        # Application entry point & server setup
├── models/
│   └── resume.go                  # Data models and structures
|── node_modules                   # Node modules
├── templates/
│   └── resume.templ               # Templ UI templates
├── repositories/
│   └── resume_repository.go       # Supabase Query
├── supabase/
│   └── client.go                  # Supabase Client
├── static/
|   └── icon.png
│   └── styles.css                 # Custom CSS & print styles
│   └── input.css
│   └── output.css
│   └── qr-code.png
├── go.mod                         # Go module dependencies
├── .air.toml                      # Air toml
├── Makefile                       # Make file
├── tailwind.config.js             # Tailwindcss config
├── nixpacks.toml                  # Toml file if you want to deploy it on Railway Cloud Services
└── README.md                      # You are here!
```

## 🎨 Tech Stack

| Technology      | Purpose                        |
| --------------- | ------------------------------ |
| **Go**          | Backend language & server      |
| **Templ**       | Type-safe HTML templating      |
| **Echo**        | Fast, minimalist web framework |
| **TailwindCSS** | Utility-first CSS framework    |
| **Supabase**    | Optional PostgreSQL database   |

## 🔧 Configuration

### Using Supabase (Optional)

1. **Create a Supabase project** at [supabase.com](https://supabase.com)

2. **Run the SQL schema** (found in `main.go` comments):

```sql
CREATE TABLE resumes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    title TEXT NOT NULL,
    email TEXT NOT NULL,
    phone TEXT,
    location TEXT,
    linkedin TEXT,
    github TEXT,
    about TEXT,
    skills JSONB,
    experience JSONB,
    education JSONB,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);
```

3. **Set environment variables**:

```bash
export SUPABASE_URL="https://your-project.supabase.co"
export SUPABASE_KEY="your-anon-key"
```

4. **Restart the application**:

```bash
go run main.go
```

### Environment Variables

| Variable            | Required | Description                   |
| ------------------- | -------- | ----------------------------- |
| `SUPABASE_URL`      | No       | Your Supabase project URL     |
| `SUPABASE_ANON_KEY` | No       | Your Supabase anon/public key |

_Note: App works without these using sample data_

## 📝 Customizing Your Resume

### Option 1: Edit Sample Data (Quick)

In `main.go`, modify the `getSampleResume()` function:

```go
func getSampleResume() models.Resume {
    return models.Resume{
        Name:     "Your Name",
        Title:    "Your Title",
        Email:    "your.email@example.com",
        Phone:    "+1 (555) 123-4567",
        Location: "Your City, State",
        LinkedIn: "...",
        GitHub: "...",
        About:  "Your professional summary here...",
        Skills: []string{
            "Skill 1", "Skill 2", "Skill 3",
        },
        // ... add your experience and education
    }
}
```

### Option 2: Use Supabase (Recommended)

Insert your data directly in Supabase dashboard or via SQL:

```sql
INSERT INTO resumes (name, title, email, phone, location, linkedin, github, about, skills, experience, education)
VALUES (
    'Your Name',
    'Your Title',
    'your.email@example.com',
    '+1 (555) 123-4567',
    'Your City, State',
    'linkedin',
    'github',
    'Your professional summary...',
    '["Go", "JavaScript", "React"]'::jsonb,
    '[{"company": "Company Name", "position": "Your Position", ...}]'::jsonb,
    '[{"institution": "University", "degree": "Your Degree", ...}]'::jsonb
);
```

## 🖨️ Printing Your Resume

### Keyboard Shortcut

Press `Ctrl+P` (Windows/Linux) or `Cmd+P` (Mac) anywhere on the page.

### Print Button

Click the "Print Resume" button at the bottom of the page.

### Print Settings

- **Enable "Background graphics"** in your browser's print dialog
- Choose **"Save as PDF"** for a digital copy
- All colors, icons, and badges will be preserved!

## 🎯 Features in Detail

### Responsive Design

- Mobile-first approach
- Breakpoints for tablets and desktops
- Optimized typography for all screen sizes

### Print Optimization

- Colors and backgrounds preserved with `print-color-adjust: exact`
- Clean page breaks
- Hidden UI elements (buttons, etc.)
- Professional PDF output

### Type Safety

- Templ provides compile-time template checking
- Catch errors before runtime
- Better IDE support and autocomplete

## 🚀 Deployment

### Deploy to a VPS or Cloud Provider

1. **Build the binary**:

```bash
go build -o resume-app
```

2. **Run the binary**:

```bash
./resume-app
```

3. **Use a process manager** (PM2, systemd, etc.):

```bash
# Example with systemd
sudo systemctl start resume-app
```

### Docker (Optional)

Create a `Dockerfile`:

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o resume-app

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/resume-app .
COPY --from=builder /app/static ./static
EXPOSE 8080
CMD ["./resume-app"]
```

Build and run:

```bash
docker build -t resume-app .
docker run -p 8080:8080 resume-app
```

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Echo Framework](https://echo.labstack.com/) - Minimalist web framework
- [Templ](https://templ.guide/) - Type-safe HTML templating
- [TailwindCSS](https://tailwindcss.com/) - Utility-first CSS
- [Supabase](https://supabase.com/) - Open source Firebase alternative

## 📧 Contact

Email - parrajulio33@gmail.com

Project Link: [https://github.com/parrajulio33/golang-resume](https://github.com/parrajulio33/golang-resume)

---

⭐ **Star this repo** if you find it helpful!

## Function Example to getResume

```
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
```

Made with Golang
