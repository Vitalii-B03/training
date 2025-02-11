package main

import "fmt"

const (
	ProductCocaCola = iota
	ProductPepsi
	ProductSprite
)

type Product struct {
	ProductID     int
	Sells         []float64
	Buys          []float64
	CurrentPrice  float64
	ProfitPercent float64
}

type Profitable interface {
	SetProduct(p *Product)
	GetAverageProfit() float64
	GetAverageProfitPercent() float64
	GetCurrentProfit() float64
	GetDifferenceProfit() float64
	GetAllData() []float64
	Average(prices []float64) float64
	Sum(prices []float64) float64
}

func (s *StatisticProfit) SetProduct(p *Product) {
	s.product = p
}
func (s *StatisticProfit) GetAverageProfit() float64 {
	if s.getAverageProfit != nil {
		return s.getAverageProfit()
	}
	return 0
}
func (s *StatisticProfit) GetAverageProfitPercent() float64 {
	if s.getAverageProfitPercent != nil {
		return s.getAverageProfitPercent()
	}
	return 0
}
func (s *StatisticProfit) GetCurrentProfit() float64 {
	if s.getCurrentProfit != nil {
		return s.getCurrentProfit()
	}
	return 0
}
func (s *StatisticProfit) GetDifferenceProfit() float64 {
	if s.getDifferenceProfit != nil {
		return s.getDifferenceProfit()
	}
	return 0
}
func (s *StatisticProfit) GetAllData() []float64 {
	if s.getAllData != nil {
		return s.getAllData()
	}
	return nil
}
func (s *StatisticProfit) Average(prices []float64) float64 {
	if len(prices) == 0 {
		return 0
	}
	return s.Sum(prices) / float64(len(prices))
}
func (s *StatisticProfit) Sum(prices []float64) float64 {
	var sum float64
	for _, price := range prices {
		sum += price
	}
	return sum
}

type StatisticProfit struct {
	product                 *Product
	getAverageProfit        func() float64
	getAverageProfitPercent func() float64
	getCurrentProfit        func() float64
	getDifferenceProfit     func() float64
	getAllData              func() []float64
}

func NewStatisticProfit(opts ...func(*StatisticProfit)) Profitable {
	statPrifit := &StatisticProfit{}
	for _, opt := range opts {
		opt(statPrifit)
	}
	return statPrifit
}

// WithAverageProfit sets the average profit of the product
//
// Average Profit = Average Sells - Average Buys
func WithAverageProfit(s *StatisticProfit) {
	var resBuy float64
	var resSell float64
	s.getAverageProfit = func() float64 {
		resBuy = s.Average(s.product.Buys)
		resSell = s.Average(s.product.Sells)
		return resSell - resBuy
	}
}

// WithAverageProfitPercent sets the average profit percent of the product
//
// Average Profit Percent = Average Profit / Average Buys * 100
func WithAverageProfitPercent(s *StatisticProfit) {
	s.getAverageProfitPercent = func() float64 {
		avgProfit := s.getAverageProfit()
		avgBuys := s.Average(s.product.Buys)
		if avgBuys == 0 {
			return 0
		}
		return (avgProfit / avgBuys) * 100
	}
}

// WithCurrentProfit sets the current profit of the product
//
// Current Profit = Current Price - Current Price * (100 - Profit Percent) / 100
func WithCurrentProfit(s *StatisticProfit) {
	s.getCurrentProfit = func() float64 {
		return s.product.CurrentPrice - s.product.CurrentPrice*(100-s.product.ProfitPercent)/100
	}
}

// WithDifferenceProfit sets the difference profit of the product
//
// Difference Profit = Current Price - Average Sells
func WithDifferenceProfit(s *StatisticProfit) {
	s.getDifferenceProfit = func() float64 {
		return s.product.CurrentPrice - s.Average(s.product.Sells)
	}
}

func WithAllData(s *StatisticProfit) {
	s.getAllData = func() []float64 {
		res := make([]float64, 0, 4)
		if s.getAverageProfit != nil {
			res = append(res, s.getAverageProfit())
		}
		if s.getAverageProfitPercent != nil {
			res = append(res, s.getAverageProfitPercent())
		}
		if s.getCurrentProfit != nil {
			res = append(res, s.getCurrentProfit())
		}
		if s.getDifferenceProfit != nil {
			res = append(res, s.getDifferenceProfit())
		}
		return res
	}
}

func main() {
	product := &Product{
		ProductID:     ProductCocaCola,
		Sells:         []float64{10, 20, 30},
		Buys:          []float64{5, 15, 25},
		CurrentPrice:  35,
		ProfitPercent: 10,
	}
	product.ProductID = ProductCocaCola
	product.ProductID = ProductPepsi
	product.ProductID = ProductSprite

	statProfit := NewStatisticProfit(
		WithAverageProfit,
		WithAverageProfitPercent,
		WithCurrentProfit,
		WithDifferenceProfit,
		WithAllData,
	).(*StatisticProfit)

	statProfit.SetProduct(product)
	fmt.Println(statProfit.GetAverageProfit())
	fmt.Println(statProfit.GetAverageProfitPercent())
	fmt.Println(statProfit.GetCurrentProfit())
	fmt.Println(statProfit.GetDifferenceProfit())
	fmt.Println(statProfit.GetAllData())
}
