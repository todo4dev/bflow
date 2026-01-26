package image

import (
	"io"
)

// Format defines the output formats supported by the domain
type Format string

const (
	PNG  Format = "png"
	JPEG Format = "jpeg"
	WEBP Format = "webp"
)

type Processor interface {
	// Resize resizes the image to the specified dimensions and output format
	Resize(
		input io.Reader,
		width int,
		height int,
		out Format,
	) (io.Reader, error)
}
