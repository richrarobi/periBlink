// blink5.go
// Rich Robinson
// Sept 2018

package main

import (
    "os"
    "fmt"
    "time"
    "os/signal"
    "syscall"
    "math/rand"
    "github.com/richrarobi/periBlink"
)

func delay(ms int) {
    time.Sleep(time.Duration(ms) * time.Millisecond)
}

func main() {
    running := true
// initialise getout
    signalChannel := make(chan os.Signal, 2)
    signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
    go func() {
        sig := <-signalChannel
        switch sig {
        case os.Interrupt:
            fmt.Println("Stopping on Interrupt")
            running = false
            return
        case syscall.SIGTERM:
            fmt.Println("Stopping on Terminate")
            running = false
            return
        }
    }()

    periBlink.Setup()
    periBlink.SetLuminance(1)
    periBlink.Clear()
    periBlink.Show()

    for running {
        pixel := rand.Intn(8)
// note the int parameter for brightness 0 to 31 (not a float)
// also reduced brightness looks better hence only up to 5 here
        periBlink.SetPixel( pixel, rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(3) )
        periBlink.Show()
        r,g,b,l := periBlink.GetPixel( pixel)
        fmt.Println("getPixel", pixel, r,g,b,l )
        delay(60)
    }
    
    fmt.Println("Stopping")
    periBlink.Exit()
}
