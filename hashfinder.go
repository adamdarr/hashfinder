package main

import (
	"fmt"
	"crypto/sha512"
	"strconv"
	"encoding/hex"
	"os"
)

func main() {
	goal := "f04031af5ee9939a931a062b836c385eb2ee681fd704ed919534b78329ae34e1" + 
			"73e13c31ebfd03db1eb9e9759cd3365195d0d7d53028256038b24975ace8c83d"

	hash := sha512.New()

    for i := 0; i < 256; i++ {
    	for j := 0; j < 256; j++ {
    		ip := "172.16." + strconv.Itoa(i) + "." + strconv.Itoa(j)
    		
    		hash.Write([]byte(ip))
    		digest := hex.EncodeToString(hash.Sum(nil))

    		if goal == digest {
    			fmt.Printf("%s\n", ip)
    			os.Exit(0)
    		}

    		hash.Reset()
    	}
    }
}