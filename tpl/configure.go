package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"
)

type Repository struct {
	Name       string
	Private    bool
	Visibility string
	TplName    string
	Archived   bool
}

func main() {
	//	org := flag.String("org", "owls-nest-farm", "The name of the organization.")
	n := flag.Int("n", 25, "The number of repositories to create in the given organization.")
	filename := flag.String("out", "github.yaml", "The name of the generated config file.")
	flag.Parse()

	t := template.Must(template.ParseFiles("github.tpl"))
	repos := make([]Repository, *n)

	for i := 0; i < *n; i++ {
		var isArchived bool
		if i%2 == 0 {
			isArchived = true
		} else {
			isArchived = false
		}

		var isPrivate bool
		if i%3 == 0 {
			isPrivate = true
		} else {
			isPrivate = false
		}

		repos[i] = Repository{
			Name:       fmt.Sprintf("foo%d", i),
			Private:    isPrivate,
			Visibility: "private",
			TplName:    "kilgore-trout",
			Archived:   isArchived,
		}
	}

	f, err := os.Create(*filename)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(f, repos)
	if err != nil {
		panic(err)
	}
}
