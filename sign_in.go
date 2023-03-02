package main

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"net/mail"
	"os"
	"strconv"
	"strings"
)

// ----------------- func main ----------------------

func main() {
	Welcome()
	OptPromptSignUp()
	fmt.Println(listUser)
}

// ----------------- func welcome ----------------------

func Welcome() {
	fmt.Println("Welcome To Meraki Chain")
	fmt.Println("=======================")
}

// ----------------- func getInput ----------------------

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

// ----------------- khai báo Struct ----------------------

type User struct {
	email   string   `json:"email"`
	wallets []Wallet `json:"wallets"`
}
type Wallet struct {
	address string  `json:"address"`
	tokens  []Token `json:"tokens"`
}
type Token struct {
	symbol  string  `json:"Symbol"`
	balance float64 `json:"Balance"`
}

var listUser []User

// ----------------- SignUpMenu ----------------------

func OptPromptSignUp() {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("\nLựa Chọn Chức Năng Ví.\n   1 - Tạo Tài Khoản.\n   2 - Đăng Nhập Tài Khoản.\n   3 - Thoát Chương Trình.", reader)
	switch opt {
	case "1":
		fmt.Println("Bạn Chọn Tạo Tài Khoản.")
		SignUp()
	case "2":
		fmt.Println("Bạn Chọn Đăng Nhập Vào Tài Khoản Hiện Có.")
	case "3":
		fmt.Println("Chương Trình Đang Thoát.")
	default:
		fmt.Println("Lựa chọn đó không có - Hãy Chọn Lại.")
		fmt.Println(".............................")
		OptPromptSignUp()
	}
}

// ----------------- SignUp ----------------------

func SignUp() {
	reader := bufio.NewReader(os.Stdin)
	email, _ := getInput(">> Hãy Nhập Email Chưa Được Liên Kết Với MerakiChain. <<", reader)

	//checkDinhdangEmailvaEmailExist
	if !validFormEmail(email) {
		fmt.Println("Sai Định Dạng Email - Hãy Nhập Lại.")
		SignUp()
	} else if !checkEmailExist(email, listUser) {
		fmt.Println("Tài Khoản Này Đã Được Tạo Rồi. Hãy Đăng Nhập.")
		return
	}
	var listToken []Token
	for {
		tokenmenu, _ := getInput("1-Nhập Token\n2-Thoát", reader)
		if tokenmenu == "1" {
			symbol, _ := getInput("Nhập Symbol:", reader)
			balance, _ := getInput("Nhập Balance:", reader)
			b, err := strconv.ParseFloat(balance, 64)
			if err != nil {
				fmt.Println("error")
			}
			listToken = append(listToken, Token{symbol: symbol, balance: b})
		} else if tokenmenu == "2" {
			fmt.Println("Thoát")
			break
		} else {
			fmt.Println("Lựa chọn đó không có - Hãy Chọn Lại.")
			continue
		}
	}
	var initWallet = Wallet{
		address: uuid.NewString(),
		tokens:  listToken,
	}
	listUser = append(listUser, User{
		email:   email,
		wallets: []Wallet{initWallet},
	})
	fmt.Println(listUser)
	OptPromptSignUp()
}

// ----------------- CheckFormEmail ----------------------
func validFormEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// ----------------- CheckEmailExist ----------------------
func checkEmailExist(email string, listUsers []User) bool {
	for _, user := range listUsers {
		if email == user.email {
			return false
		}
	}
	return true
}
