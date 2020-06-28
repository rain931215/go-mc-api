package mcfalloutcmd

import (
	"fmt"
	"strings"

	"github.com/Tnze/go-mc/chat"
	"github.com/fsnotify/fsnotify"
	"github.com/rain931215/go-mc-api/api"
	"github.com/spf13/viper"
)

//Func is the type of command's method
type Func = func(c *api.Client, Text string)

//Command is contained command's name and command's method
type Command struct {
	name   string
	method Func
}

//McfalloutCmd _
type McfalloutCmd struct {
	client    *api.Client
	whiteList []string
	cmdList   []*Command
}

// New _
func New(c *api.Client) *McfalloutCmd {
	p := new(McfalloutCmd)
	p.client = c
	c.Event.AddEventHandler(p.main, "chat")

	file := viper.New()
	//Load whiteList
	file.SetConfigName("whiteList")
	file.SetConfigType("yaml")
	file.AddConfigPath("./plugin/mcfalloutcmd")
	file.WatchConfig()
	err := file.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}
	p.whiteList = file.GetStringSlice("admin")
	//熱插拔
	file.OnConfigChange(func(e fsnotify.Event) {
		p.whiteList = file.GetStringSlice("admin")
		fmt.Println("White List Change")
	})
	//Load defaultCommand
	p.AddCmd("say", say)
	p.AddCmd("addadmin", addAdmin)
	return p
}

func (p *McfalloutCmd) main(msg chat.Message) (bool, error) {
	var text = msg.ClearString()
	for id := 0; id < len(p.whiteList); id++ {
		if strings.Index(text, "[收到私訊 "+p.whiteList[id]) == 0 {
			text = strings.TrimPrefix(text, "[收到私訊 "+p.whiteList[id]+"] : ")
			for i := 0; i < len(p.cmdList); i++ {
				if strings.Index(text, p.cmdList[i].name) == 0 {
					text = strings.TrimPrefix(text, p.cmdList[i].name+" ")
					p.cmdList[i].method(p.client, text)
					return false, nil
				}
			}
			return false, nil
		}
	}
	return false, nil
}

// AddCmd _
func (p *McfalloutCmd) AddCmd(name string, command Func) {
	newCommand := new(Command)
	newCommand.name = name
	newCommand.method = command
	p.cmdList = append(p.cmdList, newCommand)
}
