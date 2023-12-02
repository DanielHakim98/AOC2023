/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/DanielHakim98/aoc/day1"
	"github.com/spf13/cobra"
)

// dayOneCmd represents the dayOne command
var dayOneCmd = &cobra.Command{
	Use:   "dayOne",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("No filename is passed in")
		}
		fmt.Println("sum of calibration values: ", day1.PartTwo(args[0], day1.GetInput))
	},
}

func init() {
	rootCmd.AddCommand(dayOneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dayOneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dayOneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
