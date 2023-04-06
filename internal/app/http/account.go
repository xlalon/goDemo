package http

import (
	"fmt"

	"github.com/xlalon/golee/pkg/net/http/server"
)

type accountHandler struct {
	server.Handler
}

func (ah *accountHandler) newAccount(c *server.Context) {

	type UidChain struct {
		Chain string `json:"chain"`
	}

	var uc UidChain
	if err := ah.BindJSON(c, &uc); err != nil {
		fmt.Printf("wrong post data")
	}

	label := "DEPOSIT"
	resp, _ := walletSvc.NewAccount(uc.Chain, label)

	ah.JSON(c, resp)
}

func (ah *accountHandler) getAccountDetail(c *server.Context) {

	chainCode, _ := ah.Query(c, "chain")
	address := ah.Param(c, "address")

	resp, _ := walletSvc.GetAccount(chainCode, address)

	ah.JSON(c, resp)
}

func (ah *accountHandler) getAccounts(c *server.Context) {

	chainCode, _ := ah.Query(c, "chain")

	resp, _ := walletSvc.GetAccountsByChain(chainCode)

	ah.JSON(c, resp)
}

func (ah *accountHandler) getAccountBalance(c *server.Context) {

	chainCode, _ := ah.Query(c, "chain")
	address := ah.Param(c, "address")

	if chainCode == "" || address == "" {
		ah.JSON(c, "bad request query")
	}

	resp, _ := walletSvc.GetAccountBalance(chainCode, address)

	ah.JSON(c, resp)
}
