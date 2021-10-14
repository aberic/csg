/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package chainmaker_sdk_go

import (
	"chainmaker.org/chainmaker/pb-go/v2/common"
	"chainmaker.org/chainmaker/pb-go/v2/syscontract"
	"chainmaker.org/chainmaker/sdk-go/v2/utils"
)

func (cc *ChainClient) createPubkeyManagePayload(method string, kvs []*common.KeyValuePair) *common.Payload {
	cc.logger.Debugf("[SDK] create PubkeyManagePayload, method: %s", method)
	payload := cc.createPayload("", common.TxType_INVOKE_CONTRACT, syscontract.SystemContract_PUBKEY_MANAGE.String(),
		method, kvs, defaultSeq)
	return payload
}

func (cc *ChainClient) CreatePubkeyAddPayload(pubkey string, org_id string, role string) (*common.Payload, error) {

	pairs := []*common.KeyValuePair{
		{
			Key:   utils.KeyPubkey,
			Value: []byte(pubkey),
		},
		{
			Key:   utils.KeyPubkeyOrgId,
			Value: []byte(org_id),
		},
		{
			Key:   utils.KeyPubkeyRole,
			Value: []byte(role),
		},
	}

	return cc.createPubkeyManagePayload(syscontract.PubkeyManageFunction_PUBKEY_ADD.String(), pairs), nil
}

func (cc *ChainClient) CreatePubkeyDelPayload(pubkey string, org_id string) (*common.Payload, error) {
	pairs := []*common.KeyValuePair{
		{
			Key:   utils.KeyPubkey,
			Value: []byte(pubkey),
		},
		{
			Key:   utils.KeyPubkeyOrgId,
			Value: []byte(org_id),
		},
	}

	return cc.createPubkeyManagePayload(syscontract.PubkeyManageFunction_PUBKEY_DELETE.String(), pairs), nil
}

func (cc *ChainClient) CreatePubkeyQueryPayload(pubkey string) (*common.Payload, error) {
	pairs := []*common.KeyValuePair{
		{
			Key:   utils.KeyPubkey,
			Value: []byte(pubkey),
		},
	}

	return cc.createPayload("", common.TxType_QUERY_CONTRACT, syscontract.SystemContract_PUBKEY_MANAGE.String(),
		syscontract.PubkeyManageFunction_PUBKEY_QUERY.String(), pairs, defaultSeq), nil
}


func (cc *ChainClient) SendPubkeyManageRequest(payload *common.Payload, endorsers []*common.EndorsementEntry,
	timeout int64, withSyncResult bool) (*common.TxResponse, error) {
	return cc.sendContractRequest(payload, endorsers, timeout, withSyncResult)
}