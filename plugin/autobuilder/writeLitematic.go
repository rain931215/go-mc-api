package autobuilder

import (
	"io/ioutil"

	nbt2Json "github.com/Lirsty/nbt2json"
	nbtTool "github.com/rain931215/go-mc-api/nbt"
)

func (l *Litematic) WriteFile(path string) {
	nbt2Json.UseJavaEncoding()
	nbt := nbtTool.NewNbt()
	fileNameTag := nbtTool.NewCompoundTag(path)
	Metadata := nbtTool.NewCompoundTag("Metadata")
	l.writeMetaData(Metadata)
	Regions := nbtTool.NewCompoundTag("Regions")
	l.writeRegions(Regions)
	fileNameTag.AddCompoundTag(Metadata)
	fileNameTag.AddCompoundTag(Regions)
	nbt.AddCompoundTag(fileNameTag)
	json, err := nbt.ToJson()
	if err != nil {
		panic(err)
	}
	data, err := nbt2Json.Json2Nbt(json)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(path, data, 0622)
	if err != nil {
		panic(err)
	}
}

func (l *Litematic) writeMetaData(tag *nbtTool.CompoundTag) {
	EnclosingSize := nbtTool.NewCompoundTag("EnclosingSize")
	EnclosingSize.AddNewValue("x", l.Metadata.EnclosingSize.X)
	EnclosingSize.AddNewValue("y", l.Metadata.EnclosingSize.Y)
	EnclosingSize.AddNewValue("z", l.Metadata.EnclosingSize.Z)
	tag.AddCompoundTag(EnclosingSize)
	tag.AddNewValue("Auther", l.Metadata.Auther)
	tag.AddNewValue("Description", l.Metadata.Description)
	tag.AddNewValue("Name", l.Metadata.Name)
	tag.AddNewValue("RegionCount", l.Metadata.RegionCount)
	tag.AddNewValue("TimeCreated", l.Metadata.TimeCreated)
	tag.AddNewValue("TimeModified", l.Metadata.TimeModified)
	tag.AddNewValue("TotalBlocks", l.Metadata.TotalBlocks)
	tag.AddNewValue("TotalVolume", l.Metadata.TotalVolume)
}

func (l *Litematic) writeRegions(tag *nbtTool.CompoundTag) {
	for name, region := range l.Regions.Regions {
		regionTag := nbtTool.NewCompoundTag(name)
		tag.AddCompoundTag(regionTag)
		Position := nbtTool.NewCompoundTag("Position")
		Size := nbtTool.NewCompoundTag("Size")
		Position.AddNewValue("x", region.Position.X)
		Position.AddNewValue("y", region.Position.Y)
		Position.AddNewValue("z", region.Position.Z)
		Size.AddNewValue("x", region.Size.X)
		Size.AddNewValue("y", region.Size.Y)
		Size.AddNewValue("z", region.Size.Z)
		regionTag.AddCompoundTag(Position)
		regionTag.AddCompoundTag(Size)
		BlockStatePalette := nbtTool.NewListTag("BlockStatePalette", nbtTool.TagCompound)
		for _, block := range region.BlockStatePalette.Blocks {
			Block := nbtTool.NewCompoundTag("none")
			Properties := nbtTool.NewCompoundTag("Properties")
			for k, v := range block.Properties {
				Properties.AddNewValue(k, v)
			}
			Block.AddCompoundTag(Properties)
			Block.AddNewValue("Name", block.Name)
			BlockStatePalette.AddCompoundTag(Block)
		}
		regionTag.AddListTag(BlockStatePalette)

		regionTag.AddNewValue("BlockStates", region.BlockStates)
	}
}
