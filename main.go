package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	go func() {
		for {
			status := generateStatus()
			writeToFile(status)
			time.Sleep(15 * time.Second)
		}
	}()

	go func() {
		for {
			status := readFromFile()
			displayStatus(status)
			time.Sleep(15 * time.Second)
		}
	}()

	select {}
}

func generateStatus() Status {
	return Status{
		Water: rand.Intn(100) + 1,
		Wind:  rand.Intn(100) + 1,
	}
}

func writeToFile(status Status) {
	data, err := json.Marshal(status)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = ioutil.WriteFile("data.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func readFromFile() Status {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return Status{}
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return Status{}
	}

	var status Status
	err = json.Unmarshal(data, &status)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return Status{}
	}

	return status
}

func displayStatus(status Status) {
	waterState := determineState(status.Water)
	windState := determineState(status.Wind)

	fmt.Printf("Water: %d meters (%s), Wind: %d meters/second (%s)\n", status.Water, waterState, status.Wind, windState)
}

func determineState(value int) string {
	switch {
	case value < 5:
		return "Aman"
	case value <= 8:
		return "Siaga"
	default:
		return "Bahaya"
	}
}
