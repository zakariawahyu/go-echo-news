package dependency

type FooBarServices struct {
	fooServ *FooServices
	barServ *BarServices
}

func NewFooBarServices(fooServ *FooServices, barServ *BarServices) *FooBarServices {
	return &FooBarServices{
		fooServ: fooServ,
		barServ: barServ,
	}
}
