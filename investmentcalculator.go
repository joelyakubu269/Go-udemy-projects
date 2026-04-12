package main
import(
	"math"
)

func main() {
	var investedAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10
	var futureValue = float64(investedAmount) * math.Pow(1 + expectedReturnRate/100, float64(years))
}
