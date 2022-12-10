package benchmark

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/go-bamboo/contrib/contracts/flattened"
	"github.com/go-bamboo/pkg/tools"
	"github.com/sony/sonyflake"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

const (
	_prefixChainNonce = "nonce_%s_%s" // chain -> x:
)

func keyChainNonce(chain string, addr string) string {
	return fmt.Sprintf(_prefixChainNonce, chain, addr)
}

type stat struct {
	send    int32
	success int32
	fail    int32
}

type biz struct {
	nonces sync.Map
	db     *leveldb.DB
	ks     *keystore.KeyStore
	sf     *sonyflake.Sonyflake
	stat   stat
	hashs  sync.Map
	hashx  sync.Map
}

func newBiz() (*biz, error) {
	db, err := leveldb.OpenFile("data", &opt.Options{})
	if err != nil {
		return nil, err
	}
	sf := sonyflake.NewSonyflake(sonyflake.Settings{
		MachineID: func() (uint16, error) {
			return uint16(1), nil
		},
	})
	ks := keystore.NewKeyStore("keystore", int(math.Pow(2, 7)), int(math.Pow(2, 9)))
	return &biz{
		db: db,
		ks: ks,
		sf: sf,
	}, nil
}

func (uc *biz) close() error {
	return uc.db.Close()
}

func (uc *biz) GetNonce(ctx context.Context, chain, address string) (nonce int64, err error) {
	key := keyChainNonce(chain, address)
	v, ok := uc.nonces.Load(key)
	if ok {
		return v.(int64), nil
	} else {
		buf, err := uc.db.Get([]byte(key), &opt.ReadOptions{})
		if err != nil {
			if err == leveldb.ErrNotFound {
				return 0, nil
			}
			return 0, err
		}
		return strconv.ParseInt(string(buf), 10, 64)
	}
}

func (uc *biz) SetNonce(ctx context.Context, chain, address string, nonce int64) (err error) {
	key := keyChainNonce(chain, address)
	uc.nonces.Store(key, nonce)
	return nil
}

func (uc *biz) FlushNonce(ctx context.Context) (err error) {
	uc.nonces.Range(func(key, value interface{}) bool {
		k := []byte(key.(string))
		v := []byte(fmt.Sprint(value))
		if err = uc.db.Put(k, v, nil); err != nil {
			log.Printf("err : %v", err)
		}
		return true
	})
	return nil
}

// func (uc *biz) Transfer(ctx context.Context, from, pk, to string) (err error) {
// 	nonce, _ := uc.GetNonce(ctx, "ETH", from)
// 	if nonce > 0 {
// 		nonce = nonce + 1
// 	}
// 	logger := log.NewStdLogger(os.Stdout)
// 	client := eth.New(logger, "http://121.36.71.137:7545")
// 	no, err := client.EthGetTransactionCount(from, "latest")
// 	if err != nil {
// 		return
// 	}
// 	// 本地nonce值与链上比较
// 	if nonce < int64(no) {
// 		nonce = int64(no)
// 	}
// 	fmt.Printf("nonce: %v\n", nonce)

// 	agent := gorequest.New()
// 	agent = agent.Timeout(30 * time.Second)
// 	var x Web3TransferReply
// 	req := Web3TransferRequest{
// 		From:  from,
// 		Pk:    pk,
// 		Nonce: nonce,
// 		To:    to,
// 	}
// 	resp, _, errs := agent.Post("http://192.168.110.99:7002/api/v1/account/transfer").Send(req).EndStruct(&x)
// 	if len(errs) > 0 {
// 		return errs[0]
// 	} else if resp.StatusCode != http.StatusOK {
// 		err = ecode.InternalServer(x.Reason, x.Message)
// 		return
// 	}

// 	// gasPrice, err := client.EthGasPrice()
// 	// if err != nil {
// 	// 	return
// 	// }
// 	// tokenAddress := common.HexToAddress("0x71c61bD5bB80d61d98D3538098ba32Af0B9E48ff")
// 	// tokenAddress := common.HexToAddress(to)
// 	// value := big.NewInt(int64(math.Pow10(18) * 10))
// 	// data := []byte("")
// 	// tx := types.NewTransaction(uint64(nonce), tokenAddress, value, 21000, &gasPrice, data)

