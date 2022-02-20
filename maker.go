package main

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"os"
)

var DepthMap = []string{
	"body",
	"head",
	"eyes",
	"brows",
	"glasses",
	"antennas",
	"claws",
	"arm",
}

type Maker struct {
	baseDir string
}

func (m *Maker) validate() error {
	if _, err := os.Stat(m.baseDir); os.IsNotExist(err) {
		return err
	}

	return nil
}

func NewMaker(baseDir string) (*Maker, error) {
	m := &Maker{baseDir}
	err := m.validate()
	if err != nil {
		return nil, err
	}

	return m, nil
}

//set image offset

func (m *Maker) Make(g *Genome) (image.Image, error) {
	tmp := image.NewRGBA(image.Rect(0, 0, 1024, 1024))
	for _, k := range DepthMap {
		c, ok := g.chromosomes[k]
		if !ok {
			return nil, errors.New(fmt.Sprintf("no chromosome in preset %s", k))
		}

		r, err := os.Open(m.baseDir + "/" + k + "/" + c.Squeeze() + ".png")
		if err != nil {
			return nil, err
		}

		img, _, err := image.Decode(r)
		if err != nil {
			return nil, err
		}

		err = r.Close()
		if err != nil {
			return nil, err
		}

		draw.Draw(tmp, img.Bounds(), img, image.Point{}, draw.Over)
	}

	return tmp, nil
}
