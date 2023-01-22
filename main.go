package main

import "fmt"

func main()  {
	// variabel
	var namaDepan string = "saeful"

	var namaBelakang string
	namaBelakang = "mujab"

	namaTengah := "jono"

	//output ke layar
	fmt.Println("NAMA ANDA \n", namaDepan, namaTengah, namaBelakang + "!")

	var point = 2

	if point == 10 {
		fmt.Printf("LULUS SEMPURNA %d\n", point)
	}else if point > 5 {
		fmt.Printf("lulus %d\n", point)
	}else if point == 4 {
		fmt.Printf("hampir lulus %d\n", point)
	}else {
		fmt.Printf("tidak luluas %d\n", point)
	}

	//loping
	for i := 0; i < 10; i++ {
		fmt.Println("loop",i)
	}
}

/*
ini komentar
panjang banget
*/
