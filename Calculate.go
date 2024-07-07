package calculateSubnets

import (
	"fmt"
	"net"
)

// Error messages
const (
	InvalidParentCidr = "invalid parent CIDR: "
	InvalidSubnetSize = "invalid subnet size: "
	Details           = " details: %v"
)

// Calculate calculates subnets based on a parentCIDR and subnetSize.
func Calculate(parentCIDR string, subnetSize uint8) (subnets []string, err error) {
	var ipNet *net.IPNet

	if _, ipNet, err = net.ParseCIDR(parentCIDR); err != nil {
		return nil, fmt.Errorf(InvalidParentCidr+Details, err)
	}

	parentMaskSize, _ := ipNet.Mask.Size()
	newMaskSize := int(subnetSize)

	if newMaskSize > 32 || newMaskSize < parentMaskSize {
		return nil, fmt.Errorf(InvalidSubnetSize+Details, subnetSize)
	}

	ip := ipNet.IP.Mask(ipNet.Mask)

	for ipNet.Contains(ip) {
		subnet := &net.IPNet{
			IP:   ip.Mask(net.CIDRMask(newMaskSize, 32)),
			Mask: net.CIDRMask(newMaskSize, 32),
		}
		subnets = append(subnets, subnet.String())
		ip = nextSubnetIP(ip, newMaskSize)
	}

	return subnets, nil
}

// nextSubnetIP calculates the next subnet IP address based on the subnet size.
func nextSubnetIP(ip net.IP, maskSize int) net.IP {
	ip = append(ip[:0:0], ip...)
	increment := 1 << (32 - maskSize)
	for i := len(ip) - 1; i >= 0 && increment > 0; i-- {
		value := int(ip[i]) + increment
		ip[i] = byte(value % 256)
		increment = value / 256
	}
	return ip
}
