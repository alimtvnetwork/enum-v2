package messages

import (
	"github.com/alimtvnetwork/core-v9/corecomparator"
	"github.com/alimtvnetwork/core-v9/errcore"
)

var (
	ComparatorOutOfRangeMessage = errcore.RangeNotMeet(
		errcore.ComparatorShouldBeWithinRangeType.String(),
		corecomparator.Min(),
		corecomparator.Max(),
		corecomparator.Ranges())
)
