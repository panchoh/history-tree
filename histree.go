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

type commitment struct {
	hash
}

type hisTree struct {
	version     uint64
	height      uint64
	rootPos     pos
	nextRootPos pos
	nodes       map[pos]hash
	hasher      func(data []byte) [sha256.Size]byte
}

// NewHisTree returns a new (emtpy) hisTree
func NewHisTree() *hisTree {
	return &hisTree{
		version: -1,
		height:  1,
		rootPos: {
			i: 0,
			r: 1,
		},
		nextRootPos: {
			i: 0,
			r: 2,
		},
		nodes:  make(map[pos]hash),
		hasher: sha256.New(),
	}
}

func (ht *hisTree) bumpVersion() {
	ht.version++
}

func (ht *hisTree) GetHasher() func(data []byte) [sha256.Size]byte {
	return ht.hasher
}

func (ht *hisTree) Add(e *Event) (commitment, err) {
	//	Identify current root
}

func (ht *hisTree) add(e *Event) (commitment, err) {

}

// https://en.wikipedia.org/wiki/Tree_traversal
