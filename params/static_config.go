package params

import "math/big"

type ChainConfig struct {
	ChainID        uint64
	NodeID         uint64
	ShardID        uint64
	Nodes_perShard uint64
	ShardNums      uint64
	BlockSize      uint64
	BlockInterval  uint64
	InjectSpeed    uint64

	// used in transaction relaying, useless in brokerchain mechanism
	MaxRelayBlockSize uint64
}

var (
	DeciderShard          = uint64(0xffffffff)
	Init_Balance, _       = new(big.Int).SetString("100000000000000000000000000000000000000000000", 10)
	IPmap_nodeTable       = make(map[uint64]map[uint64]string)
	IPmap_ListenNodeTable = make(map[uint64]map[uint64]string)
	CommitteeMethod       = []string{"CLPA_Broker", "CLPA", "Broker", "Relay"}
	MeasureBrokerMod      = []string{"TPS_Broker", "TCL_Broker", "CrossTxRate_Broker", "TxNumberCount_Broker"}
	MeasureRelayMod       = []string{"TPS_Relay", "TCL_Relay", "CrossTxRate_Relay", "TxNumberCount_Relay"}
	IPmap_shardTable      = []string{"172.19.79.91:", "172.19.79.92:", "172.19.79.93:", "172.19.79.94:",
		"8.134.115.119:", "8.134.123.158:", "47.96.178.79:", "8.140.204.54:"}
)
