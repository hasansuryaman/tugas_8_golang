package main

import (
    "fmt"
    "runtime"
    "time"
    "math/rand"
)

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	// var pesan = []string{"Apa Kabar Teman-Teman"}
	var pesan = "Apa Kabar Teman-teman"
	fmt.Println("Kirim Pesan :", pesan)
	fmt.Println("")
	var waktu = make(chan int)

	go kirim_pesan(pesan, waktu)
	terima_pesan(pesan, waktu)
}

func kirim_pesan(pesan string, channel_1 chan<- int) {
	for i := 0; true; i++ {
		channel_1 <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}
func terima_pesan(pesan string, channel_1 <-chan int) {
loop:
	for {
		select {
		case data := <-channel_1:
			fmt.Print("Pesan Diterima :" , pesan, " -> ", data ,"\n")
		case <-time.After(time.Second * 5):
			fmt.Println("\nTimeout, Tidak ada aktivitas dalam 5 detik\n")
			break loop
		}
	}
}