// 	// privateKey, err := crypto.HexToECDSA(pk)
// 	// if err != nil {
// 	// 	err = ecode.WrapError(err)
// 	// 	return
// 	// }
// 	// signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
// 	// if err != nil {
// 	// 	err = ecode.WrapError(err)
// 	// 	return
// 	// }

// 	_, err = client.EthSendRawTransaction(x.Tx.RawTransaction)
// 	if err != nil {
// 		fmt.Printf("---------------------------------------\n")
// 		fmt.Printf("reason: %v, message: %v \n", se.Reason, se.Message)
// 		return
// 	}
// 	if err = SetNonce(ctx, db, "ETH", from, nonce); err != nil {
// 		return
// 	}
// 	return
// }

func (uc *biz) contractsMint(ctx context.Context, wg *sync.WaitGroup, from accounts.Account) {
	defer func() {
		wg.Done()
		// log.Printf("ContractsMint done")
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			pl := fmt.Sprintf("ContractsMint call panic: %v\n%s\n", err, buf)
			log.Printf("%s", pl)
		}
	}()
	nonce, _ := uc.GetNonce(ctx, "ETH", from.Address.Hex())
	ctrctAddr := common.HexToAddress(contractAddress)
	for i := 0; i < n; i++ {
		tokenURI := tools.GetUUID()
		contentHash := tools.GetUUID()
		tokenId, _ := uc.sf.NextID()
		c := pool.Get().(*ethclient.Client)
		media, err := flattened.NewMedia(ctrctAddr, c)
		if err != nil {
			log.Printf("err: %v", err)
			return
		}
		hash, rawTx, err := media.MintKs(ctx, big.NewInt(int64(chainID)), uc.ks, from, big.NewInt(nonce), big.NewInt(int64(tokenId)), tokenURI, contentHash)
		pool.Put(c)
		if err != nil {
			log.Printf("err: %v", err)
			return
		}

		if _, ok := uc.hashs.Load(hash); ok {
			log.Printf("err: 已经存在%v", hash)
			return
		}
		uc.hashs.Store(hash, rawTx)
		uc.hashx.Store(hash, 0)
		atomic.AddInt32(&uc.stat.send, 1)

		nonce = nonce + 1
		if err = uc.SetNonce(ctx, "ETH", from.Address.Hex(), nonce); err != nil {
			log.Printf("err: %v", err)
			return
		}
	}
}

func (uc *biz) post(ctx context.Context, wg *sync.WaitGroup, ch chan string) {
	defer func() {
		wg.Done()
		// log.Printf("Post done")
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			pl := fmt.Sprintf("Post call panic: %v\n%s\n", err, buf)
			log.Printf("%s", pl)
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case rawTx := <-ch:
			if len(rawTx) > 0 {
				client := rpcpool.Get().(*ethclient.Client)
				hash, _, err := client.TransactionByHash(context.TODO(), common.HexToHash(rawTx))
				rpcpool.Put(client)
				if err != nil {
					log.Printf("3. warn: %v", err)
					continue
				}
				if _, ok := uc.hashs.Load(hash); ok {
					log.Printf("4. err: 已经存在%v", hash)
					return
				}
				uc.hashs.Store(hash, rawTx)
				uc.hashx.Store(hash, 0)
				atomic.AddInt32(&uc.stat.send, 1)
			} else {
				log.Printf("err: 发现空hash")
			}
		default:
		}
	}
}

