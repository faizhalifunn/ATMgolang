package main

import (
	"fmt"
	"os"
	"time"
)

// var users = map[string]string{
// 	"Faiz": "2406425312",
// }

type Account struct {
	Username string
	Owner    string
	Password string
	Saldo    []float64
}

var account = Account{
	Username: "Faiz",
	Owner:    "Faiz Zhalifun Nafzi",
	Password: "2406425312",
	Saldo:    []float64{3500000},
}

var transactionHistory []string

func getTimestamp() string {
	return time.Now().Format("02-Jan-2006 15:04:05")
}

func login() {
	var username, password string
	fmt.Print("Enter username: ")
	fmt.Scanln(&username)

	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	// Check if the username exists and if the password is correct
	if username == account.Username && password == account.Password {
		fmt.Println("---------------------------- Welcome -----------------------------------")
		return // Exit the program successfully
	} else {
		fmt.Println("Password or Username is Wrong !")
		os.Exit(1)
	}
}

func menu() {
	for {
		// Display menu
		fmt.Println("-------------------------- Menu transaksi ------------------------------")
		fmt.Println("1. Lihat Saldo")
		fmt.Println("2. Tambah Saldo")
		fmt.Println("3. Tarik Saldo")
		fmt.Println("4. Lihat Akun")
		fmt.Println("5. History Transaksi")
		fmt.Println("6. Exit")
		fmt.Print("Pilih opsi: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			Lihatsaldo()
		case 2:
			Tambahsaldo()
		case 3:
			Tariksaldo()
		case 4:
			Lihatakun()
		case 5:
			Historytransaksi()
		case 6:
			fmt.Println("Keluar dari program.")
			os.Exit(0)
		default:
			fmt.Println("Opsi tidak valid. Silakan coba lagi.")
		}
	}
}

func Lihatsaldo() {
	fmt.Println("--------------------------- Account Balance ----------------------------")
	fmt.Printf("\nSaldo anda saat ini adalah Rp.%.2f\n\n\n", account.Saldo)
	lanjut()
}

func Tambahsaldo() {
	var tambah float64
	fmt.Println("--------------------------- Add Balance -------------------------------")
	fmt.Print("Masukkan jumlah yang ingin ditambahkan ke saldo: Rp. ")
	fmt.Scanln(&tambah)

	if tambah <= 0 {
		fmt.Println("Jumlah yang dimasukkan tidak valid. Harap masukkan jumlah yang valid.")
	} else {
		account.Saldo = append(account.Saldo, account.Saldo[len(account.Saldo)-1]+tambah)
		transactionHistory = append(transactionHistory, fmt.Sprintf("Penambahan saldo : Rp. %.2f", tambah))
		fmt.Printf("Saldo berhasil ditambahkan, Saldo anda sekarang adalah Rp. %.2f.\n", account.Saldo)
	}

	lanjut()
}

func Tariksaldo() {
	var tarik float64
	fmt.Println("------------------------ Withdraw Money ----------------------------")
	fmt.Print("Masukkan jumlah yang ingin ditarik (multiples of 50.000): Rp. ")
	fmt.Scanln(&tarik)

	if tarik <= 0 {
		fmt.Println("Jumlah yang dimasukkan tidak valid. Harap masukkan jumlah yang positif.")
	} else if int(tarik)%50000 != 0 {
		fmt.Println("Jumlah yang ingin ditarik harus kelipatan Rp. 50.000 atau kelipatan Rp. 100.000.")
	} else if tarik > account.Saldo[len(account.Saldo)-1] {
		fmt.Println("Saldo tidak cukup untuk menarik jumlah tersebut.")
	} else {
		account.Saldo = append(account.Saldo, account.Saldo[len(account.Saldo)-1]-tarik)
		transactionHistory = append(transactionHistory, fmt.Sprintf("Penarikan uang: Rp. %.2f", tarik))
		fmt.Printf("Rp. %.2f berhasil ditarik. Saldo baru anda adalah Rp. %.2f.\n", tarik, account.Saldo)
	}

	lanjut()
}

func lanjut() {
	fmt.Println("\napakah anda ingin melanjutkan transaksi ?")
	fmt.Println("\ntekan apa saja untuk kembali ke menu\ntekan '1' untuk keluar")

	var input string
	fmt.Scanln(&input)

	switch input {
	case "1":
		fmt.Println("Keluar dari program.")
		os.Exit(0)
	default:
		menu()
	}
}

func Lihatakun() {
	fmt.Println("--------------------------- Account Information -------------------------")
	fmt.Printf("Username: %s\n", account.Username)
	fmt.Printf("Owner: %s\n", account.Owner)
	fmt.Printf("Saldo: Rp. %.2f\n", account.Saldo)
	lanjut()
}

func Historytransaksi() {
	fmt.Println("------------------------ Transaction History ----------------------------")
	if len(transactionHistory) == 0 {
		fmt.Println("Tidak ada riwayat transaksi.")
	} else {
		for i, transaction := range transactionHistory {
			fmt.Printf("%d. %s\n", i+1, transaction)
		}
	}
	lanjut()
}

func main() {

	login()

	menu()

}
