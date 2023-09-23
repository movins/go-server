package main

import (
	"context"
)

// MonitorImp servant implementation
type MonitorImp struct {
}

// Init servant init
func (imp *MonitorImp) Init() error {
	//initialize servant here:
	//...
	return nil
}

// Destroy servant destroy
func (imp *MonitorImp) Destroy() {
	//destroy servant here:
	//...
}

func (imp *MonitorImp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
func (imp *MonitorImp) Sub(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