func (uc *biz) transfer(ctx context.Context, wg *sync.WaitGroup, from common.Address, priv *ecdsa.PrivateKey, to common.Address, v int64, ch chan string) {
	defer func() {
		wg.Done()
		// log.Printf("Transfer done")
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			pl := fmt.Sprintf("Transfer call panic: %v\n%s\n", err, buf)
			log.Printf("%s", pl)
		}
	}()
	nonce, _ := uc.GetNonce(ctx, "ETH", from.Hex())

	str := "save,"
	for i := 0; i < 32*1024; i++ {
		str = str + "a"
	}

	strBytes := []byte(str)
	hex := hexutil.Encode(strBytes)
	// log.Printf("hex: %v", hex)

	for i := 0; i < n; i++ {
		// 1.10.8
		// log.Printf("nonce %v", nonce)
		// baseTx := &types.Transaction{} LegacyTx{
		// 	Nonce:    uint64(nonce),
		// 	GasPrice: big.NewInt(1),
		// 	Gas:      23000000,
		// 	To:       &to,
		// 	Value:    big.NewInt(v),
		// 	Data:     []byte(hex),
		// }
		// tx := types.NewTx(baseTx)
		// signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(int64(chainID))), priv)
		// if err != nil {
		// 	log.Printf("1. err: %v", err)
		// 	return
		// }
		// data, err := signedTx.MarshalBinary()
		// if err != nil {
		// 	log.Printf("2. err: %v", err)
		// 	return
		// }

		// 1.9.25
		tx := types.NewTransaction(uint64(nonce), to, big.NewInt(v), 2300000, big.NewInt(1), []byte(hex))
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(int64(chainID))), priv)
		if err != nil {
			log.Printf("1. err: %v", err)
			return
		}
		data, err := rlp.EncodeToBytes(signedTx)
		if err != nil {
			return
		}
		rawTx := hexutil.Encode(data)
		ch <- rawTx
		nonce = nonce + 1
		if err = uc.SetNonce(ctx, "ETH", from.Hex(), nonce); err != nil {
			log.Printf("5. err: %v", err)
			return
		}
	}
}

func (uc *biz) transferV(ctx context.Context, wg *sync.WaitGroup, from accounts.Account, to accounts.Account, v int64, ch chan string) {
	defer func() {
		wg.Done()
		// log.Printf("transferV done")
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			pl := fmt.Sprintf("transferV call panic: %v\n%s\n", err, buf)
			log.Printf("%s", pl)
		}
	}()
	nonce, _ := uc.GetNonce(ctx, "ETH", from.Address.Hex())

	str := "你好,hello"
	strBytes := []byte(str)
	hex := hexutil.Encode(strBytes)

	for i := 0; i < n; i++ {
		// 1.10.8
		// log.Printf("address %v nonce: %v", from.Address.Hex(), nonce)
		// baseTx := &types.LegacyTx{
		// 	Nonce:    uint64(nonce),
		// 	GasPrice: big.NewInt(1),
		// 	Gas:      23000,
		// 	To:       &to.Address,
		// 	Value:    big.NewInt(v),
		// 	Data:     []byte(hex),
		// }
		// tx := types.NewTx(baseTx)
		// signedTx, err := uc.ks.SignTx(from, tx, big.NewInt(int64(chainID)))
		// if err != nil {
		// 	log.Printf("err: %v", err)
		// 	return
		// }
		// data, err := signedTx.MarshalBinary()
		// if err != nil {
		// 	log.Printf("err: %v", err)
		// 	return
		// }

		// 1.9.25
		tx := types.NewTransaction(uint64(nonce), to.Address, big.NewInt(v), 23000000, big.NewInt(1), []byte(hex))
		signedTx, err := uc.ks.SignTx(from, tx, big.NewInt(int64(chainID)))
		if err != nil {
			log.Printf("err: %v", err)
			return
		}
		data, err := rlp.EncodeToBytes(signedTx)
		if err != nil {
			return
		}
		rawTx := hexutil.Encode(data)
		ch <- rawTx
		nonce = nonce + 1
		if err = uc.SetNonce(ctx, "ETH", from.Address.Hex(), nonce); err != nil {
			log.Printf("err: %v", err)
			return
		}
	}
	return
}

func (uc *biz) runTest(ctx context.Context) error {
	privKey, err := crypto.HexToECDSA("cca5c1a080858a7d59b6c2246a6186db5e47bc9c7de184b24f9b1a6d03989d24")
	if err != nil {
		return err
	}
	pubKey := privKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return errInvalidContractAddress
	}
	fromAddr := crypto.PubkeyToAddress(*pubKeyECDSA)

	// 获取交易数
	client := rpcpool.Get().(*ethclient.Client)
	no, err := client.NonceAt(context.TODO(), fromAddr, nil)
	rpcpool.Put(client)
	if err != nil {
		log.Printf("err: %v", err)
		return err
	}

	// 本地nonce值与链上比较
	// nonce
	nonce, _ := uc.GetNonce(ctx, "ETH", fromAddr.Hex())
	if nonce < int64(no) {
		uc.SetNonce(ctx, "ETH", fromAddr.Hex(), int64(no))
	}

	// 发送协程
	ch := make(chan string)
	var pwg sync.WaitGroup
	pCtx, pCancel := context.WithCancel(ctx)
	for i := 0; i < 1; i++ {
		pwg.Add(1)
		go uc.post(pCtx, &pwg, ch)
	}

	// mp := uc.ks.Accounts()
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1; i++ {
		to := fromAddr
		wg.Add(1)
		uc.transfer(ctx, &wg, fromAddr, privKey, to, 0, ch)
	}
	wg.Wait()
	pCancel()
	pwg.Wait()
	cost := time.Since(start)
	log.Printf("send: %v, cost : %vs", 1*n, cost.Seconds())
	return nil
}

