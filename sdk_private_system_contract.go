/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package chainmaker_sdk_go

import (
	"chainmaker.org/chainmaker-sdk-go/pb/protogo/common"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func (cc *ChainClient) SaveCert(userCert, enclaveCert, txId string, withSyncResult bool, timeout int64) ([]byte, error) {
	if txId == "" {
		txId = GetRandTxId()
	}

	cc.logger.Infof("[SDK] begin to save cert , [contract:%s]/[method:%s]/[txId:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_SAVE_CERT.String(),
		txId,
	)

	// 构造Payload
	pairs := paramsMap2KVPairs(map[string]string{
		"user_cert":    userCert,
		"enclave_cert": enclaveCert,
	})

	payloadBytes, err := constructSystemContractPayload(
		cc.chainId,
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_SAVE_CERT.String(),
		pairs,
		defaultSequence,
	)
	if err != nil {
		return nil, fmt.Errorf("construct save cert payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequestWithTimeout(common.TxType_INVOKE_SYSTEM_CONTRACT, GetRandTxId(), payloadBytes, timeout)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_INVOKE_SYSTEM_CONTRACT.String(), err.Error())
	}

	if resp.Code == common.TxStatusCode_SUCCESS {
		if !withSyncResult {
			resp.ContractResult = &common.ContractResult{
				Code:    common.ContractResultCode_OK,
				Message: common.ContractResultCode_OK.String(),
				Result:  []byte(txId),
			}
		} else {
			contractResult, err := cc.getSyncResult(txId)
			if err != nil {
				return nil, fmt.Errorf("get sync result failed, %s", err.Error())
			}

			if contractResult.Code != common.ContractResultCode_OK {
				resp.Code = common.TxStatusCode_CONTRACT_FAIL
				resp.Message = contractResult.Message
			}

			resp.ContractResult = contractResult
		}
	}

	if err = checkProposalRequestResp(resp, true); err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_INVOKE_SYSTEM_CONTRACT.String(), err.Error())
	}

	return resp.ContractResult.Result, nil
}

