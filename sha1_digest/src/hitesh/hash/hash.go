package main

import (
    "fmt"
    "crypto/sha1"
    "io"
    "os"
)

func main () {
    if os.Args == nil || len (os.Args) < 2 {
        fmt.Println ("Add string to be converted")
        os.Exit (1)
    }

    cert_path := os.Args[1]
    
    f, err := os.Open(cert_path)
    defer f.Close()
    if err != nil {
        os.Exit (1)
    }

    b1 := make([]byte, 10000000)
    n1, err := f.Read(b1)
    
    if err != nil {
        os.Exit (1)
    }
    
    fmt.Printf("%d bytes: %s\n", n1, string(b1))

    hash_val := computeHash(string(b1)) //plain_string)
    fmt.Printf ("The hash is % x\n", hash_val)
}

func computeHash (cert_str string) []byte {
    hash_var := sha1.New()
    io.WriteString(hash_var, cert_str)
    return hash_var.Sum(nil)
}
