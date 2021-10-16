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
	"github.com/spf13/cobra"
)

var (
	filein, fileout, password string
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Use to encrypt file",
	Long: `encrypt is used to encrypt file.
	Make sure not to use same file name for input and output file, 
	your original file and be overridden`,
	Run: func(cmd *cobra.Command, args []string) {
		encryptFile(filein, password, fileout)
	},
}

// -------- -------- -------- -------- -------- -------- -------- -------- -------- --------
func init() {
	rootCmd.AddCommand(encryptCmd)

	// Here you will define your flags and configuration settings.

	encryptCmd.Flags().StringVar(&filein, "in", "", "Path of file, to be encrypted")
	encryptCmd.Flags().StringVar(&fileout, "out", "", "Path of encrypted file")
	encryptCmd.Flags().StringVar(&password, "password", "", "Password for encryption")

	encryptCmd.MarkFlagRequired("in")
	encryptCmd.MarkFlagRequired("out")
	encryptCmd.MarkFlagRequired("password")
}
