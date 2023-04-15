package main

// func main() {
// 	ch := make(chan string, 1)

// 	ch <- "cat"
// 	for i := 0; i < 9; i++ {
// 		switch <-ch {
// 		case "cat":
// 			fmt.Println("cat")
// 			ch <- "fish"
// 			break
// 		case "fish":
// 			fmt.Println("fish")
// 			ch <- "dog"
// 			break
// 		case "dog":
// 			fmt.Println("dog")
// 			ch <- "cat"
// 			break
// 			// default:
// 			// 	ch <- "cat"
// 		}
// 	}

// 	// ch <- "cat"
// 	time.Sleep(3 * time.Second)
// }
