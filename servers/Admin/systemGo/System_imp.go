package main

import (
	"context"
)

// SystemImp servant implementation
type SystemImp struct {
}

// Init servant init
func (imp *SystemImp) Init() error {
	//initialize servant here:
	//...
	return nil
}

// Destroy servant destroy
func (imp *SystemImp) Destroy() {
	//destroy servant here:
	//...
}

func (imp *SystemImp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
func (imp *SystemImp) Sub(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
