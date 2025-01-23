package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func getBrowserBytes(bytes []byte) []byte {
	for i, val := range bytes {
		if val == ']' && bytes[i-1] == '"' {
			return bytes[0 : i+1]
		}
	}
	return nil
}

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	/*
		!!! !!! !!!
		обратите внимание - в задании обязательно нужен отчет
		делать его лучше в самом начале, когда вы видите уже узкие места, но еще не оптимизировалм их
		так же обратите внимание на команду в параметром -http
		перечитайте еще раз задание
		!!! !!! !!!
	*/
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer file.Close() // закрываем, дабы избежать утечки ресурсов.

	defaultAt := "@"
	replaceAt := " [at] "
	scanner := bufio.NewScanner(file)
	seenBrowsers := make(map[string]bool, 0)
	var buf bytes.Buffer

	for i := 0; scanner.Scan(); i++ {
		var user User
		rawBytes := scanner.Bytes()
		user.UnmarshalJSON(getBrowserBytes(rawBytes))

		isAndroid := false
		isMSIE := false

		for _, browser := range user.Browsers {
			if strings.Contains(browser, "Android") {
				isAndroid = true
				_, exist := seenBrowsers[browser]
				if !exist {
					seenBrowsers[browser] = true
				}
			}

			if strings.Contains(browser, "MSIE") {
				isMSIE = true
				_, exist := seenBrowsers[browser]
				if !exist {
					seenBrowsers[browser] = true
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		user.UnmarshalJSON(rawBytes)
		fmt.Fprintf(&buf, "[%d] %s <%s>\n", i, user.Name, user.Email)
	}

	fmt.Fprintln(out, "found users:")
	fmt.Fprintln(out, strings.ReplaceAll(buf.String(), defaultAt, replaceAt))
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}

func main() {
	out := os.Stdout
	FastSearch(out)
}
