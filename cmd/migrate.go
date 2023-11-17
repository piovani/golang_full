package cmd

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/piovani/go_full/infra/database"
	"github.com/piovani/go_full/infra/logger"
	"github.com/spf13/cobra"
)

var (
	Migrate = &cobra.Command{
		Use:     "migrate",
		Short:   "Run migrate in database",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logger.NewLogger("")
			db := database.NewDatabase()

			if err := db.Migrate(); err != nil {
				logger.Error(err)
			}

			fmt.Println("Migrate executed successfully")
		},
	}
)
