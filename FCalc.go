package main 

import "fmt"

func main() {
    
    var a float64
    var x rune
    var b float64
    fmt.Scanf("%f%c%f", &a, &x, &b)
    
    if x == '/' && b != 0  {
        fmt.Println("Ваш ответ:", float64(a / b))
    } else if x == '/' && b == 0 {
        fmt.Println("Делить на ноль нельзя!")
    } else if x == '-' {
        fmt.Println("Ваш ответ:", float64(a - b))
    } else if x == '+' {
        fmt.Println("Ваш ответ:", float64(a + b))
    } else if x == '*' {
        fmt.Println("Ваш ответ:", float64(a * b))
    } else if x == '%' {
        fmt.Println("Ваш ответ:", int(a) % int(b))
    } else {
        fmt.Println("Рано, дружочек, калькулятор к такому не готов")
        }
    }
}
