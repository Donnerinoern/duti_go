package main

import (
    "fmt"
    "os"
    "os/exec"
    // "bufio"
    "strconv"
)

func parseFloats(x, y string) (float64, float64, error) {
    firstNum, err := strconv.ParseFloat(x, 64)
    secondNum, err := strconv.ParseFloat(y, 64)
    if err == nil {
        return firstNum, secondNum, err
    } else {
        return 0.0, 0.0, err
    }
}

func add(x, y float64) float64 {
    return x + y
}

func subtract(x, y float64) float64 {
    return x - y
}

func multiply(x, y float64) float64 {
    return x * y
}

func divide(x, y float64) float64 {
    return x / y
}

func main() {
    switch os.Args[1] {
    case "remember":
        var input string
        fmt.Scan(&input)
        fmt.Println(input)
    case "calc":
        firstNum, secondNum, err := parseFloats(os.Args[2], os.Args[4])
        if err != nil { // Print out error and exit
            fmt.Println(err)
            os.Exit(22)
        }
        var result float64
        switch operator := os.Args[3]; operator {
        case "+":
            result = add(firstNum, secondNum)
        case "-":
            result = subtract(firstNum, secondNum)
        case "x":
            result = multiply(firstNum, secondNum)
        case "/":
            result = divide(firstNum, secondNum)
        }
        fmt.Println(result)
    case "ip":
        cmd := exec.Command("curl", "-4", "ifconfig.co")
        out, err := cmd.Output()
        if err != nil {
            fmt.Println(err)
            os.Exit(5)
        }
        fmt.Print(string(out))
    // case "add":
    //     file, err := os.Open("/home/donnan/nixos/hosts/configuration.nix")
    //     if err != nil {
    //         fmt.Println(err)
    //         os.Exit(5)
    //     }
    //     reader := bufio.NewReader(file)
    //     for _, file := range reader {
    //         fmt.Println(reader.ReadString)
    //     }
    }        
}
