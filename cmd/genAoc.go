/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/DanielHakim98/aoc/utils/scaffold"
	"github.com/spf13/cobra"
)

var day int

// genAocCmd represents the genAoc command
var genAocCmd = &cobra.Command{
	Use:   "genAoc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if day <= 0 {
			log.Fatal("invalid 'day'")
		}

		if day < 6 {
			log.Fatal("Day 1, 2, 3, 4, 5 are aleady generated")
		}

		fmt.Println("day:", day)
		s := scaffold.New(false)
		if genErr := s.Generate(".", day); genErr != nil {
			log.Fatal("Failed to generate files")
		}

		fmt.Println("File generation completed. It's time to Go!")
	},
}

func init() {
	rootCmd.AddCommand(genAocCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genAocCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genAocCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	genAocCmd.Flags().IntVarP(&day, "day", "d", 0, "day to generate")
}
