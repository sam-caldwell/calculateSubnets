package calculateSubnets

const (
	//ArgParentCIDR - cli Argument position
	ArgParentCIDR = 1

	//ArgSubnetSize - cli Argument position
	ArgSubnetSize = 2

	//ArgResultCount - cli Argument position
	ArgResultCount = 3

	// ErrInvalidParentCIDR - standard error
	ErrInvalidParentCIDR = "invalid parent CIDR:%s"

	// ErrInvalidSubnetSize - standard error
	ErrInvalidSubnetSize = "invalid subnet size:%d"

	//MsgIpv4CIDR - standard CIDR format string
	MsgIpv4CIDR = "%d.%d.%d.%d/%d\n"
)
