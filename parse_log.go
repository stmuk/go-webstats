package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {

	fName := os.Args[1]

	vhostRegex, err := regexp.Compile(`\[(.*?)\]`)

	var dates = map[string]int{}
	var hosts = map[string]int{}

	file, err := os.Open(fName)

	if err != nil {
		log.Fatal(err)
	}

	fileRead := bufio.NewReader(file)

	// keep track of insert order
	var m = map[int]string{}
	//var m map[int]string XXX

	i := 0

	for {
		line, err := fileRead.ReadString('\n')
		if err != nil {
			break
		}

		host := strings.Fields(line)[0]
		hosts[host]++

		date := vhostRegex.FindAllStringSubmatch(line, -1)[0][1]
		date = strings.Split(date, ":")[0]
		dates[date]++

		m[i] = date
		i++
	}

	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	seen := map[string]int{}
	for _, k := range keys {

		if seen[m[k]] == 0 {
			fmt.Println("Date:", m[k], "Hits:", dates[m[k]])
			seen[m[k]]++
		}
	}
}