func (uc *biz) runTest1(ctx context.Context) error {
	privKey, err := crypto.HexToECDSA("cca5c1a080858a7d59b6c2246a6186db5e47bc9c7de184b24f9b1a6d03989d24")
	if err != nil {
		return err
	}
	pubKey := privKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return errInvalidContractAddress
	}
	fromAddr := crypto.PubkeyToAddress(*pubKeyECDSA)

	// 获取线上nonce值
	client := rpcpool.Get().(*ethclient.Client)
	no, err := client.NonceAt(context.TODO(), fromAddr, nil)
	rpcpool.Put(client)
	if err != nil {
		log.Printf("err: %v", err)
		return err
	}
	nonce, err := uc.GetNonce(ctx, "ETH", fromAddr.Hex())
	if err != nil {
		log.Printf("err : %v", err)
		return err
	}
	if nonce < int64(no) {
		uc.SetNonce(ctx, "ETH", fromAddr.Hex(), int64(no))
	}

	// 发送协程
	ch := make(chan string)
	var pwg sync.WaitGroup
	pCtx, pCancel := context.WithCancel(ctx)
	for i := 0; i < 1; i++ {
		pwg.Add(1)
		go uc.post(pCtx, &pwg, ch)
	}

	// 打包交易
	mp := uc.ks.Accounts()
	start := time.Now()
	ti := time.NewTimer(time.Second)
	i, n := 0, 0
	for {
		select {
		case <-ctx.Done():
			log.Printf("runTest1 done")
			goto handlePost
		case <-ti.C:
			// log.Printf("---------------------run 1 sec")
			start = time.Now()
			ti.Reset(time.Second)
			n = 0
		default:
			if n < 50 {
				if i >= len(mp) {
					i = 0
				}
				var wg sync.WaitGroup
				to := mp[i]
				wg.Add(1)
				go uc.transfer(ctx, &wg, fromAddr, privKey, to.Address, 1, ch)
				wg.Wait()
				i++
				n++
			} else {
				d := time.Since(start)
				sec1 := time.Second
				if sec1 > d {
					du := sec1 - d
					time.Sleep(du)
					// log.Printf("sleep: %vms", du.Microseconds())
				}
			}
		}
	}
handlePost:
	pCancel()
	pwg.Wait()
	// cost := time.Since(start)
	// log.Printf("send: %v cost : %vs", len(mp)*n, cost.Seconds())
	return nil
}

func (uc *biz) runTest2(ctx context.Context) error {
	privKey, err := crypto.HexToECDSA("cca5c1a080858a7d59b6c2246a6186db5e47bc9c7de184b24f9b1a6d03989d24")
	if err != nil {
		return err
	}
	pubKey := privKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return errInvalidContractAddress
	}
	fromAddr := crypto.PubkeyToAddress(*pubKeyECDSA)

	// 获取线上nonce值
	client := rpcpool.Get().(*ethclient.Client)
	no, err := client.NonceAt(context.TODO(), fromAddr, nil)
	rpcpool.Put(client)
	if err != nil {
		log.Printf("err: %v", err)
		return err
	}
	nonce, err := uc.GetNonce(ctx, "ETH", fromAddr.Hex())
	if err != nil {
		log.Printf("err : %v", err)
		return err
	}
	if nonce < int64(no) {
		uc.SetNonce(ctx, "ETH", fromAddr.Hex(), int64(no))
	}

	// 发送协程
	ch := make(chan string)
	var pwg sync.WaitGroup
	pCtx, pCancel := context.WithCancel(ctx)
	for i := 0; i < 1; i++ {
		pwg.Add(1)
		go uc.post(pCtx, &pwg, ch)
	}

	mp := uc.ks.Accounts()
	var wg sync.WaitGroup
	to := mp[0]
	wg.Add(1)
	go uc.transfer(ctx, &wg, fromAddr, privKey, to.Address, 1, ch)
	wg.Wait()

	pCancel()
	pwg.Wait()
	// cost := time.Since(start)
	// log.Printf("send: %v cost : %vs", len(mp)*n, cost.Seconds())
	return nil
}

