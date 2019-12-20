package cmd

import (
	"github.com/pkg/errors"
	"os"
	"strings"
)

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

func makeDirs(dirs ...string) error {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0700); err != nil {
			return errors.Wrapf(err, "can't make directory %s", dir)
		}
	}
	return nil
}
