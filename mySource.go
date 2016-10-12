package test

import (
	"bytes"
	"math/rand"
	"time"

	"gopkg.in/sensorbee/sensorbee.v0/bql"
	"gopkg.in/sensorbee/sensorbee.v0/core"
	"gopkg.in/sensorbee/sensorbee.v0/data"
)

type Device struct {
	ID         string
	num        int
	sensorData [2]SensorData
}

type SensorData struct {
	ID    string
	value float64
}

type SourceCreator struct {
	interval time.Duration
}

func (t *SensorData) MakeData(name string, min, max int) error {
	t.ID = name
	t.value = rand.Float64() * float64(max-min)
	return nil
}

func SimulatedSensor(name string, min, max int) []byte {
	buf := bytes.NewBuffer(nil) // create empty buffer
	//	value := rand.Intn(max - min)
	value := rand.Float64() * float64((max - min))

	//buf.WriteString(name + ": " + strconv.Itoa(value))
	buf.WriteByte(byte(value))
	//	value := rand.Intn(max+min) - min
	//	tuple := map[string]int{name: value}
	//	tuple := map[string]int{"name": name, "value": value}
	//	tupleM, _ := json.Marshal(tuple)
	//	return string(tupleM)v
	return buf.Bytes()
}

func (d *Device) MakeDevice(ID string) {
	d.ID = ID
	//	d.sensorData[0].MakeData("temp", 0, 30)
	//	d.sensorData[1].MakeData("humid", 70, 100)
}

func (s *SourceCreator) GenerateStream(ctx *core.Context, w core.Writer) error {
	device := new(Device)
	devName := []string{"dev1", "dev2", "dev3", "dev4", "dev5"}
	devProb := []float64{0.4, 0.3, 0.15, 0.1, 0.05}
	pickDev := func() string {
		r := rand.Float64()
		for i, p := range devProb {
			if r < p {
				return devName[i]
			}
			r -= p
		}
		return devName[len(devName)-1]
	}

	//	device.MakeDevice(pickDev())
	device.num = 0
	temp := &device.sensorData[0]
	humid := &device.sensorData[1]

	for {
		device.ID = pickDev()
		device.num += 1
		temp.MakeData("temp", 0, 30)
		humid.MakeData("humid", 0, 100)

		t := core.NewTuple(data.Map{
			"deviceID": data.String(device.ID),
			"num":      data.Int(device.num),
			"time":     data.Int(time.Now().Second()),
			temp.ID:    data.Float(float64(temp.value)),
			humid.ID:   data.Float(float64(humid.value)),
		})
		if err := w.Write(ctx, t); err != nil {
			return err
		}
		time.Sleep(s.interval)
	}
}

func CreateMySource(ctx *core.Context, ioParams *bql.IOParams, params data.Map) (core.Source, error) {
	interval := 1 * time.Nanosecond
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
