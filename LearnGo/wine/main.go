package main

import "fmt"

type Monster struct {
	Name string
}

func NewMonster() *Monster {
	return &Monster{Name: "Goblin"}
}

type Player struct {
	Name string
}

func NewPlayer(name string) *Player {
	return &Player{Name: name}
}

type Mission struct {
	Player  *Player
	Monster *Monster
}

func NewMission(player *Player, monster *Monster) *Mission {
	return &Mission{Player: player, Monster: monster}
}

func (m *Mission) Start() {
	fmt.Printf("Mission started: %s vs %s\n", m.Player.Name, m.Monster.Name)
}

func main() {
	mission := InitMission("Hero")
	mission.Start()
}
