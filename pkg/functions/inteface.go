package functions

type FunctionInterface interface {
	Store(f *Function) (string, error)
	GetAndExecute(path string) error
	Execute(function *Function) error
}
