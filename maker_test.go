package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"image/png"
	"os"
	"testing"
)

func TestMaker_Save(t *testing.T) {
	m, err := NewMaker("./static/data")
	assert.NoError(t, err)
	g1Str := `
{
  "antennas": "000-000-000-000",
  "arm": "000-000-000-000",
  "body": "000-000-000-000",
  "brows": "000-000-000-000",
  "claws": "000-000-000-000",
  "eyes": "000-000-000-000",
  "glasses":  "000-000-000-000",
  "head": "000-000-000-000"
}`

	//"{\"antennas\": \"000-000-000-000\",\"arm": \"000-000-000-000\",\"body": \"000-000-000-000\",\"brows": \"000-000-000-000\",\"claws": \"000-000-000-000\",\"eyes": \"000-000-000-000\",\"glasses":  \"000-000-000-000\",\"head": \"000-000-000-000\"}"
	var g1 *Genome
	err = json.Unmarshal([]byte(g1Str), &g1)
	assert.NoError(t, err)
	for _, v := range g1.chromosomes {
		assert.Equal(t, "000-000-000-000", v.String())
	}

	g2Str := `
{
  "antennas": "111-111-111-111",
  "arm": "111-111-111-111",
  "body": "111-111-111-111",
  "brows": "111-111-111-111",
  "claws": "111-111-111-111",
  "eyes": "111-111-111-111",
  "glasses":  "111-111-111-111",
  "head": "111-111-111-111"
}`
	var g2 *Genome
	err = json.Unmarshal([]byte(g2Str), &g2)
	assert.NoError(t, err)
	for _, v := range g2.chromosomes {
		assert.Equal(t, "111-111-111-111", v.String())
	}

	g3 := Breed(g1, g2)

	img, err := m.Make(g3)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, img)

	name := ""
	for k, v := range g2.chromosomes {
		name += k + "_" + v.String()
		assert.Equal(t, "111-111-111-111", v.String())
	}
	name += ".png"
	out, err := os.Create(name)
	assert.NoError(t, err)

	err = png.Encode(out, img)
	assert.NoError(t, err)
}
