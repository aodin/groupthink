package reddit

import (
	"time"
)

// Reddit seems to pack their timestamps as floats...
type Timestamp float64

func (t Timestamp) AsTime() time.Time {
	// TODO Does the UTC change anything?
	return time.Unix(int64(t), 0).UTC()
}
