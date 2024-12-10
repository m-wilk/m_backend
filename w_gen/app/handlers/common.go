package handlers

import (
	"net/http"
	"regexp"
	"time"

	"github.com/labstack/echo/v4"
)

type Staff struct {
	Name     string
	PhotoURL string
	JobTitle string
	Socials  []SocialLink
}

// SocialLink represents a single social media link
type SocialLink struct {
	Platform string
	Icon     string
	URL      string
}

var CREATIVE_STAFF = []Staff{
	{
		Name:     "Jan Kowalski",
		PhotoURL: "https://example.com/photos/jan-kowalski.jpg",
		JobTitle: "Software Engineer",
		Socials: []SocialLink{
			{Platform: "LinkedIn", Icon: "linkedin-icon", URL: "https://linkedin.com/in/jankowalski"},
			{Platform: "GitHub", Icon: "github-icon", URL: "https://github.com/jankowalski"},
			{Platform: "Twitter", Icon: "twitter-icon", URL: "https://twitter.com/jankowalski"},
		},
	},
	{
		Name:     "Anna Nowak",
		PhotoURL: "https://example.com/photos/anna-nowak.jpg",
		JobTitle: "Product Manager",
		Socials: []SocialLink{
			{Platform: "LinkedIn", Icon: "linkedin-icon", URL: "https://linkedin.com/in/annanowak"},
			{Platform: "Medium", Icon: "medium-icon", URL: "https://medium.com/@annanowak"},
			{Platform: "Instagram", Icon: "instagram-icon", URL: "https://instagram.com/annanowak"},
		},
	},
	{
		Name:     "Ralph Edwards",
		PhotoURL: "https://example.com/photos/ralph-edwards.jpg",
		JobTitle: "UI/UX Designer",
		Socials: []SocialLink{
			{Platform: "LinkedIn", Icon: "linkedin-icon", URL: "https://linkedin.com/in/ralph"},
			{Platform: "Medium", Icon: "medium-icon", URL: "https://medium.com/@ralph"},
			{Platform: "Instagram", Icon: "instagram-icon", URL: "https://instagram.com/ralph"},
		},
	},
	{
		Name:     "Esther Howard",
		PhotoURL: "https://example.com/photos/esther-howard.jpg",
		JobTitle: "CTO & Founder",
		Socials: []SocialLink{
			{Platform: "LinkedIn", Icon: "linkedin-icon", URL: "https://linkedin.com/in/esther"},
			{Platform: "Medium", Icon: "medium-icon", URL: "https://medium.com/@esther"},
			{Platform: "Instagram", Icon: "instagram-icon", URL: "https://instagram.com/esther"},
		},
	},
}

func (h *Handler) CreativeStaff(c echo.Context) error {
	time.Sleep(3 * time.Second)
	return c.JSON(http.StatusOK, CREATIVE_STAFF)
}

// Struct do przechowywania danych formularza
type ContactForm struct {
	FullName    string `json:"full_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	ServiceType string `json:"service_type" validate:"required"`
	Message     string `json:"message" validate:"required"`
}

func isValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(regex).MatchString(email)
}

func (h *Handler) ContactForm(c echo.Context) error {
	var form ContactForm

	if err := c.Bind(&form); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	if !isValidEmail(form.Email) {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid email address",
		})
	}

	if form.Message == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Add message",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "Thank you for your submission!",
	})

}
