package params

var (
	Block_Interval      = 5000  // generate new block interval
	MaxBlockSize_global = 1200  // the block contains the maximum number of transactions
	InjectSpeed         = 2400  // the transaction inject speed
	TotalDataSize       = 48000 // the total number of txs
	BatchSize           = 4800  // supervisor read a batch of txs then send them, it should be larger than inject speed
	BrokerNum           = 10
	NodesInShard        = 4
	ShardNum            = 4
	DataWrite_path      = "./result/"           // measurement data result output path
	LogWrite_path       = "./log"               // log output path
	SupervisorAddr      = "127.0.0.1:18800"     //supervisor ip address
	FileInput           = `../txs_0.0cross.csv` //the raw BlockTransaction data path
)