func (uc *biz) runMint(ctx context.Context) error {
	if contractAddress == "" {
		return errInvalidContractAddress
	}
	if chainID <= 0 {
		return errInvalidChainID
	}
	var wg sync.WaitGroup
	mp := uc.ks.Accounts()
	for i := 0; i < len(mp); i++ {
		acc := mp[i]
		if err := uc.ks.Unlock(acc, "123456"); err != nil {
			return err
		}
		client := rpcpool.Get().(*ethclient.Client)
		no, err := client.NonceAt(context.TODO(), acc.Address, nil)
		rpcpool.Put(client)
		if err != nil {
			return err
		}
		nonce, err := uc.GetNonce(ctx, "ETH", acc.Address.Hex())
		if err != nil {
			return err
		}
		if nonce < int64(no) {
			if err = uc.SetNonce(ctx, "ETH", acc.Address.Hex(), int64(no)); err != nil {
				return err
			}
		}
	}
	start := time.Now()
	for i := 0; i < len(mp); i++ {
		acc := mp[i]
		wg.Add(1)
		go uc.contractsMint(ctx, &wg, acc)
	}
	wg.Wait()
	cost := time.Since(start)
	log.Printf("send: %v cost : %vs", len(mp)*n, cost.Seconds())
	return nil
}

func (uc *biz) runInitV(ctx context.Context) error {
	privKey, err := crypto.HexToECDSA("cca5c1a080858a7d59b6c2246a6186db5e47bc9c7de184b24f9b1a6d03989d24")
	if err != nil {
		return err
	}
	pubKey := privKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return errInvalidContractAddress
	}
	fromAddr := crypto.PubkeyToAddress(*pubKeyECDSA)

	// 获取线上nonce值
	client := rpcpool.Get().(*ethclient.Client)
	no, err := client.NonceAt(context.TODO(), fromAddr, nil)
	rpcpool.Put(client)
	if err != nil {
		log.Printf("err: %v", err)
		return err
	}
	nonce, err := uc.GetNonce(ctx, "ETH", fromAddr.Hex())
	if err != nil {
		log.Printf("err : %v", err)
		return err
	}
	if nonce < int64(no) {
		uc.SetNonce(ctx, "ETH", fromAddr.Hex(), int64(no))
	}

	// 发送协程
	ch := make(chan string)
	var pwg sync.WaitGroup
	pCtx, pCancel := context.WithCancel(ctx)
	for i := 0; i < 1; i++ {
		pwg.Add(1)
		go uc.post(pCtx, &pwg, ch)
	}

	// 打包交易
	mp := uc.ks.Accounts()
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < len(mp); i++ {
		to := mp[i]
		wg.Add(1)
		uc.transfer(ctx, &wg, fromAddr, privKey, to.Address, 10000000000000000, ch)
	}
	wg.Wait()
	pCancel()
	pwg.Wait()
	cost := time.Since(start)
	log.Printf("send: %v cost : %vs", len(mp)*n, cost.Seconds())
	return nil
}

