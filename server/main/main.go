// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	"github.com/golobby/config/v3/pkg/feeder"
	coreconf "github.com/zilliztech/milvus-cdc/core/config"
	"github.com/zilliztech/milvus-cdc/server"
)

type config struct {
	Address   string
	Port      int
	Endpoints []string `yaml:"etcd.endpoints"`
	RootPath  string   `yaml:"etcd.rootpath"`
	Source    struct {
		Endpoints            []string `yaml:"etcd.endpoints"`
		RootPath             string   `yaml:"etcd.rootpath"`
		MetaPath             string   `yaml:"etcd.meta.path"`
		ReaderBufferLen      int      `yaml:"reader.buffer.len"`
		MQType               string   `yaml:"mqtype"`
		DefaultPartitionName string   `yaml:"default_partition_name"`
		Pulsar               struct {
			Address        string `yaml:"address"`
			Port           int    `yaml:"port"`
			WebAddress     string `yaml:"web.address"`
			WebPort        int    `yaml:"web.port"`
			MaxMessageSize int64  `yaml:"max.message.size"`
			Tenant         string `yaml:"tenant"`
			Namespace      string `yaml:"namespace"`
		} `yaml:"pulsar"`
		Kafka struct {
			BrokerList string `yaml:"broker_list"`
		} `yaml:"kafka"`
	} `yaml:"source"`
}

func main() {
	s := &server.CDCServer{}

	// parse config file
	conf := config{}
	f := feeder.Yaml{Path: "./configs/cdc.yaml"}
	if err := f.Feed(&conf); err != nil {
		panic(err)
	}

	// build mq configs
	var pulsarConfig coreconf.PulsarConfig
	var kafkaConfig coreconf.KafkaConfig
	if conf.Source.MQType == "pulsar" {
		pulsarConf := conf.Source.Pulsar
		pulsarConfig = coreconf.NewPulsarConfig(
			coreconf.PulsarAddressOption(fmt.Sprintf("pulsar://%s:%d", pulsarConf.Address, pulsarConf.Port)),
			coreconf.PulsarWebAddressOption(pulsarConf.WebAddress, pulsarConf.WebPort),
			coreconf.PulsarMaxMessageSizeOption(pulsarConf.MaxMessageSize),
			coreconf.PulsarTenantOption(pulsarConf.Tenant, pulsarConf.Namespace),
		)
	} else if conf.Source.MQType == "kafka" {
		kafkaConf := conf.Source.Kafka
		kafkaConfig = coreconf.NewKafkaConfig(
			coreconf.KafkaAddressOption(kafkaConf.BrokerList))
	} else {
		panic("Unknown mq type:" + conf.Source.MQType)
	}

	s.Run(&server.CDCServerConfig{
		Address: fmt.Sprintf("%s:%d", conf.Address, conf.Port),
		EtcdConfig: server.CDCEtcdConfig{
			Endpoints: conf.Endpoints,
			RootPath:  conf.RootPath,
		},
		SourceConfig: server.MilvusSourceConfig{
			EtcdAddress:          conf.Source.Endpoints,
			EtcdRootPath:         conf.Source.RootPath,
			EtcdMetaSubPath:      conf.Source.MetaPath,
			ReadChanLen:          conf.Source.ReaderBufferLen,
			DefaultPartitionName: conf.Source.DefaultPartitionName,
			Pulsar:               pulsarConfig,
			Kafka:                kafkaConfig,
		},
	})
}
