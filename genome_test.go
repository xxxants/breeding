package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChromosome_String(t *testing.T) {
	chStr := "000-000-000-000"
	ch, err := NewChromosomeFromString(chStr)
	exp := Chromosome{dat: [4]Gene{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, 0}}}
	newStr := exp.String()
	assert.NoError(t, err)
	assert.Equal(t, &exp, ch)
	assert.Equal(t, chStr, newStr)

	chStr = "111-111-111-111"
	ch, err = NewChromosomeFromString(chStr)
	exp = Chromosome{dat: [4]Gene{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}}}
	newStr = exp.String()
	assert.NoError(t, err)
	assert.Equal(t, ch, &exp)
	assert.Equal(t, chStr, newStr)

	chStr = "101-010-000-011"
	ch, err = NewChromosomeFromString(chStr)
	exp = Chromosome{dat: [4]Gene{{1, 0, 1}, {0, 1, 0}, {0, 0, 0}, {0, 1, 1}}}
	newStr = exp.String()
	assert.NoError(t, err)
	assert.Equal(t, ch, &exp)
	assert.Equal(t, chStr, newStr)
}

func TestChromosome_Squeeze(t *testing.T) {
	chStr := "000-011-111-100"
	ch, err := NewChromosomeFromString(chStr)
	assert.NoError(t, err)
	assert.Equal(t, "0_0_1_1", ch.Squeeze())
}

func TestGenomeChromosome_JSON(t *testing.T) {
	type testGenome struct {
		Test1 *Chromosome `json:"test1"`
		Test2 *Chromosome `json:"test2"`
	}
	ch, _ := NewChromosomeFromString("101-010-000-011")
	tc := testGenome{Test1: ch, Test2: ch}

	a, err := json.Marshal(tc)
	assert.NoError(t, err)
	assert.Equal(t, string(a), "{\"test1\":\"101-010-000-011\",\"test2\":\"101-010-000-011\"}")

	var mTc testGenome
	err = json.Unmarshal(a, &mTc)
	assert.NoError(t, err)
	assert.Equal(t, tc, mTc)
}

func TestGenome_Breed(t *testing.T) {
	g1Str := `
{
  "beard": "000-000-000-000",
  "hat":  "000-000-000-000",
  "hair": "000-000-000-000",
  "eyes":  "000-000-000-000",
  "glasses": "000-000-000-000",
  "body": "000-000-000-000",
  "hand": "000-000-000-000",
  "antennas": "000-000-000-000",
  "accessories": "000-000-000-000"
}`
	var g1 *Genome
	err := json.Unmarshal([]byte(g1Str), &g1)
	assert.NoError(t, err)
	for _, v := range g1.chromosomes {
		assert.Equal(t, "000-000-000-000", v.String())
	}

	g2Str := `
{
  "beard": "111-111-111-111",
  "hat":  "111-111-111-111",
  "hair": "111-111-111-111",
  "eyes":  "111-111-111-111",
  "glasses": "111-111-111-111",
  "body": "111-111-111-111",
  "hand": "111-111-111-111",
  "antennas": "111-111-111-111",
  "accessories": "111-111-111-111"
}`
	var g2 *Genome
	err = json.Unmarshal([]byte(g2Str), &g2)
	assert.NoError(t, err)
	for _, v := range g2.chromosomes {
		assert.Equal(t, "111-111-111-111", v.String())
	}

	g3 := Breed(g1, g2)

	j, err := json.Marshal(g3)
	assert.NoError(t, err)

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, j, "", "\t")
	assert.NoError(t, err)

	fmt.Println(string(prettyJSON.Bytes()))
}
