// package main

// import (
//     "bufio"
//     "fmt"
//     "os"
//     "strings"
// )

// func main() {
//     reader := bufio.NewReader(os.Stdin)

//     fmt.Print("> ")

//     name, _ := reader.ReadString('\n')

//     parts := strings.Split(name, " ")


//     for i := 0; i < len(parts); i++ {
//         parts[i] = strings.Title(parts[i])
//     }

//     fmt.Println(parts)
//     fmt.Println("So this is real ", name)

    
// }

// package main

// import (
//     "fmt"
//     "os"
//     "bufio"
//     "strings"
//     "strconv"
// )

// func main(){
//     reader := bufio.NewReader(os.Stdin)

//     start:

//     fmt.Print("> ")

//     input, _ := reader.ReadString('\n')

//     parts := strings.Fields(input)

//     cmd := parts[0]

//     if cmd == "help" {
//         fmt.Println("Guidelines")
//         fmt.Println("For Addition (e.g. add 4 5)")
//         fmt.Println("For Subtraction (e.g. sub 9 5)")
//         fmt.Println("For Multiplication (e.g. mul 9 5)")
//         fmt.Println("For Division (e.g. div 9 5)")
//         goto start        
//     }

//     if len(parts) != 3 {
//         fmt.Println("Invalid input. Type 'help'")
//         goto start
//     }

    

//     a, err1 := strconv.ParseFloat(parts[1], 64)
//     b, err2 := strconv.ParseFloat(parts[2], 64)


//     if err1 != nil || err2 != nil {
//     fmt.Println("Invalid numbers. Type 'help'")
//     goto start
// }

        

//     switch cmd {
//     case "add": 
//     fmt.Println("✦ Result:", a + b)
//     goto start
//     case "sub":
//         fmt.Println("✦ Result:", a - b)
//         goto start
//     case "mul":
//         fmt.Println("✦ Result:", a * b)
//         goto start
//     case "div":
//         if b == 0 {
//             fmt.Println("Can't divide by zero")
//             goto start
//         }
//         fmt.Println("✦ Result:", a / b)
//         goto start
//     }


//     // for i, part := range parts {
//     //     fmt.Println(i, part)
//     // }

//    //
//    //  fmt.Println(parts)
// }


