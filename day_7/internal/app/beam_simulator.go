package app

import (
	"day_7/internal/abstractions"
	"fmt"
)

func Simulate(
	manifold *abstractions.Manifold,
) {
	beamDirection := abstractions.Direction{
		RowDelta: 1,
		ColDelta: 0,
	}

	movingTachyons := len(manifold.Tachyons)

	for movingTachyons > 0 {

		for _, tachyon := range manifold.Tachyons {

			if tachyon.Position.RowIndex == 140 {
				fmt.Println("Tachyon", tachyon)
			}

			if !tachyon.IsMoving() {
				continue
			}

			if manifold.CanMove(*tachyon, beamDirection) == false {
				tachyon.Stop()
				movingTachyons--
				continue
			}

			if manifold.IsNextLocationEmpty(*tachyon, beamDirection) {
				tachyon.Move(manifold, beamDirection)
			} else {
				tachyon.Stop()
				movingTachyons--
				continue
			}

		}
	}
}
