//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitMission(name string) *Mission {
	wire.Build(
		NewPlayer,
		NewMonster,
		NewMission,
	)
	return &Mission{}
}
