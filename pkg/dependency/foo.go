package dependency

type FooRepository struct {
}

func NewFooRepository() *FooRepository {
	return &FooRepository{}
}

type FooServices struct {
	fooRepo *FooRepository
}

func NewFooServices(fooRepo *FooRepository) *FooServices {
	return &FooServices{
		fooRepo: fooRepo,
	}
}

type Foo struct {
}

func NewFoo() *Foo {
	return &Foo{}
}
