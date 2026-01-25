package main

import (
	"fmt"
	"math"
	"sort"
)

type Koordinat struct {
	Lat, Lon float64
}

type Lokasi struct {
	Nama         string
	Alamat       string
	Posisi       Koordinat
	Jarak        float64
	Probabilitas float64 
}

// Fungsi Haversine untuk menghitung jarak geografis
func itungJarak(asal, tujuan Koordinat) float64 {
	const jariBumi = 6371.0
	dLat := (tujuan.Lat - asal.Lat) * (math.Pi / 180)
	dLon := (tujuan.Lon - asal.Lon) * (math.Pi / 180)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(asal.Lat*(math.Pi/180))*math.Cos(tujuan.Lat*(math.Pi/180))*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return jariBumi * c
}

func main() {
	rumahSaya := Lokasi{
		Nama:   "Rumah (Ciluar Asri)",
		Alamat: "Perumahan Ciluar Asri Blok A4 No.3",
		Posisi: Koordinat{Lat: -6.5398, Lon: 106.8202},
	}

	daftarToko := []Lokasi{
		{
			Nama:   "Kopi Kenangan - Kedunghalang",
			Alamat: "Jl. Raya Pemda No.1, Kedunghalang",
			Posisi: Koordinat{Lat: -6.5601, Lon: 106.8115},
		},
		{
			Nama:   "Kopi Kenangan - Pom Bensin Pajajaran",
			Alamat: "Jl. Raya Pajajaran No.127",
			Posisi: Koordinat{Lat: -6.5785, Lon: 106.8088},
		},
		{
			Nama:   "Kopi Kenangan - Sudirman Bogor",
			Alamat: "Jl. Jend. Sudirman No.24",
			Posisi: Koordinat{Lat: -6.5892, Lon: 106.7971}, 
		},
	}

	var totalBobot float64


	for i := range daftarToko {
		daftarToko[i].Jarak = itungJarak(rumahSaya.Posisi, daftarToko[i].Posisi)

		// Rumus: Semakin kecil jarak, semakin besar peluang 
		daftarToko[i].Probabilitas = 1.0 / daftarToko[i].Jarak
		totalBobot += daftarToko[i].Probabilitas
	}

	for i := range daftarToko {
		daftarToko[i].Probabilitas = (daftarToko[i].Probabilitas / totalBobot) * 100
	}

	sort.Slice(daftarToko, func(i, j int) bool {
		return daftarToko[i].Jarak < daftarToko[j].Jarak
	})

	fmt.Println("======= HASIL ANALISIS RUTE & PELUANG REKOMENDASI =======")
	fmt.Printf("TITIK AWAL : %s\n", rumahSaya.Nama)
	fmt.Printf("ALAMAT     : %s\n", rumahSaya.Alamat)
	fmt.Println("-----------------------------------------------------")
	fmt.Printf("%-35s | %-8s | %-10s\n", "NAMA CABANG", "JARAK", "PELUANG")
	fmt.Println("-----------------------------------------------------")

	for _, toko := range daftarToko {
		fmt.Printf("%-35s | %4.2f km | %5.2f%%\n", toko.Nama, toko.Jarak, toko.Probabilitas)
	}

	fmt.Println("================================================")

	// Kesimpulan
	pilihan := daftarToko[0]
	menit := (pilihan.Jarak / 25) * 60 

	fmt.Printf("\nKESIMPULAN:")
	fmt.Printf("\nCabang terdekat adalah %s", pilihan.Nama)
	fmt.Printf("\nEstimasi perjalanan: ±%.0f Menit", menit)
	fmt.Printf("\nRUTE: [RUMAH] -------- ☕ [%s]\n", pilihan.Nama)
}
