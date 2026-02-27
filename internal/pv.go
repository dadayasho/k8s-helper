package internal

import (
	"fmt"
	"os/exec"
	"strings"
)

func CheckPV() {
	fmt.Println("🔍 Поиск свободных PV...")
	cmd := exec.Command("kubectl", "get", "pv", "--no-headers")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("❌ kubectl error: %v\n", err)
		return
	}

	lines := strings.Split(string(output), "\n")
	failed := 0
	available := 0

	for _, line := range lines {
		if strings.Contains(line, "Failed") {
			parts := strings.Fields(line)
			fmt.Printf("🔴 %s/ %s\n", parts[0], parts[2])
			failed++
		} else if strings.Contains(line, "Available") {
			parts := strings.Fields(line)
			fmt.Printf("🟡 %s (%s) - свободен\n", parts[0], parts[1])
			available++
		}

	}

	fmt.Printf("🔴 %d - упавших pv\n", failed)

	fmt.Printf("🟢 %d - доступных\n", available)

}
