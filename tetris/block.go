package tetris

import "github.com/eiannone/keyboard"

// Block ... have position, atom of game blocks
type Block struct {
	x, y      int
	isBlocked bool
}

// BlockController controll all of blocks in game
type BlockController struct {
	blocks map[int]map[int]*Block
	asdasd bool

	mapSizeX, mapSizeY int

	currentBlock []*Block
}

// New is Create BlockCtrl
func New(mapSizeX, mapSizeY int, engine *GameEngine) *BlockController {
	blockCtrl := &BlockController{mapSizeX: mapSizeX, mapSizeY: mapSizeY}
	engine.updates = append(engine.updates, blockCtrl)

	return blockCtrl
}

// AddBlock is to update or add that positions block
func (ctrl *BlockController) AddBlock(_x, _y int) {
	ctrl.blocks[_y][_x] = &Block{x: _x, y: _y, isBlocked: true}
}

// RemoveBlock will remove that block
func (ctrl *BlockController) RemoveBlock(_x, _y int) {
	ctrl.blocks[_y][_x] = &Block{x: _x, y: _y, isBlocked: false}
}

// Draw will Render all blocks
func (ctrl *BlockController) Draw() string {
	draw := make(map[int]map[int]*Block)
	for _, blocks := range ctrl.blocks {
		for _, block := range blocks {
			if draw[block.y] == nil {
				draw[block.y] = make(map[int]*Block)
			}

			draw[block.y][block.x] = block
		}
	}

	for _, block := range ctrl.currentBlock {
		if draw[block.y] == nil {
			draw[block.y] = make(map[int]*Block)
		}
		draw[block.y][block.x] = block
	}

	uiString := ""
	for y := 0; y < ctrl.mapSizeY; y++ {
		for x := 0; x < ctrl.mapSizeX; x++ {
			uiString += " "
			if block, ok := draw[y][x]; ok {
				if block.isBlocked {
					uiString += "[X]"
					continue
				}
			}

			//original map
			uiString += "[ ]"
		}
		uiString += "\n"
	}
	return uiString
}

// Update will updates that all blocks states
func (ctrl *BlockController) Update(engine *GameEngine) {
	if ctrl.currentBlock == nil {
		ctrl.currentBlock = generateBlock()
	}

	for _, v := range ctrl.currentBlock {
		v.y++

		if engine.input == keyboard.KeyArrowLeft && v.x > 0 {
			v.x--
		} else if engine.input == keyboard.KeyArrowRight && v.x < ctrl.mapSizeX {
			v.x++
		}
	}
}

// generateBlock will create blocks from random
func generateBlock() []*Block {
	blocks := make([]*Block, 3)
	blocks[0] = &Block{x: 0, y: 0, isBlocked: true}
	blocks[2] = &Block{x: 2, y: 0, isBlocked: true}
	blocks[1] = &Block{x: 1, y: 0, isBlocked: true}

	return blocks
}
