package cmd

import "strings"

type CommonOptionsCommander interface {
	SetCommon(commonOpts CommonOpts)
	Execute(args []string) error
}

type CommonOpts struct {
	CardinalURL string
	Revision    string
}

func (c *CommonOpts) SetCommon(commonOpts CommonOpts) {
	c.CardinalURL = strings.TrimSuffix(commonOpts.CardinalURL, "/")
	c.Revision = commonOpts.Revision
}
