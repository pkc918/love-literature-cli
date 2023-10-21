/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "爱情在文学家笔下的样子",
	Long:  `爱情在文学家笔下的样子！`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		file, _ := exec.LookPath(os.Args[0])
		path, _ := filepath.Abs(file)
		index := strings.LastIndex(path, string(os.PathSeparator))
		path = path[:index]
		fmt.Println(fmt.Sprintf("%s%s", path, "/cmd/libirary.txt"))
		dat, err := os.ReadFile(fmt.Sprintf("%s%s", path, "/cmd/libirary.txt"))
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dat))
		for i := 0; i < len(dat); i++ {
			if dat[i] == 10 {
				fmt.Println(dat[i])
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
