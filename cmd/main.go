package main

import (
	"github.com/sam-caldwell/ansi"
	"github.com/sam-caldwell/arg"
	"github.com/sam-caldwell/calculateSubnets"
	"github.com/sam-caldwell/errors"
	"github.com/sam-caldwell/exit"
)

func main() {

	var err error
	var parent *arg.Cidr
	var size *arg.Uint
	var subnets []string

	parent, err = arg.NewCidr("cidr", "10.1.0.0/16", "define the parent CIDR within which subnets must exist.")
	if err != nil {
		ansi.Red().Println(err.Error()).Fatal(exit.InvalidInput).Reset()
	}

	size, err = arg.NewUint("size", 24, 0, 32, "define the subnet size (in bits)")
	if err != nil {
		ansi.Red().Println(err.Error()).Fatal(exit.InvalidInput).Reset()
	}

	arg.Parse()

	if err = parent.Verify(); err != nil {
		ansi.Red().Printf(errors.InvalidInput+errors.Details, err).LF().Fatal(exit.InvalidInput).Reset()
	}

	if err = size.Verify(); err != nil {
		ansi.Red().Printf(errors.InvalidInput+errors.Details, err).LF().Fatal(exit.InvalidInput).Reset()
	}

	if *size.Value() > 32 {
		ansi.Red().Printf(errors.InvalidSubnetSize+errors.Details, *size.Value())
	}

	if subnets, err = calculateSubnets.Calculate(*parent.Value(), uint8(*size.Value())); err != nil {
		ansi.Red().Println(err.Error()).Fatal(exit.GeneralError).Reset()
	}
	for _, network := range subnets {
		ansi.Printf("%s\n", network)
	}
}
