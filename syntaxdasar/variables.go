package syntaxdasar  

import "fmt"

func DemoVariables() {
    fmt.Println("\n--- VARIABLES ---")
    
    
    var name string = "naff"      // Eksplisit
    age := 25                     // Implisit (type inference)
    var city string               // Zero value
    country := "Indonesia"        // Short declaration
    
    city = "Jakarta"              // Assign nilai
    
    // Multiple declaration
    var x, y, z int = 1, 2, 3
    a, b, c := 10, 20.5, "hello"
    
    // Constants
    const pi = 3.14159
    const appName = "Belajar Go"
    
    fmt.Printf("Name: %s, Age: %d\n", name, age)
    fmt.Printf("City: %s, Country: %s\n", city, country)
    fmt.Printf("X:%d, Y:%d, Z:%d\n", x, y, z)
    fmt.Printf("A:%d, B:%f, C:%s\n", a, b, c)
    fmt.Printf("PI: %f, App: %s\n", pi, appName)
}