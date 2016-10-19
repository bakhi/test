package test

import (
	"time"

	"gopkg.in/sensorbee/sensorbee.v0/bql/udf"
	"gopkg.in/sensorbee/sensorbee.v0/core"
	"gopkg.in/sensorbee/sensorbee.v0/data"
)

type Collector struct {
	interval time.Duration
	stopped  int32
	check    int
}

func (c *Collector) Process(ctx *core.Context, tuple *core.Tuple, w core.Writer) error {
	curTime := int(time.Now().Nanosecond() / 1e+8)
	if curTime == c.check {
		return nil
	} else {
		c.check = curTime
		if err := w.Write(ctx, tuple); err != nil {
			return err
		}
	}
	//	time.Sleep(c.interval)
	return nil
}

func (c *Collector) Terminate(ctx *core.Context) error {
	return nil
}

func CreateCollector(decl udf.UDSFDeclarer, inputStream, field string, i data.Value) (udf.UDSF, error) {
	interval, err := data.ToDuration(i)
	if err != nil {
		return nil, err
	}

	// cannot understand yet
	if err := decl.Input(inputStream, nil); err != nil {
		return nil, err
	}

	return &Collector{
		interval: interval,
		check:    int(time.Now().Nanosecond() / 1e+8),
	}, nil
}
