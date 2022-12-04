/*
Beste Geerke,


*/
package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type twoLines struct {
	lines  []string
	words  [][]string
	rhymes map[string][]string
}

func (lines twoLines) String() string {

	word1 := lines.words[0][rand.Intn(len(lines.words[0]))]
	word2 := lines.words[1][rand.Intn(len(lines.words[1]))]
	rhymeWord1 := fmt.Sprintf("%s", reflect.ValueOf(lines.rhymes).MapKeys()[rand.Intn(len(lines.rhymes))])
	rhymeWord2 := lines.rhymes[rhymeWord1][rand.Intn(len(lines.rhymes[rhymeWord1]))]

	return fmt.Sprintf("%s\n%s",
		fmt.Sprintf(lines.lines[0], word1, rhymeWord1),
		fmt.Sprintf(lines.lines[1], word2, rhymeWord2))
}

var RHYME = []twoLines{
	{
		lines: []string{
			"%s zat eens te %s",
			"wat zal hij %s eens %s",
		},
		words: [][]string{
			{"Sint", "Piet", "Enge man met grote baard"},
			{"Geerke", "Rebecca", "Leids studentje"},
		},
		rhymes: map[string][]string{
			"denken":   {"schenken", "herdenken", "krenken"},
			"vervelen": {"aanbevelen", "mailen"},
		},
	},
	{
		lines: []string{
			"%s had werkelijk waar geen %s",
			"dan maar een %s of misschien %s",
		},
		words: [][]string{
			{"Hulpsinterklaas", "Kerstman", "Diewertje Blok"},
			{"boek", "stapel faxen", "flutroman"},
		},
		rhymes: map[string][]string{
			"enkel idee":        {"een ouderwetse cd", "een met leer behangen barbie fay", "SARS-CoV-2"},
			"fatsoenlijke clue": {"wilde kangoeroe", "flinke lading doekoe", "welgemeende 'fuck you'"},
		},
	},
}

func main() {

	rand.Seed(time.Now().UnixNano())
	for _, r := range RHYME {
		fmt.Println(r)
	}

}
