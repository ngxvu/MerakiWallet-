package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type wallet struct {
	WalletName string
	TokenName  map[string]float64
}

// make new wallet

func NewWallet(WalletName string) wallet {
	w := wallet{
		WalletName: WalletName,
		TokenName:  map[string]float64{},
	}
	return w
}

// Format the Wallet

func (w wallet) format() string {
	// formatstring
	fs := "Wallet Portfolio: \n"
	var total float64 = 0

	// List Tokens
	for k, v := range w.TokenName {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	// Total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)
	return fs
}

// add token to the wallet
func (w wallet) addTokens(name string, amount float64) {
	w.TokenName[name] = amount
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}
func createWallet() wallet {
	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Print("Create a new wallet: ")
	// 	Walletname, _ := reader.ReadString('\n')
	// 	Walletname = strings.TrimSpace(Walletname)
	Walletname, _ := getInput("Create A New Wallet: ", reader)

	w := NewWallet(Walletname)
	fmt.Println("Created The Wallet - ", w.WalletName)

	return w
}

func promptoptions(w wallet) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose Your Option (A - Add Token, P - Save & Exit): ", reader)

	switch opt {
	case "A":
		tokenname, _ := getInput("Token Name: ", reader)
		amount, _ := getInput("Amount($): ", reader)
		// am : amount
		am, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			fmt.Println("The Amount Must Be A Number")
			promptoptions(w)
		}
		w.addTokens(tokenname, am)
		fmt.Println("Token Added -", tokenname, "$", amount)
		promptoptions(w)
	case "P":
		fmt.Println("You Chose Save & Exit")
	default:
		fmt.Println("That Was Not A Valid Option. Let's Choose Again!")
		promptoptions(w)
	}
}
func main() {
	Mywallet := createWallet()
	promptoptions(Mywallet)
	fmt.Println(Mywallet)
}
