package main

import (
	"fmt"
	"sort"
)

// START OMIT
type People struct{ first, last string }

func (p People) String() string { return p.first + " " + p.last }

type ByFirstName []People

func (s ByFirstName) Len() int           { return len(s) }
func (s ByFirstName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByFirstName) Less(i, j int) bool { return s[i].first < s[j].first }

func main() {
	people := []People{
		{"Bill", "Clinton"}, {"Romeo", "Capulet"},
		{"Bill", "Cosby"}, {"Juliet", "Montague"},
		{"Bill", "Gates"}, {"Walter", "White"},
		{"Allan", "Poe"},
	}
	sort.Sort(ByFirstName(people))
	fmt.Println(people)
}

// END OMIT
