package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/mail"
	"os"
	"strconv"
	"strings"
)

// ----------------- func main ----------------------

func main() {

	data, err := read()
	if err != nil {
		fmt.Println("Read file error!")
		listUser = nil
	}
	listUser = data

	Welcome()
	MenuChinh()
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

// ----------------- MenuChinh ----------------------

func MenuChinh() {
	reader := bufio.NewReader(os.Stdin)
	for {
		opt, _ := getInput("\nLựa Chọn Chức Năng Ví.\n   0 - Xem list User\n   1 - Tạo Tài Khoản.\n   2 - Đăng Nhập Tài Khoản.\n   3 - Lưu Và Thoát Chương Trình.", reader)
		switch opt {
		case "0":
			fmt.Println(listUser)
			continue
		case "1":
			fmt.Println("Bạn Chọn Tạo Tài Khoản.")
			SignUp()
			break
		case "2":
			fmt.Println("Bạn Chọn Đăng Nhập Vào Tài Khoản Hiện Có.")
			SignIn()
			break
		case "3":
			fmt.Println("Lưu Thành Công ... Thoát Chương Trình.")
			err := save(listUser)
			if err != nil {
				fmt.Println(" ERROR: cannot save data: ", err)
			}
			break
		default:
			fmt.Println("Lựa chọn đó không có - Hãy Chọn Lại.")
			fmt.Println(".............................")
			MenuChinh()
		}
		break
	}
}

// ----------------- MenuPhu ----------------------

func MenuUser() {
	for {
		reader := bufio.NewReader(os.Stdin)
		opt, _ := getInput("\nLựa Chọn Chức Năng.\n   1 - Tạo Thêm Wallet.\n   2 - Xoá Wallet.\n   3 - Sửa Tên Đăng Nhập.\n   4 - Thêm Token Vào Address\n   5 - Quay Về Menu Chính.", reader)
		switch opt {
		case "1":
			fmt.Println("Bạn Chọn Tạo Thêm Wallet.")
			addWallet()
			break
		case "2":
			fmt.Println("Bạn Chọn Xoá Wallet.")
			deleteWallet()
			break
		case "3":
			fmt.Println("Bạn Chọn Sửa Tên Email Đăng Nhập.")
			changeEmail()
			break
		case "4":
			fmt.Println("Bạn Chọn Thêm Token Vào Address.")
			addTokenIntoWallet()
			break
		case "5":
			fmt.Println("Bạn Chọn Quay Về Menu Chính.")
			MenuChinh()
			break
		default:
			fmt.Println("Lựa chọn đó không có - Hãy Chọn Lại.")
			fmt.Println(".............................")
			MenuUser()
		}
		break
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
		MenuChinh()
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
	MenuChinh()
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

// ----------------- SignIn ----------------------

func SignIn() {
	fmt.Println(listUser)
	reader := bufio.NewReader(os.Stdin)
	email, _ := getInput(">> Hãy Nhập Email Được Liên Kết Với MerakiChain. <<", reader)
	if !validFormEmail(email) {
		fmt.Println("Sai Định Dạng Email - Hãy Nhập Lại.")
		SignIn()
	}
	isLogin := false
	currentUserIndex := -1
	for index, user := range listUser {
		if email == user.email {
			fmt.Println("Đăng Nhập Thành Công.")
			currentUserIndex = index
			fmt.Println("Index của ban là: ", currentUserIndex)
			isLogin = true
			ThongTinUser(user)
			MenuUser()
			break
		}
	}
	if isLogin == false {
		fmt.Println("Đăng Nhập Thất Bại Vì Tài Khoản Chưa Được Khởi Tạo.\nMời Bạn Tạo Tài Khoản Mới.")
		MenuChinh()
		return
	}
}

// ----------------- ThongTinUser ----------------------

func ThongTinUser(user User) {
	fmt.Println("Thông Tin Của User: ")
	fmt.Println("- Email: ", user.email)
	for _, wallet := range user.wallets {
		fmt.Println("- Address: ", wallet.address)
		for _, token := range wallet.tokens {
			fmt.Println("  - Symbol: ", token.symbol)
			fmt.Println("  - Balance: ", token.balance)
		}
	}
}

// ----------------- TaoWalletMoi ----------------------

func addWallet() {

	reader := bufio.NewReader(os.Stdin)
	email, _ := getInput("Nhập Lại Email Của Bạn: ", reader)
	if !validFormEmail(email) {
		fmt.Println("Sai Định Dạng Email.")
		addWallet()
	}
	isLogin := false
	currentUserIndex := -1
	for index, user := range listUser {
		if email == user.email {
			currentUserIndex = index
			isLogin = true
			break
		}
	}
	if isLogin == false {
		fmt.Println("Đăng Nhập Thất Bại Vì Tài Khoản Chưa Được Khởi Tạo.\nMời Bạn Tạo Tài Khoản Mới.")
		return
	}

	var listToken []Token
	for {
		quest, _ := getInput("1 - Nhập Token\n2 - Thoát", reader)
		if quest == "1" {
			symbol, _ := getInput("Nhập Symbol:", reader)
			balance, _ := getInput("Nhập Balance:", reader)
			b, err := strconv.ParseFloat(balance, 64)
			if err != nil {
				fmt.Println("error")
			}
			listToken = append(listToken, Token{symbol: symbol, balance: b})
		} else if quest == "2" {
			fmt.Println("Thoát")
			break
		} else {
			fmt.Println("Lựa chọn đó không có - Hãy Chọn Lại.")
			continue
		}
	}

	newWallet := Wallet{
		address: uuid.NewString(),
		tokens:  listToken,
	}
	listUser[currentUserIndex].wallets = append(listUser[currentUserIndex].wallets, newWallet)
	fmt.Println("Bạn Đã Tạo Ví Mới Thành Công!")
	ThongTinUser(listUser[currentUserIndex])
}

// ----------------- DeleteWallet ----------------------

func deleteWallet() {
	reader := bufio.NewReader(os.Stdin)
	email, _ := getInput("Nhập Lại Email Của Bạn: ", reader)
	if !validFormEmail(email) {
		fmt.Println("Sai Định Dạng Email.")
		addWallet()
	}
	isLogin := false
	currentUserIndex := -1
	for index, user := range listUser {
		if email == user.email {
			currentUserIndex = index
			isLogin = true
			ThongTinUser(listUser[currentUserIndex])
			break
		}
	}
	if isLogin == false {
		fmt.Println("Đăng Nhập Thất Bại Vì Tài Khoản Chưa Được Khởi Tạo.\nMời Bạn Tạo Tài Khoản Mới.")
		return
	}
	quest, _ := getInput("Nhập Địa Chỉ Ví Muốn Xoá: ", reader)
	isDetele := false
	for i, wallet := range listUser[currentUserIndex].wallets {
		if quest == wallet.address {
			listUser[currentUserIndex].wallets = append(listUser[currentUserIndex].wallets[:i], listUser[currentUserIndex].wallets[i+1:]...)
			isDetele = true
			break
		}
		if !isDetele {
			fmt.Println("Không Tìm Thấy Địa Chỉ Để Xoá.")
			break
		}
	}
	fmt.Println("Ban da xoa wallet: ", quest)
	ThongTinUser(listUser[currentUserIndex])
	MenuUser()
	return
}

// ----------------- ThayDoiTenEmail ----------------------

func changeEmail() {
	reader := bufio.NewReader(os.Stdin)
	email, _ := getInput("Nhập Lại Email Của Bạn: ", reader)
	if !validFormEmail(email) {
		fmt.Println("Sai Định Dạng Email.")
		addWallet()
	}
	isLogin := false
	currentUserIndex := -1
	for index, user := range listUser {
		if email == user.email {
			currentUserIndex = index
			isLogin = true
			ThongTinUser(listUser[currentUserIndex])
			break
		}
	}
	if isLogin == false {
		fmt.Println("Đăng Nhập Thất Bại Vì Tài Khoản Chưa Được Khởi Tạo.\nMời Bạn Tạo Tài Khoản Mới.")
		return
	}
	for {
		quest, _ := getInput("Nhập tên Email mới của bạn: ", reader)
		if !checkEmailExist(quest, listUser) {
			fmt.Println("Tài Khoản Này Đã Được Tạo Rồi. Hãy Su Dung Email khac.")
		} else {
			listUser[currentUserIndex].email = quest
			fmt.Println("Bạn Đã Đổi Email Thành: ", quest)
			ThongTinUser(listUser[currentUserIndex])
			break
		}
	}
	MenuUser()
	return
}

// ----------------- addThemTokenVaoWallet ----------------------

func addTokenIntoWallet() {
	reader := bufio.NewReader(os.Stdin)
	email, _ := getInput("Nhập Lại Email Của Bạn: ", reader)
	if !validFormEmail(email) {
		fmt.Println("Sai Định Dạng Email.")
		addWallet()
	}
	isLogin := false
	currentUserIndex := -1
	for index, user := range listUser {
		if email == user.email {
			currentUserIndex = index
			isLogin = true
			ThongTinUser(listUser[currentUserIndex])
			break
		}
	}
	if isLogin == false {
		fmt.Println("Đăng Nhập Thất Bại Vì Tài Khoản Chưa Được Khởi Tạo.\nMời Bạn Tạo Tài Khoản Mới.")
		return
	}
	quest, _ := getInput("Ban them token vao xoa address nao?: ", reader)
	isFindAddress := false
	for i, wallet := range listUser[currentUserIndex].wallets {
		if quest == wallet.address {
			// bat dau nhap token
			var listToken []Token
			for {
				quest, _ := getInput("1 - Nhập Token\n2 - Thoát", reader)
				if quest == "1" {
					symbol, _ := getInput("Nhập Symbol:", reader)
					balance, _ := getInput("Nhập Balance:", reader)
					b, err := strconv.ParseFloat(balance, 64)
					if err != nil {
						fmt.Println("error")
					}
					listToken = append(listToken, Token{symbol: symbol, balance: b})
				}
				if quest == "2" {
					isFindAddress = true
					break
				}
			}
			// append token vao day
			newWallet := Wallet{
				address: wallet.address,
				tokens:  listToken,
			}
			listUser[currentUserIndex].wallets = append(listUser[currentUserIndex].wallets, newWallet)
			listUser[currentUserIndex].wallets = append(listUser[currentUserIndex].wallets[:i], listUser[currentUserIndex].wallets[i+1:]...)
			// ket thuc nhap token
		}
	}
	if !isFindAddress {
		fmt.Println("Khong tim thay dia chi")
		return
	}
	ThongTinUser(listUser[currentUserIndex])
	MenuUser()
	return
}

func read() ([]User, error) {

	var listUser []User
	dat, err := os.ReadFile("./data.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(dat, &listUser)
	if err != nil {
		return nil, err
	}

	return listUser, nil
}

func save(data []User) error {
	tmp, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile("./data.json", tmp, 0644)
	if err != nil {
		return err
	}

	return nil
}
