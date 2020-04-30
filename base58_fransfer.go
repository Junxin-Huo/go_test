package main

import (
	"fmt"
	"strings"
)

var base58DCRString string = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
var base58TRXString string = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
var base58BTCString string = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
var base58XRPString string = "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"

// BTCAddress: 1DTXLQUZKZVKz88zJbHjgVS4j3B5qMB4mE
// XRPAddress: rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN

func main() {
	addressBTC := "1DTXLQUZKZVKz88zJbHjgVS4j3B5qMB4mE"
	fmt.Println(BTCToXRP(addressBTC))

	addressXRP := "rAPERVgXZavGgiGv6xBgtiZurirW2yAmY"
	fmt.Println(XRPToBTC(addressXRP))
}

func BTCToXRP(BTCAddress string) string {
	XRPAddress := ""
	for _, value := range BTCAddress {
		index := strings.Index(base58BTCString, string(value))
		XRPAddress = XRPAddress + string(base58XRPString[index])
	}
	return XRPAddress
}

func XRPToBTC(XRPAddress string) string {
	BTCAddress := ""
	for _, value := range XRPAddress {
		index := strings.Index(base58XRPString, string(value))
		BTCAddress = BTCAddress + string(base58BTCString[index])
	}
	return BTCAddress
}
