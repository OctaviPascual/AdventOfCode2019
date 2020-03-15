package day08

import (
	"errors"
	"fmt"
	"html"
	"math"
	"strconv"
	"strings"
)

const (
	black       = 0
	white       = 1
	transparent = 2
)

var (
	imageWidth  = 25
	imageHeight = 6

	blackEmoji = html.UnescapeString("&#" + strconv.Itoa(11035) + ";")
	whiteEmoji = html.UnescapeString("&#" + strconv.Itoa(11036) + ";")
)

type encodedImage struct {
	width  int
	height int
	layers []layer
}

type decodedImage struct {
	width  int
	height int
	pixels []int
}

type layer struct {
	pixels []int
}

// Day holds the data needed to solve part one and part two
type Day struct {
	image *encodedImage
}

// NewDay returns a new Day that solves part one and two for the given input
func NewDay(input string) (*Day, error) {
	image, err := parseImage(input, imageWidth, imageHeight)
	if err != nil {
		return nil, fmt.Errorf("could not parse encodedImage %s: %w", input, err)
	}

	return &Day{
		image: image,
	}, nil
}

// SolvePartOne solves part one
func (d Day) SolvePartOne() (string, error) {
	layer := fewestZeroDigits(d.image.layers)
	numberOfOnes := layer.numberOfDigits(1)
	numberOfTwos := layer.numberOfDigits(2)
	return fmt.Sprintf("%d", numberOfOnes*numberOfTwos), nil
}

// SolvePartTwo solves part two
func (d Day) SolvePartTwo() (string, error) {
	decodedImage := d.image.decode()
	return decodedImage.render()
}

func parseImage(pixels string, width, height int) (*encodedImage, error) {
	size := width * height

	if len(pixels)%size != 0 {
		return nil, fmt.Errorf("invalid pixels length: %d (should be multiple of %d)", len(pixels), size)
	}

	lenLayers := len(pixels) / size
	layers := make([]layer, 0, lenLayers)
	for i := 0; i < lenLayers; i++ {
		start := i * size
		end := start + size
		layer, err := parseLayer(pixels[start:end])
		if err != nil {
			return nil, fmt.Errorf("could not parse layer %s: %w", pixels[start:end], err)
		}

		layers = append(layers, *layer)
	}
	return &encodedImage{
		width:  width,
		height: height,
		layers: layers,
	}, nil
}

func parseLayer(pixels string) (*layer, error) {
	layer := &layer{
		pixels: make([]int, 0, len(pixels)),
	}
	for _, pixelRune := range pixels {

		pixel, err := strconv.Atoi(string(pixelRune))
		if err != nil {
			return nil, fmt.Errorf("invalid pixel %c: %w", pixelRune, err)
		}

		layer.pixels = append(layer.pixels, pixel)
	}
	return layer, nil
}

func fewestZeroDigits(layers []layer) layer {
	minNumberOfZeros := math.MaxInt32
	var result layer
	for _, layer := range layers {
		numberOfZeros := layer.numberOfDigits(0)
		if numberOfZeros < minNumberOfZeros {
			result = layer
			minNumberOfZeros = numberOfZeros
		}
	}
	return result
}

func (l layer) numberOfDigits(digit int) int {
	total := 0
	for _, pixel := range l.pixels {
		if pixel == digit {
			total++
		}
	}
	return total
}

func (ei *encodedImage) decode() *decodedImage {
	size := ei.width * ei.height

	decodedImage := &decodedImage{
		width:  ei.width,
		height: ei.height,
		pixels: make([]int, size),
	}

	for i := 0; i < size; i++ {
		for _, layer := range ei.layers {
			if layer.pixels[i] != transparent {
				decodedImage.pixels[i] = layer.pixels[i]
				break
			}
		}
	}
	return decodedImage
}

func (di *decodedImage) render() (string, error) {
	var sb strings.Builder
	sb.WriteString("\n")
	for i := 0; i < di.height; i++ {
		for j := 0; j < di.width; j++ {
			pixel := di.pixels[i*di.width+j]
			switch pixel {
			case black:
				sb.WriteString(blackEmoji)
			case white:
				sb.WriteString(whiteEmoji)
			default:
				return "", errors.New("cannot render transparent pixel")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String(), nil
}
