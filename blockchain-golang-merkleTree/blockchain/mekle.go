package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type MerkleTree struct {
	RootNode *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

// NewMerkleNode tạo một nút mới trong cây Merkle.
// Nếu nút là lá, dữ liệu của nút sẽ được hash.
// Nếu nút không phải là lá, dữ liệu của nút sẽ là hash của dữ liệu của 2 nút con.
// Tham số:
//   - left: nút con bên trái.
//   - right: nút con bên phải.
//   - data: dữ liệu của nút.
// Trả về:
//   - con trỏ đến nút mới tạo.
func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	node := MerkleNode{}

	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Data = hash[:]
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		node.Data = hash[:]
	}

	node.Left = left
	node.Right = right

	return &node
}

// NewMerkleTree tạo một cây Merkle từ dữ liệu giao dịch.
// Cây Merkle được tạo bằng cách tạo các nút từ dữ liệu giao dịch.
// Dữ liệu giao dịch được hash và lưu vào nút lá.
// Nút không phải là lá sẽ lưu hash của dữ liệu của 2 nút con.
// Tham số:
//   - data: danh sách dữ liệu giao dịch.
// Trả về:
//   - con trỏ đến cây Merkle mới tạo.
func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode
	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	for _, dat := range data {
		node := NewMerkleNode(nil, nil, dat)
		nodes = append(nodes, *node)
	}

	for i := 0; i < len(data)/2; i++ {
		var level []MerkleNode

		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			level = append(level, *node)
		}

		nodes = level
	}
	tree := MerkleTree{&nodes[0]}

	return &tree
}

// PrintMerkleTree in ra cây Merkle.
// Tham số:
//   - block: khối chứa cây Merkle.
// Cây Merkle được in ra từ nút gốc.
// Mỗi nút được in ra với dữ liệu của nó.
// Nút con bên trái được in ra trước nút con bên phải.
// Các nút con được in ra với khoảng cách lớn hơn so với nút cha.
// Các nút con được in ra với dấu hiệu └─ hoặc ├─ tùy thuộc vào vị trí của nút con.
// Nếu nút cha là nút cuối cùng của một nhánh, dấu hiệu └─ sẽ được in ra.
// Nếu nút cha không phải là nút cuối cùng của một nhánh, dấu hiệu ├─ sẽ được in ra.
// Nếu nút cha không có nút con, không có dấu hiệu nào được in ra.
func printMerkleNode(node *MerkleNode, level int, isLeft bool, _ bool) {
	if node == nil {
		return
	}

	indent := strings.Repeat("  ", level)

	fmt.Printf("%s", indent)
	if isLeft {
		fmt.Print("├─ ")
	} else {
		fmt.Print("└─ ")
	}
	fmt.Printf("%s\n", hex.EncodeToString(node.Data))

	if node.Left != nil || node.Right != nil {
		printMerkleNode(node.Left, level+1, true, false)
		printMerkleNode(node.Right, level+1, false, false)
	}
}

func PrintMerkleTree(block *Block) {
	transactions := make([][]byte, len(block.Transactions))
	for i, tx := range block.Transactions {
		transactions[i] = tx.Data
	}

	tree := NewMerkleTree(transactions)
	fmt.Println("Merkle Tree:")
	printMerkleNode(tree.RootNode, 0, true, true)
}




