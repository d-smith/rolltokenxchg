package main

import (
	"fmt"
	"github.com/xtraclabs/rolltokenxchg"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		println("usage: go rollTokenFromAuth0Token <auth0Token>")
		return
	}
	jwtResponse := rolltokenxchg.TradeTokenForToken(os.Args[1], "http://localhost:3000/oauth2/token")
	fmt.Println("\n", jwtResponse)
}
