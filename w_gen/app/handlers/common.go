package handlers

import (
	"net/http"
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
