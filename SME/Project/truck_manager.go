package main

import (
	"errors"
)

var ErrTruckNotFound = errors.New("truck not found")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (*Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type truckManager struct {
	trucks map[string]*Truck
}

func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}

func (m *truckManager) AddTruck(id string, cargo int) error {
	m.trucks[id] = &Truck{ID: id, Cargo: cargo}
	return nil
}

func (m *truckManager) GetTruck(id string) (*Truck, error) {
	truck, ok := m.trucks[id]
	if !ok {
		return nil, ErrTruckNotFound
	}

	return truck, nil
}

func (m *truckManager) RemoveTruck(id string) error {
	_, ok := m.trucks[id]
	if !ok {
		return ErrTruckNotFound
	}
	delete(m.trucks, id)
	return nil
}

func (m *truckManager) UpdateTruckCargo(id string, cargo int) error {
	_, ok := m.trucks[id]
	if !ok {
		return ErrTruckNotFound
	}
	m.trucks[id].Cargo = cargo
	return nil

}
