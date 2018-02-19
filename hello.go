package main

import (
	"encoding/json"
	"fmt"
)

const text = `
{
	"part1": "Hello",
	"part2": "world!"
}
`

func main() {
	var m = make(map[string]string)
	json.Unmarshal([]byte(text), &m)
	fmt.Println(m["part1"], m["part2"])
}
