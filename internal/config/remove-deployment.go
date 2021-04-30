/*
 * Flow CLI
 *
 * Copyright 2019-2021 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/onflow/flow-cli/internal/command"
	"github.com/onflow/flow-cli/pkg/flowcli/output"
	"github.com/onflow/flow-cli/pkg/flowcli/project"
	"github.com/onflow/flow-cli/pkg/flowcli/services"
)

type flagsRemoveDeployment struct{}

var removeDeploymentFlags = flagsRemoveDeployment{}

var RemoveDeploymentCommand = &command.Command{
	Cmd: &cobra.Command{
		Use:     "deployment <account> <network>",
		Short:   "Remove deployment from configuration",
		Example: "flow config remove deployment Foo testnet",
		Args:    cobra.MaximumNArgs(2),
	},
	Flags: &removeDeploymentFlags,
	Run: func(
		cmd *cobra.Command,
		args []string,
		globalFlags command.GlobalFlags,
		services *services.Services,
	) (command.Result, error) {
		p, err := project.Load(globalFlags.ConfigPaths)
		if err != nil {
			return nil, fmt.Errorf("configuration does not exists")
		}

		account := ""
		network := ""
		if len(args) == 2 {
			account = args[0]
			network = args[1]
		} else {
			account, network = output.RemoveDeploymentPrompt(p.Config().Deployments)
		}

		err = p.Config().Deployments.Remove(account, network)
		if err != nil {
			return nil, err
		}

		err = p.SaveDefault()
		if err != nil {
			return nil, err
		}

		return &ConfigResult{
			result: "deployment removed",
		}, nil
	},
}

func init() {
	RemoveDeploymentCommand.AddToParent(RemoveCmd)
}