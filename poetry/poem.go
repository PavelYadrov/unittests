package poetry

import (
	"bufio"
	"os"
)

type Line string
type Stanza []Line
type Poem []Stanza

func NewPoem() Poem {
	return Poem{}
}

func LoadPoem(name string) (Poem, error) {
	file, err := os.Open(name)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	p := Poem{}

	s := Stanza{}

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		l := scan.Text()
		if l == "" {
			p = append(p, s)
			s = Stanza{}
			continue
		}
		s = append(s, Line(l))
	}
	p = append(p, s)

	return p, err
}

//функция не изменяем пришедший объект, так что просто без *, работаем с его копией
func (p Poem) Stats() (numVowels, numConsonats int) {
	for _, stanza := range p {
		for _, line := range stanza {
			for _, char := range line {
				switch char {
				case 'a', 'i', 'e', 'o', 'y', 'u':
					numVowels++
				default:
					numConsonats++
				}
			}
		}
	}
	return numVowels, numConsonats
}
