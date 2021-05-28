package main

import (
	"convert-number-to-vnese/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if args == nil {
		fmt.Printf("Please input arg")
	}
	argsWithoutProgram := os.Args[1:]
	money := argsWithoutProgram[0]
	if money != "" {
		money = strings.ReplaceAll(money, ".", "")
		if strings.Contains(money, ",") {
			money = strings.ReplaceAll(money, ",", ".")
			number, _ := strconv.ParseFloat(money, 64)
			fmt.Printf("%s >>> %s\n", money, utils.Float2Vn(number))
		} else {
			number, _ := strconv.Atoi(money)
			fmt.Printf("%s >>> %s\n", money, utils.Int2Vn(int64(number)))
		}
	}
}
