package models

type Starter struct {
	Content string
	Arcana  string
	Pet     string
	Hire    string
}

func GenerateStarter() Starter {
	starter := Starter{
		Content: "Starting content",
		Arcana:  "Powerful Magic uuuu",
		Pet:     "A doggo wuf-wuf",
		Hire:    "Solder-man",
	}
	return starter
}
