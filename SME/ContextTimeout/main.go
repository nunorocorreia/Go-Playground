package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type contextKey string

var UserIdKey contextKey = "userID"

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 1
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 1
	e.battery = -1
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery += -1
	return nil
}

// process Truck handles the loading and unloading of a truck
func processTruck(ctx context.Context, truck Truck) error {
	fmt.Printf("Processing truck %+v \n", truck)

	//example access the userId
	//userID := ctx.Value(UserIdKey)
	//log.Println(userID)

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	//simulate a long running process
	delay := time.Second * 1
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(delay):
		break
	}

	err := truck.LoadCargo()
	if err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	err = truck.UnloadCargo()
	if err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}
	fmt.Printf("Finished processing truck %+v \n", truck)

	return nil
}

// processFleet demonstrates concurrent processing of multiple trucks
func processFleet(ctx context.Context, trucks []Truck) error {
	var wg sync.WaitGroup

	for _, t := range trucks {
		wg.Add(1) //specify how many goroutines we have, we added 5

		go func(t Truck) { //5 will de executed
			if err := processTruck(ctx, t); err != nil {
				log.Println(err)
			}
			wg.Done()
		}(t)
	}

	wg.Wait() //wait for them to finish

	return nil
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, UserIdKey, 42)

	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
	}

	//Process all trucks concurrently
	if err := processFleet(ctx, fleet); err != nil {
		fmt.Printf("Error processing fleet: %v\n", err)
		return
	}

	fmt.Printf("All trucks processed successfully!")

}
