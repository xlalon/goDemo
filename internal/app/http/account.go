package http

import (
	"fmt"

	"github.com/xlalon/golee/pkg/net/http/server"
)

type accountHandler struct {
	server.Handler
}

func (ah *accountHandler) newAccount(ctx *server.Context) {

	type UidChain struct {
		Chain string `json:"chain"`
	}

	var uc UidChain
	if err := ah.BindJSON(ctx, &uc); err != nil {
		fmt.Printf("wrong post data")
	}

	label := "DEPOSIT"
	resp, _ := walletSvc.NewAccount(ctx, uc.Chain, label)

	ah.JSON(ctx, resp)
}

func (ah *accountHandler) getAccountDetail(ctx *server.Context) {

	chainCode, _ := ah.Query(ctx, "chain")
	address := ah.Param(ctx, "address")

	resp, _ := walletSvc.GetAccount(ctx, chainCode, address)

	ah.JSON(ctx, resp)
}

func (ah *accountHandler) getAccounts(ctx *server.Context) {

	chainCode, _ := ah.Query(ctx, "chain")

	resp, _ := walletSvc.GetAccountsByChain(chainCode)

	ah.JSON(ctx, resp)
}

func (ah *accountHandler) getAccountBalance(ctx *server.Context) {

	chainCode, _ := ah.Query(ctx, "chain")
	address := ah.Param(ctx, "address")

	if chainCode == "" || address == "" {
		ah.JSON(ctx, "bad request query")
	}

	resp, _ := walletSvc.GetAccountBalance(ctx, chainCode, address)

	ah.JSON(ctx, resp)
}
