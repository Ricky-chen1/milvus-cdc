/*
 * Licensed to the LF AI & Data foundation under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 * //
 *     http://www.apache.org/licenses/LICENSE-2.0
 * //
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tag

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/zilliztech/milvus-cdc/core/log"
)

var (
	BuildTime = "unknown"
	GitCommit = "unknown"
	GoVersion = "unknown"
)

func LogInfo() {
	log.Info("base info", zap.String("BuildTime", BuildTime), zap.String("GitCommit", GitCommit), zap.String("GoVersion", GoVersion))
}

// nolint
func PrintInfo() {
	fmt.Println("# Build time:", BuildTime)
	fmt.Println("# Git commit:", GitCommit)
	fmt.Println("# Go version:", GoVersion)
}
