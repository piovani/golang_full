package cmd

import (
	"github.com/piovani/go_full/http/rest"
	"github.com/piovani/go_full/infra/logger"
	"github.com/spf13/cobra"
)

var (
	APIRest = &cobra.Command{
		Use:     "rest",
		Short:   "Start listen http type rest",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logger.NewLogger("")
			api := rest.NewRest()

			logger.Error(api.Execute())
		},
	}
)
