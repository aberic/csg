/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package chainmaker_sdk_go

import (
	"chainmaker.org/chainmaker-sdk-go/pb/protogo/common"
	"encoding/hex"
	"fmt"
	"github.com/hokaccha/go-prettyjson"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArchive(t *testing.T) {

	admin1, err := createAdmin(orgId1)
	require.Nil(t, err)

	fmt.Println("====================== 数据归档 ======================")
	var targetBlockHeight int64 = 5
	testArchiveBlock(t, admin1, targetBlockHeight)
}

func TestGetFromArchiveStore(t *testing.T) {
	client, err := createClientWithConfig()
	require.Nil(t, err)

	fmt.Println("====================== 归档查询 ======================")
	var blockHeight int64 = 8
	fullBlockInfo, err := client.GetFromArchiveStore(blockHeight)
	require.Nil(t, err)
	prettyJsonShow(t, "GetFromArchiveStore fullBlockInfo", fullBlockInfo)

	fullBlockInfo, err = client.GetArchivedFullBlockByHeight(blockHeight)
	require.Nil(t, err)
	prettyJsonShow(t, "GetArchivedFullBlockByHeight fullBlockInfo", fullBlockInfo)

	blockInfo, err := client.GetArchivedBlockByHeight(blockHeight, true)
	require.Nil(t, err)
	prettyJsonShow(t, "GetArchivedBlockByHeight with rwset", blockInfo)

	blockInfo, err = client.GetArchivedBlockByHeight(blockHeight, false)
	require.Nil(t, err)
	prettyJsonShow(t, "GetArchivedBlockByHeight without rwset", blockInfo)

	blockInfo, err = client.GetArchivedBlockByHash(hex.EncodeToString(blockInfo.Block.Header.BlockHash), true)
	require.Nil(t, err)
	prettyJsonShow(t, "GetArchivedBlockByHash with rwset", blockInfo)

	blockInfo, err = client.GetArchivedBlockByHash(hex.EncodeToString(blockInfo.Block.Header.BlockHash), false)
	require.Nil(t, err)
	prettyJsonShow(t, "GetArchivedBlockByHash without rwset", blockInfo)

	txId := blockInfo.Block.Txs[0].Header.TxId
	txInfo, err := client.GetArchivedTxByTxId(txId)
	require.Nil(t, err)
	prettyJsonShow(t, "GetArchivedTxByTxId", txInfo)
}

func prettyJsonShow(t *testing.T, name string, v interface{}) {
	marshal, err := prettyjson.Marshal(v)
	require.Nil(t, err)
	fmt.Printf("\n\n\n====== %s ======\n%s\n==========================\n", name, marshal)
}

func testArchiveBlock(t *testing.T, admin1 *ChainClient, targetBlockHeight int64) {
	var (
		err                error
		payload            []byte
		signedPayloadBytes []byte
		resp               *common.TxResponse
		result             string
	)

	payload, err = admin1.CreateArchiveBlockPayload(targetBlockHeight)
	require.Nil(t, err)

	signedPayloadBytes, err = admin1.SignArchivePayload(payload)
	require.Nil(t, err)

	resp, err = admin1.SendArchiveBlockRequest(signedPayloadBytes, -1, true)
	require.Nil(t, err)

	result = string(resp.ContractResult.Result)

	fmt.Printf("resp: %+v, result:%+s\n", resp, result)
}



