package src

import (
	"container/heap"
	"log"
	"sync"
	"time"
)

type CompanyPrice struct {
	company  int
	price    int
	id       int
	date     time.Time
	isTopTen bool
}

type OriginPriceContainer struct {
	N              int
	meanPrice      int
	companyToPrice map[int]*CompanyPrice
	topNPrices     *MaxHeap
	otherPrices    *MinHeap
	mutex          *sync.Mutex
}

func NewOriginPriceContainer(n int) *OriginPriceContainer {
	return &OriginPriceContainer{
		N:              n,
		meanPrice:      0,
		companyToPrice: map[int]*CompanyPrice{},
		topNPrices:     &MaxHeap{},
		otherPrices:    &MinHeap{},
		mutex:          &sync.Mutex{},
	}
}

func (this *OriginPriceContainer) InsertPrice(priceBody PriceBody) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// parse date
	var t time.Time
	t, err := time.Parse("2006-01-02", priceBody.Date)
	if err != nil {
		log.Fatal("post", "failed to parse the date", err)
	}

	// update mean price
	defer func() {
		sum := 0
		for _, v := range this.topNPrices.MinHeap {
			sum += v.price
		}
		this.meanPrice = sum / this.topNPrices.Len()
	}()

	oldPrice, ok := this.companyToPrice[priceBody.Company]
	if !ok {
		cp := &CompanyPrice{
			company: priceBody.Company,
			price:   priceBody.Price,
			id:      0,
			date:    t,
		}
		this.companyToPrice[priceBody.Company] = cp

		// for simplicity always add to topNPrices then move head of topNPrices to otherPrices
		cp.isTopTen = true
		heap.Push(this.topNPrices, cp)
		if this.topNPrices.Len() > this.N {
			p1 := heap.Pop(this.topNPrices).(*CompanyPrice)
			p1.isTopTen = false
			heap.Push(this.otherPrices, p1)
		}

		return
	}

	// in case new date is older than the old date -> do not update
	if oldPrice.date.After(t) {
		return
	}

	// update price and fix dedicated heap
	oldPrice.date = t
	oldPrice.price = priceBody.Price
	if oldPrice.isTopTen {
		heap.Fix(this.topNPrices, oldPrice.id)
	} else {
		heap.Fix(this.otherPrices, oldPrice.id)
	}

	// move head of topNPrices to otherPrices if needed
	if this.otherPrices.Len() != 0 {
		if this.topNPrices.MinHeap[0].price > (*this.otherPrices)[0].price {
			p1 := heap.Pop(this.otherPrices).(*CompanyPrice)
			p2 := heap.Pop(this.topNPrices).(*CompanyPrice)

			p1.isTopTen = true
			heap.Push(this.topNPrices, p1)

			p2.isTopTen = false
			heap.Push(this.otherPrices, p2)
		}
	}
}

type Storage struct {
	originMap map[string]*OriginPriceContainer
}

func NewStorage() *Storage {
	return &Storage{
		originMap: map[string]*OriginPriceContainer{},
	}
}

func (s *Storage) InsertPrice(priceBody PriceBody) {
	// insert price into storage
	if _, ok := s.originMap[priceBody.Origin]; !ok {
		s.originMap[priceBody.Origin] = NewOriginPriceContainer(10)
	}
	s.originMap[priceBody.Origin].InsertPrice(priceBody)
}

func (s *Storage) GetAveragePrices() map[string]int {
	// get average values for each origin from storage
	result := map[string]int{}
	for k, v := range s.originMap {
		result[k] = v.meanPrice
	}
	return result
}
