package histree

import (
	"log"
	"math"
)

type pos struct {
	i, r int64
}

func (p *pos) isLeaf() bool {
	return p.r == 0
}

func (p *pos) leftSon() *pos {
	if p.isLeaf() {
		log.Fatal("ERROR: Cannot go any deeper!")
	}

	return &pos{
		i: p.i,
		r: p.r - 1,
	}
}

func (p *pos) rightSon() *pos {
	return &pos{
		i: p.i,
		r: p.i + int64(math.Pow(2, float64(p.r-1))),
	}
}
