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
	Digest  digest
	Version int64
}

type hisTree struct {
	version int64
	nodeAt  map[pos]digest
	h       hash.Hash
}

// NewHisTree returns a new (emtpy) hisTree
func NewHisTree() *hisTree {
	return &hisTree{
		version: -1,
		nodeAt:  make(map[pos]digest),
		h:       sha256.New(),
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

// https://play.golang.org/p/HwOxHod5okU
func (ht *hisTree) getHeight() int64 {
	return int64(
		math.Ceil(
			math.Log2(
				float64(
					ht.version + 1,
				),
			),
		),
	)
}

func (ht *hisTree) Add(e *Event) *commitment {
	ht.bumpVersion()
	fmt.Printf("hisTree: '%v'\n", ht)
	rootPos := pos{
		i: 0,
		r: ht.getHeight(),
	}
	fmt.Printf("rootPos: '%v'\n", rootPos)
	ht.h.Write(e.Value)
	d := ht.h.Sum(nil)
	ht.add(
		&digest{
			algo:  "sha256",
			value: d,
		},
		&rootPos,
	)
	// fmt.Printf("Dump tree '%v'\n", ht.nodeAt)
	for k, v := range ht.nodeAt {
		fmt.Printf("Key: '%v', Value: '%v'\n", k, v)
	}
	return &commitment{
		Digest: digest{
			algo:  "sha256",
			value: ht.nodeAt[rootPos].value,
		},
		Version: ht.version,
	}
}

func (ht *hisTree) add(ed *digest, p *pos) {
	fmt.Println("add: starting...")
	if p.isLeaf() {
		fmt.Println("add: isLeaf")
		ht.nodeAt[*p] = *ed
		return
	} else if ht.version <= p.i {
		fmt.Println("add: left!")
		ht.add(ed, p.left())
	} else {
		fmt.Println("add: right!")
		ht.add(ed, p.right())
	}

	// lv := append([]byte(nil), ht.nodeAt[*p.left()].value...)
	lv := make([]byte, len(ht.nodeAt[*p.left()].value))
	copy(lv, ht.nodeAt[*p.left()].value)
	ht.h.Write(
		append(
			lv,
			ht.nodeAt[*p.right()].value...,
		),
	)
	d := ht.h.Sum(nil)
	ht.nodeAt[*p] = digest{
		algo:  "sha256",
		value: d,
	}
}

// https://en.wikipedia.org/wiki/Tree_traversal
// https://en.wikipedia.org/wiki/Binary_search_tree
