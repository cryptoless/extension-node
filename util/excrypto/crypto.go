package excrypto

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TrimAddrPre(s string) string {
	return strings.Trim(s, "0x")
}

func Addr2Bytes(addr string) []byte {
	return common.Hex2Bytes(addr)
}

func ContractSlotKey(slot int64, data string) (h common.Hash) {
	d := TrimAddrPre(data)
	return crypto.Keccak256Hash(
		common.LeftPadBytes(common.Hex2Bytes(d), 32),
		common.LeftPadBytes(big.NewInt(slot).Bytes(), 32),
	)
}

func ContractKey(data ...string) (h common.Hash) {
	byteList := [][]byte{}
	for _, d := range data {
		td := TrimAddrPre(d)
		byteList = append(byteList, common.LeftPadBytes(common.Hex2Bytes(td), 32))
	}
	return crypto.Keccak256Hash(
		byteList...,
	)
}
