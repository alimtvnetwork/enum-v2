package iptype

func CreateTypeUsingExplicit(name string) Variant {
	isIpv4 := name == ipv4 ||
		name == ipv4Alternate
	if isIpv4 {
		return V4
	}

	return V6
}
