package main

import (
	"fmt"

	"github.com/PavelYadrov/unittests/poetry"
)

func main() {
	fmt.Printf("dd")

	p := poetry.Poem{{"I am working on a go project,",
		" which has a dependency on original-project.",
		"I now want to change the behavior in this project by modifying original-project.",
		"So I cloned github.com/y/original-project to github.com/x/my-version and replaced all",
		"occurrence of github.com/y/original-project with github.com/x/my-version (including in mod.go)."}}

	v, c := p.Stats()
	fmt.Printf("Vovels %d and Consonants %d", v, c)

}
