package logging

import (
    "os"
    "log"
    "fmt"
)

func Write (data string, print bool) {
    f, _ := os.OpenFile("dex.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

    log.SetOutput(f)
    log.Println(data)

    if (print) {
        fmt.Println(data)
    }
}

func String(n int64) string {
    buf := [11]byte{}
    pos := len(buf)
    i := int64(n)
    signed := i < 0
    if signed {
        i = -i
    }
    for {
        pos--
        buf[pos], i = '0'+byte(i%10), i/10
        if i == 0 {
            if signed {
                pos--
                buf[pos] = '-'
            }
            return string(buf[pos:])
        }
    }
}
