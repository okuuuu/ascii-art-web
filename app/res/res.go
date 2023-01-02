package res

import (
	"bufio"
	"os"
	"strings"
)

// Arbitrary functions.

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadAscii(banner string) map[int][]string {

	file, err := os.Open(banner)
	Check(err)
	defer file.Close()

	mp := make(map[int][]string)

	sc := bufio.NewScanner(file)
	for c := 0; c < 95; c++ {
		sc.Scan()
		for i := 0; i < 8; i++ {
			sc.Scan()
			mp[c] = append(mp[c], sc.Text())
		}
	}

	return mp
}

func WriteAscii(input []string, mp map[int][]string) []string {

	var output []string

	for _, i := range input {
		if i == "" {
			output = append(output, "")
			continue
		}
		rword := []rune(i)
		for v, k := range rword {
			if mp[int(k)-32] != nil {
				for l := 0; l < 8; l++ {
					if v == 0 {
						output = append(output, mp[int(k)-32][l])

					} else {
						output[len(output)-8+l] += mp[int(k)-32][l]
					}
				}
			}
		}
	}

	return output
}

func PrintAscii(str1 []string, str2 []string) string {

	var res string

	if len(strings.Join(str1, "")) == 0 {
		for range str1 {
			res += "\n"
		}
	} else {
		for _, line := range str2 {
			res += line + "\n"
		}
	}
	return res
}
