package balancer

import (
	"fmt"

	"github.com/Emyrk/LendingBot/src/core/poloniex"
)

const (
	PoloniexExchange int = iota
	BitfinexExchange
)

var Currencies map[int][]string

var (
	MaxLendAmt map[int]map[string]float64
	MinLendAmt map[int]map[string]float64
)

func init() {
	Currencies = make(map[int][]string)
	Currencies[PoloniexExchange] = []string{"BTC", "BTS", "CLAM", "DOGE", "DASH", "LTC", "MAID", "STR", "XMR", "XRP", "ETH", "FCT"}
	Currencies[BitfinexExchange] = []string{"BTC", "ETH", "ETC", "ZEC", "XMR", "LTC", "DASH", "USD"}

	MaxLendAmt = make(map[int]map[string]float64)
	MaxLendAmt[PoloniexExchange] = make(map[string]float64)
	MaxLendAmt[PoloniexExchange]["BTC"] = .1
	MaxLendAmt[PoloniexExchange]["BTS"] = 1
	MaxLendAmt[PoloniexExchange]["CLAM"] = 20
	MaxLendAmt[PoloniexExchange]["DOGE"] = 1
	MaxLendAmt[PoloniexExchange]["DASH"] = 1
	MaxLendAmt[PoloniexExchange]["LTC"] = 1
	MaxLendAmt[PoloniexExchange]["MAID"] = 1
	MaxLendAmt[PoloniexExchange]["STR"] = 1
	MaxLendAmt[PoloniexExchange]["XMR"] = 1
	MaxLendAmt[PoloniexExchange]["XRP"] = 1
	MaxLendAmt[PoloniexExchange]["ETH"] = .2
	MaxLendAmt[PoloniexExchange]["FCT"] = 200

	MinLendAmt = make(map[int]map[string]float64)
	MinLendAmt[PoloniexExchange] = make(map[string]float64)
	MinLendAmt[PoloniexExchange]["BTC"] = .01
	MinLendAmt[PoloniexExchange]["BTS"] = 1
	MinLendAmt[PoloniexExchange]["CLAM"] = 10
	MinLendAmt[PoloniexExchange]["DOGE"] = 1
	MinLendAmt[PoloniexExchange]["DASH"] = 1
	MinLendAmt[PoloniexExchange]["LTC"] = 1
	MinLendAmt[PoloniexExchange]["MAID"] = 1
	MinLendAmt[PoloniexExchange]["STR"] = 1
	MinLendAmt[PoloniexExchange]["XMR"] = 1
	MinLendAmt[PoloniexExchange]["XRP"] = 1
	MinLendAmt[PoloniexExchange]["ETH"] = 1
	MinLendAmt[PoloniexExchange]["FCT"] = 100
}

func GetExchangeString(exch int) string {
	switch exch {
	case PoloniexExchange:
		return "Poloniex"
	case BitfinexExchange:
		return "Bitfinex"
	}
	return fmt.Sprintf("Unknown {%d}", exch)
}

type OrderDensity struct {
	Amount float64
	Rate   float64

	Orders []poloniex.PoloniexLoanOrder
}

func GetDensityOfLoans(orders *poloniex.PoloniexLoanOrders) []OrderDensity {
	all := make([]OrderDensity, 2002)
	for _, order := range orders.Offers {
		if int(order.Rate*100000) > 2000 {
			all[2001].AddOrder(order)
		} else {
			all[int(order.Rate*100000)].AddOrder(order)
		}
	}
	return all
}

func (od *OrderDensity) AddOrder(order poloniex.PoloniexLoanOrder) {
	prev := od.Amount
	od.Amount = od.Amount + order.Amount
	od.Rate = od.Rate * prev / od.Amount
	od.Rate = od.Rate + (order.Rate*order.Amount)/od.Amount
	//od.Rate = order.Rate
	od.Orders = append(od.Orders[:], order)
}

func (od *OrderDensity) String() string {
	str := fmt.Sprintf("Loan Density Info of %d loans - Total Coin: %.4f at AVG: %.4f%s", len(od.Orders), od.Amount, od.Rate*100, "%")
	return str
}
