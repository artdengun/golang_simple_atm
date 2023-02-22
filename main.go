package main

import "fmt"

// Fungsi untuk melakukan penarikan uang
func withdraw(amount int, balance *int) bool {
	if amount > *balance {
		return false
	}
	*balance -= amount
	return true
}

// fungsi untuk melakukan deposit uang
func deposit(amount int, balance *int) {
	*balance += amount
}

func main() {
	balance := 0 // saldo awal

	var choice int
	for {
		fmt.Println("1. Cek Saldo")
		fmt.Println("2. Tarik uang")
		fmt.Println("3. Deposit uang")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih Menu : ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Saldo Anda:", balance)
		case 2:
			var amount int
			fmt.Print("Masukkan jumlah uang yang ingin ditarik: ")
			fmt.Scanln(&amount)
			if withdraw(amount, &balance) {
				fmt.Println("Uang berhasil ditarik")
			} else {
				fmt.Println("Maaf, saldo Anda tidak mencukupi")
			}
		case 3:
			var amount int
			fmt.Print("Masukkan jumlah uang yang ingin disimpan: ")
			fmt.Scanln(&amount)
			deposit(amount, &balance)
			fmt.Println("Uang berhasil disimpan")
		case 4:
			fmt.Println("Terima kasih telah menggunakan layanan ATM")
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
