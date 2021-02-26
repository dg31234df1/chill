package indexnode

import (
	"bytes"
	"log"
	"strconv"
	"sync"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb2"
	"github.com/zilliztech/milvus-distributed/internal/util/paramtable"
)

const (
	StartParamsKey = "START_PARAMS"
)

type ParamTable struct {
	paramtable.BaseTable

	IP      string
	Address string
	Port    int

	NodeID int64

	MasterAddress string

	MinIOAddress         string
	MinIOAccessKeyID     string
	MinIOSecretAccessKey string
	MinIOUseSSL          bool
	MinioBucketName      string
}

var Params ParamTable
var once sync.Once

func (pt *ParamTable) Init() {
	once.Do(func() {
		pt.BaseTable.Init()
		pt.initParams()
	})
}

func (pt *ParamTable) initParams() {
	pt.initMinIOAddress()
	pt.initMinIOAccessKeyID()
	pt.initMinIOSecretAccessKey()
	pt.initMinIOUseSSL()
	pt.initMinioBucketName()
}

func (pt *ParamTable) LoadConfigFromInitParams(initParams *internalpb2.InitParams) error {
	pt.NodeID = initParams.NodeID

	config := viper.New()
	config.SetConfigType("yaml")
	for _, pair := range initParams.StartParams {
		if pair.Key == StartParamsKey {
			err := config.ReadConfig(bytes.NewBuffer([]byte(pair.Value)))
			if err != nil {
				return err
			}
			break
		}
	}

	for _, key := range config.AllKeys() {
		val := config.Get(key)
		str, err := cast.ToStringE(val)
		if err != nil {
			switch val := val.(type) {
			case []interface{}:
				str = str[:0]
				for _, v := range val {
					ss, err := cast.ToStringE(v)
					if err != nil {
						log.Panic(err)
					}
					if len(str) == 0 {
						str = ss
					} else {
						str = str + "," + ss
					}
				}

			default:
				log.Panicf("undefine config type, key=%s", key)
			}
		}
		err = pt.Save(key, str)
		if err != nil {
			panic(err)
		}

	}

	pt.initParams()
	return nil
}

func (pt *ParamTable) initMinIOAddress() {
	ret, err := pt.Load("_MinioAddress")
	if err != nil {
		panic(err)
	}
	pt.MinIOAddress = ret
}

func (pt *ParamTable) initMinIOAccessKeyID() {
	ret, err := pt.Load("minio.accessKeyID")
	if err != nil {
		panic(err)
	}
	pt.MinIOAccessKeyID = ret
}

func (pt *ParamTable) initMinIOSecretAccessKey() {
	ret, err := pt.Load("minio.secretAccessKey")
	if err != nil {
		panic(err)
	}
	pt.MinIOSecretAccessKey = ret
}

func (pt *ParamTable) initMinIOUseSSL() {
	ret, err := pt.Load("minio.useSSL")
	if err != nil {
		panic(err)
	}
	pt.MinIOUseSSL, err = strconv.ParseBool(ret)
	if err != nil {
		panic(err)
	}
}

func (pt *ParamTable) initMinioBucketName() {
	bucketName, err := pt.Load("minio.bucketName")
	if err != nil {
		panic(err)
	}
	pt.MinioBucketName = bucketName
}
