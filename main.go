package main

import (
	"fmt"
	"os"
	"text/template"
)

type Stats struct {
	cnt      int
	category map[string]Events
}

type Events struct {
	cnt   int
	event map[string]*Event
}

type Event struct {
	value int64
}

type Foo struct {
	I int
	S string
}

func main() {
	stats := Stats{}
	stats.cnt = 33
	stats.category = make(map[string]Events)
	e, f := stats.category["aa"]
	if !f {
		e = Events{} //f will be nil, so let e = Events{}
	}

	//Go ahead and initialize these fields so we can later access them using the "aa" index that we create in just a sec..
	e.cnt = 66
	e.event = make(map[string]*Event)

	stats.category["aa"] = e //Associate the "aa" index with the Events struct (that was created earlier)

	stats.category["aa"].event["bb"] = &Event{}
	stats.category["aa"].event["bb"].value = 99 //And finally assign the actual in64 value type (99) to the "value" field name

	fmt.Println(stats.cnt, stats.category["aa"].event["bb"].value)

	var m = map[int]Foo{
		7: {6 * 9, "Question"},
		3: {42, "Answer"},
	}

	// Get the 2nd record (index=3) and parse it to screen using Templates' New() method expression.
	tmpl, err := template.New("innovative").Parse("{{.I}} is the integer and {{.S}} is the string")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, m[3])

}
