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

func (p *pos) left() *pos {
	if p.isLeaf() {
		log.Fatal("ERROR: Cannot go any deeper!")
	}

	return &pos{
		i: p.i,
		r: p.r - 1,
	}
}

func (p *pos) right() *pos {
	return &pos{
		i: p.i + int64(math.Pow(2, float64(p.r-1))),
		r: p.r - 1,
	}
}

func (p *pos) sortKey() int64 {
	return p.i + int64(math.Pow(2, float64(p.r-1)))
}
