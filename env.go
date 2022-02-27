package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type WrpEnv struct {
	InnerHost string
	ServePort uint16
}

func (x *WrpEnv) ServrAddr() string {
	return fmt.Sprintf(":%d", x.ServePort)
}

var noEnv = WrpEnv{
	InnerHost: "",
	ServePort: 0,
}

var we *WrpEnv = &noEnv

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[ENV] Error loading .env file: %v\n", err)
	}

	env := WrpEnv{}
	env.InnerHost = os.Getenv("WRP_INNER_HOST")
	if env.InnerHost == "" {
		log.Fatalf("[ENV] WRP_INNER_HOST is empty")
	}

	val := os.Getenv("WRP_SERVE_PORT")
	num, err := strconv.Atoi(val)
	if err != nil || num <= 0 {
		log.Fatalf("[ENV] WRP_SERVE_PORT is invalid: %s", val)
	}
	env.ServePort = uint16(num)

	we = &env
}

func GetEnv() *WrpEnv {
	return we
}
