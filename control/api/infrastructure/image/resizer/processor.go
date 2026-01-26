// infrastructure/image/resizer/processor.go
package resizer

import (
	"bytes"
	"fmt"
	"io"

	port "src/port/image"

	"github.com/disintegration/imaging"
)

type Processor struct{}

var _ port.Processor = (*Processor)(nil)

func New() *Processor {
	return &Processor{}
}

func (p *Processor) Resize(
	input io.Reader,
	width int,
	height int,
	out port.Format,
) (io.Reader, error) {
	src, err := imaging.Decode(input)
	if err != nil {
		return nil, fmt.Errorf("error decoding image: %w", err)
	}

	dst := imaging.Fit(src, width, height, imaging.Lanczos)

	buf := new(bytes.Buffer)
	var format imaging.Format

	switch out {
	case port.PNG:
		format = imaging.PNG
	case port.JPEG:
		format = imaging.JPEG
	default:
		format = imaging.PNG
	}

	err = imaging.Encode(buf, dst, format)
	if err != nil {
		return nil, fmt.Errorf("error encoding image to %s: %w", out, err)
	}

	return buf, nil
}
