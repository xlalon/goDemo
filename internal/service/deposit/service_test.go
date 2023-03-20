package deposit

import (
	"fmt"
	"testing"

	chainSvc "github.com/xlalon/golee/internal/service/chain"
	"github.com/xlalon/golee/internal/service/wallet"
	"github.com/xlalon/golee/pkg/json"
)

var (
	testSvc = NewService(testRepository.DepositRepository(), chainSvc.NewService(testRepository.ChainRepository()), wallet.NewService(testRepository.WalletRepository()))
)

func TestService_GetDeposit(t *testing.T) {
	deps, err := testSvc.GetDeposits()
	if err != nil {
		fmt.Println("GetDeposit err:", err)
	}
	json.PPrint("deposit", deps)
}
