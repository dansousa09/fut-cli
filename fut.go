package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Match struct {
	Teams string `json:"teams"`
	Time  string `json:"time"`
}

var matches []Match

func main() {
	c := colly.NewCollector()

	c.OnHTML("div[class=match-wrap-content] div div div", func(h *colly.HTMLElement) {
		match := Match{
			Teams: h.ChildText("a[class=match-content-score] div div div"),
			Time:  h.ChildText("div[class=match-footer] div div div"),
		}
		if match.Teams != "" && match.Time != "" {
			matches = append(matches, match)
		}
	})

	currentTime := time.Now()

	c.Visit("https://www.uol.com.br/esporte/futebol/central-de-jogos/#/" + currentTime.Format("02-01-2006"))

	for i, match := range matches {
		if i < 50 {
			teamsStr := fmt.Sprintf("%+v", match.Teams)
			timeStr := fmt.Sprintf("%+v", match.Time)

			parts := strings.Split(teamsStr, "-()-")
			team1 := strings.TrimSpace(parts[0])
			team2 := strings.TrimSpace(parts[1])

			fmt.Println(team1 + " x " + team2 + " / " + timeStr)
		}
	}
}
