package mapcommon

import (
	"fmt"
)

type MapStringer interface {
	comparable
	fmt.Stringer
}
