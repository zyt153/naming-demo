package main

import (
	"fmt"
	namingV1 "github.com/zyt153/baby-naming/naming"
	namingV2 "github.com/zyt153/baby-naming/v2/naming"
	namingV3 "github.com/zyt153/baby-naming/v3/naming"
)

func main() {
	fmt.Println("Create a name in naming-v1: ", namingV1.CreateBabyName())

	fmt.Println("Create a boy's name in naming-v2: ", namingV2.CreateBabyName(true))
	fmt.Println("Create a girl's name in naming-v2: ", namingV2.CreateBabyName(false))

	minLen := 4
	name, err := namingV3.CreateBabyName(true, minLen)
	if err == nil {
		fmt.Printf("Create a boy's name with min length %v in naming-v3: %v", minLen, name)
	} else {
		fmt.Println("Invalid input in naming-v3: ", err)
	}
}