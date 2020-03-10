package rlog

import (
  "fmt"
)

type Config struct {
    fname string
    maxsize int32
}

func init() {
    fmt.Println("RLOG - rotate log")
}



