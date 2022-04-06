//Exemplo para somar numeros inteiros com loop 50x

package main
import "fmt"

func main() {
    var n, sum = 10, 0
		n = 50;
  
    for i := 1 ; i <= n; i++ {
      sum += i    // sum = sum + i  
    }

    fmt.Println("sum =", sum)
}