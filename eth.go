package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type RpcTx struct {
	BlockNumber *hexutil.Big    `json:"blockNumber" gencodec:"required"`
	BlockHash   *common.Hash    `json:"blockHash" rlp:"-"`
	TxHash      *common.Hash    `json:"hash"     rlp:"-"`
	TxIndex     *hexutil.Uint   `json:"transactionIndex"`
	CallFrom    *common.Address `json:"from"`
	CallTo      *common.Address `json:"to"       rlp:"nil"`
	GasPrice    *hexutil.Big    `json:"gasPrice" gencodec:"required"`
	EthAmount   *hexutil.Big    `json:"value"    gencodec:"required"`
	//Payload   *hexutil.Bytes  `json:"input"    gencodec:"required"`
}

type RpcHeader struct {
	Number    *hexutil.Big  `json:"number"`
	BlockHash *common.Hash  `json:"hash"`
	Timestamp *hexutil.Uint `json:"timestamp"`
}

type RpcBlock struct {
	Transactions []*RpcTx        `json:"transactions"`
	BlockNumber  *hexutil.Big    `json:"number"`
	BlockHash    *common.Hash    `json:"hash"`
	Timestamp    *hexutil.Big    `json:"timestamp"`
	Miner        *common.Address `json:"miner"`
}

func main() {
	client, err := rpc.Dial("http://localhost:8545")
	//client, err := rpc.Dial("http://localhost:8548")
	if err != nil {
		fmt.Println("rpc.Dial err", err)
		return
	}

	// 11111
	var tmp interface{}
	err = client.Call(&tmp, "eth_getBlockByNumber", "0x1b4", true)
	if err != nil {
		fmt.Println("client.Call err", err)
		return
	}
	fmt.Println(tmp)

	// 22222
	var raw json.RawMessage
	//err = client.CallContext(context.Background(), &raw, "eth_getBlockByNumber", "0x1b4", true)
	t := strconv.FormatInt(436, 16)
	fmt.Println("strconv.FormatInt(436, 16): ", t)
	err = client.CallContext(context.Background(), &raw, "eth_getBlockByNumber", "0x" + t, true)
	if err != nil {
		fmt.Println("11111")
		return
	} else if len(raw) == 0 {
		fmt.Println("22222")
		return
	}
	var block *RpcBlock
	if err := json.Unmarshal(raw, &block); err != nil {
		fmt.Println("33333")
		return
	}
	fmt.Println("block:", block)


	// 33333
	c := ethclient.NewClient(client)
	block_2, err := c.BlockByNumber(context.Background(), big.NewInt(123))

	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("block_2:", block_2)
	}

	var bn *hexutil.Big
	err = client.CallContext(context.Background(), &bn, "eth_blockNumber")
	if err != nil {
		fmt.Println("eth_blockNumber Call err", err)
		return
	}

	fmt.Println("bn:", bn)
	if bn.ToInt().IsUint64() {
		fmt.Println("IsUint64:", bn.ToInt().Uint64())
	}
	if bn.ToInt().IsInt64() {
		fmt.Println("IsInt64:", bn.ToInt().Int64())
	}

	var txRec *types.Receipt
	err = client.CallContext(context.Background(), &txRec, "eth_getTransactionReceipt",
		"0x2be49d7312cfc0eae3a3f5698b421d2a491c77708dc52b25621c63300667edd1")
	if err != nil {
		fmt.Println("4444")
		return
	}
	fmt.Println("txr:", txRec)

	var rpcTx *RpcTx
	err = client.CallContext(context.Background(), &rpcTx, "eth_getTransactionByHash",
		"0x2be49d7312cfc0eae3a3f5698b421d2a491c77708dc52b25621c63300667edd1")
	if err != nil {
		fmt.Println("5555")
		return
	}
	fmt.Println("rpcTx:", rpcTx)
	fmt.Println("rpcTx.BlockNumber:", rpcTx.BlockNumber)
	fmt.Println("rpcTx.BlockNumber.String():", rpcTx.BlockNumber.String())

	var tmp2 interface{}
	err = client.Call(&tmp2, "eth_getTransactionByHash",
		"0x2be49d7312cfc0eae3a3f5698b421d2a491c77708dc52b25621c63300667edd1")
	if err != nil {
		fmt.Println("client.Call err", err)
		return
	}
	fmt.Println(tmp2)
}