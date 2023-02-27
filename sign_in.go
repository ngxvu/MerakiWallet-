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

func main() {

	Welcome()
	OptPromptSignIn()
}

func Welcome() {
	fmt.Println("Welcome To Meraki Chain")
	fmt.Println("=======================")
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func OptPromptSignIn() {
	reader := bufio.NewReader(os.Stdin)
	mainOpt, _ := getInput("Lựa Chọn Chức Năng Ví.\n   1 - Tạo Tài Khoản.\n   2 - Đăng Nhập Tài Khoản.\n   3 - Thoát Chương Trình.", reader)
	switch mainOpt {
	case "1":
		fmt.Println("Bạn Chọn Đăng Kí Tài Khoản.")
		SignUp()
	case "2":
		fmt.Println("Bạn Chọn Đăng Nhập Vào Tài Khoản Hiện Có.")
		SignIn()
	case "3":
		fmt.Println("Chương Trình Đang Thoát.")
	default:
		fmt.Println("Lựa chọn đó không có - Hãy Chọn Lại.\n")
		OptPromptSignIn()
	}
}

// Khai bao struct

type User struct {
	email   string   `json:"email"`
	wallets []Wallet `json:"wallets"`
}

type Wallet struct {
	address string  `json:"address"`
	tokens  []Token `json:"tokens"`
}
type Token struct {
	symbol  string
	balance float64
}

func SignIn() {
	reader := bufio.NewReader(os.Stdin)
	email, _ := getInput("Hãy Nhập Email Có Liên Kết Với MerakiChain.", reader)
	if !validFormEmail(email) {
		fmt.Println("Sai Định Dạng Email - Hãy Nhập Lại.")
		SignIn()
	} else if !checkEmailExist(email, listUser) {
		fmt.Println("Sửa Wallet")
	} else {
		fmt.Println("\nEmail Không Có Trong Hệ Thống, Hãy Tạo Tài Khoản Mới !")
		OptPromptSignUp()
	}
}

// check form email
func validFormEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func checkEmailExist(email string, listUsers []User) bool {
	for _, user := range listUsers {
		if email == user.email {
			return false
		}
	}
	return true
}

func checkAddressExist(address string, listWallet []Wallet) bool {
	for _, wallet := range listWallet {
		if address == wallet.address {
			return false
		}
	}
	return true
}

func YouChoseCreateAccount() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Bạn Chọn Tạo Tài Khoản ")
	optcase1, _ := getInput(" 1 - Tiếp Tục Tạo Tài Khoản \n 2 - Quay Lại Menu Chương Trình ", reader)
	switch optcase1 {
	case "1":
		fmt.Println("Bạn Chọn 1 - Hãy Nhập Email Của Bạn")
		SignUp()
	case "2":
		fmt.Println("Bạn Chọn 2 - Quay Lại Menu Chương Trình")
		OptPromptSignUp()
	default:
		fmt.Println("Lựa chọn đó không có - Hãy Chọn Lại.\n")
		fmt.Println(".............................")
		YouChoseCreateAccount()
	}

}

func OptPromptSignUp() {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("\nLựa Chọn Chức Năng Ví.\n   1 - Tạo Tài Khoản.\n   2 - Đăng Nhập Tài Khoản.\n   3 - Thoát Chương Trình.", reader)
	switch opt {
	case "1":
		YouChoseCreateAccount()
	case "2":
		fmt.Println("Bạn Chọn Đăng Nhập Vào Tài Khoản Hiện Có.")
		SignIn()
	case "3":
		fmt.Println("Chương Trình Đang Thoát.")
	default:
		fmt.Println("Lựa chọn đó không có - Hãy Chọn Lại.")
		fmt.Println(".............................")
		OptPromptSignUp()
	}
}

var listToken []Token
var listUser []User
var listWallet []Wallet

func SignUp() {

	reader := bufio.NewReader(os.Stdin)
	email, _ := getInput(">> Hãy Nhập Email Chưa Được Liên Kết Với MerakiChain. <<", reader)
	//checkDinhdangEmailvaEmailExist
	if !validFormEmail(email) {
		fmt.Println("Sai Định Dạng Email - Hãy Nhập Lại.")
		SignUp()
	} else if checkEmailExist(email, listUser) {
		fmt.Println("\n** Tài Khoản Của Bạn Đã Tạo Thành Công. **")
	} else {
		fmt.Println("Tài Khoản Này Đã Được Tạo Rồi. Hãy Đăng Nhập.")
		OptPromptSignUp()
	}

	//addtheToken
	for {
		TokenOpt, _ := getInput("Chọn Options Của Bạn\n  1 - Nhập Token\n  2 - Thoát Và Quay Về Menu Chính", reader)
		switch TokenOpt {
		case "1":
			symbol, _ := getInput("Nhập Tên Token: ", reader)
			balance, _ := getInput("Nhập Số Dư: ", reader)
			b, err := strconv.ParseFloat(balance, 64)
			if err != nil {
				fmt.Println("Số Dư Phải Nhập Dạng Số")
			}
			listToken = append(listToken, Token{
				symbol:  symbol,
				balance: b,
			})
		case "2":
			fmt.Println("Bạn Chọn Thoát Và Quay Về Menu Chính")
		}
		if TokenOpt == "2" {
			break
		}
	}

	//symbol, _ := getInput("Token Name: ", reader)
	//balance, _ := getInput("Token Balance($): ", reader)
	//b, err := strconv.ParseFloat(balance, 64)
	//if err != nil {
	//	fmt.Println("The balance must be a number")
	//}

	//autoGenerateAddress&CheckAddressUnique
	address := uuid.NewString()
	if checkAddressExist(address, listWallet) {
		fmt.Println("Địa Chỉ Ví Của Bạn Là: ", address)
	}

	listUser = append(listUser, User{
		email: email,
		wallets: []Wallet{
			{
				address: address,
				tokens:  listToken,
			},
		},
	},
	)
	fmt.Println(listUser)

}
