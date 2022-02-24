package constant

import (
	"os"
	"strings"
)

var names []string

func LoadNamesData() {
	fstring, _ := os.ReadFile("./data/People-Name-List/Crawled-People-Names/names")
	data := strings.Split(string(fstring), "\n")
	res := []string{}
	for _, name := range data {
		if len(name) > 3 {
			res = append(res, name[:len(name)-3])
		}
	}
	names = res
}

func GetName(idx int) string {
	return names[idx]
}
