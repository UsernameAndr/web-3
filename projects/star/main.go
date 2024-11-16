package main
import "fmt"
func main() {
	var stroka string
	fmt.Scan(&stroka)
	for i, v := range stroka {
	if i != 0 {
		fmt.Print("*")
	}
		fmt.Printf("%c", v)
	}
}