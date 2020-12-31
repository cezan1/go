package main

import (
	"fmt"
	"myCalc/simplemath"
	"os"
	"strconv"
)

//Usage 输出信息
var Usage = func() {
	fmt.Println("USAGE: mycalc command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\tAddition of two values.\n\tsqrt\tSquare root of a ono-negative value.")
}

func main() {
	args := os.Args //接受命令行参数
	if args == nil || len(args) < 3 {
		Usage()
		return
	}

	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("USAGE:calcs add <integer1> <interger2>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calcs add <integer1> <integer2>")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result: ", ret)
	case "sqrt":
		if len(args) != 3 {
			fmt.Println("USAGE: calcs sqrt <integer>")
		}
		v, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("USAGE: calcs sqrt <integer>")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("Result: ", ret)
	default:
		Usage()
	}
}
