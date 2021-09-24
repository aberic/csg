/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package chainmaker_sdk_go

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"chainmaker.org/chainmaker/pb-go/v2/common"
	"chainmaker.org/chainmaker/pb-go/v2/syscontract"
	"chainmaker.org/chainmaker/sdk-go/v2/utils"
)

func (cc *ChainClient) CreateContractCreatePayload(contractName, version, byteCode string, runtime common.RuntimeType,
	kvs []*common.KeyValuePair) (*common.Payload, error) {

	cc.logger.Debugf("[SDK] create [ContractCreate] to be signed payload")
	return cc.createContractManageWithByteCodePayload(contractName,
		syscontract.ContractManageFunction_INIT_CONTRACT.String(), version, byteCode, runtime, kvs)
}

func (cc *ChainClient) CreateContractUpgradePayload(contractName, version, byteCode string, runtime common.RuntimeType,
	kvs []*common.KeyValuePair) (*common.Payload, error) {

	cc.logger.Debugf("[SDK] create [ContractUpgrade] to be signed payload")
	return cc.createContractManageWithByteCodePayload(contractName,
		syscontract.ContractManageFunction_UPGRADE_CONTRACT.String(), version, byteCode, runtime, kvs)
}

func (cc *ChainClient) CreateContractFreezePayload(contractName string) (*common.Payload, error) {
	cc.logger.Debugf("[SDK] create [ContractFreeze] to be signed payload")
	return cc.createContractManagePayload(contractName, syscontract.ContractManageFunction_FREEZE_CONTRACT.String())
}

func (cc *ChainClient) CreateContractUnfreezePayload(contractName string) (*common.Payload, error) {
	cc.logger.Debugf("[SDK] create [ContractUnfreeze] to be signed payload")
	return cc.createContractManagePayload(contractName, syscontract.ContractManageFunction_UNFREEZE_CONTRACT.String())
}

func (cc *ChainClient) CreateContractRevokePayload(contractName string) (*common.Payload, error) {
	cc.logger.Debugf("[SDK] create [ContractRevoke] to be signed payload")
	return cc.createContractManagePayload(contractName, syscontract.ContractManageFunction_REVOKE_CONTRACT.String())
}

func (cc *ChainClient) createContractManagePayload(contractName, method string) (*common.Payload, error) {
	kvs := []*common.KeyValuePair{
		{
			Key:   syscontract.GetContractInfo_CONTRACT_NAME.String(),
			Value: []byte(contractName),
		},
	}
	return cc.createPayload("", common.TxType_INVOKE_CONTRACT, syscontract.SystemContract_CONTRACT_MANAGE.String(),
		method, kvs, defaultSeq), nil
}

func (cc *ChainClient) createNativeContractAccessPayload(method string,
	toAddContractList ...string) (*common.Payload, error) {
	val, _ := json.Marshal(toAddContractList)
	kvs := []*common.KeyValuePair{
		{
			Key:   syscontract.ContractAccess_NATIVE_CONTRACT_NAME.String(),
			Value: val,
		},
	}
	return cc.createPayload("", common.TxType_INVOKE_CONTRACT, syscontract.SystemContract_CONTRACT_MANAGE.String(),
		method, kvs, defaultSeq), nil
}

func (cc *ChainClient) createContractManageWithByteCodePayload(contractName, method, version, byteCode string,
	runtime common.RuntimeType, kvs []*common.KeyValuePair) (*common.Payload, error) {
	var (
		err       error
		codeBytes []byte
	)

	exists := utils.Exists(byteCode)
	if exists {
		codeBytes, err = ioutil.ReadFile(byteCode)
		if err != nil {
			return nil, fmt.Errorf("read from byteCode file %s failed, %s", byteCode, err)
		}
	} else {
		byteCode = strings.TrimSpace(byteCode)

		if codeBytes, err = hex.DecodeString(byteCode); err != nil {
			if codeBytes, err = base64.StdEncoding.DecodeString(byteCode); err != nil {
				return nil, fmt.Errorf("decode byteCode failed, %s", err)
			}
		}
	}

	if !cc.checkKeyValuePair(kvs) {
		return nil, fmt.Errorf("use reserved word")
	}

	payload := cc.createPayload("", common.TxType_INVOKE_CONTRACT,
		syscontract.SystemContract_CONTRACT_MANAGE.String(), method, kvs, defaultSeq)

	payload.Parameters = append(payload.Parameters, &common.KeyValuePair{
		Key:   syscontract.InitContract_CONTRACT_NAME.String(),
		Value: []byte(contractName),
	})

	payload.Parameters = append(payload.Parameters, &common.KeyValuePair{
		Key:   syscontract.InitContract_CONTRACT_VERSION.String(),
		Value: []byte(version),
	})

	payload.Parameters = append(payload.Parameters, &common.KeyValuePair{
		Key:   syscontract.InitContract_CONTRACT_RUNTIME_TYPE.String(),
		Value: []byte(runtime.String()),
	})

	payload.Parameters = append(payload.Parameters, &common.KeyValuePair{
		Key:   syscontract.InitContract_CONTRACT_BYTECODE.String(),
		Value: codeBytes,
	})

	return payload, nil
}

