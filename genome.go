package main

import (
	"encoding/json"
	"errors"
	"math/rand"
	"strconv"
	"strings"
)

const ChromosomeLength = 4
const GeneLength = 3

type Gene [GeneLength]int
type Chromosome struct {
	dat [ChromosomeLength]Gene
}

func (c *Chromosome) Squeeze() string {
	ans := ""
	for i := range c.dat {
		tmp := 0
		for _, j := range c.dat[i] {
			tmp = tmp&0x01 ^ j
		}
		ans = ans + "_" + strconv.Itoa(tmp)
	}

	return ans[1:]
}

func IsSplit(i int) bool {
	return (i+1)%(GeneLength+1) == 0
}

func validateChromosome(s string) error {
	if len(s) != ChromosomeLength*GeneLength+ChromosomeLength-1 {
		return errors.New("length not 12")
	}

	for i, c := range s {
		if IsSplit(i) {
			if c != '-' {
				return errors.New("error split")
			}
		} else {
			if !(c == '1' || c == '0') {
				return errors.New("not valid symbols in gene")
			}
		}

	}

	return nil
}

func NewChromosomeFromString(s string) (*Chromosome, error) {
	err := validateChromosome(s)
	if err != nil {
		return nil, err
	}

	cr := Chromosome{}
	k := 0
	for i := 0; i < len(cr.dat); i++ {
		for j := 0; j < len(cr.dat[i]); j++ {

			cr.dat[i][j] = int(s[k] - '0')
			k++
			if IsSplit(k) {
				k++
			}
		}
	}

	return &cr, nil
}

func (c *Chromosome) MarshalJSON() ([]byte, error) {
	return []byte("\"" + c.String() + "\""), nil
}

func (c *Chromosome) UnmarshalJSON(data []byte) error {
	rawStr := string(data)
	ch, err := NewChromosomeFromString(strings.Trim(rawStr, "\""))
	if err != nil {
		return err
	}
	c.dat = ch.dat
	return nil
}

func (c *Chromosome) String() string {
	ans := ""
	for _, g := range c.dat {
		for _, i := range g {
			ans += strconv.Itoa(i)
		}
		ans += "-"
	}

	return ans[:len(ans)-1]
}

type Genome struct {
	chromosomes map[string]*Chromosome

	//Beard       *Chromosome `json:"beard"`
	//Hat         *Chromosome `json:"hat"`
	//Hair        *Chromosome `json:"hair"`
	//Eyes        *Chromosome `json:"eyes"`
	//Glasses     *Chromosome `json:"glasses"`
	//Body        *Chromosome `json:"body"`
	//Hand        *Chromosome `json:"hand"`
	//Antennas    *Chromosome `json:"antennas"`
	//Accessories *Chromosome `json:"accessories"`
}

func (g *Genome) MarshalJSON() ([]byte, error) {
	a, err := json.Marshal(g.chromosomes)
	return a, err
}

func (g *Genome) UnmarshalJSON(data []byte) error {
	var tmp map[string]*Chromosome
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	g.chromosomes = tmp
	return nil
}

func crossover(c1 *Chromosome, c2 *Chromosome) Chromosome {
	split := rand.Intn(ChromosomeLength-1) + 1 // [1, ChromosomeLength-1]

	var left *Chromosome
	var right *Chromosome

	coin := rand.Intn(2)
	if coin == 0 {
		left = c1
		right = c2
	} else {
		left = c2
		right = c1
	}

	var d [4]Gene
	for i := range d {
		if i < split {
			d[i] = left.dat[i]
		} else {
			d[i] = right.dat[i]
		}
	}

	return Chromosome{dat: d}
}

const MaxMutates = 8

func swap(i int) int {
	if i == 0 {
		return 1
	}
	return 0
}

func mutate(c1 *Chromosome) Chromosome {
	newCh := *c1
	nMutate := rand.Intn(MaxMutates + 1)
	for i := 0; i < nMutate; i++ {
		j := rand.Intn(ChromosomeLength)
		k := rand.Intn(GeneLength)
		newCh.dat[j][k] = swap(newCh.dat[j][k])
	}

	return newCh
}

func Breed(g1 *Genome, g2 *Genome) *Genome {
	ng := Genome{chromosomes: map[string]*Chromosome{}}
	for k := range g1.chromosomes {
		newCh := crossover(g1.chromosomes[k], g2.chromosomes[k])
		newCh = mutate(&newCh)
		ng.chromosomes[k] = &newCh
	}

	return &ng
}
