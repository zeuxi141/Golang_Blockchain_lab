// Block trong block.go
package blockchain

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	index   int 
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	MerkleRoot    []byte
}

type Transaction struct {
	Data []byte
}

// SetHash tính toán và lưu hash của khối.
// Hash của khối được tạo bằng cách nối các thông tin sau:
//   - Hash của khối trước đó.
//   - Dữ liệu của các giao dịch trong khối.
//   - Thời gian tạo khối.
// Hash được tạo bằng thuật toán SHA-256.
// Hash được lưu vào trường Hash của khối.
// Trường PrevBlockHash của khối tiếp theo sẽ trỏ đến hash của khối hiện tại.
// Trả về:
//   - error: lỗi nếu có.
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	var transactionsData [][]byte
	for _, tx := range b.Transactions {
		transactionsData = append(transactionsData, tx.Data)
	}

	// Nối các mảng byte lại với nhau
	headers := bytes.Join([][]byte{b.PrevBlockHash, bytes.Join(transactionsData, []byte{}), timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
// NewBlock tạo một khối mới với thông tin cơ bản.
// Tham số:
//   - index: số thứ tự của khối trong blockchain.
//   - transactions: danh sách các giao dịch trong khối.
//   - prevBlockHash: hash của khối trước đó.
// Trả về:
//   - con trỏ đến khối mới tạo.
func NewBlock(index int, transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{index, time.Now().Unix(), transactions, prevBlockHash, []byte{}, []byte{}}
	block.SetHash()
	block.MerkleRoot = block.CalculateMerkleRoot()
	return block
}


// CalculateMerkleRoot tính toán Merkle Root của các giao dịch trong khối.
func (b *Block) CalculateMerkleRoot() []byte {
	var transactionsData [][]byte
	for _, tx := range b.Transactions {
		transactionsData = append(transactionsData, tx.Data)
	}
	merkleTree := NewMerkleTree(transactionsData)
	return merkleTree.RootNode.Data
}


// UpdateTransactionData cập nhật dữ liệu giao dịch mới cho một giao dịch trong khối.
// Tham số:
//   - blockchain: con trỏ đến blockchain.
//   - index: số thứ tự của khối trong blockchain.
//   - transactionIndex: số thứ tự của giao dịch trong khối.
//   - newTransactionData: dữ liệu giao dịch mới.
// Trả về:
//   - error: lỗi nếu có.
func UpdateTransactionData(blockchain *Blockchain, index, transactionIndex int, newTransactionData string) error {
	if index < 0 || index >= len(blockchain.Blocks) {
		return errors.New("invalid block number")
	}
	if transactionIndex < 0 || transactionIndex >= len(blockchain.Blocks[index].Transactions) {
		return errors.New("invalid transaction index")
	}

	blockCopy := *blockchain.Blocks[index]
	blockCopy.Transactions[transactionIndex].Data = []byte(newTransactionData)
	blockchain.Blocks[index] = &blockCopy
	return nil
}

// PrintBlock in ra thông tin của một khối.
func PrintBlock(block *Block){
	fmt.Printf("Block Number: %d\n", block.index)
	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	fmt.Println("Transactions:")
	for _, tx := range block.Transactions {
		fmt.Printf("  - %s\n", string(tx.Data))
	}
	fmt.Printf("Hash: %x\n", block.Hash)
	fmt.Printf("Timestamp: %d\n", block.Timestamp)
	fmt.Printf("MerkleRoot: %x\n", block.MerkleRoot)
}
	