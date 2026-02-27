package internal

import (
	"fmt"
	"os/exec"
	"strings"
)

func CheckPVC() {
	fmt.Println("🔍 Поиск свободных PVC...")
	cmd := exec.Command("kubectl", "get", "pvc", "-A", "--no-headers")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("❌ kubectl error: %v\n", err)
		return
	}

	lines := strings.Split(string(output), "\n")
	failed := 0
	released := 0
	bound := 0
	pending := 0
	status_lists := map[string][]string{
		"failed_list":   make([]string, 0),
		"released_list": make([]string, 0),
		"bound_list":    make([]string, 0),
		"pending_list":  make([]string, 0),
	}

	for _, line := range lines {
		if strings.Contains(line, "Failed") {
			parts := strings.Fields(line)
			failed++
			status_lists["failed_list"] = append(status_lists["failed_list"], fmt.Sprintf("🔴 %s (%s) %s %s", parts[0], parts[1], parts[2], parts[3]))
		} else if strings.Contains(line, "Released") {
			parts := strings.Fields(line)
			released++
			status_lists["released_list"] = append(status_lists["released_list"], fmt.Sprintf("🟡 %s (%s) %s %s", parts[0], parts[1], parts[2], parts[3]))
		} else if strings.Contains(line, "Bound") {
			parts := strings.Fields(line)
			bound++
			status_lists["bound_list"] = append(status_lists["bound_list"], (fmt.Sprintf("%s (%s)", parts[0], parts[1])))
		} else if strings.Contains(line, "Pending") {
			parts := strings.Fields(line)
			pending++
			status_lists["pending_list"] = append(status_lists["pending_list"], (fmt.Sprintf("%s (%s)", parts[0], parts[1])))
		}
	}

	if failed == 0 && released == 0 && bound == 0 && pending == 0 {
		fmt.Println("Нету PVC")
	} else {
		for status, pvcs := range status_lists {
			fmt.Printf("=== %s ===\n", status)
			for j := 0; j < len(pvcs); j++ {
				fmt.Println(pvcs[j])
			}
			fmt.Println()
		}
		fmt.Println("Итого \n Failed - %d \n Released - %d \n Bound - %d \n Pending - %d", failed, released, bound, pending)
	}

}
