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

	label := "deposit"
	resp, _ := accountSvc.NewAccountByChain(uc.Chain, label)

	ah.JSON(c, resp)
}

func (ah *accountHandler) getAccountDetail(c *server.Context) {

	chainCode, _ := ah.Query(c, "chain")
	address := ah.Param(c, "address")

	resp, _ := accountSvc.GetAccountDetail(chainCode, address)

	ah.JSON(c, resp)
}

func (ah *accountHandler) getAccounts(c *server.Context) {

	chainCode, _ := ah.Query(c, "chain")

	resp, _ := accountSvc.GetAccountsByChain(chainCode)

	ah.JSON(c, resp)
}

func (ah *accountHandler) getAccountBalance(c *server.Context) {

	chainCode, _ := ah.Query(c, "chain")
	address := ah.Param(c, "address")
	identity, _ := ah.Query(c, "identity")

	if chainCode == "" || address == "" {
		ah.JSON(c, "bad request query")
	}

	resp, _ := accountSvc.GetAccountBalance(chainCode, address, identity)

	ah.JSON(c, resp)
}
