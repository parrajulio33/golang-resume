package models

type Resume struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	Title      string       `json:"title"`
	Email      string       `json:"email"`
	Phone      string       `json:"phone"`
	Location   string       `json:"location"`
	LinkedIn   string       `json:"linkedin"`
	GitHub     string       `json:"github"`
	About      string       `json:"about"`
	Skills     []string     `json:"skills"`
	Experience []Experience `json:"experience"`
	Projects   []Projects   `json:"projects"`
	Education  []Education  `json:"education"`
}

type Experience struct {
	Company     string `json:"company"`
	WebPage     string `json:"webpage"`
	Position    string `json:"position"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Description string `json:"description"`
}

type Projects struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type Education struct {
	Institution string `json:"institution"`
	WebPage     string `json:"webpage"`
	Degree      string `json:"degree"`
	Year        string `json:"year"`
}
