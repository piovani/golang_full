package cmd

import (
	"log"

	"github.com/piovani/go_full/infra/config"
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
		APIRest,
	)
}

func (c *Cmd) errCheck(err error) {
	if err != nil {
		log.Fatalln("DEU RUIM: ", err)
	}
}
