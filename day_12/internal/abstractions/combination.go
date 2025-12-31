package abstractions

type Combination struct {
	/* Index of the first present in the combination */
	Index PresentIndex
	/* Index of the other present in the combination */
	OtherIndex PresentIndex
	/* Overall shape of the combination */
	Shape Shape
}
