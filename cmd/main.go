package main

import (
	"encoding/hex"
	"fmt"
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

	fmt.Println(bytes)
	result := &Result{}

	for i := 0; i < len(bytes); i++ {
		temperatureBytes := bytes[2:4] // для подсчета значения температуры
		temperatureInt := int(temperatureBytes[0])<<8 + int(temperatureBytes[1])
		switch bytes[i] {
		case 0x03: // Температура

			if bytes[i+1] == 0x67 { // если тип канала Temperature соответствует
				if bytes[i+2] == 0xFF { // отрицательное значение температуры

					temperatureInt = ^temperatureInt + 1
					result.Temperature = float64((temperatureInt ^ 0xFFFF + 1))
				} else { // положительное значение температуры
					result.Temperature = float64(int16(bytes[3])<<8|int16(bytes[2])) / 10.0
				}
			} else {
				fmt.Println("ERROR: Wrong channel type for Temperature")
			}

		case 0x04: // Влажность

			if bytes[i+1] == 0x68 { // если тип канала Humidity соответствует
				result.Humidity = float64(int(bytes[6])) / 2.0 // Формула: Humidity = 130 / 2.0
			} else {
				fmt.Println("ERROR: Wrong channel type for Humidity")
			}

		case 0x06: // Магнитный статус
			if bytes[i+1] == 0x00 {
				if bytes[i+2] == 0x01 { // проверяем значение канала
					result.MagneticStatus = "Open"
				} else {
					result.MagneticStatus = "Closed"
				}
			} else {
				fmt.Println("ERROR: Wrong channel type for Magnetic Status")
			}
		}
	}
	return result, nil
}

func main() {

	hexStr := "0367F600046882060001" //
	// hexStr := "0367FFE1046882060001" // hex код с отрицательной температурой (-31)
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
