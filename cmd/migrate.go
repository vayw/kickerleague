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
	"github.com/spf13/cobra"
	"github.com/vayw/kickerleague/database"
	"github.com/vayw/kickerleague/migrations"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		database.ConnectDB()
		defer database.DBCon.Close()
		if err := migrations.Migrate(database.DBCon); err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
