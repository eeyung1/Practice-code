package main

import "os"

func main() {
	if len(os.Args) != 4 {
		return
	}

	s, old, new := os.Args[1], os.Args[2], os.Args[3]

	if len(old) != 1 || len(new) != 1 {
		println(s)
		return
	}

	found := false
	for i := 0; i < len(s); i++ {
		if string(s[i]) == old {
			found = true
			break
		}
	}

	if !found {
		println(s)
		return
	}
	res := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == old {
			res += new
		} else {
			res += string(s[i])
		}
	}
	println(res)
}
