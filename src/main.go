package main

import(
	"fmt"
)

func main() {
	var nilai int 
	var nama string

	fmt.Print("Masukkan nama siswa : ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan nilai siswa : ")
	fmt.Scanln(&nilai)

	if nilai<=74{
		fmt.Println("Dear "+nama+", nilai kamu di bawah rata-rata.")
	}else if nilai==75{
		fmt.Println("Dear "+nama+", nilai kamu cukup.")
	}else if nilai>=76 && nilai<=100{
		fmt.Println("Dear "+nama+", nilai kamu sangat baik!")
	}else if nilai>=101{
		fmt.Println("Dear "+nama+", nilai kamu melebihi standar.")
	}else{
		fmt.Println("Dear "+nama+", nilaimu tidak valid.")
	}
}