package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/avneeshatri/goapp/annpurna/util"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

var hlfUtil = &util.HlfUtil{}
var crytoUtil = &util.CryptoUtil{}

var (
	Log *log.Logger
)

type WalletContract struct {
	contractapi.Contract
}

type AnnpurnaWallet struct {
	Id        string `json:"id"`
	CreatedOn int    `json:"createdOn"`
	CreatedBy string `json:"createdBy"`
	Owner     string `json:"owner,omitempty"`
	Value     int    `json:"value"`
	Secret    string `json:"secret,omitempty"`
}

const (
	ZUDEXO_MSP          string = "ZudexoMSP"
	CREATE_WALLET_EVENT string = "CREATE_WALLET"
	ADD_FUND_EVENT      string = "ADD_FUND"
)

func (s *WalletContract) CreateWallet(ctx contractapi.TransactionContextInterface,
	walletJson string, signature string) (string, error) {
	Log.Println("Request to Create Wallet")

	initiator, err := hlfUtil.GetTxInitatorUserID(ctx)
	if err != nil {
		panic(err)
	}
	wallet := AnnpurnaWallet{}
	err = json.Unmarshal([]byte(walletJson), &wallet)

	wallet.CreatedBy = initiator

	if err != nil {
		panic(err)
	}

	if hlfUtil.StateExists(ctx, wallet.Id) {
		panic(fmt.Errorf("wallet with id  %s already exist", wallet.Id))
	}
	assertSignature(wallet, signature)
	wallet.Value = 0
	wallet.CreatedOn = 0

	walletAsBytes, _ := json.Marshal(wallet)
	ctx.GetStub().SetEvent(CREATE_WALLET_EVENT, walletAsBytes)
	Log.Println("Wallet Bytes size:", len(walletAsBytes))
	Log.Println("Wallet Json:", string(walletAsBytes))
	if hlfUtil.UpdateStateInLeger(ctx, wallet.Id, wallet) {

		//wallet.Secret = ""
		//wallet.Owner = ""
		walletAsBytes, _ = json.Marshal(wallet)
		jsonStr := string(walletAsBytes)
		Log.Println("Response:", jsonStr)

		return jsonStr, nil
	}
	panic("Wallet creation failed")
}

func (s *WalletContract) CreatePartnerWallet(ctx contractapi.TransactionContextInterface,
	orgMspId string) string {
	Log.Println("Request to create parnter wallet ", orgMspId)
	clientOrgMspId, err := hlfUtil.GetClientOrgId(ctx)
	if err != nil {
		panic(err)
	}
	if ZUDEXO_MSP != clientOrgMspId {
		errorMessage := "Only %s is allowed to create Parter Wallet not %s" + ZUDEXO_MSP + orgMspId
		fmt.Println(errorMessage)
	}
	wallet := AnnpurnaWallet{}
	wallet.Id = orgMspId
	wallet.Owner = orgMspId
	wallet.Value = 0
	wallet.CreatedBy = clientOrgMspId

	walletAsBytes, _ := json.Marshal(wallet)
	Log.Println("Wallet Bytes size:", len(walletAsBytes))
	ctx.GetStub().SetEvent(CREATE_WALLET_EVENT, walletAsBytes)
	jsonStr := string(walletAsBytes)
	Log.Println("Response:", jsonStr)
	if hlfUtil.UpdateStateInLeger(ctx, wallet.Id, wallet) {
		return jsonStr
	}
	panic("Couldnt save asset")
}

func (s *WalletContract) ReadWallet(ctx contractapi.TransactionContextInterface,
	walletId string) AnnpurnaWallet {

	return readAnnpurnaWallet(ctx, walletId)
}

func readAnnpurnaWallet(ctx contractapi.TransactionContextInterface,
	walletId string) AnnpurnaWallet {
	Log.Println("Geting state for Id:", walletId)
	walletBytes, err := hlfUtil.ReadState(ctx, walletId)
	jsonStr := string(walletBytes)
	Log.Println("Response:", jsonStr)
	if err != nil {
		panic(err)
	} else if len(walletBytes) == 0 {
		panic("Wallet not found with id " + walletId)
	}
	wallet := AnnpurnaWallet{}
	err = json.Unmarshal(walletBytes, &wallet)

	if nil != err {
		panic(err)
	}

	return wallet
}

