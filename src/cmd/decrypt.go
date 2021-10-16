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

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Use to decrypt file",
	Long: `decrypt is used to decrypt file.
	Make sure not to use same file name for input and output file, 
	your original file and be overridden`,
	Run: func(cmd *cobra.Command, args []string) {
		decryptFile(filein, password, fileout)
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	// Here you will define your flags and configuration settings.

	decryptCmd.PersistentFlags().StringVar(&filein, "in", "", "Path of file, to decrypt")
	decryptCmd.PersistentFlags().StringVar(&fileout, "out", "", "Path for decrypted file")
	decryptCmd.Flags().StringVar(&password, "password", "", "Password for decryption")

	decryptCmd.MarkFlagRequired("in")
	decryptCmd.MarkFlagRequired("out")
	decryptCmd.MarkFlagRequired("password")
}
