package main

import (
	"fmt"
	namingV1 "github.com/zyt153/baby-naming/naming"
	// namingV2 "github.com/zyt153/baby-naming/v2/naming"

)

func main() {
	fmt.Println("v1: ", namingV1.CreateBabyName())
	// fmt.Println("v2: ", namingV2.CreateBabyName(true))
}