// Copyright 2017 Xiaomi, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


package g

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/toolkits/file"
)

type HttpConfig struct {
	Enable bool   `json:"enable"`
	Listen string `json:"listen"`
}

type IndexConfig struct {
	Enable     bool              `json:"enable"`
	Dsn        string            `json:"dsn"`
	MaxIdle    int               `json:"maxIdle"`
	AutoDelete bool              `json:"autoDelete"`
	Cluster    map[string]string `json:"cluster"`
}

type CollectorConfig struct {
	Enable    bool     `json:"enable"`
	DestUrl   string   `json:"destUrl"`
	SrcUrlFmt string   `json:"srcUrlFmt"`
	Cluster   []string `json:"cluster"`
}

type PluginConfig struct {
	ApiUrlFmt      string `json:"apiUrlFmt"`
	Interval       int32  `json:"interval"`
	Concurrent     int32  `json:"concurrent"`
	ConnectTimeout int32  `json:"connectTimeout"`
	RequestTimeout int32  `json:"requestTimeout"`
}

type CleanerConfig struct {
	Interval int32 `json:"interval"`
}

type AgentConfig struct {
	Enable  bool           `json:"enable"`
	Dsn     string         `json:"dsn"`
	MaxIdle int32          `json:"maxIdle"`
	Plugin  *PluginConfig  `json:"plugin"`
	Cleaner *CleanerConfig `json:"cleaner"`
}

type GlobalConfig struct {
	Debug     bool             `json:"debug"`
	Http      *HttpConfig      `json:"http"`
	Index     *IndexConfig     `json:"index"`
	Collector *CollectorConfig `json:"collector"`
	Agent     *AgentConfig     `json:"agent"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent. maybe you need `mv cfg.example.json cfg.json`")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	configLock.Lock()
	defer configLock.Unlock()
	config = &c

	log.Println("g:ParseConfig ok, ", cfg)
}