func (cc *ChainClient) checkKeyValuePair(kvs []*common.KeyValuePair) bool {
	for _, kv := range kvs {
		if kv.Key == syscontract.InitContract_CONTRACT_NAME.String() ||
			kv.Key == syscontract.InitContract_CONTRACT_RUNTIME_TYPE.String() ||
			kv.Key == syscontract.InitContract_CONTRACT_VERSION.String() ||
			kv.Key == syscontract.InitContract_CONTRACT_BYTECODE.String() ||
			kv.Key == syscontract.UpgradeContract_CONTRACT_NAME.String() ||
			kv.Key == syscontract.UpgradeContract_CONTRACT_RUNTIME_TYPE.String() ||
			kv.Key == syscontract.UpgradeContract_CONTRACT_VERSION.String() ||
			kv.Key == syscontract.UpgradeContract_CONTRACT_BYTECODE.String() {
			return false
		}
	}

	return true
}

func (cc *ChainClient) SignContractManagePayload(payload *common.Payload) (*common.EndorsementEntry, error) {
	return cc.SignPayload(payload)
}

func (cc *ChainClient) SendContractManageRequest(payload *common.Payload, endorsers []*common.EndorsementEntry,
	timeout int64, withSyncResult bool) (*common.TxResponse, error) {
	return cc.sendContractRequest(payload, endorsers, timeout, withSyncResult)
}

func (cc *ChainClient) InvokeContract(contractName, method, txId string, kvs []*common.KeyValuePair, timeout int64,
	withSyncResult bool) (*common.TxResponse, error) {

	cc.logger.Debugf("[SDK] begin to INVOKE contract, [contractName:%s]/[method:%s]/[txId:%s]/[params:%+v]",
		contractName, method, txId, kvs)

	payload := cc.createPayload(txId, common.TxType_INVOKE_CONTRACT, contractName, method, kvs, defaultSeq)

	return cc.sendContractRequest(payload, nil, timeout, withSyncResult)
}

func (cc *ChainClient) QueryContract(contractName, method string, kvs []*common.KeyValuePair,
	timeout int64) (*common.TxResponse, error) {

	cc.logger.Debugf("[SDK] begin to QUERY contract, [contractName:%s]/[method:%s]/[params:%+v]",
		contractName, method, kvs)

	payload := cc.createPayload("", common.TxType_QUERY_CONTRACT, contractName, method, kvs, defaultSeq)

	resp, err := cc.proposalRequestWithTimeout(payload, nil, timeout)
	if err != nil {
		return nil, fmt.Errorf("send %s failed, %s", payload.TxType.String(), err.Error())
	}

	return resp, nil
}

func (cc *ChainClient) GetTxRequest(contractName, method, txId string,
	kvs []*common.KeyValuePair) (*common.TxRequest, error) {
	if txId == "" {
		txId = utils.GetRandTxId()
	}

	cc.logger.Debugf("[SDK] begin to create TxRequest, [contractName:%s]/[method:%s]/[txId:%s]/[params:%+v]",
		contractName, method, txId, kvs)

	payload := cc.createPayload(txId, common.TxType_INVOKE_CONTRACT, contractName, method, kvs, defaultSeq)

	req, err := cc.generateTxRequest(payload, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (cc *ChainClient) SendTxRequest(txRequest *common.TxRequest, timeout int64,
	withSyncResult bool) (*common.TxResponse, error) {

	resp, err := cc.sendTxRequest(txRequest, timeout)
	if err != nil {
		return resp, fmt.Errorf("%s failed, %s", txRequest.Payload.TxType.String(), err.Error())
	}

	if resp.Code == common.TxStatusCode_SUCCESS {
		if !withSyncResult {
			resp.TxId = txRequest.Payload.TxId
		} else {
			contractResult, err := cc.getSyncResult(txRequest.Payload.TxId)
			if err != nil {
				return nil, fmt.Errorf("get sync result failed, %s", err.Error())
			}

			if contractResult.Code != utils.SUCCESS {
				resp.Code = common.TxStatusCode_CONTRACT_FAIL
				resp.Message = contractResult.Message
			}

			resp.ContractResult = contractResult
		}
	}

	return resp, nil
}
