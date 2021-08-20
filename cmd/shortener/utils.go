package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getId(s string) (int, error){
	id := strings.Split(s, "/")[1]

	fmt.Println("parsed id", id)

	return strconv.Atoi(id)
}