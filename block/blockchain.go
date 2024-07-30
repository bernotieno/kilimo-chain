package blockchain

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"time"
)

type Block struct {
	TimeStamp     int64
	Data          BlockData
	PrevBlockHash []byte
	Hash          []byte
}

type Blockchain struct {
	blocks []*Block
}

type BlockData struct {
	Transactions      []Transaction
	QualityAssesments []QualityAssesment
	PriceUpdates      []PriceUpdate
}

type Transaction struct {
	ProductID  string
	Type       string
	SenderID   string
	ReceiverID string
	Quantity   float64
	Price      float64
	Timestamp  int64
}

type QualityAssesment struct {
	AssesmentID string
	ProductID   string
	AssessorID  string
	Grade       string
	Metrics     map[string]float64
	Timestamp   int64
}
type PriceUpdate struct {
	ProductType string
	Location    string
	Price       float64
	Timestamp   int64
}

func (b *Block) SetHash() {
	// Convert timestamp to byte slice
	timestampBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(timestampBytes, uint64(b.TimeStamp))

	// Marshal BlockData to JSON
	data, _ := json.Marshal(b.Data)

	// Concatenate PrevBlockHash, data, and timestampBytes
	headers := append(b.PrevBlockHash, data...)
	headers = append(headers, timestampBytes...)

	// Calculate SHA-256 hash
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data BlockData, PrevBlockHash []byte) *Block {
	block := &Block{
		TimeStamp:     time.Now().Unix(),
		Data:          data,
		PrevBlockHash: PrevBlockHash,
		Hash:          []byte{},
	}
	block.SetHash()
	return block
}

func (bc *Blockchain) AddBlock(data string) {
	// Convert string data to BlockData
	var blockData BlockData
	err := json.Unmarshal([]byte(data), &blockData)
	if err != nil {
		// Handle error (e.g., log it, return an error, etc.)
		return
	}

	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(blockData, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	// Create a genesis block with no previous hash and empty data
	genesisBlockData := BlockData{}
	genesisBlock := NewBlock(genesisBlockData, []byte{})

	return &Blockchain{
		blocks: []*Block{genesisBlock},
	}
}
