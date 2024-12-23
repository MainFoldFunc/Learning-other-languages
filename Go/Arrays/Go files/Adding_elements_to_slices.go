package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func addingInt(lenSlice int) ([]int, error) {
	slice := make([]int, lenSlice)
	for i := 0; i < lenSlice; i++ {
		var value int
		fmt.Printf("Enter an integer value for position %d: ", i)
		_, err := fmt.Scan(&value)
		if err != nil {
			return nil, errors.New("invalid input")
		}
		slice[i] = value
	}
	return slice, nil
}

func addingString(lenSlice int) ([]string, error) {
	slice := make([]string, lenSlice)
	for i := 0; i < lenSlice; i++ {
		var value string
		fmt.Printf("Enter a string value for position %d: ", i)
		_, err := fmt.Scan(&value)
		if err != nil {
			return nil, errors.New("invalid input")
		}
		slice[i] = value
	}
	return slice, nil
}

func addingFloat(lenSlice int) ([]float64, error) {
	slice := make([]float64, lenSlice)
	for i := 0; i < lenSlice; i++ {
		var value float64
		fmt.Printf("Enter a float value for position %d: ", i)
		_, err := fmt.Scan(&value)
		if err != nil {
			return nil, errors.New("invalid input")
		}
		slice[i] = value
	}
	return slice, nil
}

func main() {
	var dataType string
	var lenSlice int
	var exit string

	fmt.Print("What type should be used in the slice (int, string, float64): ")
	fmt.Scan(&dataType)
	fmt.Print("How long should the slice be: ")
	fmt.Scan(&lenSlice)

	// Clear buffer
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	var slice interface{}
	var err error

	switch dataType {
	case "int":
		slice, err = addingInt(lenSlice)
	case "string":
		slice, err = addingString(lenSlice)
	case "float64":
		slice, err = addingFloat(lenSlice)
	default:
		fmt.Println("Unsupported type. Please use int, string, or float64.")
		return
	}

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("This is your slice:", slice)
	}

	// Wait for Enter key to exit
	fmt.Print("Press enter to exit... ")
	fmt.Scan(&exit)
}

