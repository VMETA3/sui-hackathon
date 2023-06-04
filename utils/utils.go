package utils

import (
	"fmt"
	"hash/crc32"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func GetProjectRoot() (path string, err error) {
	path, err = os.Getwd()
	return
}

func GetENV(Conf string) string {
	env := os.Getenv(Conf)
	if env == "" {
		log.Fatalln("There is no " + Conf + " in the file env")
	}
	return env
}

func GetRandomString(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

func GetHash32String() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.FormatUint(uint64(crc32.ChecksumIEEE([]byte(GetRandomString(rand.Intn(10000000))))), 10)
}
