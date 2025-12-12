package app

import (
	"day_7/internal/abstractions"
)

func Simulate(
	manifold *abstractions.Manifold,
) {
	beamDirection := abstractions.Direction{
		RowDelta: 1,
		ColDelta: 0,
	}

	movingTachyons := true

	for movingTachyons {

		for _, tachyon := range manifold.Tachyons {

			if !tachyon.IsMoving() {
				continue
			}

			if manifold.CanMove(*tachyon, beamDirection) == false {
				tachyon.Stop()
				continue
			}

			if manifold.IsNextLocationEmpty(*tachyon, beamDirection) {
				tachyon.Move(manifold, beamDirection)
			} else if manifold.IsNextLocationSplitter(*tachyon, beamDirection) {
				manifold.SplitBeamAt(tachyon, beamDirection)
			} else {
				tachyon.Stop()
				continue
			}

		}

		movingTachyons = false

		for _, tachyon := range manifold.Tachyons {
			movingTachyons = movingTachyons || tachyon.IsMoving()
		}
	}
}
