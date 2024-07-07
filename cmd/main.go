package main

import (
	"github.com/sam-caldwell/ansi"
	"github.com/sam-caldwell/calculateSubnets"
	"github.com/sam-caldwell/errors"
	"github.com/sam-caldwell/exit"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		ansi.Red().Println(errors.MissingArguments).Reset().Fatal(exit.MissingArg)
	}
	parentCIDR := os.Args[calculateSubnets.ArgParentCIDR]
	subnetSize := func() int {
		var err error
		var n int64
		s := os.Args[calculateSubnets.ArgSubnetSize]
		if n, err = strconv.ParseInt(s, 10, 32); err != nil {
			ansi.Red().Println(err.Error()).Fatal(exit.ParseError).Reset()
		}
		return int(n)
	}()

	//Optional result count
	resultCount := 0
	if len(os.Args) == 4 {
		resultCount = func() int {
			var err error
			var n int64
			s := os.Args[calculateSubnets.ArgResultCount]
			if n, err = strconv.ParseInt(s, 10, 32); err != nil {
				ansi.Red().Println(errors.InvalidResultCount).Fatal(exit.InvalidResult).Reset()
			}
			return int(n)
		}()
	}

	if subnets, err := calculateSubnets.CalculateSubnets(parentCIDR, subnetSize); err != nil {
		ansi.Red().Println(err.Error()).Fatal(exit.GeneralError).Reset()
	} else {
		if resultCount == 0 {
			resultCount = len(subnets)
		}
		for _, network := range subnets[:resultCount] {
			ansi.Printf("%s", network)
		}
	}
}
