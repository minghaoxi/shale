# shale
SHALE: An Efficient Algorithm for Allocation of Guaranteed Display Advertising Guaranteed Display Advertising

### mock example
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

### test result 30 iter

func TestShale(t *testing.T) {
	MockSampleShale()

	fmt.Println("begin")

	algo := MockSampleShale()

	algo.Shale(30)

	x := algo.Dual2Primal()

	fmt.Println(x)
}


### result

=== RUN   TestShale
begin
[0.6666667 0.8333333]

[0 0.33333334 0 0]

[0.16666672 0.16666663]

[0 0.5 0 0]
[0.25 0.25]
[0 0.5833333 0.049999952 0]
[0.29166672 0.31666663]
[0 0.6388889 0.11666665 0]
[0.31944442 0.37777784]
[0 0.68518513 0.17777781 0]
[0.34259263 0.43148145]
[0 0.7253087 0.23148137 0]
[0.36265442 0.47839502]
[0 0.76028806 0.278395 0]
[0.38014403 0.5193415]
[0 0.79080933 0.3193415 0]
[0.39540467 0.55507547]
[0 0.817444 0.35507542 0]
[0.40872207 0.5862598]
[0 0.84068745 0.3862597 0]
[0.4203438 0.61347365]
[0 0.86097145 0.4134736 0]
[0.43048584 0.6372225]
[0 0.8786729 0.4372224 0]
[0.43933654 0.65794766]
[0 0.8941205 0.45794764 0]
[0.44706023 0.6760339]
[0 0.9076012 0.4760339 0]
[0.45380065 0.6918175]
[0 0.9193656 0.4918176 0]
[0.45968276 0.70559156]
[0 0.9296322 0.5055916 0]
[0.46481612 0.7176118]
[0 0.9385915 0.5176117 0]
[0.4692958 0.72810155]
[0 0.9464102 0.5281015 0]
[0.47320512 0.7372559]
[0 0.9532334 0.53725576 0]
[0.47661683 0.7452446]
[0 0.9591878 0.54524463 0]
[0.47959396 0.7522163]
[0 0.9643841 0.55221635 0]
[0.48219207 0.7583002]
[0 0.9689188 0.5583002 0]
[0.48445937 0.76360947]
[0 0.9728761 0.56360954 0]
[0.48643798 0.7682428]
[0 0.9763295 0.56824267 0]
[0.4881648 0.7722861]
[0 0.97934324 0.5722861 0]
[0.48967162 0.7758147]
[0 0.98197335 0.57581466 0]
[0.49098665 0.77889407]
[0 0.9842685 0.57889396 0]
[0.49213424 0.78158116]
[0 0.9862714 0.58158106 0]
[0.4931357 0.7839261]
[0 0.9880193 0.5839261 0]
[0.4940097 0.78597265]
[0 0.9895447 0.5859727 0]
[0.49477234 0.78775865]
[[0.9965149 0] [0.33681843 0.66517824] [0 1.0014881] [0 0]]
--- PASS: TestShale (0.00s)

