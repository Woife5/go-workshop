package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var input = flag.String("i", "points.json", "Input File")
var output = flag.String("o", "output.yml", "Output File")

func handleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	flag.Parse()
	fmt.Println(*input)

	jsonFile, err := os.Open(*input)
	handleError(err)
	defer jsonFile.Close()

	bytes, err := io.ReadAll(jsonFile)
	handleError(err)

	var points []Point
	json.Unmarshal(bytes, &points)

	for _, point := range points {
		fmt.Println(point.PrettyPrint())
	}

	yamlData, err := yaml.Marshal(&points)
	handleError(err)

	yamlFile, err := os.OpenFile(*output, os.O_WRONLY|os.O_CREATE, 0600)
	handleError(err)
	defer yamlFile.Close()

	yamlFile.Write(yamlData)
}
