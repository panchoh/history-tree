package histree

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"math"
)

type digest struct {
	algo  string
	value []byte
}

type commitment struct {
	digest
}

type hisTree struct {
	version     int64
	height      int64
	rootPos     pos
	nextRootPos pos
	nodeAt      map[pos]digest
	h           hash.Hash
}

// NewHisTree returns a new (emtpy) hisTree
func NewHisTree() *hisTree {
	return &hisTree{
		version: -1,
		height:  0,
		rootPos: pos{
			i: 0,
			r: 0,
		},
		nextRootPos: pos{
			i: 0,
			r: 1,
		},
		nodeAt: make(map[pos]digest),
		h:      sha256.New(),
	}
}

func (ht *hisTree) bumpVersion() {
	ht.version++
}

func (ht *hisTree) GetHash() hash.Hash {
	return ht.h
}

type Event struct {
	Value []byte
}

// https://play.golang.org/p/pwfSvHSebzR
func (ht *hisTree) getHeight() int64 {
	return 1 + int64(
		math.Ceil(
			math.Log2(
				float64(
					ht.height+1,
				),
			),
		),
	)
}

func (ht *hisTree) Add(e *Event) *commitment {
	//	TODO: Identify current root
	ht.bumpVersion()
	ht.rootPos = pos{
		i: 0,
		r: ht.getHeight(),
	}
	ht.h.Write(e.Value)
	d := ht.h.Sum(nil)
	ht.add(
		&digest{
			algo:  "sha256",
			value: d,
		},
		&ht.rootPos,
	)
	fmt.Printf("Dump tree '%v'\n", ht.nodeAt)
	return &commitment{
		digest{
			algo:  "sha256",
			value: ht.nodeAt[ht.rootPos].value,
		},
	}
}

func (ht *hisTree) add(ed *digest, p *pos) {
	if p.isLeaf() {
		ht.nodeAt[*p] = *ed
		return
	} else if ht.version <= p.i {
		ht.add(ed, p.leftSon())
	} else {
		ht.add(ed, p.rightSon())
	}
	ht.h.Write(
		append(
			ht.nodeAt[*p.leftSon()].value,
			ht.nodeAt[*p.rightSon()].value...,
		),
	)
	d := ht.h.Sum(nil)
	ht.nodeAt[*p] = digest{
		algo:  "sha256",
		value: d,
	}
}

func cat(l, r []byte) (buf []byte) {
	for i, v := range l {
		buf[i] = v
	}
	offset := len(l)
	for i, v := range r {
		buf[offset+i] = v
	}
	return
}

// https://en.wikipedia.org/wiki/Tree_traversal
// https://en.wikipedia.org/wiki/Binary_search_tree
