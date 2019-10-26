// Time : 2019/10/19 20:28
// Author : MashiroC

// card
package card

import (
	"fmt"
	"mashiroc.fun/dragongame/game"
	"mashiroc.fun/dragongame/game/effect"
)

// mycard.go something

type WildBoar struct {
	game.BaseCard
}

func NewWildBoar() game.Card {
	return &WildBoar{
		*game.NewCard("石牙野猪", 1, 1, 1, effect.Charge{}),
	}
}

type ElvenArcher struct {
	game.BaseCard
}

func NewElvenArcher() game.Card {
	return &ElvenArcher{
		*game.NewCard("精灵弓箭手",
			1,
			1,
			1,
			effect.NewBattleCry(
				func(card game.Card, player game.Character, dragon game.Character) {
					fmt.Println("请选择一个对方目标：")
					var flag int
					fmt.Scanf("%d", &flag)
					var target game.Card
					if flag == 0 {
						target = dragon
					} else {
						target = dragon.Servant()[flag-1]
					}
					target.ReduceHp(1)
				}),
		),
	}
}

func NewDoctor() game.Card {
	return &ElvenArcher{
		*game.NewCard("巫医",
			1,
			1,
			2,
			effect.NewBattleCry(
				func(card game.Card, player game.Character, dragon game.Character) {
					fmt.Println("请选择一个对方目标：")
					var flag int
					fmt.Scanf("%d", &flag)
					var target game.Card
					if flag == 0 {
						target = player
					} else {
						target = dragon.Servant()[flag-1]
					}
					target.ReduceHp(-2)
					fmt.Println("有人要找医生吗？")
				}),
		),
	}
}

func NewDragon() game.Card {
	return &WildBoar{
		*game.NewCard("暴王龙克鲁什", 8, 9, 8, effect.Charge{}),
	}
}