package models

import (
	"encoding/json"
	"fmt"
)

type Bag struct {
	Model

	Title   string `validate:"required,max=255"`
	Volume  uint   `validate:"gt=0"`
	Disable bool

	Cuboids []Cuboid
}

func (b *Bag) PayloadVolume() uint {
	var v uint
	for _, c := range b.Cuboids {
		v += c.PayloadVolume()
	}
	return v
}

func (b *Bag) AvailableVolume() uint {
	return b.Volume - b.PayloadVolume()
}

func (b *Bag) SetDisabled(value bool) {
	b.Disable = value
}

func (b *Bag) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		ID              uint     `json:"id"`
		Title           string   `json:"title"`
		Volume          uint     `json:"volume"`
		PayloadVolume   uint     `json:"payloadVolume"`
		AvailableVolume uint     `json:"availableVolume"`
		Disable         bool     `json:"disable"`
		Cuboids         []Cuboid `json:"cuboids"`
	}{
		b.ID, b.Title, b.Volume, b.PayloadVolume(), b.AvailableVolume(), b.Disable, b.Cuboids,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Bag. %w", err)
	}

	return j, nil
}