func (uc *biz) runTransV(ctx context.Context) error {
	// 发送协程
	ch := make(chan string)
	var pwg sync.WaitGroup
	pCtx, pCancel := context.WithCancel(ctx)
	for i := 0; i < 1; i++ {
		pwg.Add(1)
		go uc.post(pCtx, &pwg, ch)
	}

	// 解锁
	var wg sync.WaitGroup
	mp := uc.ks.Accounts()
	for i := 0; i < len(mp); i++ {
		from := mp[i]
		if err := uc.ks.Unlock(from, "123456"); err != nil {
			log.Printf("err : %v", err)
			return err
		}

		// 拿去nonce
		client := rpcpool.Get().(*ethclient.Client)
		no, err := client.NonceAt(context.TODO(), from.Address, nil)
		rpcpool.Put(client)
		if err != nil {
			return err
		}
		nonce, err := uc.GetNonce(ctx, "ETH", from.Address.Hex())
		if err != nil {
			log.Printf("err : %v", err)
			return err
		}
		if nonce < int64(no) {
			if err = uc.SetNonce(ctx, "ETH", from.Address.Hex(), int64(no)); err != nil {
				return err
			}
		}
	}

	// 发送交易
	start := time.Now()
	for i := 0; i < len(mp); i++ {
		from := mp[i]
		next := i + 1
		if next >= len(mp) {
			next = 0
		}
		to := mp[next]
		wg.Add(1)
		go uc.transferV(ctx, &wg, from, to, 1, ch)
	}
	wg.Wait()
	pCancel()
	pwg.Wait()
	cost := time.Since(start)
	log.Printf("send: %v cost : %vs", len(mp)*n, cost.Seconds())
	return nil
}

func (uc *biz) showResult(ctx context.Context) error {
	// 获取交易数
	client := rpcpool.Get().(*ethclient.Client)
	for uc.stat.fail+uc.stat.success < uc.stat.send {
		//
		start := time.Now()
		uc.hashs.Range(func(key, value interface{}) bool {
			state, ok := uc.hashx.Load(key)
			if !ok {
				log.Printf("不存在: %v", key)
				return true
			}
			switch state {
			case 0:
				// test
				receipt, err := client.TransactionReceipt(context.TODO(), common.HexToHash(key.(string)))
				if err != nil {
					log.Printf("err: %v", err)
					return true
				}
				if receipt.Status == 1 {
					atomic.AddInt32(&uc.stat.success, 1)
					uc.hashx.Store(key, 1)

					// 成功后，测试数据
					tx, _, err := client.TransactionByHash(context.TODO(), common.HexToHash(key.(string)))
					if err != nil {
						log.Printf("err: %v", err)
						return true
					}
					input, err := hexutil.Decode(string(tx.Data()))
					if err != nil {
						log.Printf("err: %v", err)
						return true
					}
					str, err := hexutil.Decode(string(input))
					if err != nil {
						log.Printf("err: %v", err)
						return true
					}
					log.Printf("input data : %v", string(str))
				} else if receipt.Status == 0 {
					atomic.AddInt32(&uc.stat.fail, 1)
					uc.hashx.Store(key, 2)
				} else {
					if retry == 1 {
						//client.EthSendRawTransaction(value.(string))
					}
				}
			case 2:
				if retry == 1 {
					//client.EthSendRawTransaction(value.(string))
				}
			}
			return true
		})
		cost := time.Since(start)
		log.Printf("query: %vs, send: %v, success: %v, fail : %v, sleep 1 sec", cost.Seconds(), uc.stat.send, uc.stat.success, uc.stat.fail)
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (uc *biz) run(ctx context.Context) error {
	switch mint {
	case "transfer":
		if contractAddress == "" {
			return errInvalidContractAddress
		}
		// mp := uc.ks.Accounts()
		// for _, acc := range mp {
		// 	uc.ks.SignHashWithPassphrase()
		// 	if err := uc.Transfer(ctx, rdb, "0xEC05F40e44e36CB18fc47F3CEe036d79D4a0Dd31", "cca5c1a080858a7d59b6c2246a6186db5e47bc9c7de184b24f9b1a6d03989d24", addr); err != nil {
		// 		return err
		// 	}
		// }
	case "mint":
		if err := uc.runMint(ctx); err != nil {
			return err
		}
	case "initv":
		if err := uc.runInitV(ctx); err != nil {
			return err
		}
	case "transv":
		if err := uc.runTransV(ctx); err != nil {
			return err
		}
	case "test":
		if err := uc.runTest(ctx); err != nil {
			return err
		}
	case "test1":
		if err := uc.runTest1(ctx); err != nil {
			return err
		}
	case "test2":
		if err := uc.runTest2(ctx); err != nil {
			return err
		}
	default:
	}

	// flush
	if err := uc.FlushNonce(ctx); err != nil {
		log.Printf("err :%v", err)
		return err
	}

	// 显示最后的查询结果
	if show == 1 {
		if err := uc.showResult(ctx); err != nil {
			log.Printf("err :%v", err)
			return err
		}
	}

	return nil
}
