package main

import (
	"context"
)

// ToolsImp servant implementation
type ToolsImp struct {
}

// Init servant init
func (imp *ToolsImp) Init() error {
	//initialize servant here:
	//...
	return nil
}

// Destroy servant destroy
func (imp *ToolsImp) Destroy() {
	//destroy servant here:
	//...
}

func (imp *ToolsImp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
func (imp *ToolsImp) Sub(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
