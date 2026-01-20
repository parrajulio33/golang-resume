# 📄 Go Resume Builder

A modern, responsive resume application built with Go, featuring beautiful UI and seamless database integration.

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-green)
![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)

## ✨ Features

- 🎨 **Beautiful UI** - Modern, professional design with TailwindCSS
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
git clone https://github.com/yourusername/go-resume-builder.git
cd go-resume-builder
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

That's it! 🎉 Your resume is now running with sample data.

## 🏗️ Project Structure

```
resume-app/
├── main.go                 # Application entry point & server setup
├── models/
│   └── resume.go          # Data models and structures
|── node_modules           # Node modules
├── templates/
│   └── resume.templ       # Templ UI templates
├── static/
|   └── icon.png
│   └── styles.css         # Custom CSS & print styles
│   └── input.css
│   └── output.css
│   └── qr-code.png
├── go.mod                 # Go module dependencies
├── .air.toml                 # Air toml
├── Makefile                 # Make file
├── netlify.toml                 # Netlify toml
├── tailwind.config.js      # Tailwindcss config
├── build.sh                 # Build sh
└── README.md              # You are here!
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
    summary TEXT,
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

| Variable       | Required | Description                   |
| -------------- | -------- | ----------------------------- |
| `SUPABASE_URL` | No       | Your Supabase project URL     |
| `SUPABASE_KEY` | No       | Your Supabase anon/public key |

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
        Summary:  "Your professional summary here...",
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
INSERT INTO resumes (name, title, email, phone, location, summary, skills, experience, education)
VALUES (
    'Your Name',
    'Your Title',
    'your.email@example.com',
    '+1 (555) 123-4567',
    'Your City, State',
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

Your Name - [@yourtwitter](https://twitter.com/yourtwitter) - your.email@example.com

Project Link: [https://github.com/yourusername/go-resume-builder](https://github.com/yourusername/go-resume-builder)

---

⭐ **Star this repo** if you find it helpful!

Made with Golang
