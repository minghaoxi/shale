package shale

import (
	"fmt"
	"testing"
)

func MockSampleShale() *Shale {
	opt := new(Shale)

	opt.demand_size = 2

	opt.supply_size = 4

	demand_size := opt.demand_size

	supply_size := opt.supply_size

	//设置供应节点
	opt.supply = NewVector(supply_size)

	opt.supply.Set(0, 300)

	opt.supply.Set(1, 300)

	opt.supply.Set(2, 300)

	opt.supply.Set(3, 300)

	//设置需求节点
	opt.demands = NewVector(demand_size)

	opt.demands.Set(0, 400)

	opt.demands.Set(1, 500)

	//设置供应节点对应的需求
	opt.supplyLinks = make([]*vector, supply_size)

	opt.supplyLinks[0] = NewVector(1)

	opt.supplyLinks[0].Set(0, 0)

	opt.supplyLinks[1] = NewVector(2)

	opt.supplyLinks[1].Set(0, 0)

	opt.supplyLinks[1].Set(1, 1)

	opt.supplyLinks[2] = NewVector(1)

	opt.supplyLinks[2].Set(0, 1)

	opt.supplyLinks[3] = NewVector(0)

	//设置需求节点对应的供应

	opt.demandLinks = make([]*vector, demand_size)

	opt.demandLinks[0] = NewVector(2)

	opt.demandLinks[0].Set(0, 0)

	opt.demandLinks[0].Set(1, 1)

	opt.demandLinks[1] = NewVector(2)

	opt.demandLinks[1].Set(0, 1)

	opt.demandLinks[1].Set(1, 2)

	opt.alpha = NewVector(demand_size)

	opt.alpha.Clear()

	opt.beta = NewVector(supply_size)

	opt.beta.Clear()

	return opt

}

func TestShale(t *testing.T) {
	MockSampleShale()

	fmt.Println("begin")

	algo := MockSampleShale()

	algo.Shale(30)

	x := algo.Dual2Primal()

	fmt.Println(x)
}
