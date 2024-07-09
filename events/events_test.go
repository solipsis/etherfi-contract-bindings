package events

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestInfuraSuggestedRange(t *testing.T) {

	err := errors.New("query returned more than 10000 results. Try with this block range [0x106F23B, 0x1247590].")
	fmt.Println(InfuraSuggestedRange(err))
}

func TestGetLog(t *testing.T) {
	sig := "TransferSingle(address,address,address,uint256,uint256)"
	contractAddr := common.HexToAddress("0xb49e4420eA6e35F98060Cd133842DbeA9c27e479")
	abiStr := MembershipNFTABI

	if err := GetLogTest(sig, contractAddr, abiStr, nil); err != nil {
		t.Fatal(err)
	}

}

func TestGetLogGeneric(t *testing.T) {

	rpcClient, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatal(err)
	}

	abiStr := MembershipNFTABI
	contractAddr := common.HexToAddress("0xb49e4420eA6e35F98060Cd133842DbeA9c27e479")
	contract := MustBind(contractAddr, abiStr, rpcClient)

	sig := "TransferSingle(address,address,address,uint256,uint256)"

	var results []MembershipNFTTransferSingle
	if err := GetLogGeneric(contract, sig, &results); err != nil {
		t.Fatal(err)
	}
	for _, res := range results {
		fmt.Printf("from: %s, to: %s, value: %d\n", res.From, res.To, res.Value.Uint64())
	}

}
