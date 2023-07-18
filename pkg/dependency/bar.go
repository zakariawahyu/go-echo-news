package dependency

type BarRepository struct {
}

func NewBarRepository() *BarRepository {
	return &BarRepository{}
}

type BarServices struct {
	barRepo *FooRepository
}

func NewBarServices(barRepo *FooRepository) *BarServices {
	return &BarServices{
		barRepo: barRepo,
	}
}

type Bar struct {
}

func NewBar() *Bar {
	return &Bar{}
}
