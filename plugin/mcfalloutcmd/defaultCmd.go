package mcfalloutcmd

import (
	"github.com/rain931215/go-mc-api/api"
	"github.com/spf13/viper"
)

func (p *McfalloutCmd) loadDefaultCmd() {
	p.AddCmd("say", say)
	p.AddCmd("addadmin", addAdmin)
	p.AddCmd("respawn", respawn)
	p.AddCmd("throwall", throwall)
}

func addAdmin(c *api.Client, sender string, text string, args []string) {
	file := viper.New()
	file.SetConfigName("whiteList")
	file.SetConfigType("yaml")
	file.AddConfigPath("./plugin/mcfalloutcmd")
	file.ReadInConfig()
	file.Set("admin", append(file.GetStringSlice("admin"), args[0]))
	file.WriteConfig()
}
func throwall(c *api.Client, sender string, text string, args []string) {
	for slot := 9; slot < 46; slot++ {
		c.ClickWindow(0, int16(slot), 1, 4)
	}
}
func say(c *api.Client, sender string, text string, args []string) {
	c.Chat(text)
}
func respawn(c *api.Client, sender string, text string, args []string) {
	c.ReSpawn()
}