func (s *WalletContract) GetWalletHistory(ctx contractapi.TransactionContextInterface,
	walletId string, signature string) []AnnpurnaWallet {

	if !hlfUtil.StateExists(ctx, walletId) {
		panic(fmt.Errorf("wallet with id  %s doesnot exist", walletId))
	}

	userWallet := readAnnpurnaWallet(ctx, walletId)

	assertSignature(userWallet, signature)
	historyIer, err := ctx.GetStub().GetHistoryForKey(walletId)

	if err != nil {
		errMsg := "Coudnt Fetch history of wallet:" + walletId + ",Err:" + err.Error()
		fmt.Println(errMsg)
		panic(err)
	}
	history := []AnnpurnaWallet{}
	i := 0
	for historyIer.HasNext() {
		modification, err := historyIer.Next()
		i++
		if err != nil {
			errMsg := "Coudnt Fetch Next history of wallet:" + walletId + ",Err:" + err.Error()
			fmt.Println(errMsg)
			panic(errMsg)
		}
		wallet := AnnpurnaWallet{}
		json.Unmarshal(modification.GetValue(), &wallet)
		history = append(history, wallet)

	}
	historyIer.Close()
	Log.Println("History update count:", i)

	if err != nil {
		panic(err.Error())
	}
	return history
}

func (s *WalletContract) BalanceOf(ctx contractapi.TransactionContextInterface,
	walletId string, signature string) int {
	wallet := readAnnpurnaWallet(ctx, walletId)
	clientOrgMspId, err := hlfUtil.GetClientOrgId(ctx)
	if err != nil {
		panic(err)
	}
	if clientOrgMspId != walletId {
		assertSignature(wallet, signature)
	}

	return wallet.Value
}

func (s *WalletContract) Transfer(ctx contractapi.TransactionContextInterface, signature string,
	senderWalletId string, recipientWalletId string, amount int) {
	walletSender := readAnnpurnaWallet(ctx, senderWalletId)

	assertSignature(walletSender, signature)
	walleRecipent := readAnnpurnaWallet(ctx, recipientWalletId)
	walleRecipent.Value = walleRecipent.Value + amount
	walletSender.Value = walletSender.Value - amount

	updateWallet(ctx, walleRecipent)
	updateWallet(ctx, walletSender)

}

func (s *WalletContract) AddFunds(ctx contractapi.TransactionContextInterface, value int) string {

	orgMspId, err := hlfUtil.GetClientOrgId(ctx)
	if err != nil {
		panic(err)
	}

	if ZUDEXO_MSP != orgMspId {
		panic("Only %s is allowed to create Parter Wallet not " + orgMspId)

	}
	wallet := readAnnpurnaWallet(ctx, orgMspId)
	wallet.Value = wallet.Value + value
	walletBytes, err := json.Marshal(wallet)
	if err != nil {
		panic(err)
	}
	ctx.GetStub().SetEvent(ADD_FUND_EVENT, []byte("Added Fund")) //+strconv.FormatInt(int64(value), 10)))
	jsonStr := string(walletBytes)
	Log.Println("Response:", jsonStr)
	if hlfUtil.UpdateStateInLeger(ctx, wallet.Id, wallet) {
		return jsonStr
	}
	panic("Add fund failed")

}

func (s *WalletContract) TransferTo(ctx contractapi.TransactionContextInterface,
	recipientWalletId string, amount int) {
	orgMspId, err := hlfUtil.GetClientOrgId(ctx)
	if err != nil {
		panic("Couldnot evaluate MSP Id of sender")
	}
	walletSender := readAnnpurnaWallet(ctx, orgMspId)

	walleRecipent := readAnnpurnaWallet(ctx, recipientWalletId)
	walleRecipent.Value = walleRecipent.Value + amount
	walletSender.Value = walletSender.Value - amount

	updateWallet(ctx, walleRecipent)
	updateWallet(ctx, walletSender)

}

func updateWallet(ctx contractapi.TransactionContextInterface, wallet AnnpurnaWallet) AnnpurnaWallet {
	if !hlfUtil.StateExists(ctx, wallet.Id) {
		panic(fmt.Errorf("wallet with id %s already exist", wallet.Id))
	}
	if hlfUtil.UpdateStateInLeger(ctx, wallet.Id, wallet) {
		return wallet
	}
	panic("Couldnt save asset")
}

func assertSignature(wallet AnnpurnaWallet, signature string) {

	status := crytoUtil.VerifyWithPublicKeyPKCS8(crytoUtil.DecodeBase64(signature),
		wallet.Id, crytoUtil.DecodeBase64(wallet.Owner))

	if !status {
		panic("Signature verification failed")
	}
}

func setUpLogger() {
	// set location of log file
	logpath := "/home/atri/workspace_hlf/annpurna/scripts/logs/external_builder/" + os.Getenv("CORE_PEER_LOCALMSPID") + "_chaincode_go.log"
	f, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	//defer f.Close()
	//wrt := io.MultiWriter(os.Stdout, f)
	//log.SetOutput(wrt)
	//log.Println("LogFile : " + logpath)
	Log = log.New(f, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)

}

func main() {
	setUpLogger()
	Log.Println("Start")
	chaincode, err := contractapi.NewChaincode(new(WalletContract))

	if err != nil {
		fmt.Printf("Error create annpurna-wallet chaincode: %s", err.Error())
		panic(err)
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting annpurna chaincode: %s", err.Error())
	}
}
