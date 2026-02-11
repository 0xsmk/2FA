
#  2FA TOTP Tool (Go)

Minimal TOTP implementation in Go. Compatible with Google Authenticator, Authy, and any other TOTP app.

##  Current Features

- **Generate TOTP secrets** – 30-second, 6-digit codes
- **QR code generation** – easy setup with mobile apps
- **Code verification** – validate 6-digit codes
- **Single binary** – no dependencies, just download and run
- **Cross-platform** – build for Linux, macOS, Windows

##  Planned Features

- **Encrypted secret storage** – save and manage multiple accounts
- **Backup codes** – 8 one-time recovery codes
- **Interactive CLI** – `2fa list`, `2fa add`, `2fa remove`
- **Clipboard support** – copy codes directly to clipboard
- **TOTP counter** – show time remaining for current code

##  Installation

```bash
# Build from source
git clone https://github.com/0xsmk/2FA.git
cd 2fa
go mod download
go build -o 2fa

# Or just download the binary for your OS
```

##  Usage

### Generate new secret
```bash
./2fa --generate
```

### Generate QR code
```bash
./2fa --qr JBSWY3DPEHPK3PXP:alice@example.com
```

### Verify code
```bash
./2fa --verify JBSWY3DPEHPK3PXP:123456
```

##  Commands

| Flag | Description |
|------|-------------|
| `--generate` | Create new random secret |
| `--qr secret:name` | Generate QR code for account |
| `--verify secret:code` | Check if code is valid |

## 🔧 Build for different platforms

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o 2fa-linux

# macOS
GOOS=darwin GOARCH=amd64 go build -o 2fa-macos

# Windows
GOOS=windows GOARCH=amd64 go build -o 2fa.exe
```

##  Example output

```
 Secret: JBSWY3DPEHPK3PXP
 Add to Google Authenticator

Quick test:
  ./2fa --verify JBSWY3DPEHPK3PXP:123456
```

```
✅ Code 123456 is VALID
```

```
❌ Code 123456 is INVALID
```

```
 QR code saved as alice@example.com_qrcode.png
 Scan it with Google Authenticator
```

##  Disclaimer

This tool is for **educational purposes**.  
Never store real 2FA secrets in plain text. Use password manager or hardware token for production.
