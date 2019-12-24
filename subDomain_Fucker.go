package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func subDomainFuck(domain string) /*(hosts , subdomainIPaddr , websiteTitle ,httpResponeCode string)*/ {
	ospwd, _ := os.Getwd()
	subdomainDict := ospwd + "/dict/subDomain.txt"

	/*逐行读取文件内容*/

	fileOpen, err := os.Open(subdomainDict)
	if err != nil {
		fmt.Println("[❌]\t字典文件未找到\t[❌]")
		os.Exit(0)
	}
	defer fileOpen.Close()
	contextLine := bufio.NewReader(fileOpen)
	for {
		contextText, _, eofErr := contextLine.ReadLine()
		if eofErr == io.EOF {
			break
		}
		hosts := string(contextText) + "." + domain
		result, _ := net.LookupIP(hosts)
		if len(result) > 0 {
			fmt.Println(hosts, "\t", "---→", "\t", result)
		}

	}

}

func main() {
	targets := os.Args[1]
	subDomainFuck(targets)
}
