package tetris

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
)

//types

type GameEngine struct {
	updates []IUpdateAble
	isStop  bool
	input   keyboard.Key
}

type IUpdateAble interface {
	Update(engine *GameEngine)
}

// functions

func (engine *GameEngine) Update() {
	for _, iupdateable := range engine.updates {
		iupdateable.Update(engine)
	}
}

func (engine *GameEngine) GetKey() {
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}
		engine.input = event.Key
		// fmt.Printf("You pressed: rune %q, key %X\r\n", event.Rune, event.Key)
		if event.Key == keyboard.KeyEsc {
			engine.isStop = true
			break
		}
	}
}

func StartGame() {

	// engine init
	engine := GameEngine{}

	// add modules, for game engine
	blockCtrl := New(10, 20, &engine)

	// start engine's getkey
	go engine.GetKey()

	for {
		engine.Update()
		if engine.input != keyboard.KeyBackspace {
			engine.input = keyboard.KeyBackspace
		}

		uiString = blockCtrl.Draw()
		fmt.Print("\033[H\033[2J") // clear fmt
		fmt.Println(uiString)      // show ui string

		//tick
		time.Sleep(time.Second * 1)

		if engine.isStop {
			break
		}
	}
}

//values

var uiString string
