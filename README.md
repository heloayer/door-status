
# DOOR STATUS
## Hexadecimal Code to Data

The program converts the given hex code to informative data of door status upon opening or closing:
- Temperature
- Humidity
- Magnetic Status

The given hex code in a format of EM300 (Environment
Monitoring Sensor)




## Used stack

- Golang

### Libraries:
- "encoding/hex"
-	"fmt"


## Installation

Run the program using 
```bash
 go run .
```

Use one of these hexadecimal codes to take out the data
```bash
 0367F600046882060001

Output: 
    Temperature: 24.6 C
    Humidity: 65%
    MagneticStatus: Open

 0367FFFFE1046882060001

Output: 
    Temperature: -31.0 C
    Humidity: 65%
    MagneticStatus: Open
```
## Autor
Ansar Issabekov
https://www.linkedin.com/in/ansarissabekov/


## Documentation

[EM300 Series (Environment Monitoring Sensor)](https://resource.milesight-iot.com/milesight/document/em300-series-user-guide-en.pdf) 

- Section “5.2 Sensor Data” explains how to decode the measurements the sensor sends.

[EM300-MCS Overview](https://www.milesight-iot.com/lorawan/sensor/em300-mcs/)



