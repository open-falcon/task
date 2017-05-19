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


package proc

import (
	nproc "github.com/toolkits/proc"
	"log"
)

// 索引更新
var (
	IndexUpdateCnt      = nproc.NewSCounterQps("IndexUpdateCnt")
	IndexUpdateErrorCnt = nproc.NewSCounterQps("IndexUpdateErrorCnt")
	IndexDeleteCnt      = nproc.NewSCounterQps("IndexDeleteCnt")
)

// 监控数据采集
var (
	CollectorCronCnt = nproc.NewSCounterQps("CollectorCronCnt")
)

func Start() {
	log.Println("proc.Start ok")
}

func GetAll() []interface{} {
	ret := make([]interface{}, 0)

	// index
	ret = append(ret, IndexUpdateCnt.Get())
	ret = append(ret, IndexUpdateErrorCnt.Get())
	ret = append(ret, IndexDeleteCnt.Get())

	// collector
	ret = append(ret, CollectorCronCnt.Get())

	return ret
}
