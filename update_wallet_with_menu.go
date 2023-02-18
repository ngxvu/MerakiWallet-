package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func welcome() {
	fmt.Println("WELCOME TO CHAIN MERAKI")
	fmt.Println("=======================")
	fmt.Println("  Choose Your Option  ")
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func promptOptions() {

	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput(" a - Create your wallet \n b - Sign In To Your Wallet \n c - Search A Wallet By Adrress  \n d - Exit", reader)
	switch opt {
	case "a":
		fmt.Println("You Chose a - Create A Wallet.")
		createWallet()
	case "b":
		fmt.Println("You Chose b - Sign In To Your Wallet.")
	case "c":
		fmt.Println("You Chose c - Search A Wallet By Address.")
	case "d":
		fmt.Println("You Chose d - Exit.")
	default:
		fmt.Println("That's Not A Valid Option.")
		fmt.Println("Please Select A Valid Options")
		fmt.Println(".............................")
		promptOptions()
	}

}

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

func createWallet() wallet {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create A Wallet Name: ", reader)
	email, _ := getInput("Your Email: ", reader)
	address, _ := fmt.Println("Your Wallet Address Is:", generateAdrress)
	w := newWallet(name, email, string(address))
	fmt.Println("Created the wallet -", w.name, w.email, w.address, w.currentTime)

	return w
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

func main() {
	welcome()
	promptOptions()
}
