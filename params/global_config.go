package params

var (
	Block_Interval      = 5000  // generate new block interval
	MaxBlockSize_global = 2000  // the block contains the maximum number of transactions
	InjectSpeed         = 2000  // the transaction inject speed
	TotalDataSize       = 20000 // the total number of txs
	BatchSize           = 4000  // supervisor read a batch of txs then send them, it should be larger than inject speed
	BrokerNum           = 10
	NodesInShard        = 4
	ShardNum            = 4
	DataWrite_path      = "./result/"           // measurement data result output path
	LogWrite_path       = "./log"               // log output path
	SupervisorAddr      = "47.111.226.10:18800" //supervisor ip address
	FileInput           = `../txs_0.0cross.csv` //the raw BlockTransaction data path
)
