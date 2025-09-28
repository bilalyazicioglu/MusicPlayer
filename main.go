package main

import (
	"bufio"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, err := os.Open("Sagopa Kajmer - 24.mp3")
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	volume := &effects.Volume{
		Streamer: beep.Loop(-1, streamer),
		Base:     2,
		Volume:   0, // 0 = normal ses
		Silent:   false,
	}

	// Duraklama kontrolü
	ctrl := &beep.Ctrl{Streamer: volume, Paused: false}
	speaker.Play(ctrl)

	fmt.Println("Müzik Kontrolleri:")
	fmt.Println("p - Duraklat/Devam Et")
	fmt.Println("+/= - Sesi Aç")
	fmt.Println("-/_ - Sesi Kıs")
	fmt.Println("v [sayı] - Belirli ses seviyesi (örn: v -2)")
	fmt.Println("q - Çık")
	fmt.Printf("Mevcut ses seviyesi: %.1f\n\n", volume.Volume)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Komut girin: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		input = strings.TrimSpace(input)

		switch strings.ToLower(input) {
		case "p":
			speaker.Lock()
			ctrl.Paused = !ctrl.Paused
			speaker.Unlock()
			if ctrl.Paused {
				fmt.Println("Müzik duraklatıldı")
			} else {
				fmt.Println("Müzik devam ediyor")
			}

		case "+", "=":
			speaker.Lock()
			volume.Volume += 0.5
			if volume.Volume > 5 {
				volume.Volume = 5 // Max ses
			}
			speaker.Unlock()
			fmt.Printf("Ses seviyesi: %.1f\n", volume.Volume)

		case "-", "_":
			speaker.Lock()
			volume.Volume -= 0.5
			if volume.Volume < -10 {
				volume.Volume = -10 // Min ses
			}
			speaker.Unlock()
			fmt.Printf("Ses seviyesi: %.1f\n", volume.Volume)

		case "q":
			fmt.Println("Çıkılıyor...")
			return

		default:
			if strings.HasPrefix(strings.ToLower(input), "v ") {
				parts := strings.Split(input, " ")
				if len(parts) == 2 {
					if vol, err := strconv.ParseFloat(parts[1], 64); err == nil {
						if vol >= -10 && vol <= 5 {
							speaker.Lock()
							volume.Volume = vol
							speaker.Unlock()
							fmt.Printf("Ses seviyesi: %.1f\n", volume.Volume)
						} else {
							fmt.Println("Ses seviyesi -10 ile 5 arasında olmalı")
						}
					} else {
						fmt.Println("Geçersiz ses seviyesi")
					}
				} else {
					fmt.Println("Kullanım: v [sayı] (örn: v -2)")
				}
			} else {
				fmt.Println("Bilinmeyen komut. Yardım için kontrollere bakın.")
			}
		}
	}
}
