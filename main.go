package main

import (
	"errors"
	"fmt"
)

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}

	return a / b, nil
}

func main() {
	fmt.Println("Hello, 世界")

	o, err := Division(8.0, 2)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}
	fmt.Println("output : ", o)
	return

}
