package autobuilder

type template struct {
	Region []region
}

type region struct {
	BlockStatePalette blockStatePalette
	BlockStates       []int64
}

type blockStatePalette struct {
	Blocks []block
}

type block struct {
	Properties []string
	Name       string
}
