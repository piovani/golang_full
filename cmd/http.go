package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	APIRest = &cobra.Command{
		Use:     "rest",
		Short:   "Start listen http type rest",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("AUQI")
		},
	}
)
