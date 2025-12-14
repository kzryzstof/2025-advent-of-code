package abstractions

type MovieTheater interface {
	GetRedTiles() []*Tile
	IsValidTile(x uint, y uint) bool
}
