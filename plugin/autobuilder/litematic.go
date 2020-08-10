package autobuilder

func NewLitematic() *Litematic {
	l := new(Litematic)
	return l
}

type Litematic struct {
	Metadata struct {
		EnclosingSize struct {
			X, Y, Z int32
		}
		Auther       string
		Description  string
		Name         string
		RegionCount  int32
		TimeCreated  int64
		TimeModified int64
		TotalBlocks  int32
		TotalVolume  int32
	}
	Regions struct {
		Regions map[string]Region
	}
}

type Region struct {
	Position struct {
		X, Y, Z int32
	}
	Size struct {
		X, Y, Z int32
	}
	BlockStatePalette BlockStatePalette
	BlockStates       []int64
}

type BlockStatePalette struct {
	Blocks []Blocktype
}

type Blocktype struct {
	Properties map[string]string
	Name       string
}
