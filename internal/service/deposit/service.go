package deposit

import (
	"errors"

	"github.com/xlalon/golee/internal/onchain"
	"github.com/xlalon/golee/internal/service/chain"
	chainConf "github.com/xlalon/golee/internal/service/chain/conf"
	"github.com/xlalon/golee/internal/service/deposit/conf"
	"github.com/xlalon/golee/internal/service/deposit/domain"
	"github.com/xlalon/golee/internal/service/deposit/repoimpl/dao"
	"github.com/xlalon/golee/internal/service/wallet"
	walletConf "github.com/xlalon/golee/internal/service/wallet/conf"
)

type Service struct {
	depositRepo domain.DepositRepository
	onchainSvc  *onchain.Service
	chainSvc    *chain.Service
	walletSvc   *wallet.Service
}

func NewService(conf *conf.Config) *Service {
	return &Service{
		depositRepo: dao.New(conf),
		onchainSvc:  onchain.NewService(),
		chainSvc: chain.NewService(&chainConf.Config{
			Mysql: conf.Mysql,
			Redis: conf.Redis,
		}),
		walletSvc: wallet.NewService(&walletConf.Config{
			Mysql: conf.Mysql,
			Redis: conf.Redis,
		}),
	}
}

func (s *Service) NewDeposit(depDTO *domain.DepositDTO) error {
	if !depDTO.Amount.GreaterThanOrEqualZero() {
		return errors.New("amount should greater than 0")
	}
	return s.depositRepo.Save(domain.DepositFactory(depDTO))
}

func (s *Service) GetDeposit(depositId int64) (*domain.DepositDTO, error) {
	return s.depositDMToDTO(s.depositRepo.GetDepositById(depositId))
}

func (s *Service) GetDeposits() ([]*domain.DepositDTO, error) {
	return s.depositsDMToDTOs(s.depositRepo.GetDeposits())
}

func (s *Service) depositDMToDTO(dep *domain.Deposit, err error) (*domain.DepositDTO, error) {
	if err != nil {
		return nil, err
	}
	return dep.ToDepositDTO(), nil
}

func (s *Service) depositsDMToDTOs(deps []*domain.Deposit, err error) ([]*domain.DepositDTO, error) {
	if err != nil {
		return nil, err
	}
	var depsDTOs []*domain.DepositDTO
	for _, dep := range deps {
		depsDTOs = append(depsDTOs, dep.ToDepositDTO())
	}
	return depsDTOs, nil
}
