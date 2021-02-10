package main

import (
    "fmt"
    "log"
    "net"
)

func localAddresses() {
    ifaces, err := net.Interfaces()
    if err != nil {
        log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
        return
    }
    for _, i := range ifaces {
        addrs, err := i.Addrs()
        if err != nil {
            log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
            continue
        }
        for _, a := range addrs {
            log.Printf("%v %v\n", i.Name, a)
        }
    }
}


func getMacAddr() ([]string, error) {
    ifas, err := net.Interfaces()
    if err != nil {
        return nil, err
    }
    var as []string
    for _, ifa := range ifas {
        a := ifa.HardwareAddr.String()
        if a != "" {
            as = append(as, a)
        }
    }
    return as, nil
}


func main() {
    as, err := getMacAddr()
    if err != nil {
        log.Fatal(err)
    }
    for _, a := range as {
        fmt.Println(a)
    }
    localAddresses()

}