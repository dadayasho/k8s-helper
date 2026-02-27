package main

import (
	"fmt"
	"k8s-helper/internal"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "pods":
			internal.CheckStuckPods()
		case "pv":
			internal.CheckPV()
		case "pvc":
			internal.CheckPVC()
		default:
			fmt.Printf("❌ Нет команды '%s'\n", os.Args[1])
		}
	} else {
		fmt.Printf("")
	}
}
