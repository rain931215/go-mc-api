package mcfalloutcmd

import (
	"github.com/rain931215/go-mc-api/api"
	"github.com/spf13/viper"
)

func addAdmin(c *api.Client, Text string) {
	file := viper.New()
	file.SetConfigName("whiteList")
	file.SetConfigType("yaml")
	file.AddConfigPath("./plugin/mcfalloutcmd")
	file.Set("admin", append(file.GetStringSlice("admin"), Text))
	file.WriteConfig()
}

func say(c *api.Client, Text string) {
	c.Chat(Text)
}
