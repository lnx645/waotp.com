package utils

import (
	"fmt"
	"image/color"
	"log"

	"github.com/skip2/go-qrcode"
)

func GenerateAndSaveQrCode(content string, filename string) {
	whatsappGreen := color.RGBA{37, 211, 102, 255}
	white := color.White
	err := qrcode.WriteColorFile(content, qrcode.Medium, 256, white, whatsappGreen, filename)
	if err != nil {
		log.Fatal("Gagal generate QR Code:", err)
	}
	fmt.Printf("QR Code berhasil disimpan sebagai %s\n", filename)
}
