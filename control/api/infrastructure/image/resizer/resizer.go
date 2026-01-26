// infrastructure/image/resizer/resizer.go
package resizer

import (
	"src/port/image"

	"github.com/leandroluk/gox/di"
)

func Provide() {
	di.SingletonAs[image.Processor](New)
}
