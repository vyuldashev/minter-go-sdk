package grpc_client_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/v2/api/grpc_client"
	"github.com/MinterTeam/minter-go-sdk/v2/transaction"
	"github.com/MinterTeam/minter-go-sdk/v2/wallet"
	"github.com/MinterTeam/node-grpc-gateway/api_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"math/big"
	"strings"
)

func Example() {
	client, _ := grpc_client.New("localhost:8842")
	_ = client.CheckVersion("1.2", true)
	coinID, _ := client.CoinID("SYMBOL")
	w, _ := wallet.Create("1 2 3 4 5 6 7 8 9 10 11 12", "")
	data := transaction.NewSendData().SetCoin(coinID).SetValue(transaction.BipToPip(big.NewInt(1))).MustSetTo(w.Address)
	transactionsBuilder := transaction.NewBuilder(transaction.TestNetChainID)
	tx, _ := transactionsBuilder.NewTransaction(data)
	sign, _ := tx.SetNonce(1).SetGasPrice(1).Sign(w.PrivateKey)
	encode, _ := sign.Encode()
	hash, _ := sign.Hash()
	subscribeClient, _ := client.Subscribe(fmt.Sprintf("tx.hash = '%s'", strings.ToUpper(hash[2:])))
	defer subscribeClient.CloseSend()

	res, err := client.SendTransaction(encode)
	if err != nil {
		_, _, _ = client.ErrorBody(err)
	}
	if res.Code != 0 {
		panic(res.Log)
	}

	_, err = subscribeClient.Recv()
	if err == io.EOF {
		return
	}
	if code := status.Code(err); code != codes.OK {
		if code == codes.DeadlineExceeded || code == codes.Canceled {
			return
		}
		panic(err)
	}

	response, _ := client.Transaction(hash)
	_, _ = client.Marshal(response)
	sendData := new(api_pb.SendData)
	_ = response.Data.UnmarshalTo(sendData)
}
