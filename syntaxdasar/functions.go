package syntaxdasar

import "fmt"

func DemoFunctions() {
    fmt.Println("\n--- FUNGSI ---")
    
    result := add(10, 20)
    fmt.Printf("10 + 20 = %d\n", result)
    
    sum, difference := calculate(50, 30)
    fmt.Printf("Sum: %d, Difference: %d\n", sum, difference)
    
    // Variadic function
    total := sumNumbers(1, 2, 3, 4, 5)
    fmt.Printf("Sum of 1-5: %d\n", total)
}


func add(a int, b int) int {
    return a + b
}


func calculate(x, y int) (int, int) {
    return x + y, x - y
}


func sumNumbers(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}