package shale

import (
	"fmt"
	"math"
)

type Shale struct {
	demands *vector

	supply *vector

	supplyLinks []*vector

	demandLinks []*vector

	alpha *vector

	beta *vector

	theta *vector

	demand_size int

	supply_size int
}

//获取平均概率
func (opt *Shale) GetTheta() *vector {
	//eligibles := NewVector(opt.demand_size)

	eligibles := make([]float32, opt.demand_size)

	for i := 0; i < opt.demand_size; i++ {

		eligibles[i] = 0

		for j := 0; j < int(opt.demandLinks[i].Length()); j++ {

			idx := int(opt.demandLinks[i].Get(j))

			eligibles[i] += opt.supply.Get(idx)

		}

	}

	for i := 0; i < opt.demand_size; i++ {

		eligibles[i] = opt.demands.Get(i) / eligibles[i]

	}

	ret := NewVector(opt.demand_size)

	ret.SetValues(eligibles)

	opt.theta = ret

	return ret

}

func (opt *Shale) Dual2Primal() [][]float32 {

	opt.GetTheta()

	x := make([][]float32, opt.supply_size)

	for i := 0; i < opt.supply_size; i++ {

		x[i] = make([]float32, opt.demand_size)

	}

	for i := 0; i < opt.supply_size; i++ {

		for j := 0; j < opt.supplyLinks[i].Length(); j++ {

			idx := int(opt.supplyLinks[i].Get(j))

			rate := opt.theta.Get(idx) * (1 + opt.alpha.Get(idx) - opt.beta.Get(i))

			x[i][idx] = float32(math.Max(float64(0), float64(rate)))

		}

		/*for j := 0; j < opt.demand_size; j++ {





			x[i][j] = float32(math.Max(float64(0), float64(rate)))
		}*/

	}

	return x

}

func (opt *Shale) GetContractDual() *vector {

	opt.GetTheta()

	for i := 0; i < opt.demand_size; i++ {
		a := opt.supply.VectorByIndex(opt.demandLinks[i].Values()).Sum()
		b1 := opt.supply.VectorByIndex(opt.demandLinks[i].Values())
		b2 := opt.beta.VectorByIndex(opt.demandLinks[i].Values())
		b := VecDotProduct(b1, b2)

		tmp1 := opt.demands.Get(i) + opt.theta.Get(i)*b - opt.theta.Get(i)*a

		tmp2 := opt.theta.Get(i) * a

		if tmp2 == 0 {
			opt.alpha.Set(i, 0)
		} else {
			opt.alpha.Set(i, tmp1/tmp2)
		}

	}

	return opt.alpha
}

func (opt *Shale) DumpAlpha() {
	fmt.Println(opt.alpha.Values())
}

func (opt *Shale) DumpBeta() {
	fmt.Println(opt.beta.Values())
}

func (opt *Shale) DumpTheta() {
	fmt.Println(opt.theta.Values())
}

func (opt *Shale) GetSupplyDual() *vector {

	opt.GetTheta()

	beta := NewVector(opt.supply_size)

	beta.Clear()

	for i := 0; i < opt.supply_size; i++ {

		var tmp2 float32 = 0

		var tmp1 float32 = 0

		for j := 0; j < opt.supplyLinks[i].Length(); j++ {
			demand_idx := int(opt.supplyLinks[i].Get(j))

			tmp2 += opt.theta.Get(demand_idx)

		}

		//fmt.Println("begin ", i, "  tmp2:", tmp2)

		s_theta := opt.theta.VectorByIndex(opt.supplyLinks[i].Values())

		a_theta := opt.alpha.VectorByIndex(opt.supplyLinks[i].Values())

		//fmt.Println("s_theta", s_theta.Values())

		//fmt.Println("a_theta", a_theta.Values())

		//fmt.Println("theta of supplyLinks dot product", VecDotProduct(s_theta, a_theta))

		tmp1 = tmp2 + VecDotProduct(s_theta, a_theta) - 1

		//fmt.Println("tmp1:", tmp1, "tmp2:", tmp2)

		if tmp2 == 0 {
			beta.Set(i, 0)
		} else {
			beta.Set(i, tmp1/tmp2)
		}

		if beta.Get(i) < 0 {
			beta.Set(i, 0)
		}

	}

	opt.beta = beta

	return opt.beta

}

func (opt *Shale) Shale(iter int) {

	opt.GetTheta()

	opt.DumpTheta()

	//opt.DumpAlpha()

	//opt.DumpBeta()

	for i := 0; i < iter; i++ {

		opt.GetSupplyDual()

		opt.GetContractDual()

		opt.DumpBeta()

		opt.DumpAlpha()

	}

}
