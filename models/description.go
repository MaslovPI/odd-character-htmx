package models

import "fmt"

type Starter struct {
	Content string
	Arcana  string
	Pet     string
	Hire    string
}

func GenerateStarter(max int, hp int) Starter {
	starter := Starter{
		Content: fmt.Sprintf("Starting content, hp: %d, max: %d", hp, max),
		Arcana:  "Powerful Magic uuuu",
		Pet:     "A doggo wuf-wuf",
		Hire:    "Solder-man",
	}
	return starter
}
