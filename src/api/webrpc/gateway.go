package webrpc

import (
	"hgeekl.com/gkicoin/src/cipher"
	"hgeekl.com/gkicoin/src/coin"
	"hgeekl.com/gkicoin/src/daemon"
	"hgeekl.com/gkicoin/src/visor"
	"hgeekl.com/gkicoin/src/visor/historydb"
)

//go:generate goautomock -template=testify Gatewayer

// Gatewayer provides interfaces for getting gkicoin related info.
type Gatewayer interface {
	GetLastBlocks(num uint64) (*visor.ReadableBlocks, error)
	GetBlocks(start, end uint64) (*visor.ReadableBlocks, error)
	GetBlocksInDepth(vs []uint64) (*visor.ReadableBlocks, error)
	GetUnspentOutputs(filters ...daemon.OutputsFilter) (*visor.ReadableOutputSet, error)
	GetTransaction(txid cipher.SHA256) (*visor.Transaction, error)
	InjectBroadcastTransaction(tx coin.Transaction) error
	GetAddrUxOuts(addr []cipher.Address) ([]*historydb.UxOut, error)
	GetTimeNow() uint64
}
