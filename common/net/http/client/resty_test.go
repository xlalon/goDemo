package client

import (
	"context"
	"fmt"
	"testing"
)

var _testRestyClient = NewRestyClient("https://wax.greymass.com")

func TestRestyClient_Get(t *testing.T) {

	resp, err := _testRestyClient.Get(context.Background(), "/v1/chain/get_info", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}

func TestRestyClient_Post(t *testing.T) {
	body := map[string]interface{}{
		"id": "7b89fdd2b27ced1c36c2e3a1e785f9253769db8ee2b4c2fc076e7d7cd30d7f04",
	}
	resp, err := _testRestyClient.Post(context.Background(), "/v1/history/get_transaction", body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
