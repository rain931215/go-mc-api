package mcfalloutcmd

import (
	"github.com/rain931215/go-mc-api/api"
	"github.com/spf13/viper"
)

func addAdmin(c *api.Client, Text string) {
	viper.Set("admin", append(whiteList, Text))
	viper.WriteConfig()
}
func say(c *api.Client, Text string) {
	c.Chat(Text)
}
