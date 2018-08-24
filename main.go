package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"

	"github.com/hashicorp/terraform/helper/hashcode"
)

type Elements struct {
	Elements []Element `json:"elements"`
}

type Element struct {
	Order int    `json: "order"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

func main() {
	var buf bytes.Buffer
	filePtr := flag.String("file", "elements.json", "a JSON file of buffer elements")
	flag.Parse()

	jsonFile, err := os.Open(*filePtr)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var elements Elements

	err2 := json.Unmarshal(byteValue, &elements)

	if err2 != nil {
		panic(err2)
	}

	sort.Slice(elements.Elements[:], func(i, j int) bool {
		return elements.Elements[i].Order < elements.Elements[j].Order
	})

	for _, e := range elements.Elements {
		if e.Type == "bool" {
			b, err := strconv.ParseBool(e.Value)

			if err != nil {
				panic(err)
			}

			buf.WriteString(fmt.Sprintf("%t-", b))
		} else if e.Type == "string" {
			buf.WriteString(fmt.Sprintf("%s-", string(e.Value)))
		} else {
			panic("Unknown element type! Only bool and string supported.")
		}
	}

	fmt.Println(hashcode.String(buf.String()))
}
