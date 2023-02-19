package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type wallet struct {
	name        string
	email       string
	address     string
	currentTime time.Time
	tokens      map[string]float64 //token: Symbol(string) - balance(float64)
}

func newWallet(name, email, address string) wallet {
	w := wallet{
		name:        name,
		email:       email,
		address:     address,
		currentTime: time.Now(),
		tokens:      map[string]float64{},
	}
	return w
}

// format the wallet
func (w *wallet) format() string {
	fmt.Println("Wallet name:", w.name)
	fmt.Println("Wallet email:", w.email)
	fmt.Println("Waller address:", generateAdrress)
	fmt.Println("Time created:", w.currentTime)
	fs := "Wallet Breakdown: \n"
	var total float64 = 0
	// list token
	for k, v := range w.tokens {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}
	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}

// update wallet name and gmail

func (w *wallet) updateNameEmailAddress(name, email string) {
	w.name = name
	w.email = email
}

// update tokens
func (w *wallet) addTokens(token string, balance float64) {
	w.tokens[token] = balance
}

func generateAdrress() {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits + specials
	length := 8
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	str := string(buf) // E.g. "3i[g0|)z"

	fmt.Println(str)
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createWallet() wallet {
	reader := bufio.NewReader(os.Stdin)
	// reade wallet name
	//fmt.Println("Create A Wallet Name: ")
	//name, _ := reader.ReadString('\n')
	//name = strings.TrimSpace(name)
	// reade email
	//fmt.Println("Your Email: ")
	//email, _ := reader.ReadString('\n')
	//email = strings.TrimSpace(email)
	// autogen wallet's address
	//address, _ := fmt.Println("Your Wallet Address Is:", generateAdrress)

	name, _ := getInput("Create A Wallet Name: ", reader)
	email, _ := getInput("Your Email: ", reader)
	address, _ := fmt.Println("Your Wallet Address Is:", generateAdrress)
	w := newWallet(name, email, string(address))
	fmt.Println("Created the wallet -", w.name, w.email, w.address, w.currentTime)

	return w
}

func promptOptions(w wallet) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose Your Option ( a - Create your wallet | b - Sign In To Your Wallet | c - Search A Wallet By Adrress  | d - Exit ): ", reader)

	switch opt {
	case "a":
		fmt.Println("You Chose a - Create A Wallet.")
	case "b":
		fmt.Println("You Chose b - Sign In To Your Wallet.")
	case "c":
		fmt.Println("You Chose c - Search A Wallet By Address.")
	case "d":
		fmt.Println("You Chose d - Exit.")
	default:
		fmt.Println("That's Not A Valid Option.")
		promptOptions(w)
	}
}
func main() {
	fmt.Println("WELCOME TO CHAIN MERAKI")
	fmt.Println("=======================")

	mywallet := createWallet()
	fmt.Println(mywallet)
	promptOptions(mywallet)

}
