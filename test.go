package main

import (
	"fmt"
	"github.com/tarm/goserial"
	"strconv"
	"time"
)

func main() {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	 s, err := serial.OpenPort(c)

    if err != nil {
            fmt.Println(err)
    }
	    buf := make([]byte, 10)
	for {
	    n, err := s.Read(buf)
	    if err != nil {
		    fmt.Println(err)
	    }
	    _ = n
		low25, err := strconv.ParseInt(fmt.Sprintf("%d",buf[2]), 10, 32)
		high25, err := strconv.ParseInt(fmt.Sprintf("%d",buf[3]), 10, 32)
		calc25 := ((high25*256)+low25)/10
		fmt.Print("pm2.5: ")
		fmt.Println(calc25)
		low10, err := strconv.ParseInt(fmt.Sprintf("%d",buf[4]), 10, 32)
		high10, err := strconv.ParseInt(fmt.Sprintf("%d",buf[5]), 10, 32)
		calc10 := ((high10*256)+low10)/10
		fmt.Print("pm10: ")
		fmt.Println(calc10)
		_ = err
		time.Sleep(1 * time.Second)
	}
}
