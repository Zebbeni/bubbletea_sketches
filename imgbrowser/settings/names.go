package settings

import "github.com/nfnt/resize"

var interpolationNames = map[resize.InterpolationFunction]string{
	resize.NearestNeighbor:   "Nearest Neighbor",
	resize.Bilinear:          "Bilinear",
	resize.Bicubic:           "Bicubic",
	resize.MitchellNetravali: "MitchellNetravali",
	resize.Lanczos2:          "Lanczos2",
	resize.Lanczos3:          "Lanczos3",
}
