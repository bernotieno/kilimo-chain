package blockchain

import (
	"bytes"
	"crypto/sha256"
	"json"
	"strconv"
	"time"
	
)

type Block struct {
	TimeStamp     int64
	Data          BlockData
	PrevBlockHash []byte
	Hash          []byte
}

type BlockData struct {
	Transactions []Transaction
	QualityAssesments []QualityAssesment
	PriceUpdates []PriceUpdate
}

type Transaction struct {
	ProductID string
	Type string
	SenderID string
	ReceiverID string
	Quantity float64
	Price float64
	Timestamp int64
}

type QualityAssesment struct {
	AssesmentID string
	ProductID string
	AssessorID string
	Grade string
	Metrics map[string]float64
	Timestamp int64
}
type PriceUpdate struct {
	ProductType string
	Location string
	Price float64
	Timestamp int64
}

func (b *Block) SetHash() {
	timestamp := []byte(string(b.TimeStamp))
	data, _ := json.Marshal(b.Data)
	headers := append(b.PrevBlockHash, data...)
	headers = append(headers,timestamp...)
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]

}

func NewBlock(data BlockData, PrevBlockHash []byte) *Block {
	block := &Block {
		Timestamp: time.Now().Unix(),
		Data: data,
		PrevBlockHash: PrevBlockHash,
		Hash: []byte{},
	}
	block.SetHash()
	return block
}