// Copyright 2019 The Meshery Authors
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

package cmd

import (
	// "os"
	// "os/exec"
	// log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	testURL = ""
	testName = ""
	testMesh = ""
	testFile = ""
	qps = ""
	parallelRequests = ""
	duration = ""
	loadGenerator = ""
)

/*
http://localhost:9081/api/load-test?name=test1&mesh=istio&
url=http%3A%2F%2F192.168.99.103%3A30176%2Fproductpage&qps=100&c=5&t=10&dur=s&
uuid=134883e9-4eb5-4857-b985-e7786b92a70a&loadGenerator=fortio
*/

// updateCmd represents the update command
var perfCmd = &cobra.Command{
	Use:   "perf",
	Short: "Performance testing and benchmarking",
	Long:  `Performance testing and benchmarking.`,
	Run: func(cmd *cobra.Command, args []string) {
		//Check prerequisite
		preReqCheck()
		
	},
}

func init() {
	perfCmd.Flags().StringVar(&testURL, "url", "YOUR NAME", "DESCRIPTION")
	perfCmd.Flags().StringVar(&testName, "name", "YOUR NAME", "DESCRIPTION")
	perfCmd.Flags().StringVar(&testMesh, "mesh", "YOUR NAME", "DESCRIPTION")
	perfCmd.Flags().StringVar(&testFile, "file", "YOUR NAME", "DESCRIPTION")
	perfCmd.Flags().StringVar(&qps, "qps", "YOUR NAME", "DESCRIPTION")
	perfCmd.Flags().StringVar(&parallel_requests, "parallel_requests", "YOUR NAME", "DESCRIPTION")
	perfCmd.Flags().StringVar(&duration, "duration", "YOUR NAME", "DESCRIPTION")
	perfCmd.Flags().StringVar(&load_generator, "load_generator", "YOUR NAME", "DESCRIPTION")
	rootCmd.AddCommand(perfCmd)
}