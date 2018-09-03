package main

import (
	"fmt"

	"github.com/GoesToEleven/GolangTraining/02_package/stringutil"
)

func main() {
	fmt.Println("Hi my name is", stringutil.Reverse(stringutil.MyName))
}
