package test

import (
	"fmt"
	"testing"
	"os"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
)

func TestAnsibleEc2Instance(t *testing.T) {
	t.Parallel()

	data, err := os.ReadFile("endpoint_test_ip.txt")
	if err != nil {
		fmt.Println("Error reading text file:", err)
		return
	}

	fmt.Printf("Found IP Address of Jenkins server %s\n", string(data))

	url := fmt.Sprintf("http://%s:8080", data)

	http_helper.HttpGet(t, url, nil)
}
