/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// divCmd represents the div command
var divCmd = &cobra.Command{
	Use:   "div",
	Short: "Divide two integers",
	Long: `Divide two integers.`,
	Run: func(cmd *cobra.Command, args []string) {
		n1,err1 := strconv.Atoi(args[0])
		n2,err2 :=  strconv.Atoi(args[1])
		if err1 != nil || err2 != nil{
			fmt.Println("Error:Put valid numbers")
		}
		if n2 == 0{
			fmt.Println("Error: Cannot divide by zero")
		}
		fmt.Printf("%d / %d = %d\n",n1,n2,n1/n2)

	},
}

func init() {
	rootCmd.AddCommand(divCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// divCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// divCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
