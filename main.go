package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Barang struct {
	Kode  string
	Nama  string
	Harga float64
}

var DataBarang = map[string]Barang{
	"01": {"001", "Buku", 5.00},
	"02": {"002", "Pulpen", 2.00},
	"03": {"003", "Pencil", 2.50},
	"04": {"004", "Penghapus", 1.00},
	"05": {"005", "Tipp-Ex", 3.00},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Selamat datang di Program Kasir")
		keranjang := belanja(scanner)
		tampilkanDetailBelanja(keranjang)
		simpanKeFile(keranjang)
		fmt.Print("Apakah Anda ingin berbelanja lagi? (y/n): ")
		scanner.Scan()
		jawaban := scanner.Text()
		if jawaban != "y" {
			break
		}
	}

	fmt.Println("Terima kasih telah berbelanja!")
}

func belanja(scanner *bufio.Scanner) map[string]int {
	keranjang := make(map[string]int)

	for {
		fmt.Println("\nDaftar Barang:")
		tampilkanDaftarBarang()

		fmt.Print("Masukkan kode barang (atau 'selesai' untuk selesai belanja): ")
		scanner.Scan()
		kodeBarang := scanner.Text()

		if kodeBarang == "selesai" {
			break
		}

		barang, ada := DataBarang[kodeBarang]
		if !ada {
			fmt.Println("Kode barang tidak valid.")
			continue
		}

		fmt.Printf("Barang: %s - %s\n", kodeBarang, barang.Nama)
		fmt.Printf("Harga per item: %.2f\n", barang.Harga)

		fmt.Print("Masukkan jumlah yang ingin dibeli: ")
		scanner.Scan()
		jumlahStr := scanner.Text()
		jumlah, err := strconv.Atoi(jumlahStr)
		if err != nil || jumlah <= 0 {
			fmt.Println("Jumlah tidak valid.")
			continue
		}

		keranjang[kodeBarang] += jumlah
	}

	return keranjang
}

func tampilkanDaftarBarang() {
	for kode, barang := range DataBarang {
		fmt.Printf("%s - %s - %.2f\n", kode, barang.Nama, barang.Harga)
	}
}

func tampilkanDetailBelanja(keranjang map[string]int) {
	fmt.Println("\n== Detail Belanja ==")
	totalHarga := 0.0
	for kodeBarang, jumlah := range keranjang {
		barang := DataBarang[kodeBarang]
		subtotal := float64(jumlah) * barang.Harga
		totalHarga += subtotal
		fmt.Printf("%s - %s (x%d): %.2f\n", kodeBarang, barang.Nama, jumlah, subtotal)
	}
	fmt.Printf("Total Biaya: %.2f\n", totalHarga)
}

func simpanKeFile(keranjang map[string]int) {
	file, err := os.Create("invoice.txt")
	if err != nil {
		fmt.Println("Gagal membuat file invoice.")
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	writer.WriteString("== Invoice Belanja ==\n")
	for kodeBarang, jumlah := range keranjang {
		barang := DataBarang[kodeBarang]
		subtotal := float64(jumlah) * barang.Harga
		writer.WriteString(fmt.Sprintf("%s - %s (x%d): %.2f\n", kodeBarang, barang.Nama, jumlah, subtotal))
	}
	writer.WriteString(fmt.Sprintf("Total Biaya: %.2f\n", hitungTotalBiaya(keranjang)))
	writer.WriteString("Terima kasih telah berbelanja!")
}

func hitungTotalBiaya(keranjang map[string]int) float64 {
	totalHarga := 0.0
	for kodeBarang, jumlah := range keranjang {
		barang := DataBarang[kodeBarang]
		subtotal := float64(jumlah) * barang.Harga
		totalHarga += subtotal
	}
	return totalHarga
}
