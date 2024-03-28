package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
    // Uruchomienie zewnętrznego programu i przechwycenie stdout
    cmd := exec.Command("cat", "/etc/passwd")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Output:")
        fmt.Println(out.String())
    }

    // Przekazanie danych na stdin zewnętrznego programu
    input_data := "To jest testowy tekst\nLinia druga\n"
    cmd = exec.Command("grep", "testowy")
    cmd.Stdin = bytes.NewBufferString(input_data)
    var grepOut bytes.Buffer
    cmd.Stdout = &grepOut
    err = cmd.Run()
    if err != nil {
        fmt.Println("Nie znaleziono dopasowania:", err)
    } else {
        fmt.Println("Znaleziono dopasowanie:")
        fmt.Println(grepOut.String())
    }



	cmd = exec.Command("sleep", "5")
    err = cmd.Start()
    if err != nil {
        fmt.Print(err)
        return
    }

	
    fmt.Printf("PID %d\n", cmd.Process.Pid)
    err = cmd.Wait()
    if err != nil {
        fmt.Print(err)
    }
}
