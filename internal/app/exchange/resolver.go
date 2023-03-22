package exchange

import (
	"fmt"
)

type MarketFactoryResolver interface {
	Resolve(exchange string) (MarketFactory, error)
}

type marketFactoryResolver struct {
	factories []MarketFactory
}

func (r *marketFactoryResolver) Resolve(exchange string) (MarketFactory, error) {
	for _, f := range r.factories {
		if f.Exchange() == exchange {
			return f, nil
		}
	}

	return nil, fmt.Errorf("exchange %s is not supported", exchange)
}

func NewMarketFactoryResolver(f []MarketFactory) MarketFactoryResolver {
	return &marketFactoryResolver{factories: f}
}
