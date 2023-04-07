package main

import (
	"fmt"
	"time"
)

func main ()  {
	// menampilkan waktu saat ini
	now := time.Now()
	fmt.Println("Waktu saat ini adalah", now)

	// mengubah string menjadi waktu
	layout := "2006-01-02"
	str := "2022-12-31"
	t, _ := time.Parse(layout, str)
	fmt.Println("waktu yang diubah dari string adalah", t)

	// mengubah waktu menjadi string
	layout = "02 Januari 2006"
	str2 := t.Format(layout) 
	fmt.Println("waktu yang diubah menjadi string adalah", str2)

	// menambah durasi ke waktu yang sudah ada
	duration, _ := time.ParseDuration("1h")
	oneHoursLater := now.Add(duration)
	fmt.Println("waktu 1 jam setelah waktu saat ini adalah", oneHoursLater)

	// mengurangi waktu dengan waktu yang lain
	newYear := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	diff := newYear.Sub(now)
	fmt.Println("Selisih antara waktu saat ini dengan 1 Januari 2024 adalah", diff)

	// memotong waktu menjadi waktu yang lebih kecil
	truncate := now.Truncate(24 * time.Hour)
	fmt.Println("waktu saat ini yang sudah dipotong menjadi waktu hari ini adalah", truncate)
}