package util

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type HlfUtil struct {
}

func (h *HlfUtil) StateExists(ctx contractapi.TransactionContextInterface, stateId string) bool {
	stub := ctx.GetStub()
	stateBytes, err := stub.GetState(stateId)
	if err != nil || stateBytes == nil || len(stateBytes) == 0 {
		return false
	}
	return true
}

/*
func (h *HlfUtil) GetTxInitatorUserID(ctx contractapi.TransactionContextInterface) (string, error) {
	creator, err := ctx.GetStub().GetCreator()
	if err != nil {
		return "", err
	}
	return string(creator), nil
}
*/
// Unmarshals the bytes returned by ChaincodeStubInterface.GetCreator method and
// returns the resulting msp.SerializedIdentity object
func (h *HlfUtil) GetTxInitatorUserID(ctx contractapi.TransactionContextInterface) (string, error) {
	client, err := cid.New(ctx.GetStub())
	fmt.Println("client", client)
	if err != nil {
		return "", err
	}
	id, err := client.GetMSPID()
	if err != nil {
		return "", err
	}
	return id, nil
}

func (h *HlfUtil) GetClientOrgId(ctx contractapi.TransactionContextInterface) (string, error) {
	return ctx.GetClientIdentity().GetMSPID()
}

func (h *HlfUtil) UpdateStateInLeger(ctx contractapi.TransactionContextInterface,
	stateId string, state interface{}) bool {
	stub := ctx.GetStub()
	stateBytes, _ := json.Marshal(state)
	stub.PutState(stateId, stateBytes)
	return true
}

func (h *HlfUtil) ReadState(ctx contractapi.TransactionContextInterface, stateId string) ([]byte, error) {
	stub := ctx.GetStub()
	state, err := stub.GetState(stateId)
	if err != nil {
		return nil, err
	}
	return state, nil
}
