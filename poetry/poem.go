package poetry

type Line string
type Stanza []Line
type Poem []Stanza

func NewPoem() Poem {
	return Poem{}
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
