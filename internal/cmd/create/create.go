package create

import (
	"fmt"
	"log"
	"time"

	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
	"github.com/urfave/cli/v2"
	"google.golang.org/api/sqladmin/v1"
)

func Create(c *cli.Context) error {
	ctx := c.Context
	sqladminService, err := sqladmin.NewService(ctx)
	if err != nil {
		return err
	}

	project := c.String("project")
	instance := c.String("instance")
	database := c.String("database")
	bucket := c.String("bucket")
	filename := fmt.Sprintf("Cloud_SQL_Export_%s_%s_daily.sql", database, time.Now().Format("2006-01-02"))

	req := sqladminService.Instances.Export(project, instance, &sqladmin.InstancesExportRequest{
		ExportContext: &sqladmin.ExportContext{
			Databases: []string{database},
			FileType:  "SQL",
			Offload:   false,
			Uri:       fmt.Sprintf("gs://%s/%s", bucket, filename),
		},
	})

	op, err := req.Context(ctx).Do()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(1000 * time.Millisecond)
	done := make(chan bool)
	bar := progressbar.NewOptions(100,
		progressbar.OptionSetRenderBlankState(false),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionSetDescription(fmt.Sprintf("[cyan][1/3][reset] Exporting '%s'...", database)),
	)

	go func() {
		for {
			select {
			case <-done:
				ticker.Stop()
				return
			case <-ticker.C:
				bar.Add(1)
				op, err = sqladminService.Operations.Get(project, op.Name).Context(ctx).Do()
				if err != nil {
					log.Println(err)
					done <- true
				}
				if op.Status == "DONE" {
					bar.Set(100)
					bar.Finish()
					done <- true
				}
			}
		}
	}()

	<-done

	return nil
}
