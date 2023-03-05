package eth

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestEventTopic(t *testing.T) {
	hash1 := common.BytesToHash(crypto.Keccak256([]byte("Transfer(address,address,uint256)")))
	t.Logf("hex: %v", hash1.Hex())

	hash2 := common.BytesToHash(crypto.Keccak256([]byte("eCreateTransferInfo(uint256,address,uint256,address,uint256,uint256,address)")))
	t.Logf("hex: %v", hash2.Hex())

	hash3 := common.BytesToHash(crypto.Keccak256([]byte("eCreateTransferInfo(uint256,address,uint256,address,uint256,uint256,address)")))
	t.Logf("hex: %v", hash3.Hex())

	hash4 := common.BytesToHash(crypto.Keccak256([]byte("eRoleWithdraw(address,uint256,address,uint256)")))
	t.Logf("hex: %v", hash4.Hex())

	hash5 := common.BytesToHash(crypto.Keccak256([]byte("eCost1155Info(uint256,address,address,uint256[],uint256[],uint256,address)")))
	t.Logf("hex: %v", hash5.Hex())

}
