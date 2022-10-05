package assignment01bca

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Block struct {
	transaction string
	nonce       int
	prev_hash   string
	block_hash  string
}

type BlockChain struct {
	list []*Block
}

func Create_B() *BlockChain {
	a := new(BlockChain)
	return a
}

func CalculateHash(stringToHash string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))

}

func NewBlock(chain *BlockChain, trans string, n int, prev_h string) *Block {
	b := new(Block)
	b.transaction = trans
	b.nonce = n
	b.prev_hash = prev_h
	b.block_hash = CalculateHash(b.transaction + strconv.Itoa(b.nonce) + b.prev_hash)

	chain.list = append(chain.list, b)

	fmt.Printf("\n")
	return b
}

func DisplayBlocks(chain *BlockChain) {

	for index, val := range chain.list {
		fmt.Printf("%s Block :  %d %s\n", strings.Repeat("=", 25), index, strings.Repeat("=", 25))

		fmt.Printf(" Transection: %s \n Nonce: %d \n Previous Block Hash: %s \n Current Block Hash %s \n \n ", val.transaction, val.nonce, val.prev_hash, val.block_hash)
	}
}

func ChangeBlock(chain *BlockChain) {
	var ind int
	fmt.Println(" to edit transections, Enter the index of block")
	fmt.Scan(&ind)
	fmt.Println(ind)

	if ind < len(chain.list) {
		fmt.Printf("current transection is : %s \n ", chain.list[ind].transaction)

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Enter corrupt transection : ")

		scanner.Scan()

		text := scanner.Text()

		chain.list[ind].transaction = text

		fmt.Println("block corrupted")
	}

}

func VerifyChain(chain *BlockChain) {
	var check = 0

	for _, val := range chain.list {

		hash := CalculateHash(val.transaction + strconv.Itoa(val.nonce) + val.prev_hash)

		if hash != val.block_hash {
			check = 1
			break
		}

	}

	if check == 0 {
		fmt.Printf("no changes made inany block \n")
	} else {
		fmt.Printf("chain is not in original form, changes are made \n")
	}

}
