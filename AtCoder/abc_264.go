package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	L := scanInt()
	R := scanInt()
	str := "atcoder"
	fmt.Println(str[L-1 : R])
}
