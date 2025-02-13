package coinpayments_test

import (
	"testing"

	"github.com/tryvium-travels/go-coinpayments"
)

func TestCallGetTxInfo(t *testing.T) {
	client, err := testClient()
	if err != nil {
		t.Fatal("should have instantiated a new client with valid configuration")
	}

	resp, err := client.CallGetTxInfo(&coinpayments.TxInfoRequest{TxID: "CPHB0ZYSLG1TRNVDGKPMINAYFW", Full: "0"})
	if err != nil {
		t.Fatal("error getting tx info ", err)
	}

	if coin, exists := resp.Result["coin"].(string); !exists || coin != "SYS" {
		t.Fatalf("error in tx: Expected coin to be SYS, found %s", coin)
	}
}

func TestCallCreateTransaction(t *testing.T) {
	client, err := testClient()
	if err != nil {
		t.Fatalf("Should have instantiated a new client with valid config and http client, but it threw error: %s", err.Error())
	}

	resp, err := client.CallCreateTransaction(&coinpayments.TransactionRequest{Amount: "100", Currency1: "USD", Currency2: "BTC", BuyerEmail: "noreply@tryvium.io"})
	if err != nil {
		t.Fatalf("Could not call create transaction: %s", err.Error())
	}

	_, err = client.CallGetTxInfo(&coinpayments.TxInfoRequest{TxID: resp.TxnID})
	if err != nil {
		t.Fatalf("Could not call get tx info: %s", err.Error())
	}

}

func TestCallGetCallbackAddress(t *testing.T) {
	client, err := testClient()
	if err != nil {
		t.Fatalf("Should have instantiated a new client with valid config and http client, but it threw error: %s", err.Error())
	}

	_, err = client.CallGetCallbackAddress(&coinpayments.CallbackAddressRequest{Currency: "BTC"})

	if err != nil {
		t.Fatalf("Could not call get callback address: %s", err.Error())
	}

}

func TestCallGetDepositAddress(t *testing.T) {
	client, err := testClient()
	if err != nil {
		t.Fatalf("Should have instantiated a new client with valid config and http client, but it threw error: %s", err.Error())
	}

	_, err = client.CallGetDepositAddress(&coinpayments.DepositAddressRequest{Currency: "SYS"})

	if err != nil {
		t.Fatalf("Could not call get callback address: %s", err.Error())
	}

}

func TestCallGetTxList(t *testing.T) {
	client, err := testClient()
	if err != nil {
		t.Fatalf("Should have instantiated a new client with valid config and http client, but it threw error: %s", err.Error())
	}

	_, err = client.CallGetTxList(&coinpayments.TxListRequest{})

	if err != nil {
		t.Fatalf("Could not call get tx list: %s", err.Error())
	}

}
