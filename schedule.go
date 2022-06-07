package nogisched

import (
	"bytes"
	"fmt"
	"strings"
)

type Schedule struct {
	Date        string
	Appearances []Appearance
}

type Appearance struct {
	Category string
	Time     string
	Title    string
}

func (s Schedule) String() string {
	wr := &bytes.Buffer{}
	fmt.Fprintf(wr, "%s\n", s.Date)
	for _, a := range s.Appearances {
		fmt.Fprintf(wr, "%s\n", a.String())
	}
	return wr.String()
}

func (a Appearance) String() string {
	target := []string{}
	if a.Category != "" {
		target = append(target, a.Category)
	}
	if a.Time != "" {
		target = append(target, a.Time)
	}
	if a.Title != "" {
		target = append(target, a.Title)
	}
	return strings.Join(target, " ")
}
