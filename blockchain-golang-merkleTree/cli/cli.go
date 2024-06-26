package cli

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"golang-blockchain/blockchain"
	"os"
	"strconv"
	"strings"
)

func Run() {
	var bc *blockchain.Blockchain
	bc = &blockchain.Blockchain{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("##############################################")
		fmt.Printf("1. Create Blockchain \n2. Add Block \n3. Print Blockchain \n4. Update Transaction \n5. Print Merkle Tree \n6. Validate Block \n7. Exit\n")
		
		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSuffix(choice, "\r\n")
		intchoice, err := strconv.Atoi(choice)
		if err != nil {
			choice = strings.TrimSuffix(choice, "\n")
			intchoice, err = strconv.Atoi(choice)
		}
		switch intchoice {
		case 1:
			bc = blockchain.NewBlockchain()
			blockchain.PrintBlock(bc.Blocks[0])
		case 2:
			if len(bc.Blocks) == 0 {
				fmt.Println("Create Blockchain first!")
				break
			}
			fmt.Println("Create transaction, one per line. Press ctrl + enter to finish!")
			transactions := []*blockchain.Transaction{}
			for {
				scanner.Scan()
				input := scanner.Text()
				if len(input) > 0 {
					fmt.Println("Transaction added!")
					transactions = append(transactions, &blockchain.Transaction{Data: []byte(input)})				
				}else if len(transactions) > 0 {
					bc.AddBlock(transactions)
					fmt.Println("Block added!")
					break
				} else {
					fmt.Println("No transactions entered!")
					break
				}
			}	
		case 3:
			if len(bc.Blocks) == 0 {
				fmt.Println("Create Blockchain first!")
				break
			}
			for _, block := range bc.Blocks {
				blockchain.PrintBlock(block)
			}
		case 4:
			if len(bc.Blocks) == 0 {
				fmt.Println("Create Blockchain first!")
				break
			}
			fmt.Printf("Enter Block Number: ")
			scanner.Scan()
			blockNumber := scanner.Text()
			intblockNumber, _ := strconv.Atoi(blockNumber)
			fmt.Printf("Enter Transaction Number: ")
			scanner.Scan()
			transactionNumber := scanner.Text()
			inttransactionNumber, _ := strconv.Atoi(transactionNumber)
			fmt.Printf("Enter New Transaction Data :")
			scanner.Scan()
			input := scanner.Text()
			blockchain.UpdateTransactionData(bc, intblockNumber, inttransactionNumber-1, input)
			fmt.Println("Transaction Updated!")
		case 5:
			if len(bc.Blocks) == 0 {
				fmt.Println("Create Blockchain first!")
				break
			}
			fmt.Printf("Enter Block Number: ")
			scanner.Scan()
			blockNumber := scanner.Text()
			intblockNumber, _ := strconv.Atoi(blockNumber)
			blockchain.PrintMerkleTree(bc.Blocks[intblockNumber])
		case 6:
			if len(bc.Blocks) == 0 {
				fmt.Println("Create Blockchain first!")
				break
			}
			fmt.Printf("Enter Block Number: ")
			scanner.Scan()
			blockNumber := scanner.Text()
			intblockNumber, _ := strconv.Atoi(blockNumber)
			block := bc.Blocks[intblockNumber]
			if hex.EncodeToString(block.CalculateMerkleRoot()) == hex.EncodeToString(block.MerkleRoot) {
				fmt.Println("Valid Block!")
			}else {
				fmt.Println("Invalid Block!")
			}
		case 7:
			fmt.Println("Exiting...")
			os.Exit(0)
		}	

	}
}

