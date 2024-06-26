
package blockchain

type Blockchain struct {
	Blocks []*Block
}

// NewGenesisBlock tạo một khối Genesis với dữ liệu khởi tạo.
func NewGenesisBlock(initData string) *Block {
	initTransaction := &Transaction{Data: []byte(initData)}
	return NewBlock(0, []*Transaction{initTransaction}, []byte{})
}

// NewBlockchain tạo một blockchain mới với khối Genesis.
func NewBlockchain() *Blockchain {
	initData := "Init Genesis Block"
	genesisBlock := NewGenesisBlock(initData)
	return &Blockchain{[]*Block{genesisBlock}}
}

// AddBlock thêm một khối mới vào blockchain.
// Tham số:
//   - transactions: danh sách các giao dịch trong khối.
// Trả về:
//   - error: lỗi nếu có.
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newindex := prevBlock.index + 1
	newBlock := NewBlock(newindex, transactions, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}