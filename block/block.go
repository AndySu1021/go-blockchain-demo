package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce 		  int64
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10)) // 把 Timestamp 換算成 10 進位
	payload := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{}) // 把 Block 的 Data 都放進去
	hash := sha256.Sum256(payload) // 產生 Hash
	b.Hash = hash[:] // 設置 Hash
}

func CreateBlock(data string, prevHash []byte) *Block{
	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		prevHash,
		[]byte{},
		0,
	}

	return block
}
