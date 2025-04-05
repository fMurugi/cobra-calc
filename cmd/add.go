/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/fatih/color"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add two integers",
	Long: `Add two integers`,
	Run: func(cmd *cobra.Command, args []string) {
		n1,err1 := strconv.Atoi(args[0])
		n2,err2 := strconv.Atoi(args[1])
		if err1 != nil || err2 != nil{
			color.Red("Please provide two valid integers")
			return
		}
		color.Green("%d + %d = %d\n", n1,n2,n1+n2)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
