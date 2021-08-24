package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Dict struct {
	elems [][]byte
}

func (d *Dict) set(full []byte) string {
	d.elems = append(d.elems, full)

	return "http://" + addr + "/" + strconv.Itoa(len(d.elems) - 1)
}

func (d *Dict) get(id int) ([]byte, error) {
	if len(d.elems) == 0 {
		return nil, errors.New("not found")
	}

	el := d.elems[id]

	fmt.Println("founded el", el)

	if len(el) != 0 {
		return el, nil
	}

	return nil, errors.New("not found")
}
