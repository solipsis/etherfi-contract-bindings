package events

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func IsExceededLimitError(err error) bool {
	str := err.Error()
	// quick heuristic for Infura and alchemy respectively
	if strings.Contains(str, "query returned more than") || strings.Contains(str, "Log response size exceeded") {
		return true
	}
	return false
}

// parse infura error returns 80% of the difference between the suggested range
// that the infura API gives when querying an event range with over 10k events.
// It is the responsibility of the caller to only call this with the intended error
// e.g.
//
//	"query returned more than 10000 results. Try with this block range [0x106F23B, 0x1247590]."
func InfuraSuggestedRange(err error) int64 {

	msg := err.Error()
	fmt.Println("msg:", msg)

	// find the range section
	lIdx := strings.Index(msg, "[")
	rIdx := strings.Index(msg, "]")
	hexArr := strings.Split(msg[lIdx+1:rIdx], ", ")

	lHex, rHex := hexArr[0], hexArr[1]
	lVal, err := strconv.ParseInt(lHex[2:], 16, 64)
	if err != nil {
		panic("invalid range value")
	}
	rVal, err := strconv.ParseInt(rHex[2:], 16, 64)
	if err != nil {
		panic("invalid range value")
	}

	return max(rVal-lVal, 1)
	/*
		// apply 80% buffer because it is likely that traffic increases over time
		fDiff := float64(rVal-lVal) * 0.8

		return max(int64(fDiff), 1)
	*/
}

// if e.Type.Kind() != reflect.Pointer {

type Contract struct {
	rpcClient *ethclient.Client
	abiStr    string
	abi       abi.ABI
	addr      common.Address
}

func MustBind(contract common.Address, abiStr string, rpcClient *ethclient.Client) *Contract {
	contractAbi, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		panic(err)
	}

	return &Contract{
		rpcClient: rpcClient,
		abiStr:    abiStr,
		abi:       contractAbi,
		addr:      contract,
	}
}

func GetLogGeneric[T any](contract *Contract, sig string, results *[]T) error {

	sigHash := crypto.Keccak256Hash([]byte(sig))
	filterQuery := [][]any{{sigHash}}

	topics, err := abi.MakeTopics(filterQuery...)
	if err != nil {
		panic(err)
	}

	query := ethereum.FilterQuery{
		//FromBlock: big.NewInt(19173587),
		//ToBlock:   big.NewInt(20173587),
		Addresses: []common.Address{
			contract.addr,
		},
		//Topics: [][]common.Hash{{sigHash}},
		Topics: topics,
	}
	if query.FromBlock == nil {
		query.FromBlock = big.NewInt(0)
	}
	if query.ToBlock == nil {
		endBlock, err := contract.rpcClient.BlockNumber(context.Background())
		if err != nil {
			panic(err)
		}
		query.ToBlock = big.NewInt(int64(endBlock))
	}

	targetBlock := query.ToBlock.Uint64()

	// This will be set to an appropriate value determined by RPC if limit exceeded
	blocksPerChunk := int64(999999999)

	// query events until we have processed all events in range.
	// If a given query returns more than 10k events, it must be split into subqueries
	var allLogs []types.Log
	for query.FromBlock.Uint64() < targetBlock {
		fmt.Println(query.FromBlock, query.ToBlock)
		logs, err := contract.rpcClient.FilterLogs(context.Background(), query)
		if err != nil {
			if !IsExceededLimitError(err) {
				panic(err)
			}

			// set new query range
			blocksPerChunk = InfuraSuggestedRange(err)
			query.ToBlock = new(big.Int).Add(query.FromBlock, big.NewInt(int64(blocksPerChunk)))
			fmt.Println("newFrom:", query.FromBlock, "newTo:", query.ToBlock)

			continue
		}

		allLogs = append(allLogs, logs...)
		query.FromBlock.Add(query.FromBlock, big.NewInt(blocksPerChunk))
		query.ToBlock.Add(query.ToBlock, big.NewInt(blocksPerChunk))
	}

	/*
		logs, err := contract.rpcClient.FilterLogs(context.Background(), query)
		if err != nil {
			if !IsExceededLimitError(err) {
				panic(err)
			}

			// set new query range

			// keep going until new range

			// if limit is exceeded I need to further shrink
		}
	*/

	for _, vLog := range allLogs {

		var testEvent T

		// TODO: name from signature
		if err := UnpackLog(contract.abi, &testEvent, "TransferSingle", vLog); err != nil {
			panic(err)
		}
		*results = append(*results, testEvent)
		//		fmt.Println(testEvent)
	}

	return nil
}

func GetLogTest(sig string, contract common.Address, abiStr string, results []any) error {

	rpcClient, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatal(err)
	}

	sigHash := crypto.Keccak256Hash([]byte(sig))
	filterQuery := [][]any{{sigHash}}

	topics, err := abi.MakeTopics(filterQuery...)
	if err != nil {
		panic(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(19173587),
		ToBlock:   big.NewInt(20173587),
		Addresses: []common.Address{
			contract,
		},
		//Topics: [][]common.Hash{{sigHash}},
		Topics: topics,
	}

	logs, err := rpcClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {

		var testEvent MembershipNFTTransferSingle

		// TODO: name from signature
		if err := UnpackLog(contractAbi, &testEvent, "TransferSingle", vLog); err != nil {
			panic(err)
		}
		fmt.Println(testEvent)
	}

	/*
			// TODO: name from signature
			err := contractAbi.Unpack(&testEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			/*

		/*
			contractAbi, err := abi.JSON(strings.NewReader(string(etherfiNodesManager.EtherfiNodesManagerMetaData.ABI)))
			if err != nil {
				log.Fatal(err)
			}
	*/

	return nil
}

// UnpackLog unpacks a retrieved log into the provided output structure.
func UnpackLog(contractABI abi.ABI, out interface{}, event string, log types.Log) error {
	// Anonymous events are not supported.
	if len(log.Topics) == 0 {
		return fmt.Errorf("anonymous events not supported")
	}
	if log.Topics[0] != contractABI.Events[event].ID {
		return fmt.Errorf("event signature does not match expected signature")
	}
	if len(log.Data) > 0 {
		if err := contractABI.UnpackIntoInterface(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range contractABI.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopics(out, indexed, log.Topics[1:])
}

/*
func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol (ZRX) token address
	contractAddress := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(6383820),
		ToBlock:   big.NewInt(6383840),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(etherfiNodesManager.EtherfiNodesManagerMetaData.ABI)))
	if err != nil {
		log.Fatal(err)
	}

	adminChangedSig := []byte("AdminChanged(address,address)")
	adminChangedSigHash := crypto.Keccak256Hash(adminChangedSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case adminChangedSigHash.Hex():

			var adminEvent etherfiNodesManager.EtherfiNodesManagerAdminChanged

			err := contractAbi.Unpack(&adminEvent, "AdminChanged", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			/*
				adminEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
				adminEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

				fmt.Printf("From: %s\n", transferEvent.From.Hex())
				fmt.Printf("To: %s\n", transferEvent.To.Hex())
				fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())

		}
	}
}
*/
