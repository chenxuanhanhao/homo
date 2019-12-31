//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, August 2019
//

package cmd

import (
	"github.com/countstarlight/homo/logger"
	"github.com/urfave/cli/v2"
	"os"
)

// Execute execute
func Execute() {
	app := &cli.App{
		Name:    AppName,
		Version: Version,
		Usage:   "Expand the combination of artificial intelligence applications and the IoT",
		Flags:   flags,
		Commands: []*cli.Command{
			{
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "start homo",
				Action:  startInternal,
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "show the version of homo",
				Action:  version,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		logger.S.Fatal(err)
	}
}