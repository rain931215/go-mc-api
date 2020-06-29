package mcfalloutcmd

import (
	"github.com/rain931215/go-mc-api/api"
	"github.com/spf13/viper"
)

func addAdmin(c *api.Client, sender string, text string, args []string) {
	file := viper.New()
	file.SetConfigName("whiteList")
	file.SetConfigType("yaml")
	file.AddConfigPath("./plugin/mcfalloutcmd")
	file.ReadInConfig()
	file.Set("admin", append(file.GetStringSlice("admin"), args[0]))
	file.WriteConfig()
}

func say(c *api.Client, sender string, text string, args []string) {
	c.Chat(text)
}
