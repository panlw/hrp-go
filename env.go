package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type HrpEnv struct {
	InnerHost string
	ServePort uint16
}

func (x *HrpEnv) ServrAddr() string {
	return fmt.Sprintf(":%d", x.ServePort)
}

var he *HrpEnv

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("[ENV] failed to load .env file: %v\n", err)
	}

	env := HrpEnv{}
	env.InnerHost = os.Getenv("HRP_INNER_HOST")
	if env.InnerHost == "" {
		log.Fatalf("[ENV] HRP_INNER_HOST is empty")
	}

	val := os.Getenv("HRP_SERVE_PORT")
	num, err := strconv.Atoi(val)
	if err != nil || num <= 0 {
		log.Fatalf("[ENV] HRP_SERVE_PORT is invalid: %s", val)
	}
	env.ServePort = uint16(num)

	he = &env
}

func GetEnv() *HrpEnv {
	return he
}
