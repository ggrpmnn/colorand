package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var r, g, b int

func main() {
	rnd := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	r = rnd.Intn(256)
	g = rnd.Intn(256)
	b = rnd.Intn(256)

	baseHexStr := fmt.Sprintf(`%02X%02X%02X`, r, g, b)
	baseHexNum, err := strconv.ParseUint(baseHexStr, 16, 32)
	if err != nil {
		log.Printf("failed to convert hex string to num!")
	}
	invertHexNum := 0xFFFFFF - baseHexNum
	invertHexStr := fmt.Sprintf("%06X", invertHexNum)
	content := fmt.Sprintf("<html style=\"background: #%s; font-size: 256px; font-family: Verdana; color: #ffffff; -webkit-text-stroke: 2px #%s; text-align: center; padding-top: 15%%;\">#%s</html>", baseHexStr, invertHexStr, baseHexStr)
	fileName := fmt.Sprintf("%s.html", baseHexStr)

	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("failed to create output file!")
	}
	file.Write([]byte(content))

	openCmd := ""
	switch runtime.GOOS {
	case "darwin":
		openCmd = "open"
	case "linux":
		openCmd = "xdg-open"
	case "windows":
		openCmd = "start"
	}
	err = exec.Command(openCmd, fileName).Run()
	if err != nil {
		log.Printf("failed to run open command!")
	}

	// sleep to prevent file from being deleted too early
	time.Sleep(5 * time.Second)

	err = os.Remove(fileName)
	if err != nil {
		log.Printf("failed to remove temp file!")
	}
}
