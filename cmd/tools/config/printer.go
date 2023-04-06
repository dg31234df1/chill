package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/milvus-io/milvus/pkg/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func ShowYaml(filepath string) {
	reader := viper.New()
	reader.SetConfigFile(filepath)
	if err := reader.ReadInConfig(); err != nil {
		log.Warn("read config failed", zap.Error(err))
		os.Exit(-3)
	}
	keys := reader.AllKeys()
	sort.Strings(keys)
	for _, key := range keys {
		v := reader.GetString(key)
		fmt.Fprintln(os.Stdout, key, "=", v)
	}
}
