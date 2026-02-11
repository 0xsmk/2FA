package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

func main() {
	// CLI flags
	generateCmd := flag.Bool("generate", false, "Generate new TOTP secret")
	qrCmd := flag.String("qr", "", "Generate QR code (format: secret:account_name)")
	verifyCmd := flag.String("verify", "", "Verify TOTP code (format: secret:123456)")
	flag.Parse()

	switch {
	case *generateCmd:
		generateSecret()
	case *qrCmd != "":
		generateQR(*qrCmd)
	case *verifyCmd != "":
		verifyCode(*verifyCmd)
	default:
		printHelp()
	}
}

func generateSecret() {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "2FA App",
		AccountName: "user@example.com",
	})
	if err != nil {
		log.Fatal("Failed to generate secret:", err)
	}

	fmt.Println(" Secret:", key.Secret())
	fmt.Println(" Add to Google Authenticator")
	fmt.Printf("\nQuick test:\n")
	fmt.Printf("  ./2fa --verify %s:123456\n", key.Secret())
}

func generateQR(input string) {
	parts := strings.SplitN(input, ":", 2)
	if len(parts) != 2 {
		log.Fatal("Invalid format. Use --qr secret:account_name")
	}

	secret, account := parts[0], parts[1]

	// Generate TOTP URI
	uri := fmt.Sprintf("otpauth://totp/%s?secret=%s&issuer=2FA%%20App", account, secret)

	// Generate QR code
	filename := fmt.Sprintf("%s_qrcode.png", account)
	err := qrcode.WriteFile(uri, qrcode.Medium, 256, filename)
	if err != nil {
		log.Fatal("Failed to generate QR code:", err)
	}

	fmt.Printf(" QR code saved as %s\n", filename)
	fmt.Println(" Scan it with Google Authenticator")
}

func verifyCode(input string) {
	parts := strings.SplitN(input, ":", 2)
	if len(parts) != 2 {
		log.Fatal("Invalid format. Use --verify secret:123456")
	}

	secret, code := parts[0], parts[1]

	valid := totp.Validate(code, secret)
	if valid {
		fmt.Printf("✅ Code %s is VALID\n", code)
	} else {
		fmt.Printf("❌ Code %s is INVALID\n", code)
	}
}

func printHelp() {
	fmt.Println(" Simple 2FA TOTP Tool (Go version)")
	fmt.Println("\nUsage:")
	fmt.Println("  Generate new secret:")
	fmt.Println("    ./2fa --generate")
	fmt.Println("\n  Generate QR code:")
	fmt.Println("    ./2fa --qr JBSWY3DPEHPK3PXP:alice@example.com")
	fmt.Println("\n  Verify code:")
	fmt.Println("    ./2fa --verify JBSWY3DPEHPK3PXP:123456")
}