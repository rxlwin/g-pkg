package util

import "fmt"

func P(vals ...interface{}) {
	fmt.Printf("===start===\n")
	for _, val := range vals {
		fmt.Printf("(%T)%v\n", val, val)
	}
	fmt.Printf("===end===\n")
}

func PP(location int, val interface{}) {
	fmt.Printf("%v 地址:%p 值:%v\n", location, &val, val)
}

func PMap(location int, val map[string]interface{}) {
	fmt.Printf("%v start ", location)
	for k, v := range val {
		fmt.Printf("(%p) %v => (%p) %v , ", &k, k, &v, v)
	}
	fmt.Printf("\n")
}
