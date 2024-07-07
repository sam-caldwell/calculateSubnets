Subnetting Tools
================

## Description

This command will calculate a list of subnet networks within the parent CIDR block.

## Commands

**Syntax:** `calculateSubnet ${parentCIDR} ${subnetSize}`

* parentCidr = CIDR string (e.g. 10.11.0.0/16)
* subnetSize = size of the subnet (e.g. integer 0-32)
  but value must be within parent subnet.