func (cc *ChainClient) SaveDir(userCert, orderId, dirHash, dirSign, txId string,
	privateDir *common.StrSlice, withSyncResult bool, timeout int64) ([]byte, error) {

	if txId == "" {
		txId = GetRandTxId()
	}

	cc.logger.Infof("[SDK] begin to save dir , [contract:%s]/[method:%s]/[txId:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_SAVE_DIR.String(),
		txId,
	)

	// 构造Payload
	pairs := paramsMap2KVPairs(map[string]string{
		"user_cert":   userCert,
		"order_id":    orderId,
		"dir_hash":    dirHash,
		"dir_sign":    dirSign,
		"private_dir": privateDir.String(),
	})

	payloadBytes, err := constructSystemContractPayload(
		cc.chainId,
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_SAVE_DIR.String(),
		pairs,
		defaultSequence,
	)
	if err != nil {
		return nil, fmt.Errorf("construct save dir payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequestWithTimeout(common.TxType_INVOKE_SYSTEM_CONTRACT, GetRandTxId(), payloadBytes, timeout)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_INVOKE_SYSTEM_CONTRACT.String(), err.Error())
	}

	if resp.Code == common.TxStatusCode_SUCCESS {
		if !withSyncResult {
			resp.ContractResult = &common.ContractResult{
				Code:    common.ContractResultCode_OK,
				Message: common.ContractResultCode_OK.String(),
				Result:  []byte(txId),
			}
		} else {
			contractResult, err := cc.getSyncResult(txId)
			if err != nil {
				return nil, fmt.Errorf("get sync result failed, %s", err.Error())
			}

			if contractResult.Code != common.ContractResultCode_OK {
				resp.Code = common.TxStatusCode_CONTRACT_FAIL
				resp.Message = contractResult.Message
			}

			resp.ContractResult = contractResult
		}
	}

	if resp.Code != common.TxStatusCode_SUCCESS || resp.Message != "OK" {
		return nil, fmt.Errorf(errStringFormat, common.TxType_INVOKE_SYSTEM_CONTRACT.String(), err.Error())
	}

	return []byte(resp.Message), nil
}

func (cc *ChainClient) GetContract(userCert, contractName, codeHash,
	hashSign string) (*common.PrivateGetContract, error) {

	cc.logger.Infof("[SDK] begin to get contract , [contract:%s]/[method:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_CONTRACT.String(),
	)

	// 构造Payload
	pairs := paramsMap2KVPairs(map[string]string{
		"user_cert":     userCert,
		"contract_name": contractName,
		"code_hash":     codeHash,
		"hash_sign":     hashSign,
	})

	payloadBytes, err := constructQueryPayload(
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_CONTRACT.String(),
		pairs,
	)
	if err != nil {
		return nil, fmt.Errorf("marshal get contract payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequest(common.TxType_QUERY_SYSTEM_CONTRACT, GetRandTxId(), payloadBytes)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	if err = checkProposalRequestResp(resp, true); err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	contractInfo := &common.PrivateGetContract{}
	if err = proto.Unmarshal(resp.ContractResult.Result, contractInfo); err != nil {
		return nil, fmt.Errorf("GetContract unmarshal contract info payload failed, %s", err.Error())
	}

	return contractInfo, nil
}

func (cc *ChainClient) SaveData(computeResult, contractName, gas, reportSign, userCert, txId string, rwSet *common.TxRWSet,
	events *common.StrSlice, withSyncResult bool, timeout int64) ([]byte, error) {
	if txId == "" {
		txId = GetRandTxId()
	}

	cc.logger.Infof("[SDK] begin to save data , [contract:%s]/[method:%s]/[txId:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_SAVE_DATA.String(),
		txId,
	)

	// 构造Payload
	var rwSetStr string
	if rwSet != nil {
		rwb, err := rwSet.Marshal()
		if err != nil {
			return nil, fmt.Errorf("construct save data payload failed, %s", err.Error())
		}
		rwSetStr = string(rwb)
	}

	var eventsStr string
	if events != nil {
		eb, err := events.Marshal()
		if err != nil {
			return nil, fmt.Errorf("construct save data payload failed, %s", err.Error())
		}
		eventsStr = string(eb)
	}

	pairs := paramsMap2KVPairs(map[string]string{
		"compute_result": computeResult,
		"contract_name":  contractName,
		"rw_set":         rwSetStr,
		"events":         eventsStr,
		"gas":            gas,
		"report_sign":    reportSign,
		"user_cert":      userCert,
	})

	payloadBytes, err := constructSystemContractPayload(
		cc.chainId,
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_SAVE_DATA.String(),
		pairs,
		defaultSequence,
	)
	if err != nil {
		return nil, fmt.Errorf("construct save data payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequestWithTimeout(common.TxType_INVOKE_SYSTEM_CONTRACT, GetRandTxId(), payloadBytes, timeout)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_INVOKE_SYSTEM_CONTRACT.String(), err.Error())
	}

	if resp.Code == common.TxStatusCode_SUCCESS {
		if !withSyncResult {
			resp.ContractResult = &common.ContractResult{
				Code:    common.ContractResultCode_OK,
				Message: common.ContractResultCode_OK.String(),
				Result:  []byte(txId),
			}
		} else {
			contractResult, err := cc.getSyncResult(txId)
			if err != nil {
				return nil, fmt.Errorf("get sync result failed, %s", err.Error())
			}

			if contractResult.Code != common.ContractResultCode_OK {
				resp.Code = common.TxStatusCode_CONTRACT_FAIL
				resp.Message = contractResult.Message
			}

			resp.ContractResult = contractResult
		}
	}

	if resp.Code != common.TxStatusCode_SUCCESS || resp.Message != "OK" {
		return nil, fmt.Errorf(errStringFormat, common.TxType_INVOKE_SYSTEM_CONTRACT.String(), err.Error())
	}

	return []byte(resp.Message), nil
}

func (cc *ChainClient) GetData(contractName, privateKey, userCert, dirSign string) ([]byte, error) {
	cc.logger.Infof("[SDK] begin to get data , [contract:%s]/[method:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_DATA.String(),
	)

	// 构造Payload
	pairs := paramsMap2KVPairs(map[string]string{
		"contract_name": contractName,
		"private_key":   privateKey,
		"user_cert":     userCert,
		"dir_sign":      dirSign,
	})

	payloadBytes, err := constructQueryPayload(
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_GET_DATA.String(),
		pairs,
	)
	if err != nil {
		return nil, fmt.Errorf("marshal get data payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequest(common.TxType_QUERY_SYSTEM_CONTRACT, GetRandTxId(), payloadBytes)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	if err = checkProposalRequestResp(resp, true); err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_QUERY_SYSTEM_CONTRACT.String(), err.Error())
	}

	return resp.ContractResult.Result, nil
}

func (cc *ChainClient) SaveContract(contractCode []byte, codeHash, name, txId, version string, withSyncResult bool,
	timeout int64) (*common.ContractResult, error) {
	if txId == "" {
		txId = GetRandTxId()
	}

	cc.logger.Infof("[SDK] begin to save contract code , [contract:%s]/[method:%s]/[txId:%s]",
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_SAVE_CONTRACT.String(),
		txId,
	)

	// 构造Payload
	pairs := paramsMap2KVPairs(map[string]string{
		"contract_code": string(contractCode),
		"code_hash":     codeHash,
		"name ":         name,
		"version":       version,
	})

	payloadBytes, err := constructSystemContractPayload(
		cc.chainId,
		common.ContractName_SYSTEM_CONTRACT_PRIVATE_COMPUTE.String(),
		common.PrivateComputeContractFunction_SAVE_CONTRACT.String(),
		pairs,
		defaultSequence,
	)
	if err != nil {
		return nil, fmt.Errorf("construct save contract code payload failed, %s", err.Error())
	}

	resp, err := cc.proposalRequestWithTimeout(common.TxType_INVOKE_SYSTEM_CONTRACT, GetRandTxId(), payloadBytes, timeout)
	if err != nil {
		return nil, fmt.Errorf(errStringFormat, common.TxType_INVOKE_SYSTEM_CONTRACT.String(), err.Error())
	}

	if resp.Code == common.TxStatusCode_SUCCESS {
		if !withSyncResult {
			resp.ContractResult = &common.ContractResult{
				Code:    common.ContractResultCode_OK,
				Message: common.ContractResultCode_OK.String(),
				Result:  []byte(txId),
			}
		} else {
			contractResult, err := cc.getSyncResult(txId)
			if err != nil {
				return nil, fmt.Errorf("get sync result failed, %s", err.Error())
			}

			if contractResult.Code != common.ContractResultCode_OK {
				resp.Code = common.TxStatusCode_CONTRACT_FAIL
				resp.Message = contractResult.Message
			}

			resp.ContractResult = contractResult
		}
	}

	return resp.ContractResult, nil
}
