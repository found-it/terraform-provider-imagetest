package inventory

import (
	"context"
)

type Inventory interface {
	Create(context.Context) error
	Open(context.Context) error
	AddHarness(context.Context, Harness) (bool, error)
	AddFeature(context.Context, Feature) (bool, error)
	GetFeatures(context.Context, Harness) ([]Feature, error)
	RemoveHarness(context.Context, Harness) error
	RemoveFeature(context.Context, Feature) ([]Feature, error)
}

type Harness string

type Feature struct {
	Id      string            `json:"id"`
	Labels  map[string]string `json:"labels"`
	Harness Harness           `json:"harness"`
}

type InventoryModel map[Harness][]Feature
