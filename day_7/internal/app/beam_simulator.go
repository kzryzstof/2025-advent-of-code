package app

import (
	"day_7/internal/abstractions"
	"fmt"
)

func Simulate(
	manifold *abstractions.Manifold,
	drawProgress bool,
) {
	beamDirection := abstractions.Direction{
		RowDelta: 1,
		ColDelta: 0,
	}

	for manifold.AreTachyonsMoving() {

		if drawProgress {
			manifold.Draw()
		}

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
			} else if manifold.IsNextLocationOtherTachyon(*tachyon, beamDirection) {
				/* No need to simulate the same tachyon multiple times: we simulate one but
				   keep track of the fact that there is more than one tachyon */
				manifold.Merge(tachyon, beamDirection)
			} else if manifold.IsNextLocationSplitter(*tachyon, beamDirection) {
				manifold.SplitBeamAt(tachyon, beamDirection)
			} else {
				tachyon.Stop()
				continue
			}
		}

		if drawProgress {
			fmt.Print("\033[2J\033[H")
		}
	}
}
