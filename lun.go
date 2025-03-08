package main

import (
    "crypto/rand"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "math/big"
    "os"
    "time"
)

func getRandomNumber() *big.Int {
    min := big.NewInt(1000000000000000)
    max := big.NewInt(9000000000000000)
    n, _ := rand.Int(rand.Reader, max)
    return n.Add(n, min)
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: lun <seconds>")
        os.Exit(1)
    }

    duration, success := new(big.Int).SetString(os.Args[1], 10)
    if !success {
        fmt.Println("Please provide a valid number of seconds")
        os.Exit(1)
    }

    startTime := time.Now()
    fmt.Printf("\nStarted: %s\n\n", startTime.Format("2006-01-02 15:04:05"))

    var firstSum, lastSum *big.Int
    first := true
    targetTime := startTime

    for i := new(big.Int).Set(duration); i.Sign() > 0; i.Sub(i, big.NewInt(1)) {
        randomNum := getRandomNumber()
        epochTime := time.Now().Unix()
        sum := new(big.Int).Add(randomNum, big.NewInt(epochTime))

        if first {
            firstSum = new(big.Int).Set(sum)
            first = false
        }
        lastSum = new(big.Int).Set(sum)

        fmt.Printf("%s %d %s\n", randomNum, epochTime, sum)

        targetTime = targetTime.Add(time.Second)
        time.Sleep(time.Until(targetTime))
    }

    endTime := time.Now()
    adjustedEndTime := endTime.Add(-1 * time.Second)
    fmt.Printf("\nEnded: %s\n", adjustedEndTime.Format("2006-01-02 15:04:05"))

    elapsed := endTime.Sub(startTime)
    hours := int(elapsed.Hours())
    minutes := int(elapsed.Minutes()) % 60
    seconds := int(elapsed.Seconds()) % 60
    fmt.Printf("Time elapsed: %02d:%02d:%02d\n", hours, minutes, seconds)

    if !first {
        lucky := new(big.Int).Mul(firstSum, lastSum)
        fmt.Printf("Multiplying: %s * %s\n", firstSum.String(), lastSum.String())
        fmt.Printf("Your lucky number is: %s\n", lucky.String())

        luckyHash := sha256.Sum256([]byte(lucky.String()))
fmt.Printf("SHA256: %s\n", hex.EncodeToString(luckyHash[:]))
    } else {
        fmt.Printf("No numbers generated - lucky number unavailable\n")
    }
}
