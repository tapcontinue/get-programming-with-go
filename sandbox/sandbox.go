// import(
//   "fmt"
// )
//
// var x = "hello"
//
// func main(){
//   fmt.Println("Hello")
//   fmt.Printf(fmt.Sprintf("Oh %s there", x))
// }

//
// package main
// import "fmt"
//
//
// func swap(x, y string)(string, string)
// {
//   return y, x
// }
//
// func main(){
//   a, b, :=swap("hello", "world")
//   fmt.Println(a, b)
// }

// NEXY EXAMPLE

// package main
//
// import "fmt"
//
// func split(sum int) (x, y int){
//   x = sum * 4 / 9
//   y = sum - x
//   return
// }
//
// func main(){
//   fmt.Println(split(17))
// }

// package main
// import "fmt"
//
// func main(){
//   var year int = 2024
//   var age int = 43
//   fmt.Println(age)
//   fmt.Println("Hello World")
//   fmt.Printf("In %v I will be %4v\n", year, age)
// }

// How long does it take to get to Mars?
// package main
//
// import "fmt"
//
// func main() {
//     const lightSpeed = 299792 // km/s
//     var distance = 56000000   // km
//
//     fmt.Println(distance/lightSpeed, "seconds")
//
//     distance = 401000000
//     fmt.Println(distance/lightSpeed, "seconds")
//  }

// package main
//
// import (
//     "fmt"
//     "math/rand"
// )
//
// func main() {
//     var num = rand.Intn(345000001) + 5000000
//     fmt.Println(num)
//
//     num = rand.Intn(10) + 1
//     fmt.Println(num)
// }

//Malacandra
// package main
//
// import (
//   "fmt"
// )
//
// func main(){
//   var distance = 56_000_000 //km
//   var timeLimit = 28 //days
//   var speed = distance / (timeLimit * 24)
//   fmt.Println(speed)
// }

// package main
//
// import (
//     "fmt"
//     "strings"
// )
//
// func main() {
//     fmt.Println("You find yourself in a dimly lit cavern.")
//
//     var command = "walk outside"
//     var exit = strings.Contains(command, "outside")
//
//     fmt.Println("You leave the cave:", exit)
//  }

// package main
//
// import "fmt"
//
// func main(){
//   fmt.Println("There is a sign near the entrance that reads 'No Minors'.")
//
//   var age = 41
//   var minor = age < 18
//
//   fmt.Printf("At age %v, am I a minor? %v\n", age, minor)
// }

// Quick check

// package main
//
// import "fmt"
//
// func main(){
//   var location = "arcade"
//
//   if location == "cave"{
//     fmt.Println("You are in a cave")
//   } else if location == "entrance"{
//     fmt.Println("You are at the entrance")
//   } else{
//     fmt.Println("You are lost")}
// }

// Leap year checker

// package main
//
// import "fmt"
//
// func main(){
//
//   var year = 2100
//   if (year%4 == 0 && year%100 != 0) || year%400 == 0{
//     fmt.Println("Leap year")
//   } else{
//     fmt.Println("Not a leap year")
// }
// }

// Torch checker
// package main
// import "fmt"
//
// func main(){
// var haveTorch = true
// var litTorch = false
//
// if !haveTorch || !litTorch {
//     fmt.Println("Nothing to see here.")
// }
// }

// Leap year checker v2

// package main
// import "fmt"
//
// func main(){
//
//   var year = 2100
//   if year%400 == 0 || (year%4 == 0 && year%100 != 0){
//     fmt.Println("Leap year")
//   } else{
//     fmt.Println("Not a leap year")
// }
// }

// The switch statement v1
// package main
// import "fmt"
//
//
// func main(){
//   fmt.Println("There is a cavern entrance here and a path to the east.")
//   var command = "go inside"
//
//   switch command {
//   case "go east":
//       fmt.Println("You head further up the mountain.")
//   case "enter cave", "go inside":
//       fmt.Println("You find yourself in a dimly lit cavern.")
//   case "read sign":
//       fmt.Println("The sign reads 'No Minors'.")
//   default:
//       fmt.Println("Didn't quite get that.")
//   }
//  }

// switch statement v2

// package main
//
// import "fmt"
//
// func main() {
//     day := "Monday"
//
//     switch day {
//     case "Monday":
//         fmt.Println("It's Monday!")
//     case "Tuesday":
//         fmt.Println("It's Tuesday!")
//     case "Wednesday":
//         fmt.Println("It's Wednesday!")
//     case "Thursday":
//         fmt.Println("It's Thursday!")
//     case "Friday":
//         fmt.Println("It's Friday!")
//     case "Saturday":
//         fmt.Println("It's Saturday!")
//     case "Sunday":
//         fmt.Println("It's Sunday!")
//     default:
//         fmt.Println("Unknown day!")
//     }
// }

//switch fallthough example
// package main
// import "fmt"
//
// func main(){
// var room = "lake"
//
// switch {
// case room == "cave":
//     fmt.Println("You find yourself in a dimly lit cavern.")
// case room == "lake":
//     fmt.Println("The ice seems solid enough.")
//     fallthrough
// case room == "underwater":
//     fmt.Println("The water is freezing cold.")
// }
// }

// Repetion w/ loops

// package main
//
// import (
//     "fmt"
//     "time"
// )
//
// func main() {
//     var count = 10
//
//     for count > 0 {
//         fmt.Println(count)
//         time.Sleep(time.Second)
//         count--
//     }
//     fmt.Println("Liftoff!")
// }

// Example 2

// package main
// import (
//   "fmt"
//   "math/rand"
// )
//
// func main(){
// var degrees = 0
//
// for {
//     fmt.Println(degrees)
//
//     degrees++
//     if degrees >= 360 {
//         degrees = 0
//         if rand.Intn(2) == 0 {
//             break
//         }
//     }
// }
// }

// Quick check 3.6
// package main
//
// import (
//   "fmt"
//   "math/rand"
//   "time"
// )
//
// func main(){
//     var count = 10
//
//     for count > 0 {
//         fmt.Println(count)
//         time.Sleep(time.Second)
//         if rand.Intn(10) == 0 {
//             break
//         }
//         count--
//
//     }
//     if count == 0 {
//         fmt.Println("Liftoff!")
//     } else {
//         fmt.Println("Launch failed.")
//     }
// }

// Make a guessing game
// package main
//
// import (
//   "fmt"
//   "math/rand"
// )
//
// func main(){
//
//     my_rand := rand.Intn(100) + 1
//     computers_guess := rand.Intn(100) + 1
//     var num_guesses = 0
//
//     for{
//     if my_rand == computers_guess{
//         fmt.Println("-------------------------------")
//         fmt.Println("My random number was: ", my_rand)
//         fmt.Println("The computer guessed the number in:", num_guesses, "guesses")
//         break
//     } else {
//         num_guesses++
//         fmt.Printf("The computer guessed - %v\n", computers_guess)
//         computers_guess = rand.Intn(100) + 1
//     }
// }
// }

// Tickets to Mars
package main

func main() {
	x := 6
	x -= 1
	print(x)
}
