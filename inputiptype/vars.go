package inputiptype

import (
	"github.com/alimtvnetwork/core-v9/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:      "Invalid",
		Ip:           "Ip",
		IpWithSubnet: "IpWithSubnet",
		SubnetMask:   "SubnetMask",
		Gateway:      "Gateway",
		IpWithPort:   "IpWithPort",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.DefaultAllCases(
		Invalid,
		Ranges[:],
	)
)
