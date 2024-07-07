package calculateSubnets

import (
	"reflect"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		parentCIDR string
		subnetSize uint8
		expected   []string
	}{
		{
			parentCIDR: "10.1.0.0/28",
			subnetSize: 29,
			expected: []string{
				"10.1.0.0/29",
				"10.1.0.8/29",
			},
		},
		{
			parentCIDR: "10.1.0.0/27",
			subnetSize: 28,
			expected: []string{
				"10.1.0.0/28",
				"10.1.0.16/28",
			},
		},
		{
			parentCIDR: "192.168.0.0/24",
			subnetSize: 25,
			expected: []string{
				"192.168.0.0/25",
				"192.168.0.128/25",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.parentCIDR, func(t *testing.T) {
			subnets, err := Calculate(tt.parentCIDR, tt.subnetSize)
			if err != nil {
				t.Fatalf("Calculate(%q, %d) returned error: %v", tt.parentCIDR, tt.subnetSize, err)
			}
			if !reflect.DeepEqual(subnets, tt.expected) {
				t.Errorf("Calculate(%q, %d) = %v, want %v", tt.parentCIDR, tt.subnetSize, subnets, tt.expected)
			}
		})
	}
}
