/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/abibby/backup/backend"
	"github.com/abibby/backup/database"
	"github.com/abibby/backup/reconcile"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// reconcileCmd represents the reconcile command
var reconcileCmd = &cobra.Command{
	Use:   "reconcile",
	Short: "Update the local database to match the remote storage backend",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := database.Open(viper.GetString("database"))
		if err != nil {
			return errors.Wrap(err, "failed to initialize database")
		}

		backends, err := getBackends()
		if err != nil {
			return err
		}
		for _, b := range backends {
			if b, ok := b.(backend.Closer); ok {
				defer b.Close()
			}
			err = reconcile.Reconcile(db, b)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(reconcileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reconcileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reconcileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
