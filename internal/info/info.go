package info

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cobra"
)

const (
	indent = "    "
	b      = 1024
	mb     = 1024 * b
	gb     = 1024 * mb
)

type info struct {
	waitTime time.Duration
}

func CommandNew() *cobra.Command {
	i := new(info)
	cmd := &cobra.Command{
		Use:   "info",
		Short: "show machine information",
		Long:  "show machine information",
		RunE:  i.run,
	}

	cmd.Flags().DurationVarP(&i.waitTime, "time", "t", time.Second, "specify time to measure cpu utilization")

	return cmd
}

func (i *info) run(cmd *cobra.Command, args []string) error {
	var buf bytes.Buffer
	v, err := mem.VirtualMemory()
	if err != nil {
		return errors.Wrap(err, "Failed to get virtual memory info")
	}
	total := v.Total / gb
	free := v.Free / mb
	fmt.Fprintf(&buf,
		"Virtual memory:\n%sTotal: %dGB\n%sFree: %dMB\n%sUsed: %f%%\n%sCached: %d\n%sSwapCached: %d\n",
		indent,
		total,
		indent,
		free,
		indent,
		v.UsedPercent,
		indent,
		v.Cached,
		indent,
		v.SwapCached,
	)

	ci, err := cpu.Info()
	if err != nil {
		return errors.Wrap(err, "Failed to get cpu info")
	}
	buf.WriteString("\nCPU:\n")
	for _, c := range ci {
		fmt.Fprintf(&buf,
			"%sID: %s\n%sVendorID: %s\n%sNums: %d\n%sCores: %d\n%sCacheSize: %d\n%sModel: %s\n\n",
			indent,
			c.CoreID,
			indent,
			c.VendorID,
			indent,
			runtime.NumCPU(),
			indent,
			c.Cores,
			indent,
			c.CacheSize,
			indent,
			c.ModelName,
		)
	}

	times, err := cpu.Percent(i.waitTime, true)
	if err != nil {
		return errors.Wrap(err, "Failed to get cpu utilization")
	}
	for idx, t := range times {
		fmt.Fprintf(&buf,
			"%sCPU[%d] Utilization: %f%%\n",
			indent,
			idx,
			t,
		)
	}

	buf.Write([]byte{0x0a})

	avg, err := load.Avg()
	if err != nil {
		return errors.Wrap(err, "Failed to get load average")
	}
	fmt.Fprintf(&buf, "Load average: %1.2f, %1.2f, %1.2f\n\n", avg.Load1, avg.Load5, avg.Load15)

	if _, err := os.Stdout.Write(buf.Bytes()); err != nil {
		return errors.Wrap(err, "Failed to display info")
	}
	return nil
}
