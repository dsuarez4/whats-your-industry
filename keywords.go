package main

import (
	"os"
	"encoding/json"
	"io/ioutil"
	"log"
	"errors"
	"fmt"
)

type Keywords struct {
	Data map[string]Value
}

type Value struct {
	Code interface{} `json:code`
	Count float64 `json:count`
	UsIndustry interface{} `json:usindustry`
}

func (kwMap *Keywords) New() *Keywords {

	// Get data
	b, err1 := ioutil.ReadFile("./naics_keyword_output_pretty.json")
	if err1 != nil {
		log.Fatal("Failed to read file")
		os.Exit(1)
	}

	// Keys are arbitrary so pull out keys into map
	var objmap = make(map[string]*json.RawMessage)
	err := json.Unmarshal(b, &objmap)
	if err != nil {
		panic(err)
	}

	// Map json.RawMessage to value struct
	var resultMap = make(map[string]Value)
	for i, j := range objmap {

		values := Value{}

		err := json.Unmarshal(*j, &values)
		if err != nil {
			panic(err)
		}
		resultMap[i] = values
	}

	return &Keywords{resultMap}
}

func (kwMap Keywords) Get(keyword string) (Value, error) {

	if val, ok := kwMap.Data[keyword]; !ok {
		log.Panic("No key found in database")
		return Value{}, errors.New("No key found in database")
	} else {
		fmt.Println(val)
	}


	return Value{}, nil

}


