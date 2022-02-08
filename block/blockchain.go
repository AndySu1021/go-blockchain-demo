package block

import "math/big"

type Blockchain struct {
	Blocks []*Block
}

func (blockchain *Blockchain) AddBlock(Data [] byte) {
	// 取出前一個 Block
	PrevBlock := blockchain.Blocks[len(blockchain.Blocks) - 1]
	// 創建 Block
	NewBlock := CreateBlock(string(Data), PrevBlock.Hash)
	// 驗證 block
	p := newProofOfWork(NewBlock)
	nonce, hash := p.Proof()
	NewBlock.Hash = hash[:]
	NewBlock.Nonce = nonce
	// 將 block 加到鏈上
	blockchain.Blocks = append(blockchain.Blocks, NewBlock)    // 把新創建的 Block 接上去
}

func newProofOfWork(b *Block) (p *ProofOfWork) {
	target := big.NewInt(1)
	target.Lsh(target, uint(256 - targetBits))

	p = &ProofOfWork{target, b}

	return
}
