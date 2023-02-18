package eth

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestXxx(t *testing.T) {
	hash1 := common.BytesToHash(crypto.Keccak256([]byte("Transfer(address,address,uint256)")))
	t.Logf("hex: %v", hash1.Hex())

	hash2 := common.BytesToHash(crypto.Keccak256([]byte("eCreateTransferInfo(uint256,address,uint256,address,uint256,uint256,address)")))
	t.Logf("hex: %v", hash2.Hex())
}
