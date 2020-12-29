package main

import (
	"fmt"
	// "io/ioutil"
	// "strings"
	"strconv"
	// "sort"
	"crypto/md5"
	"encoding/hex"
)


func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
 }

func main() {
	input := "iwrupvqb"

	i := 1
	foundP1 := false
	for {
		hash := GetMD5Hash(input + strconv.Itoa(i))

		if (! foundP1 && string(hash[:5]) == "00000") {
			fmt.Println(i)
			foundP1 = true
		}

		if (string(hash[:6]) == "000000") {
			fmt.Println(i)
			break
		}

		i++
	}
}
