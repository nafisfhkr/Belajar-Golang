package syntaxdasar  

import "fmt"

// Fungsi publik (huruf besar)
func DemoTypes() {
    fmt.Println("\n--- TIPE DATA ---")
    
    var a bool = true
    var b int8 = 5
    var c float32 = 3.14
    var d string = "Hi!"
    var e byte = 255     // alias uint8
    var f rune = 'A'     // alias int32
    
    fmt.Printf("Boolean: %v, Tipe: %T\n", a, a)
    fmt.Printf("Integer: %v, Tipe: %T\n", b, b)
    fmt.Printf("Float: %v, Tipe: %T\n", c, c)
    fmt.Printf("String: %v, Tipe: %T\n", d, d)
    fmt.Printf("Byte: %v, Tipe: %T\n", e, e)
    fmt.Printf("Rune: %v, Tipe: %T\n", f, f)
}

// Fungsi private (huruf kecil)
func internalFunction() {
    fmt.Println("Ini fungsi internal")
}