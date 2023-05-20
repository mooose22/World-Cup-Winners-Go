package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
)

var (
	worldCupWinners = make(map[string]string)
	keys            = []string{}
	values          = []string{}
)

func main() {
	readFile, _ := os.ReadFile("data.txt")
	a := regexp.MustCompile(`\r\n|\n`) // split by newline, cursor newline
	str := a.Split(string(readFile), -1)
	for _, values := range str {
		keys = append(keys, values) // creating slice
	}
	for i := 0; i < len(keys); i++ {
		remover := regexp.MustCompile(`:`) // removing unnecessary chars
		keys[i] = remover.ReplaceAllString(keys[i], "")
		findKey := regexp.MustCompile(`\d+`)                                         //  looking for year
		findValue := regexp.MustCompile(`[a-zA-Z]+`)                                 // looking for country
		worldCupWinners[findKey.FindString(keys[i])] = findValue.FindString(keys[i]) // creating map key-year and country-value

	}

	keySlice := make([]string, 0, len(worldCupWinners)) // creating slice to sort years
	for key := range worldCupWinners {
		keySlice = append(keySlice, key)
	}
	sort.Strings(keySlice) // sorting keys

	yearFlag := flag.String("year", "default", "choose year")
	flag.Parse()
	checkYear := ""
	if len(os.Args) != 1 && len(os.Args) <= 2 { // if flag is active
		*yearFlag = os.Args[1]
		regex := regexp.MustCompile(`\-+year=`)
		checkYear = regex.ReplaceAllLiteralString(os.Args[1], "")
		_, ok := worldCupWinners[checkYear]
			if ok {
				fmt.Printf("WC %s Winner: %s\n", checkYear, worldCupWinners[checkYear])
			} else {
				fmt.Printf("The FIFA World Cup didn't take place in %s\n",checkYear)
			}
		}

	if len(os.Args) == 1 { // if no flag, show all winners
		for i := 0; i < len(keySlice); i++ {
			fmt.Printf("WC %s Winner: %s\n", keySlice[i], worldCupWinners[keySlice[i]])
		}
	}
}
