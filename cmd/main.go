package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

type Result struct {
	Temperature    float64
	Humidity       float64
	MagneticStatus string
}

func Decode(hexStr string) (*Result, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}

	// fmt.Println(bytes)
	result := &Result{}

	for i := 0; i < len(bytes); i++ {
		channel := bytes[i]
		var checkType string
		var conv int

		switch channel {
		case 0x03: // Температура
			checkType = strconv.FormatInt(int64(bytes[i+1]), 16) // проверяем тип канала
			conv, err = strconv.Atoi(checkType)
			if err != nil {
				log.Fatal("couldn't convert channel type")
			}
			if conv == 67 {
				result.Temperature = float64(int16(bytes[3])<<8|int16(bytes[2])) / 10.0
			} else {
				fmt.Println("wrong channel type for temperature")
			}

		case 0x04: // Влажность
			checkType = strconv.FormatInt(int64(bytes[i+1]), 16)
			conv, err = strconv.Atoi(checkType)
			if err != nil {
				log.Fatal("couldn't convert humidity channel type")
			}
			if conv == 68 {
				result.Humidity = float64(int(bytes[6])) / 2.0 // Формула: Humidity = 130 / 2.0
			} else {
				fmt.Println("wrong channel type for humidity")
			}

		case 0x06: // Магнитный статус
			if bytes[i+2] == 1 {
				result.MagneticStatus = "Open"
			} else {
				result.MagneticStatus = "Closed"
			}
		}
	}
	return result, nil
}

func main() {

	hexStr := "0367F600046882060001"
	result, err := Decode(hexStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Temperature: %.1f C\n", result.Temperature)
	fmt.Printf("Humidity: %.0f%%\n", result.Humidity)
	fmt.Printf("MagneticStatus: %s\n", result.MagneticStatus)
}

// Опциональная функция main для ручного ввода значений

// func main() {

// 	// reader := bufio.NewReader(os.Stdin)
// 	// fmt.Print("Введите значение в шестнадцатеричном формате: ")
// 	// hexStr, _ := reader.ReadString('\n')
// 	// hexStr = strings.TrimSpace(hexStr)
// 	// result, err := Decode(hexStr)
// 	// if err != nil {
// 	// 	fmt.Println("Ошибка:", err)
// 	// 	return
// 	// }

// 	// fmt.Printf("Temperature: %.1f C\n", result.Temperature)
// 	// fmt.Printf("Humidity: %.0f%%\n", result.Humidity)
// 	// fmt.Printf("MagneticStatus: %s\n", result.MagneticStatus)
// }
