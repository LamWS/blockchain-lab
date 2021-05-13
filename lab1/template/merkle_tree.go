package main

import "crypto/sha256"

// MerkleTree represent a Merkle tree
type MerkleTree struct {
	RootNode *MerkleNode
}

// MerkleNode represent a Merkle tree node
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

// NewMerkleTree creates a new Merkle tree from a sequence of data
// implement
func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode
	for i := 0; i < len(data); i++ {
		node := *NewMerkleNode(nil, nil, data[i])
		nodes = append(nodes, node)
	}
	for len(nodes) != 1 {
		if len(nodes)%2 != 0 {
			nodes = append(nodes, nodes[len(nodes)-1])
		}
		var newNodes []MerkleNode
		for i := 0; i < len(nodes); i += 2 {
			node := NewMerkleNode(&nodes[i], &nodes[i+1], nil)
			newNodes = append(newNodes, *node)
		}
		nodes = newNodes
	}
	return &MerkleTree{&nodes[0]}
}

func NewMerkleNode(left *MerkleNode, right *MerkleNode, data []byte) *MerkleNode {
	node := MerkleNode{left, right, data}
	if node.Left == nil && node.Right == nil {
		hash := sha256.Sum256(data)
		node.Data = hash[:]
	} else {
		hash := sha256.Sum256(append(node.Left.Data, node.Right.Data...))
		node.Data = hash[:]
	}
	return &node
}
