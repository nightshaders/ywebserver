package decorator


type Handler interface {
	Handle(p Params) error
}

type WebHandler func(Params) error

func (w WebHandler) Handle(p Params) error {
	return w(p)
}




type ParamsDecorator func(Params) Params

func NewParamDecoration() ParamsDecorator {
	return func(h Params) Params {
		return h
	}
}
func (d ParamsDecorator) Next(next ParamsDecorator) ParamsDecorator {
	return func(w Params) Params {
		return d(next(w))
	}
}

func (d ParamsDecorator) Handle() Decorator {
	return func(w WebHandler) WebHandler {
		return func(p Params) error {
			return w(d(p))
		}
	}
}