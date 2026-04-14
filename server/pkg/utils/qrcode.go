package utils

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"github.com/disintegration/imaging"
	"github.com/skip2/go-qrcode"
)

func GenerateAndSaveQrCode(content string, filename string) {
	final := CreateQrCodeImage(content, filename)
	errr := imaging.Save(final, filename)
	if errr != nil {
		fmt.Println("Gagal simpan file")
	}
	fmt.Println("Qr berhasil disimpan")
}
func CreateQrCodeImage(content string, filename string) *image.NRGBA {
	qr, err := qrcode.New(content, qrcode.High)

	if err != nil {
		log.Fatal(err)
	}
	qr.BackgroundColor = color.White
	qr.ForegroundColor = color.Black
	qr.DisableBorder = true
	qrImg := qr.Image(512)

	logofile, err := os.Open("logo.png")
	if err != nil {
		log.Fatal(err)
	}
	defer logofile.Close()
	logoImg, _, err := image.Decode(logofile)
	if err != nil {
		log.Fatal(err)
	}
	logoFinal := imaging.Resize(logoImg, 70, 70, imaging.Lanczos)
	final := imaging.OverlayCenter(qrImg, logoFinal, 1.0)
	return final
}
