package main

import (
	"fmt"
	"log"
	"net/http"
)

func verifyTHJ(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w,"ParseForm() err %v",err)
		return
	}
	fmt.Fprintf(w,"Verifying the key\n")
	key := r.FormValue("key")
	decodedString := ""
	for i := 0; i < len(key); i += 2 {
		hexPair := key[i : i+2]
		asciiVal := hexPairToASCII(hexPair)
		decodedString += string(asciiVal)
	}
	fmt.Fprint(w,"Decoded string:", decodedString)
}

func signTHJ(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w,"ParseForm() err %v",err)
		return
	}
	fmt.Fprintf(w,"POST is successful\n")
	signature := r.FormValue("signature")
	fmt.Fprintf(w,"Signature %s %x",signature,signature)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/sign", signTHJ)
	http.HandleFunc("/verify", verifyTHJ)

	fmt.Println("Starting Server on Port 9876")
	if err := http.ListenAndServe(":9876", nil); err != nil {
		log.Fatal(err)
	}

}

func hexPairToASCII(hexPair string) byte {
	high := hexCharToValue(hexPair[0]) << 4
	low := hexCharToValue(hexPair[1])
	return high | low
}

func hexCharToValue(hexChar byte) byte {
	switch {
	case '0' <= hexChar && hexChar <= '9':
		return hexChar - '0'
	case 'a' <= hexChar && hexChar <= 'f':
		return hexChar - 'a' + 10
	case 'A' <= hexChar && hexChar <= 'F':
		return hexChar - 'A' + 10
	default:
		return 0
	}
}
