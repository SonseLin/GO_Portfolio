package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <package_path>")
		return
	}

	path := fmt.Sprintf("./%s/", args[1])

	// 1. go test -coverprofile
	testCmd := exec.Command("go", "test", path, "-coverprofile=coverage.out")
	testCmd.Stdout = os.Stdout
	testCmd.Stderr = os.Stderr
	if err := testCmd.Run(); err != nil {
		fmt.Println("go test failed:", err)
		return
	}

	// 2. go tool cover -html
	coverCmd := exec.Command("go", "tool", "cover", "-html=coverage.out", "-o", "coverage.html")
	coverCmd.Stdout = os.Stdout
	coverCmd.Stderr = os.Stderr
	if err := coverCmd.Run(); err != nil {
		fmt.Println("go tool cover failed:", err)
		return
	}

	// 3. open coverage.html
	openCmd := exec.Command("open", "coverage.html")
	openCmd.Stdout = os.Stdout
	openCmd.Stderr = os.Stderr
	if err := openCmd.Run(); err != nil {
		fmt.Println("Failed to open coverage report:", err)
		return
	}

	// 4. clean up except coverage.html
	cleanCmd := exec.Command("rm", "coverage.out")
	cleanCmd.Stdout = os.Stdout
	cleanCmd.Stderr = os.Stderr
	if err := cleanCmd.Run(); err != nil {
		fmt.Println("Failed to clean up coverage file:", err)
	}
}
