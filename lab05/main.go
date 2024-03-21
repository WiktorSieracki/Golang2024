package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.TrimSpace(text)

    words := strings.Fields(text)
    cmd := exec.Command(words[0], words[1:]...)
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(output))
}


// wpisać coś do programu na standardowe wejście
// wypisać coś z programu na standardowe wyjście
// popen