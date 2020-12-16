package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)


func solveP1(inputString []string) int {
	mem := make(map[int]int64, 0)
	mask := ""
	for _, l := range(inputString) {
		parts := strings.Split(l, "=")
		if (len(parts) < 2) {
			continue
		}

		if (strings.TrimSpace(parts[0]) == "mask") {
			mask = strings.TrimSpace(parts[1])
		} else {
			add := strings.TrimSpace(parts[0])
			add = strings.ReplaceAll(add, "mem[", "")
			add = strings.ReplaceAll(add, "]", "")
			address, _ := strconv.Atoi(add)
			num, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
			binary := strconv.FormatInt(int64(num), 2)

			mem[address] = applyMaskP1(mask, binary)
		}
	}
	ans := 0
	for v := range(mem) {
		ans += int(mem[v])
	}

	return ans
}

func applyMaskP1(mask string, binary string) int64 {
	res := ""

	for i := len(mask) - 1; i >= 0; i-- {
		diff := len(mask) - i - 1
		c := "0"
		if diff < len(binary) {
			c = string(binary[len(binary) - diff - 1])
		}

		if string(mask[i]) != "X" {
			c = string(mask[i])
		}

		res = c + res
	}

	num, _ := strconv.ParseInt(res, 2, 64)

	return num
}

func solveP2(inputString []string) int {
	mem := make(map[string]int, 0)
	mask := ""
	for _, l := range(inputString) {
		parts := strings.Split(l, "=")
		if (len(parts) < 2) {
			continue
		}

		if (strings.TrimSpace(parts[0]) == "mask") {
			mask = strings.TrimSpace(parts[1])
		} else {
			add := strings.TrimSpace(parts[0])
			add = strings.ReplaceAll(add, "mem[", "")
			add = strings.ReplaceAll(add, "]", "")
			address, _ := strconv.Atoi(add)

			binaryAddress := strconv.FormatInt(int64(address), 2)
			genAddr := applyMaskP2(mask, binaryAddress)
			num, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

			for v := range(genAddr) {
				mem[genAddr[v]] = int(num)
			}
		}
	}

	ans := 0
	for v := range(mem) {
		ans += int(mem[v])
	}

	return ans
}

// If the bitmask bit is 0, the corresponding memory address bit is unchanged.
// If the bitmask bit is 1, the corresponding memory address bit is overwritten with 1.
// If the bitmask bit is X, the corresponding memory address bit is floating.
func applyMaskP2(mask string, binary string) []string {
	res := ""

	for i := len(mask) - 1; i >= 0; i-- {
		diff := len(mask) - i - 1
		c := "0"
		if diff < len(binary) {
			c = string(binary[len(binary) - diff - 1])
		}

		if string(mask[i]) == "X" {
			c = string(mask[i])
		} else if (string(mask[i]) == "1") {
			c = "1"
		}

		res = c + res
	}

	gen := generateAddrs([]string{res})

	return gen
}

// generate addresses based on floating indexes
func generateAddrs(addresses []string) []string {
	done := true

	for i, addr := range(addresses) {
		for c := 0; c < len(addr); c++ {
			if string(addr[c]) == "X" {
				done = false
				addr1 := addr[0:c] + "0" + addr[c+1:]
				addr2 := addr[0:c] + "1" + addr[c+1:]

				addresses = append(addresses[:i], addresses[i+1:]...)

				addresses = append(addresses, addr1)
				addresses = append(addresses, addr2)

				return generateAddrs(addresses)
			}
		}
	}

	if (done) {
		return addresses
	}

	return []string{}
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	inputString := strings.Split(strings.TrimSpace(string(input)), "\n")

	ans := solveP1(inputString)
	fmt.Println("P1", ans)

	ans = solveP2(inputString)
	fmt.Println("P2", ans)
}
