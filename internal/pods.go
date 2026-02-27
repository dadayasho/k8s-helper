package internal

import (
	"fmt"
	"os/exec"
	"strings"
)

func CheckStuckPods() {
	fmt.Println("🔍 Проверяю зависшие поды...")

	cmd := exec.Command("kubectl", "get", "pods", "-A", "--no-headers")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("❌ kubectl error: %v\n", err)
		return
	}

	lines := strings.Split(string(output), "\n")
	stuck := 0

	for _, line := range lines {
		if strings.Contains(line, "Pending") || strings.Contains(line, "CrashLoopBackOff") {
			parts := strings.Fields(line)
			if len(parts) >= 4 {
				fmt.Printf("🔴 %s/%s %s\n", parts[0], parts[1], parts[3])
				stuck++
			}
		}
	}

	if stuck == 0 {
		fmt.Println("🟢 Все поды здоровы!")
	} else {
		fmt.Printf("💀 Найдено %d проблемных подов\n", stuck)
	}
}
