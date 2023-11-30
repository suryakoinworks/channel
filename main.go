package main

import (
	"channel/pkg"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	roles := []pkg.Role{}
	for i := 1; i <= 10; i++ {
		roles = append(roles, pkg.GetRoleByUserId(i)...)
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)

	fmt.Println(len(roles))

	start = time.Now()

	roles = []pkg.Role{}
	for i := 1; i <= 10; i++ {
		cRole := pkg.FastGetRoleByUserId(i)
		for role := range cRole {
			roles = append(roles, role)
		}
	}

	elapsed = time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)

	fmt.Println(len(roles))
}
