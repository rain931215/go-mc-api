package autobuilder

type template struct {
	region []region
}

type region struct {
	blockStatePalette blockStatePalette
	BlockStates       []int64
}

type blockStatePalette struct {
	blocks []block
}

type block struct {
	Properties []string
	Name       string
}
