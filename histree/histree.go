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
	hashFunc
}

type hashFunc struct {
	h    hash.Hash
	algo string
}

// NewHisTree returns a new (empty) hisTree
func NewHisTree() *hisTree {
	return &hisTree{
		version: -1,
		nodeAt:  make(map[pos]digest),
		hashFunc: hashFunc{
			h:    sha256.New(),
			algo: "SHA-256",
		},
	}
}

func (ht *hisTree) bumpVersion() {
	ht.version++
}

func (ht *hisTree) GetHashFun() hash.Hash {
	return ht.hashFunc.h
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
	n, err := ht.hashFunc.h.Write(e.Value)
	if err != nil {
		fmt.Println("ERROR: Add: hash.Hasher() returned error", err)
	}
	fmt.Printf("Add: hash.Hash() returned '%d' bytes.\n", n)
	d := ht.hashFunc.h.Sum(nil)
	ht.add(
		&digest{
			algo:  ht.hashFunc.algo,
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
			algo:  ht.hashFunc.algo,
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

	ht.rehash(p)
}

func (ht *hisTree) rehash(p *pos) {
	// Should never happen
	if p.isLeaf() {
		fmt.Println("rehash: Cannot rehash a leaf!")
		return
	}

	// lv := append([]byte(nil), ht.nodeAt[*p.left()].value...)
	lv := make([]byte, len(ht.nodeAt[*p.left()].value))
	copy(lv, ht.nodeAt[*p.left()].value)
	n, err := ht.hashFunc.h.Write(
		append(
			lv,
			ht.nodeAt[*p.right()].value...,
		),
	)
	fmt.Printf("rehash: hash.Hasher.Write() returned '%d' bytes.\n", n)
	if err != nil {
		fmt.Println("ERROR: hash.Hash.Write() failed with error", err)
	}
	d := ht.hashFunc.h.Sum(nil)
	ht.nodeAt[*p] = digest{
		algo:  ht.hashFunc.algo,
		value: d,
	}
}

// TODO: Ensure we are using hash.Hash correctly (when is Reset() needed?)
// https://en.wikipedia.org/wiki/Tree_traversal
// https://en.wikipedia.org/wiki/Binary_search_tree
