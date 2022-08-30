package main

import "fmt"

func main() {
	x := map[string]bool{
		"create": false,
		"update": true,
	}

	fmt.Printf("Create: %t\n", x["create"])
	fmt.Printf("Update: %t\n", x["update"])
	fmt.Printf("Not Existed: %t\n", x["NotExisted"])
}
