package tetris

import (
	"fmt"
)

type GameEngine struct {
	updates []IUpdateAble
}

type IUpdateAble interface {
	Update()
}

func (engine *GameEngine) Update() {
	for _, iupdateable := range engine.updates {
		iupdateable.Update()
	}
}

var uiString string

func StartGame() {
	// engine init
	engine := GameEngine{}

	// add modules, for game engine
	blockCtrl := New(10, 20, &engine)

	for {
		engine.Update()
		uiString = blockCtrl.Draw()
		fmt.Print("\033[H\033[2J")
		fmt.Println(uiString)
		//tick
		//time.Sleep(time.Second * 1)
	}
}
