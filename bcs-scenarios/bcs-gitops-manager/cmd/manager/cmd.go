/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package manager xxx
package manager

import (
	"fmt"
	"os"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/spf13/cobra"

	"github.com/Tencent/bk-bcs/bcs-scenarios/bcs-gitops-manager/cmd/manager/options"
)

var (
	commandName = "bcs-gitops-manager"
	configFile  = "./bcs-gitops-manager.json"
)

// NewCommand create command line for gitops manager
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   commandName,
		Short: "bcs-gitops-manager is a bcs-service for gitops",
		Long:  `bcs-gitops-manager integrates gitops sulotion into bcs-services.`,
		RunE: func(cmd *cobra.Command, args []string) error { // nolint
			if err := options.Parse(configFile); err != nil {
				fmt.Printf("option parse failed: %s\n", err.Error())
				return err
			}

			op := options.GlobalOptions()
			blog.InitLogs(op.LogConfig)
			defer blog.CloseLogs()
			// config option verification
			if err := op.Complete(); err != nil {
				fmt.Fprintf(os.Stderr, "server option complete failed, %s\n", err.Error())
				return err
			}
			if err := op.Validate(); err != nil {
				fmt.Fprintf(os.Stderr, "server option validate failed, %s\n", err.Error())
				return err
			}

			// run server
			serv := NewServer(op)
			if err := serv.Init(); err != nil {
				fmt.Fprintf(os.Stderr, "server init failed, %s", err.Error())
				return err
			}
			return serv.Run()
		},
	}
	// setting server configuration flag
	cmd.Flags().StringVarP(&configFile, "config", "c", configFile, "manager configuration json file")
	return cmd
}
