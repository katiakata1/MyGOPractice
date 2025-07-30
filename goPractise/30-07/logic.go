// Count Failures Per Service
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Service struct {
	Name      string `json:"service"`
	Status    string `json:"status"`
	TimeStamp string `json:"timestamp"`
}

func unmarshalJson(fileName string) []Service {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error opening the file:", err)
		return nil
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("error closing the file: ", err)
			return
		}
	}()

	services := []Service{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var service Service
		if err := json.Unmarshal([]byte(line), &service); err != nil {
			fmt.Println("error unmarshaling json to struct: ", err)
			return nil
		}

		services = append(services, service)
	}

	return services
}

// func failureCount(services []Service) map[string]int {
// 	failures := map[string]int{}

// 	for _, service := range services {
// 		if service.Status == "fail" {
// 			if count, exist := failures[service.Name]; exist {
// 				failures[service.Name] = count + 1
// 			} else {
// 				failures[service.Name] = 1
// 			}
// 		}
// 	}

// 	return failures
// }

func failureCount(services []Service) map[string]int {
	failures := map[string]int{}

	for _, service := range services {
		if service.Status == "fail" {
			failures[service.Name]++
		}
	}

	return failures
}

func main() {
	file := "./file.jsonl"
	services := unmarshalJson(file)
	failures := failureCount(services)

	for name, count := range failures {
		fmt.Printf("For %s number of failures is: %d\n", name, count)
	}
}
