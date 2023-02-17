package main

import (
	"fmt"
	"math/rand"
	"time"
)

type wallet struct {
	name        string
	email       string
	address     string
	currentTime time.Time
	tokens      map[string]float64 //token: Symbol(string) - balance(float64)
}

func newwallet(name, email, address string) wallet {
	w := wallet{
		name:        name,
		email:       email,
		address:     address,
		currentTime: time.Now(),
		tokens:      map[string]float64{"ADA": 2000, "CAKE": 1000},
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

// add the token to the wallet

// func (w *wallet) token(symbol string, balance float64) {
// w.token[symbol]balance
// }

func main() {
	fmt.Println("WELCOME TO MERAKI CHAIN!")
	fmt.Println("========================")
	mywallet := newwallet("Vu's wallet", "ngxvu126@gmail.com", "abcd")
	mywallet.updateNameEmailAddress("Thang's wallet", "pxthang@gmail.com")
	mywallet.addTokens("CLV", 500)
	mywallet.addTokens("NEAR", 700)
	mywallet.addTokens("AXS", 900)
	fmt.Println(mywallet.format())

}
