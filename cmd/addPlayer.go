// Copyright © 2019 Ivan Kirillov vayw@botans.org
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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vayw/kickerleague/player"
)

// addPlayerCmd represents the addPlayer command
var addPlayerCmd = &cobra.Command{
	Use:   "addPlayer <name>",
	Short: "add player",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		player, err := player.AddPlayer(args[0])
		if err == nil {
			fmt.Printf("success! Player ID is %d\n", player.ID)
		} else {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addPlayerCmd)
}
