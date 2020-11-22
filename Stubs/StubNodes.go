package Stubs

import (
	c "../Contracts"
)

func GetStubNodes() []c.Node {
	return []c.Node{
		c.Node{Name: "N1"},
		c.Node{Name: "N2"},
	}
}
