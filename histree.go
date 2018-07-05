package histree

import "crypto/sha256"

type pos struct {
	i, r uint64
}

func (p pos) isLeaf() bool {
	return p.r == 0
}

type hash struct {
	algo  string
	value [sha256.Size]byte
}

type hisTree struct {
	version uint64
	nodes   map[pos]hash
	hasher  func(data []byte) [sha256.Size]byte
}

// NewHisTree returns a new (emtpy) hisTree
func NewHisTree() *hisTree {
	return &hisTree{
		version: 0,
		nodes:   make(map[pos]hash),
		hasher:  sha256.New(),
	}
}

func (ht *hisTree) Add() hash {

}
