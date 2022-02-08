package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
)

const targetBits = 12
const maxNonce = 1000

type ProofOfWork struct {
	target *big.Int
	block  *Block
}

func (pow *ProofOfWork) prepareData(nonce int64) (data []byte) {
	data = bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			[]byte(strconv.FormatInt(pow.block.Timestamp, 16)),
			[]byte(strconv.FormatInt(targetBits, 16)),
			[]byte(strconv.FormatInt(nonce, 16)),
		},
		[]byte{},
	)

	return
}

func (pow *ProofOfWork) Proof() (int64, []byte) {
	var hashInt big.Int
	var hash [32]byte
	var n int64

	fmt.Printf("Mining : \"%s\"\n", pow.block.Data)

	for n < maxNonce {
		data := pow.prepareData(n)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			n++
		}
	}

	fmt.Printf("\n\n")

	return n, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
