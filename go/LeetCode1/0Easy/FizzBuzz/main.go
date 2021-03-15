package main
import "fmt"
import "strconv"

func fizzBuzz(n int) []string {
	ansArr := []string{}
	for i:=1; i<= n; i ++ {
	divisible3 := (i % 3 == 0)
	divisible5 := (i % 5 == 0)
	ans := ""
	if divisible3 {
		ans += "Fizz"
	}
	if divisible5 {
		ans += "Buzz"
	}

	if ans == "" {
		ans += strconv.Itoa(i)
	}
ansArr = append(ansArr, ans)


}
return ansArr
}

func main() {
	fmt.Printf("%v\n", fizzBuzz(15))
}