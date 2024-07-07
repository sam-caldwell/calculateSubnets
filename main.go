package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(ErrMissingArguments)
		os.Exit(ExitMissingArgs)
	}
	parentCIDR := os.Args[ArgParentCIDR]
	subnetSize := func() int {
		var err error
		var n int64
		s := os.Args[ArgSubnetSize]
		if n, err = strconv.ParseInt(s, 10, 32); err != nil {
			fmt.Println(err)
			os.Exit(ExitSubnettingError)
		}
		return int(n)
	}()

	//Optional result count
	resultCount := 0
	if len(os.Args) == 4 {
		resultCount = func() int {
			var err error
			var n int64
			s := os.Args[ArgResultCount]
			if n, err = strconv.ParseInt(s, 10, 32); err != nil {
				fmt.Println(ErrInvalidResultCount)
				os.Exit(ExitInvalidResultCount)
			}
			return int(n)
		}()
	}

	if subnets, err := CalculateSubnets(parentCIDR, subnetSize); err != nil {
		fmt.Printf(ErrGeneral, err)
	} else {
		if resultCount == 0 {
			resultCount = len(subnets)
		}
		for _, network := range subnets[:resultCount] {
			fmt.Printf("%s", network)
		}
	}
	os.Exit(ExitSuccess)
}
