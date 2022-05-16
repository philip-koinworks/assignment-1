package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"assignment1/model"
)

var (
	errNoArguments    = errors.New("No arguments found")
	errManyArguments  = errors.New("Please specify only 1 arguments")
	errArgumentNotInt = errors.New("Please pass integer as the argument")
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(errNoArguments)
	}

	if len(os.Args) > 2 {
		log.Fatal(errManyArguments)
	}

	idx, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(errArgumentNotInt)
	}

	s := model.GetStudent(idx)
	res, err := s.ToJson()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
