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

	for manifold.AreTachyonsMoving() {

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
	}
}
