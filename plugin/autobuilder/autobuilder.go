package autobuilder

import (
	"github.com/rain931215/go-mc-api/api"
	"github.com/rain931215/go-mc-api/plugin/navigate"
)

//AutoBuilder _
type AutoBuilder struct {
	c        *api.Client
	navigate *navigate.Navigate
}

//New _
func New(c *api.Client) *AutoBuilder {
	p := new(AutoBuilder)
	p.c = c
	p.navigate = navigate.New(c)
	return p
}
