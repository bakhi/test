package test

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"

	"gopkg.in/sensorbee/sensorbee.v0/bql"
	"gopkg.in/sensorbee/sensorbee.v0/core"
	"gopkg.in/sensorbee/sensorbee.v0/data"
)

type SensorData struct {
	ID    string
	value int
}

type SourceCreator struct {
	interval time.Duration
}

func (t *SensorData) MakeData(name string, min, max int) error {
	t.ID = name
	t.value = rand.Intn(max - min)
	return nil
}

func SimulatedSensor(name string, min, max int) []byte {
	buf := bytes.NewBuffer(nil) // create empty buffer
	value := rand.Intn(max - min)
	buf.WriteString(name + ": " + strconv.Itoa(value))
	//	value := rand.Intn(max+min) - min
	//	tuple := map[string]int{name: value}
	//	tuple := map[string]int{"name": name, "value": value}
	//	tupleM, _ := json.Marshal(tuple)
	//	return string(tupleM)v
	return buf.Bytes()
}

func (s *SourceCreator) GenerateStream(ctx *core.Context, w core.Writer) error {
	tuple := new(SensorData)
	for {
		tuple.MakeData("temperature", 0, 30)
		t := core.NewTuple(data.Map{
			tuple.ID: data.Int(tuple.value),
		})

		if err := w.Write(ctx, t); err != nil {
			return err
		}
		//		fmt.Println(tuple)
		time.Sleep(s.interval)
	}
}

func CreateMySource(ctx *core.Context, ioParams *bql.IOParams, params data.Map) (core.Source, error) {
	interval := 1 * time.Second
	if v, ok := params["interval"]; ok {
		i, err := data.ToDuration(v)
		if err != nil {
			return nil, err
		}
		interval = i
	}
	return core.ImplementSourceStop(&SourceCreator{
		interval: interval,
	}), nil

}

func (s *SourceCreator) Stop(ctx *core.Context) error {
	return nil
}
