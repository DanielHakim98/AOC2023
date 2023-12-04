/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/DanielHakim98/aoc/day4"
	"github.com/DanielHakim98/aoc/utils"
	"github.com/spf13/cobra"
)

// dayFourCmd represents the dayFour command
var dayFourCmd = &cobra.Command{
	Use:   "dayFour",
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
		d := day4.Day4{}
		switch part {
		case 1:
			fmt.Println("sum of scratch points: ", d.PartOne(args[0], utils.GetInput))
		case 2:
			fmt.Println("sum of scratch points:")
		default:
			log.Fatal("Invalid 'part' flag")
		}
	},
}

func init() {
	rootCmd.AddCommand(dayFourCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dayFourCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dayFourCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
