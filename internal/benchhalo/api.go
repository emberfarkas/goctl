package benchhalo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-bamboo/pkg/log"
	"github.com/imroc/req/v3"
	"google.golang.org/protobuf/runtime/protoimpl"
	"sync"
)

type WalletLoginReq struct {
	Address string `json:"address" binding:"required"`
	ChainId int64  `json:"chain_id" binding:"required"`
	Sign    string `json:"sign" binding:"required"`
	Wallet  string `json:"wallet"`
}

type RespLoginX struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Email       string `json:"email"`
	UserId      string `json:"user_id"`
	Header      string `json:"header"`
	Salt        string `json:"salt"`
}

type RespJsonData struct {
	Code int        `json:"code"`
	Msg  string     `json:"message"`
	Data RespLoginX `json:"data"`
}

func walletLogin(host string, ks *keystore.KeyStore, account accounts.Account) (resData *RespLoginX, err error) {
	log.Infof("----------")
	headers := map[string]string{
		//"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36",
	}
	walletSignResp := &RespJsonData{}
	walletSignReq := &WalletLoginReq{
		Address: account.Address.Hex(),
		ChainId: 56,
		Sign:    "hello",
		Wallet:  "OKX",
	}
	data := fmt.Sprintf("address=%s,chain_id=%d", walletSignReq.Address, walletSignReq.ChainId)
	message := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	hash := crypto.Keccak256Hash([]byte(message))
	signature, err := ks.SignHashWithPassphrase(account, "123456", hash.Bytes())
	if err != nil {
		return nil, err
	}
	signature[64] += 27
	walletSignReq.Sign = hexutil.Encode(signature)
	res, err := req.R().SetHeaders(headers).SetBody(walletSignReq).EnableDump().Post(fmt.Sprintf("%v", host))
	if err != nil {
		log.Error(err)
		return
	}
	if err := json.Unmarshal(res.Bytes(), walletSignResp); err != nil {
		return nil, err
	}
	if walletSignResp.Code > 0 {
		log.Error("-------------------code: %v", walletSignResp.Code)
		return nil, errors.New("resp code")
	}
	log.Debugf("dump: %v", res.Dump())
	log.Debugf("AccessToken: %v", walletSignResp.Data.AccessToken)
	log.Debugf("UserId: %v", walletSignResp.Data.UserId)
	return &walletSignResp.Data, nil
}

type NewBip44Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mnemonic string `protobuf:"bytes,1,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
	Pass     string `protobuf:"bytes,2,opt,name=pass,proto3" json:"pass,omitempty"`
	Index    uint32 `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`
	Ty       int    `protobuf:"varint,6,opt,name=ty,proto3,enum=address.v1.CreateTy" json:"ty,omitempty"`
}

func walletGen(wg *sync.WaitGroup) (resData *RespLoginX, err error) {
	defer wg.Done()
	walletSignReq := NewBip44Request{
		Mnemonic: "てんてき すわる さっきょく はせる つまる てちょう いなか じゅしん ほあん しゃおん わかやま なめらか",
	}
	_, err = req.R().SetBody(walletSignReq).Post("http://127.0.0.1:8009/chain/createAddressByMnemonic")
	if err != nil {
		log.Error(err)
		return
	}
	//log.Infof("---------%v", res)
	return nil, nil
}
