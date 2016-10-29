package main

import (
	"fmt"
	"github.com/sachaos/todoist/lib"
	"github.com/urfave/cli"
	"os"
	"text/tabwriter"
)

func Projects(config Config, c *cli.Context) error {
	var sync lib.Sync

	sync, err := lib.LoadCache(default_cache_path)
	if err != nil {
		sync, err = Sync(config, c)
		if err != nil {
			return err
		}
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 4, 1, ' ', 0)

	for _, project := range sync.Projects {
		fmt.Fprintf(w, "%d\t%s\n", project.ID, project.Name)
	}
	w.Flush()
	return nil
}