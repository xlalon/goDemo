package deposit

import (
	"github.com/xlalon/golee/core/model/deposit"
)

type Filterer interface {
	IsValid(*deposit.Deposit) bool
	Satisfied([]*deposit.Deposit) ([]*deposit.Deposit, error)
}

type FilterSpecification struct{}

func (f *FilterSpecification) IsValid(*deposit.Deposit) bool {
	return true
}

func (f *FilterSpecification) Satisfied(deps []*deposit.Deposit) ([]*deposit.Deposit, error) {
	return deps, nil
}

type AccountFilter struct {
	FilterSpecification
}

func NewAccountFilter() *AccountFilter {
	return &AccountFilter{}
}

func (af *AccountFilter) Satisfied(deps []*deposit.Deposit) ([]*deposit.Deposit, error) {

	return deps, nil
}

type AmountFilter struct {
	FilterSpecification
}

func NewAmountFilter() *AmountFilter {
	return &AmountFilter{}
}

func (af *AmountFilter) Satisfied(deps []*deposit.Deposit) ([]*deposit.Deposit, error) {
	return deps, nil
}
