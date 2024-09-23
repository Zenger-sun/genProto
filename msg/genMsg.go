package msg

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
)

// todo
func GenMsg(path string) {
	file, err := os.Open(path)
	if err != nil {
		slog.Warn("[genMsg] open proto file err:", err.Error())
		return
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		fmt.Println(scan.Text())
	}

	return
}
