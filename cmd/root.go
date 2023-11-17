package cmd

import (
	"github.com/piovani/go_full/infra/config"
	"github.com/piovani/go_full/infra/logger"
)

func (c *Cmd) Execute() {
	c.errCheck(c.initConfig())
	c.addCommands()
	c.errCheck(c.cobra.Execute())
}

func (c *Cmd) initConfig() error {
	return config.InitConfig()
}

func (c *Cmd) addCommands() {
	c.cobra.AddCommand(
		Migrate,
		APIRest,
	)
}

func (c *Cmd) errCheck(err error) {
	if err != nil {
		logger.NewLogger("structure").Error(err)
	}
}
