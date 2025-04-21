package main
import "fmt"

func main(){
	const NMAX int = 100
	var A[NMAX]int
	var B[NMAX]int
	var i, n, val, hasil int
	fmt.Scan(&n, &val)

	for i =0; i < n; i++{
		fmt.Scan(&A[i])
	}
	for i = 0; i < n; i++{
		if A[i] != val{
			B[i] = A[i]
			hasil++
		}
	}
	fmt.Print(hasil)
}