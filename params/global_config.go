package params

var (
	Block_Interval       = 5000  // generate new block interval
	MaxBlockSize_global  = 1200  // the block contains the maximum number of transactions
	InjectSpeed          = 1200  // the transaction inject speed
	TotalDataSize        = 24000 // the total number of txs
	BatchSize            = 2400  // supervisor read a batch of txs then send them, it should be larger than inject speed
	BrokerNum            = 10
	NodesInShard         = 4
	ShardNum             = 4
	DataWrite_path       = "./result/"          // measurement data result output path
	LogWrite_path        = "./log"              // log output path
	SupervisorAddr       = "172.19.79.91:18800" //supervisor ip address
	SupervisorListenAddr = "0.0.0.0:18800"
	FileInput            = `../txs_0.0cross.csv` //the raw BlockTransaction data path
)
