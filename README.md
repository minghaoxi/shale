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

	=== RUN   TestShale
	begin
	[[0.9965149 0] [0.33681843 0.66517824] [0 1.0014881] [0 0]]
	--- PASS: TestShale (0.00s)

### description

	demand side	supply side

	[0]400                	[0] 300

				[1] 300 
	[1]600                 
				[2] 300
		    
				[3] 300
	edges:
	0	->	0
	0	->	1
	1	->	1
	1	->	2
	
		  
###expected result:
	
	demand index	supply index	expected ratio	shale reulst

	0                    0                  1                0.9965149
	1                    0                  0                0
	0                    1                  1/3              0.33681843
	1                    1                  2/3              0.66517824
	0                    2                  0                0
	1                    2                  1                1.0014881
	0                    3                  0                0
	1                    3                  0                0

