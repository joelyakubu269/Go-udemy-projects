package main
import(
	"math"
	"fmt"
)

func main() {
	const inflationRate float64 = 2.5
	var investedAmount  float64 
	var expectedReturnRate float64 
	var years float64 
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investedAmount)
	fmt.Print("years: ")
	fmt.Scan(&years)
	fmt.Print("expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	var futureValue = (investedAmount) * math.Pow(1 + expectedReturnRate/100, (years))
	var futureRealValue = futureValue/ math.Pow(1 +inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
