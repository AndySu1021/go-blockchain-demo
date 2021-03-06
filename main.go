package main

import (
	"bufio"
	"fmt"
	"go-blockchain-demo/block"
	"os"
)

func CreateBlockchain() *block.Blockchain {
	// 創建創始區塊
	genesisBlock := block.CreateBlock("Genesis Block", []byte{})
	genesisBlock.SetHash()
	return &block.Blockchain{Blocks: []*block.Block{genesisBlock}}
}

func Help() {
	fmt.Println("There are 3 operations:")
	fmt.Println("Type 1 for adding a new Block")
	fmt.Println("Type 2 for printing the Blockchain")
	fmt.Println("Type 3 for exiting")
}

func main() {
	fmt.Println("Welcome to our Blockchain project.")
	fmt.Println("Enter h for help")

	var (
		op string
	)

	NewBlockchain := CreateBlockchain() // 新增一個區塊鏈
	for true {
		fmt.Scanln(&op)
		if op == "h" {
			fmt.Println("Printing the help")
			Help() // 顯示使用方法
		} else if op == "1" {
			fmt.Println("Entering your data:")
			reader := bufio.NewReader(os.Stdin)
			data, _, _:= reader.ReadLine() // 讀一整行 input
			NewBlockchain.AddBlock(data) // 利用 input 作為 data 來創建區塊鏈
		} else if op == "2"{
			for _ , block := range NewBlockchain.Blocks {
				fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
				fmt.Printf("Data: %s\n", block.Data)
				fmt.Printf("Hash: %x\n", block.Hash)
				fmt.Println()
			} // 查詢資料
		} else if op == "3"{
			break
		} else {
			fmt.Println("Please Enter h, 1, 2, 3")
		}
	}
}
