package data

import "encoding/json"

var itemIDs map[string]struct {
	ProtocolID int `json:"protocol_id"`
}

// ItemNameByID store Item's ID and
var ItemNameByID []string

func init() {
	json.Unmarshal([]byte(itemIDsJSON), &itemIDs)
	ItemNameByID = make([]string, 1000+1)
	for i, v := range itemIDs {
		ItemNameByID[v.ProtocolID] = i
	}
}

// Generate with follow steps:
// java -cp minecraft_server.1.15.jar net.minecraft.data.Main --all
// {reports/registries.json}.minecraft:block.entries
var itemIDsJSON = `{
    "minecraft:air": {
        "protocol_id": 0
      },
      "minecraft:stone": {
        "protocol_id": 1
      },
      "minecraft:granite": {
        "protocol_id": 2
      },
      "minecraft:polished_granite": {
        "protocol_id": 3
      },
      "minecraft:diorite": {
        "protocol_id": 4
      },
      "minecraft:polished_diorite": {
        "protocol_id": 5
      },
      "minecraft:andesite": {
        "protocol_id": 6
      },
      "minecraft:polished_andesite": {
        "protocol_id": 7
      },
      "minecraft:grass_block": {
        "protocol_id": 8
      },
      "minecraft:dirt": {
        "protocol_id": 9
      },
      "minecraft:coarse_dirt": {
        "protocol_id": 10
      },
      "minecraft:podzol": {
        "protocol_id": 11
      },
      "minecraft:cobblestone": {
        "protocol_id": 12
      },
      "minecraft:oak_planks": {
        "protocol_id": 13
      },
      "minecraft:spruce_planks": {
        "protocol_id": 14
      },
      "minecraft:birch_planks": {
        "protocol_id": 15
      },
      "minecraft:jungle_planks": {
        "protocol_id": 16
      },
      "minecraft:acacia_planks": {
        "protocol_id": 17
      },
      "minecraft:dark_oak_planks": {
        "protocol_id": 18
      },
      "minecraft:oak_sapling": {
        "protocol_id": 19
      },
      "minecraft:spruce_sapling": {
        "protocol_id": 20
      },
      "minecraft:birch_sapling": {
        "protocol_id": 21
      },
      "minecraft:jungle_sapling": {
        "protocol_id": 22
      },
      "minecraft:acacia_sapling": {
        "protocol_id": 23
      },
      "minecraft:dark_oak_sapling": {
        "protocol_id": 24
      },
      "minecraft:bedrock": {
        "protocol_id": 25
      },
      "minecraft:sand": {
        "protocol_id": 26
      },
      "minecraft:red_sand": {
        "protocol_id": 27
      },
      "minecraft:gravel": {
        "protocol_id": 28
      },
      "minecraft:gold_ore": {
        "protocol_id": 29
      },
      "minecraft:iron_ore": {
        "protocol_id": 30
      },
      "minecraft:coal_ore": {
        "protocol_id": 31
      },
      "minecraft:oak_log": {
        "protocol_id": 32
      },
      "minecraft:spruce_log": {
        "protocol_id": 33
      },
      "minecraft:birch_log": {
        "protocol_id": 34
      },
      "minecraft:jungle_log": {
        "protocol_id": 35
      },
      "minecraft:acacia_log": {
        "protocol_id": 36
      },
      "minecraft:dark_oak_log": {
        "protocol_id": 37
      },
      "minecraft:stripped_oak_log": {
        "protocol_id": 38
      },
      "minecraft:stripped_spruce_log": {
        "protocol_id": 39
      },
      "minecraft:stripped_birch_log": {
        "protocol_id": 40
      },
      "minecraft:stripped_jungle_log": {
        "protocol_id": 41
      },
      "minecraft:stripped_acacia_log": {
        "protocol_id": 42
      },
      "minecraft:stripped_dark_oak_log": {
        "protocol_id": 43
      },
      "minecraft:stripped_oak_wood": {
        "protocol_id": 44
      },
      "minecraft:stripped_spruce_wood": {
        "protocol_id": 45
      },
      "minecraft:stripped_birch_wood": {
        "protocol_id": 46
      },
      "minecraft:stripped_jungle_wood": {
        "protocol_id": 47
      },
      "minecraft:stripped_acacia_wood": {
        "protocol_id": 48
      },
      "minecraft:stripped_dark_oak_wood": {
        "protocol_id": 49
      },
      "minecraft:oak_wood": {
        "protocol_id": 50
      },
      "minecraft:spruce_wood": {
        "protocol_id": 51
      },
      "minecraft:birch_wood": {
        "protocol_id": 52
      },
      "minecraft:jungle_wood": {
        "protocol_id": 53
      },
      "minecraft:acacia_wood": {
        "protocol_id": 54
      },
      "minecraft:dark_oak_wood": {
        "protocol_id": 55
      },
      "minecraft:oak_leaves": {
        "protocol_id": 56
      },
      "minecraft:spruce_leaves": {
        "protocol_id": 57
      },
      "minecraft:birch_leaves": {
        "protocol_id": 58
      },
      "minecraft:jungle_leaves": {
        "protocol_id": 59
      },
      "minecraft:acacia_leaves": {
        "protocol_id": 60
      },
      "minecraft:dark_oak_leaves": {
        "protocol_id": 61
      },
      "minecraft:sponge": {
        "protocol_id": 62
      },
      "minecraft:wet_sponge": {
        "protocol_id": 63
      },
      "minecraft:glass": {
        "protocol_id": 64
      },
      "minecraft:lapis_ore": {
        "protocol_id": 65
      },
      "minecraft:lapis_block": {
        "protocol_id": 66
      },
      "minecraft:dispenser": {
        "protocol_id": 67
      },
      "minecraft:sandstone": {
        "protocol_id": 68
      },
      "minecraft:chiseled_sandstone": {
        "protocol_id": 69
      },
      "minecraft:cut_sandstone": {
        "protocol_id": 70
      },
      "minecraft:note_block": {
        "protocol_id": 71
      },
      "minecraft:powered_rail": {
        "protocol_id": 72
      },
      "minecraft:detector_rail": {
        "protocol_id": 73
      },
      "minecraft:sticky_piston": {
        "protocol_id": 74
      },
      "minecraft:cobweb": {
        "protocol_id": 75
      },
      "minecraft:grass": {
        "protocol_id": 76
      },
      "minecraft:fern": {
        "protocol_id": 77
      },
      "minecraft:dead_bush": {
        "protocol_id": 78
      },
      "minecraft:seagrass": {
        "protocol_id": 79
      },
      "minecraft:sea_pickle": {
        "protocol_id": 80
      },
      "minecraft:piston": {
        "protocol_id": 81
      },
      "minecraft:white_wool": {
        "protocol_id": 82
      },
      "minecraft:orange_wool": {
        "protocol_id": 83
      },
      "minecraft:magenta_wool": {
        "protocol_id": 84
      },
      "minecraft:light_blue_wool": {
        "protocol_id": 85
      },
      "minecraft:yellow_wool": {
        "protocol_id": 86
      },
      "minecraft:lime_wool": {
        "protocol_id": 87
      },
      "minecraft:pink_wool": {
        "protocol_id": 88
      },
      "minecraft:gray_wool": {
        "protocol_id": 89
      },
      "minecraft:light_gray_wool": {
        "protocol_id": 90
      },
      "minecraft:cyan_wool": {
        "protocol_id": 91
      },
      "minecraft:purple_wool": {
        "protocol_id": 92
      },
      "minecraft:blue_wool": {
        "protocol_id": 93
      },
      "minecraft:brown_wool": {
        "protocol_id": 94
      },
      "minecraft:green_wool": {
        "protocol_id": 95
      },
      "minecraft:red_wool": {
        "protocol_id": 96
      },
      "minecraft:black_wool": {
        "protocol_id": 97
      },
      "minecraft:dandelion": {
        "protocol_id": 98
      },
      "minecraft:poppy": {
        "protocol_id": 99
      },
      "minecraft:blue_orchid": {
        "protocol_id": 100
      },
      "minecraft:allium": {
        "protocol_id": 101
      },
      "minecraft:azure_bluet": {
        "protocol_id": 102
      },
      "minecraft:red_tulip": {
        "protocol_id": 103
      },
      "minecraft:orange_tulip": {
        "protocol_id": 104
      },
      "minecraft:white_tulip": {
        "protocol_id": 105
      },
      "minecraft:pink_tulip": {
        "protocol_id": 106
      },
      "minecraft:oxeye_daisy": {
        "protocol_id": 107
      },
      "minecraft:cornflower": {
        "protocol_id": 108
      },
      "minecraft:lily_of_the_valley": {
        "protocol_id": 109
      },
      "minecraft:wither_rose": {
        "protocol_id": 110
      },
      "minecraft:brown_mushroom": {
        "protocol_id": 111
      },
      "minecraft:red_mushroom": {
        "protocol_id": 112
      },
      "minecraft:gold_block": {
        "protocol_id": 113
      },
      "minecraft:iron_block": {
        "protocol_id": 114
      },
      "minecraft:oak_slab": {
        "protocol_id": 115
      },
      "minecraft:spruce_slab": {
        "protocol_id": 116
      },
      "minecraft:birch_slab": {
        "protocol_id": 117
      },
      "minecraft:jungle_slab": {
        "protocol_id": 118
      },
      "minecraft:acacia_slab": {
        "protocol_id": 119
      },
      "minecraft:dark_oak_slab": {
        "protocol_id": 120
      },
      "minecraft:stone_slab": {
        "protocol_id": 121
      },
      "minecraft:smooth_stone_slab": {
        "protocol_id": 122
      },
      "minecraft:sandstone_slab": {
        "protocol_id": 123
      },
      "minecraft:cut_sandstone_slab": {
        "protocol_id": 124
      },
      "minecraft:petrified_oak_slab": {
        "protocol_id": 125
      },
      "minecraft:cobblestone_slab": {
        "protocol_id": 126
      },
      "minecraft:brick_slab": {
        "protocol_id": 127
      },
      "minecraft:stone_brick_slab": {
        "protocol_id": 128
      },
      "minecraft:nether_brick_slab": {
        "protocol_id": 129
      },
      "minecraft:quartz_slab": {
        "protocol_id": 130
      },
      "minecraft:red_sandstone_slab": {
        "protocol_id": 131
      },
      "minecraft:cut_red_sandstone_slab": {
        "protocol_id": 132
      },
      "minecraft:purpur_slab": {
        "protocol_id": 133
      },
      "minecraft:prismarine_slab": {
        "protocol_id": 134
      },
      "minecraft:prismarine_brick_slab": {
        "protocol_id": 135
      },
      "minecraft:dark_prismarine_slab": {
        "protocol_id": 136
      },
      "minecraft:smooth_quartz": {
        "protocol_id": 137
      },
      "minecraft:smooth_red_sandstone": {
        "protocol_id": 138
      },
      "minecraft:smooth_sandstone": {
        "protocol_id": 139
      },
      "minecraft:smooth_stone": {
        "protocol_id": 140
      },
      "minecraft:bricks": {
        "protocol_id": 141
      },
      "minecraft:tnt": {
        "protocol_id": 142
      },
      "minecraft:bookshelf": {
        "protocol_id": 143
      },
      "minecraft:mossy_cobblestone": {
        "protocol_id": 144
      },
      "minecraft:obsidian": {
        "protocol_id": 145
      },
      "minecraft:torch": {
        "protocol_id": 146
      },
      "minecraft:end_rod": {
        "protocol_id": 147
      },
      "minecraft:chorus_plant": {
        "protocol_id": 148
      },
      "minecraft:chorus_flower": {
        "protocol_id": 149
      },
      "minecraft:purpur_block": {
        "protocol_id": 150
      },
      "minecraft:purpur_pillar": {
        "protocol_id": 151
      },
      "minecraft:purpur_stairs": {
        "protocol_id": 152
      },
      "minecraft:spawner": {
        "protocol_id": 153
      },
      "minecraft:oak_stairs": {
        "protocol_id": 154
      },
      "minecraft:chest": {
        "protocol_id": 155
      },
      "minecraft:diamond_ore": {
        "protocol_id": 156
      },
      "minecraft:diamond_block": {
        "protocol_id": 157
      },
      "minecraft:crafting_table": {
        "protocol_id": 158
      },
      "minecraft:farmland": {
        "protocol_id": 159
      },
      "minecraft:furnace": {
        "protocol_id": 160
      },
      "minecraft:ladder": {
        "protocol_id": 161
      },
      "minecraft:rail": {
        "protocol_id": 162
      },
      "minecraft:cobblestone_stairs": {
        "protocol_id": 163
      },
      "minecraft:lever": {
        "protocol_id": 164
      },
      "minecraft:stone_pressure_plate": {
        "protocol_id": 165
      },
      "minecraft:oak_pressure_plate": {
        "protocol_id": 166
      },
      "minecraft:spruce_pressure_plate": {
        "protocol_id": 167
      },
      "minecraft:birch_pressure_plate": {
        "protocol_id": 168
      },
      "minecraft:jungle_pressure_plate": {
        "protocol_id": 169
      },
      "minecraft:acacia_pressure_plate": {
        "protocol_id": 170
      },
      "minecraft:dark_oak_pressure_plate": {
        "protocol_id": 171
      },
      "minecraft:redstone_ore": {
        "protocol_id": 172
      },
      "minecraft:redstone_torch": {
        "protocol_id": 173
      },
      "minecraft:stone_button": {
        "protocol_id": 174
      },
      "minecraft:snow": {
        "protocol_id": 175
      },
      "minecraft:ice": {
        "protocol_id": 176
      },
      "minecraft:snow_block": {
        "protocol_id": 177
      },
      "minecraft:cactus": {
        "protocol_id": 178
      },
      "minecraft:clay": {
        "protocol_id": 179
      },
      "minecraft:jukebox": {
        "protocol_id": 180
      },
      "minecraft:oak_fence": {
        "protocol_id": 181
      },
      "minecraft:spruce_fence": {
        "protocol_id": 182
      },
      "minecraft:birch_fence": {
        "protocol_id": 183
      },
      "minecraft:jungle_fence": {
        "protocol_id": 184
      },
      "minecraft:acacia_fence": {
        "protocol_id": 185
      },
      "minecraft:dark_oak_fence": {
        "protocol_id": 186
      },
      "minecraft:pumpkin": {
        "protocol_id": 187
      },
      "minecraft:carved_pumpkin": {
        "protocol_id": 188
      },
      "minecraft:netherrack": {
        "protocol_id": 189
      },
      "minecraft:soul_sand": {
        "protocol_id": 190
      },
      "minecraft:glowstone": {
        "protocol_id": 191
      },
      "minecraft:jack_o_lantern": {
        "protocol_id": 192
      },
      "minecraft:oak_trapdoor": {
        "protocol_id": 193
      },
      "minecraft:spruce_trapdoor": {
        "protocol_id": 194
      },
      "minecraft:birch_trapdoor": {
        "protocol_id": 195
      },
      "minecraft:jungle_trapdoor": {
        "protocol_id": 196
      },
      "minecraft:acacia_trapdoor": {
        "protocol_id": 197
      },
      "minecraft:dark_oak_trapdoor": {
        "protocol_id": 198
      },
      "minecraft:infested_stone": {
        "protocol_id": 199
      },
      "minecraft:infested_cobblestone": {
        "protocol_id": 200
      },
      "minecraft:infested_stone_bricks": {
        "protocol_id": 201
      },
      "minecraft:infested_mossy_stone_bricks": {
        "protocol_id": 202
      },
      "minecraft:infested_cracked_stone_bricks": {
        "protocol_id": 203
      },
      "minecraft:infested_chiseled_stone_bricks": {
        "protocol_id": 204
      },
      "minecraft:stone_bricks": {
        "protocol_id": 205
      },
      "minecraft:mossy_stone_bricks": {
        "protocol_id": 206
      },
      "minecraft:cracked_stone_bricks": {
        "protocol_id": 207
      },
      "minecraft:chiseled_stone_bricks": {
        "protocol_id": 208
      },
      "minecraft:brown_mushroom_block": {
        "protocol_id": 209
      },
      "minecraft:red_mushroom_block": {
        "protocol_id": 210
      },
      "minecraft:mushroom_stem": {
        "protocol_id": 211
      },
      "minecraft:iron_bars": {
        "protocol_id": 212
      },
      "minecraft:glass_pane": {
        "protocol_id": 213
      },
      "minecraft:melon": {
        "protocol_id": 214
      },
      "minecraft:vine": {
        "protocol_id": 215
      },
      "minecraft:oak_fence_gate": {
        "protocol_id": 216
      },
      "minecraft:spruce_fence_gate": {
        "protocol_id": 217
      },
      "minecraft:birch_fence_gate": {
        "protocol_id": 218
      },
      "minecraft:jungle_fence_gate": {
        "protocol_id": 219
      },
      "minecraft:acacia_fence_gate": {
        "protocol_id": 220
      },
      "minecraft:dark_oak_fence_gate": {
        "protocol_id": 221
      },
      "minecraft:brick_stairs": {
        "protocol_id": 222
      },
      "minecraft:stone_brick_stairs": {
        "protocol_id": 223
      },
      "minecraft:mycelium": {
        "protocol_id": 224
      },
      "minecraft:lily_pad": {
        "protocol_id": 225
      },
      "minecraft:nether_bricks": {
        "protocol_id": 226
      },
      "minecraft:nether_brick_fence": {
        "protocol_id": 227
      },
      "minecraft:nether_brick_stairs": {
        "protocol_id": 228
      },
      "minecraft:enchanting_table": {
        "protocol_id": 229
      },
      "minecraft:end_portal_frame": {
        "protocol_id": 230
      },
      "minecraft:end_stone": {
        "protocol_id": 231
      },
      "minecraft:end_stone_bricks": {
        "protocol_id": 232
      },
      "minecraft:dragon_egg": {
        "protocol_id": 233
      },
      "minecraft:redstone_lamp": {
        "protocol_id": 234
      },
      "minecraft:sandstone_stairs": {
        "protocol_id": 235
      },
      "minecraft:emerald_ore": {
        "protocol_id": 236
      },
      "minecraft:ender_chest": {
        "protocol_id": 237
      },
      "minecraft:tripwire_hook": {
        "protocol_id": 238
      },
      "minecraft:emerald_block": {
        "protocol_id": 239
      },
      "minecraft:spruce_stairs": {
        "protocol_id": 240
      },
      "minecraft:birch_stairs": {
        "protocol_id": 241
      },
      "minecraft:jungle_stairs": {
        "protocol_id": 242
      },
      "minecraft:command_block": {
        "protocol_id": 243
      },
      "minecraft:beacon": {
        "protocol_id": 244
      },
      "minecraft:cobblestone_wall": {
        "protocol_id": 245
      },
      "minecraft:mossy_cobblestone_wall": {
        "protocol_id": 246
      },
      "minecraft:brick_wall": {
        "protocol_id": 247
      },
      "minecraft:prismarine_wall": {
        "protocol_id": 248
      },
      "minecraft:red_sandstone_wall": {
        "protocol_id": 249
      },
      "minecraft:mossy_stone_brick_wall": {
        "protocol_id": 250
      },
      "minecraft:granite_wall": {
        "protocol_id": 251
      },
      "minecraft:stone_brick_wall": {
        "protocol_id": 252
      },
      "minecraft:nether_brick_wall": {
        "protocol_id": 253
      },
      "minecraft:andesite_wall": {
        "protocol_id": 254
      },
      "minecraft:red_nether_brick_wall": {
        "protocol_id": 255
      },
      "minecraft:sandstone_wall": {
        "protocol_id": 256
      },
      "minecraft:end_stone_brick_wall": {
        "protocol_id": 257
      },
      "minecraft:diorite_wall": {
        "protocol_id": 258
      },
      "minecraft:oak_button": {
        "protocol_id": 259
      },
      "minecraft:spruce_button": {
        "protocol_id": 260
      },
      "minecraft:birch_button": {
        "protocol_id": 261
      },
      "minecraft:jungle_button": {
        "protocol_id": 262
      },
      "minecraft:acacia_button": {
        "protocol_id": 263
      },
      "minecraft:dark_oak_button": {
        "protocol_id": 264
      },
      "minecraft:anvil": {
        "protocol_id": 265
      },
      "minecraft:chipped_anvil": {
        "protocol_id": 266
      },
      "minecraft:damaged_anvil": {
        "protocol_id": 267
      },
      "minecraft:trapped_chest": {
        "protocol_id": 268
      },
      "minecraft:light_weighted_pressure_plate": {
        "protocol_id": 269
      },
      "minecraft:heavy_weighted_pressure_plate": {
        "protocol_id": 270
      },
      "minecraft:daylight_detector": {
        "protocol_id": 271
      },
      "minecraft:redstone_block": {
        "protocol_id": 272
      },
      "minecraft:nether_quartz_ore": {
        "protocol_id": 273
      },
      "minecraft:hopper": {
        "protocol_id": 274
      },
      "minecraft:chiseled_quartz_block": {
        "protocol_id": 275
      },
      "minecraft:quartz_block": {
        "protocol_id": 276
      },
      "minecraft:quartz_pillar": {
        "protocol_id": 277
      },
      "minecraft:quartz_stairs": {
        "protocol_id": 278
      },
      "minecraft:activator_rail": {
        "protocol_id": 279
      },
      "minecraft:dropper": {
        "protocol_id": 280
      },
      "minecraft:white_terracotta": {
        "protocol_id": 281
      },
      "minecraft:orange_terracotta": {
        "protocol_id": 282
      },
      "minecraft:magenta_terracotta": {
        "protocol_id": 283
      },
      "minecraft:light_blue_terracotta": {
        "protocol_id": 284
      },
      "minecraft:yellow_terracotta": {
        "protocol_id": 285
      },
      "minecraft:lime_terracotta": {
        "protocol_id": 286
      },
      "minecraft:pink_terracotta": {
        "protocol_id": 287
      },
      "minecraft:gray_terracotta": {
        "protocol_id": 288
      },
      "minecraft:light_gray_terracotta": {
        "protocol_id": 289
      },
      "minecraft:cyan_terracotta": {
        "protocol_id": 290
      },
      "minecraft:purple_terracotta": {
        "protocol_id": 291
      },
      "minecraft:blue_terracotta": {
        "protocol_id": 292
      },
      "minecraft:brown_terracotta": {
        "protocol_id": 293
      },
      "minecraft:green_terracotta": {
        "protocol_id": 294
      },
      "minecraft:red_terracotta": {
        "protocol_id": 295
      },
      "minecraft:black_terracotta": {
        "protocol_id": 296
      },
      "minecraft:barrier": {
        "protocol_id": 297
      },
      "minecraft:iron_trapdoor": {
        "protocol_id": 298
      },
      "minecraft:hay_block": {
        "protocol_id": 299
      },
      "minecraft:white_carpet": {
        "protocol_id": 300
      },
      "minecraft:orange_carpet": {
        "protocol_id": 301
      },
      "minecraft:magenta_carpet": {
        "protocol_id": 302
      },
      "minecraft:light_blue_carpet": {
        "protocol_id": 303
      },
      "minecraft:yellow_carpet": {
        "protocol_id": 304
      },
      "minecraft:lime_carpet": {
        "protocol_id": 305
      },
      "minecraft:pink_carpet": {
        "protocol_id": 306
      },
      "minecraft:gray_carpet": {
        "protocol_id": 307
      },
      "minecraft:light_gray_carpet": {
        "protocol_id": 308
      },
      "minecraft:cyan_carpet": {
        "protocol_id": 309
      },
      "minecraft:purple_carpet": {
        "protocol_id": 310
      },
      "minecraft:blue_carpet": {
        "protocol_id": 311
      },
      "minecraft:brown_carpet": {
        "protocol_id": 312
      },
      "minecraft:green_carpet": {
        "protocol_id": 313
      },
      "minecraft:red_carpet": {
        "protocol_id": 314
      },
      "minecraft:black_carpet": {
        "protocol_id": 315
      },
      "minecraft:terracotta": {
        "protocol_id": 316
      },
      "minecraft:coal_block": {
        "protocol_id": 317
      },
      "minecraft:packed_ice": {
        "protocol_id": 318
      },
      "minecraft:acacia_stairs": {
        "protocol_id": 319
      },
      "minecraft:dark_oak_stairs": {
        "protocol_id": 320
      },
      "minecraft:slime_block": {
        "protocol_id": 321
      },
      "minecraft:grass_path": {
        "protocol_id": 322
      },
      "minecraft:sunflower": {
        "protocol_id": 323
      },
      "minecraft:lilac": {
        "protocol_id": 324
      },
      "minecraft:rose_bush": {
        "protocol_id": 325
      },
      "minecraft:peony": {
        "protocol_id": 326
      },
      "minecraft:tall_grass": {
        "protocol_id": 327
      },
      "minecraft:large_fern": {
        "protocol_id": 328
      },
      "minecraft:white_stained_glass": {
        "protocol_id": 329
      },
      "minecraft:orange_stained_glass": {
        "protocol_id": 330
      },
      "minecraft:magenta_stained_glass": {
        "protocol_id": 331
      },
      "minecraft:light_blue_stained_glass": {
        "protocol_id": 332
      },
      "minecraft:yellow_stained_glass": {
        "protocol_id": 333
      },
      "minecraft:lime_stained_glass": {
        "protocol_id": 334
      },
      "minecraft:pink_stained_glass": {
        "protocol_id": 335
      },
      "minecraft:gray_stained_glass": {
        "protocol_id": 336
      },
      "minecraft:light_gray_stained_glass": {
        "protocol_id": 337
      },
      "minecraft:cyan_stained_glass": {
        "protocol_id": 338
      },
      "minecraft:purple_stained_glass": {
        "protocol_id": 339
      },
      "minecraft:blue_stained_glass": {
        "protocol_id": 340
      },
      "minecraft:brown_stained_glass": {
        "protocol_id": 341
      },
      "minecraft:green_stained_glass": {
        "protocol_id": 342
      },
      "minecraft:red_stained_glass": {
        "protocol_id": 343
      },
      "minecraft:black_stained_glass": {
        "protocol_id": 344
      },
      "minecraft:white_stained_glass_pane": {
        "protocol_id": 345
      },
      "minecraft:orange_stained_glass_pane": {
        "protocol_id": 346
      },
      "minecraft:magenta_stained_glass_pane": {
        "protocol_id": 347
      },
      "minecraft:light_blue_stained_glass_pane": {
        "protocol_id": 348
      },
      "minecraft:yellow_stained_glass_pane": {
        "protocol_id": 349
      },
      "minecraft:lime_stained_glass_pane": {
        "protocol_id": 350
      },
      "minecraft:pink_stained_glass_pane": {
        "protocol_id": 351
      },
      "minecraft:gray_stained_glass_pane": {
        "protocol_id": 352
      },
      "minecraft:light_gray_stained_glass_pane": {
        "protocol_id": 353
      },
      "minecraft:cyan_stained_glass_pane": {
        "protocol_id": 354
      },
      "minecraft:purple_stained_glass_pane": {
        "protocol_id": 355
      },
      "minecraft:blue_stained_glass_pane": {
        "protocol_id": 356
      },
      "minecraft:brown_stained_glass_pane": {
        "protocol_id": 357
      },
      "minecraft:green_stained_glass_pane": {
        "protocol_id": 358
      },
      "minecraft:red_stained_glass_pane": {
        "protocol_id": 359
      },
      "minecraft:black_stained_glass_pane": {
        "protocol_id": 360
      },
      "minecraft:prismarine": {
        "protocol_id": 361
      },
      "minecraft:prismarine_bricks": {
        "protocol_id": 362
      },
      "minecraft:dark_prismarine": {
        "protocol_id": 363
      },
      "minecraft:prismarine_stairs": {
        "protocol_id": 364
      },
      "minecraft:prismarine_brick_stairs": {
        "protocol_id": 365
      },
      "minecraft:dark_prismarine_stairs": {
        "protocol_id": 366
      },
      "minecraft:sea_lantern": {
        "protocol_id": 367
      },
      "minecraft:red_sandstone": {
        "protocol_id": 368
      },
      "minecraft:chiseled_red_sandstone": {
        "protocol_id": 369
      },
      "minecraft:cut_red_sandstone": {
        "protocol_id": 370
      },
      "minecraft:red_sandstone_stairs": {
        "protocol_id": 371
      },
      "minecraft:repeating_command_block": {
        "protocol_id": 372
      },
      "minecraft:chain_command_block": {
        "protocol_id": 373
      },
      "minecraft:magma_block": {
        "protocol_id": 374
      },
      "minecraft:nether_wart_block": {
        "protocol_id": 375
      },
      "minecraft:red_nether_bricks": {
        "protocol_id": 376
      },
      "minecraft:bone_block": {
        "protocol_id": 377
      },
      "minecraft:structure_void": {
        "protocol_id": 378
      },
      "minecraft:observer": {
        "protocol_id": 379
      },
      "minecraft:shulker_box": {
        "protocol_id": 380
      },
      "minecraft:white_shulker_box": {
        "protocol_id": 381
      },
      "minecraft:orange_shulker_box": {
        "protocol_id": 382
      },
      "minecraft:magenta_shulker_box": {
        "protocol_id": 383
      },
      "minecraft:light_blue_shulker_box": {
        "protocol_id": 384
      },
      "minecraft:yellow_shulker_box": {
        "protocol_id": 385
      },
      "minecraft:lime_shulker_box": {
        "protocol_id": 386
      },
      "minecraft:pink_shulker_box": {
        "protocol_id": 387
      },
      "minecraft:gray_shulker_box": {
        "protocol_id": 388
      },
      "minecraft:light_gray_shulker_box": {
        "protocol_id": 389
      },
      "minecraft:cyan_shulker_box": {
        "protocol_id": 390
      },
      "minecraft:purple_shulker_box": {
        "protocol_id": 391
      },
      "minecraft:blue_shulker_box": {
        "protocol_id": 392
      },
      "minecraft:brown_shulker_box": {
        "protocol_id": 393
      },
      "minecraft:green_shulker_box": {
        "protocol_id": 394
      },
      "minecraft:red_shulker_box": {
        "protocol_id": 395
      },
      "minecraft:black_shulker_box": {
        "protocol_id": 396
      },
      "minecraft:white_glazed_terracotta": {
        "protocol_id": 397
      },
      "minecraft:orange_glazed_terracotta": {
        "protocol_id": 398
      },
      "minecraft:magenta_glazed_terracotta": {
        "protocol_id": 399
      },
      "minecraft:light_blue_glazed_terracotta": {
        "protocol_id": 400
      },
      "minecraft:yellow_glazed_terracotta": {
        "protocol_id": 401
      },
      "minecraft:lime_glazed_terracotta": {
        "protocol_id": 402
      },
      "minecraft:pink_glazed_terracotta": {
        "protocol_id": 403
      },
      "minecraft:gray_glazed_terracotta": {
        "protocol_id": 404
      },
      "minecraft:light_gray_glazed_terracotta": {
        "protocol_id": 405
      },
      "minecraft:cyan_glazed_terracotta": {
        "protocol_id": 406
      },
      "minecraft:purple_glazed_terracotta": {
        "protocol_id": 407
      },
      "minecraft:blue_glazed_terracotta": {
        "protocol_id": 408
      },
      "minecraft:brown_glazed_terracotta": {
        "protocol_id": 409
      },
      "minecraft:green_glazed_terracotta": {
        "protocol_id": 410
      },
      "minecraft:red_glazed_terracotta": {
        "protocol_id": 411
      },
      "minecraft:black_glazed_terracotta": {
        "protocol_id": 412
      },
      "minecraft:white_concrete": {
        "protocol_id": 413
      },
      "minecraft:orange_concrete": {
        "protocol_id": 414
      },
      "minecraft:magenta_concrete": {
        "protocol_id": 415
      },
      "minecraft:light_blue_concrete": {
        "protocol_id": 416
      },
      "minecraft:yellow_concrete": {
        "protocol_id": 417
      },
      "minecraft:lime_concrete": {
        "protocol_id": 418
      },
      "minecraft:pink_concrete": {
        "protocol_id": 419
      },
      "minecraft:gray_concrete": {
        "protocol_id": 420
      },
      "minecraft:light_gray_concrete": {
        "protocol_id": 421
      },
      "minecraft:cyan_concrete": {
        "protocol_id": 422
      },
      "minecraft:purple_concrete": {
        "protocol_id": 423
      },
      "minecraft:blue_concrete": {
        "protocol_id": 424
      },
      "minecraft:brown_concrete": {
        "protocol_id": 425
      },
      "minecraft:green_concrete": {
        "protocol_id": 426
      },
      "minecraft:red_concrete": {
        "protocol_id": 427
      },
      "minecraft:black_concrete": {
        "protocol_id": 428
      },
      "minecraft:white_concrete_powder": {
        "protocol_id": 429
      },
      "minecraft:orange_concrete_powder": {
        "protocol_id": 430
      },
      "minecraft:magenta_concrete_powder": {
        "protocol_id": 431
      },
      "minecraft:light_blue_concrete_powder": {
        "protocol_id": 432
      },
      "minecraft:yellow_concrete_powder": {
        "protocol_id": 433
      },
      "minecraft:lime_concrete_powder": {
        "protocol_id": 434
      },
      "minecraft:pink_concrete_powder": {
        "protocol_id": 435
      },
      "minecraft:gray_concrete_powder": {
        "protocol_id": 436
      },
      "minecraft:light_gray_concrete_powder": {
        "protocol_id": 437
      },
      "minecraft:cyan_concrete_powder": {
        "protocol_id": 438
      },
      "minecraft:purple_concrete_powder": {
        "protocol_id": 439
      },
      "minecraft:blue_concrete_powder": {
        "protocol_id": 440
      },
      "minecraft:brown_concrete_powder": {
        "protocol_id": 441
      },
      "minecraft:green_concrete_powder": {
        "protocol_id": 442
      },
      "minecraft:red_concrete_powder": {
        "protocol_id": 443
      },
      "minecraft:black_concrete_powder": {
        "protocol_id": 444
      },
      "minecraft:turtle_egg": {
        "protocol_id": 445
      },
      "minecraft:dead_tube_coral_block": {
        "protocol_id": 446
      },
      "minecraft:dead_brain_coral_block": {
        "protocol_id": 447
      },
      "minecraft:dead_bubble_coral_block": {
        "protocol_id": 448
      },
      "minecraft:dead_fire_coral_block": {
        "protocol_id": 449
      },
      "minecraft:dead_horn_coral_block": {
        "protocol_id": 450
      },
      "minecraft:tube_coral_block": {
        "protocol_id": 451
      },
      "minecraft:brain_coral_block": {
        "protocol_id": 452
      },
      "minecraft:bubble_coral_block": {
        "protocol_id": 453
      },
      "minecraft:fire_coral_block": {
        "protocol_id": 454
      },
      "minecraft:horn_coral_block": {
        "protocol_id": 455
      },
      "minecraft:tube_coral": {
        "protocol_id": 456
      },
      "minecraft:brain_coral": {
        "protocol_id": 457
      },
      "minecraft:bubble_coral": {
        "protocol_id": 458
      },
      "minecraft:fire_coral": {
        "protocol_id": 459
      },
      "minecraft:horn_coral": {
        "protocol_id": 460
      },
      "minecraft:dead_brain_coral": {
        "protocol_id": 461
      },
      "minecraft:dead_bubble_coral": {
        "protocol_id": 462
      },
      "minecraft:dead_fire_coral": {
        "protocol_id": 463
      },
      "minecraft:dead_horn_coral": {
        "protocol_id": 464
      },
      "minecraft:dead_tube_coral": {
        "protocol_id": 465
      },
      "minecraft:tube_coral_fan": {
        "protocol_id": 466
      },
      "minecraft:brain_coral_fan": {
        "protocol_id": 467
      },
      "minecraft:bubble_coral_fan": {
        "protocol_id": 468
      },
      "minecraft:fire_coral_fan": {
        "protocol_id": 469
      },
      "minecraft:horn_coral_fan": {
        "protocol_id": 470
      },
      "minecraft:dead_tube_coral_fan": {
        "protocol_id": 471
      },
      "minecraft:dead_brain_coral_fan": {
        "protocol_id": 472
      },
      "minecraft:dead_bubble_coral_fan": {
        "protocol_id": 473
      },
      "minecraft:dead_fire_coral_fan": {
        "protocol_id": 474
      },
      "minecraft:dead_horn_coral_fan": {
        "protocol_id": 475
      },
      "minecraft:blue_ice": {
        "protocol_id": 476
      },
      "minecraft:conduit": {
        "protocol_id": 477
      },
      "minecraft:polished_granite_stairs": {
        "protocol_id": 478
      },
      "minecraft:smooth_red_sandstone_stairs": {
        "protocol_id": 479
      },
      "minecraft:mossy_stone_brick_stairs": {
        "protocol_id": 480
      },
      "minecraft:polished_diorite_stairs": {
        "protocol_id": 481
      },
      "minecraft:mossy_cobblestone_stairs": {
        "protocol_id": 482
      },
      "minecraft:end_stone_brick_stairs": {
        "protocol_id": 483
      },
      "minecraft:stone_stairs": {
        "protocol_id": 484
      },
      "minecraft:smooth_sandstone_stairs": {
        "protocol_id": 485
      },
      "minecraft:smooth_quartz_stairs": {
        "protocol_id": 486
      },
      "minecraft:granite_stairs": {
        "protocol_id": 487
      },
      "minecraft:andesite_stairs": {
        "protocol_id": 488
      },
      "minecraft:red_nether_brick_stairs": {
        "protocol_id": 489
      },
      "minecraft:polished_andesite_stairs": {
        "protocol_id": 490
      },
      "minecraft:diorite_stairs": {
        "protocol_id": 491
      },
      "minecraft:polished_granite_slab": {
        "protocol_id": 492
      },
      "minecraft:smooth_red_sandstone_slab": {
        "protocol_id": 493
      },
      "minecraft:mossy_stone_brick_slab": {
        "protocol_id": 494
      },
      "minecraft:polished_diorite_slab": {
        "protocol_id": 495
      },
      "minecraft:mossy_cobblestone_slab": {
        "protocol_id": 496
      },
      "minecraft:end_stone_brick_slab": {
        "protocol_id": 497
      },
      "minecraft:smooth_sandstone_slab": {
        "protocol_id": 498
      },
      "minecraft:smooth_quartz_slab": {
        "protocol_id": 499
      },
      "minecraft:granite_slab": {
        "protocol_id": 500
      },
      "minecraft:andesite_slab": {
        "protocol_id": 501
      },
      "minecraft:red_nether_brick_slab": {
        "protocol_id": 502
      },
      "minecraft:polished_andesite_slab": {
        "protocol_id": 503
      },
      "minecraft:diorite_slab": {
        "protocol_id": 504
      },
      "minecraft:scaffolding": {
        "protocol_id": 505
      },
      "minecraft:iron_door": {
        "protocol_id": 506
      },
      "minecraft:oak_door": {
        "protocol_id": 507
      },
      "minecraft:spruce_door": {
        "protocol_id": 508
      },
      "minecraft:birch_door": {
        "protocol_id": 509
      },
      "minecraft:jungle_door": {
        "protocol_id": 510
      },
      "minecraft:acacia_door": {
        "protocol_id": 511
      },
      "minecraft:dark_oak_door": {
        "protocol_id": 512
      },
      "minecraft:repeater": {
        "protocol_id": 513
      },
      "minecraft:comparator": {
        "protocol_id": 514
      },
      "minecraft:structure_block": {
        "protocol_id": 515
      },
      "minecraft:jigsaw": {
        "protocol_id": 516
      },
      "minecraft:composter": {
        "protocol_id": 517
      },
      "minecraft:turtle_helmet": {
        "protocol_id": 518
      },
      "minecraft:scute": {
        "protocol_id": 519
      },
      "minecraft:iron_shovel": {
        "protocol_id": 520
      },
      "minecraft:iron_pickaxe": {
        "protocol_id": 521
      },
      "minecraft:iron_axe": {
        "protocol_id": 522
      },
      "minecraft:flint_and_steel": {
        "protocol_id": 523
      },
      "minecraft:apple": {
        "protocol_id": 524
      },
      "minecraft:bow": {
        "protocol_id": 525
      },
      "minecraft:arrow": {
        "protocol_id": 526
      },
      "minecraft:coal": {
        "protocol_id": 527
      },
      "minecraft:charcoal": {
        "protocol_id": 528
      },
      "minecraft:diamond": {
        "protocol_id": 529
      },
      "minecraft:iron_ingot": {
        "protocol_id": 530
      },
      "minecraft:gold_ingot": {
        "protocol_id": 531
      },
      "minecraft:iron_sword": {
        "protocol_id": 532
      },
      "minecraft:wooden_sword": {
        "protocol_id": 533
      },
      "minecraft:wooden_shovel": {
        "protocol_id": 534
      },
      "minecraft:wooden_pickaxe": {
        "protocol_id": 535
      },
      "minecraft:wooden_axe": {
        "protocol_id": 536
      },
      "minecraft:stone_sword": {
        "protocol_id": 537
      },
      "minecraft:stone_shovel": {
        "protocol_id": 538
      },
      "minecraft:stone_pickaxe": {
        "protocol_id": 539
      },
      "minecraft:stone_axe": {
        "protocol_id": 540
      },
      "minecraft:diamond_sword": {
        "protocol_id": 541
      },
      "minecraft:diamond_shovel": {
        "protocol_id": 542
      },
      "minecraft:diamond_pickaxe": {
        "protocol_id": 543
      },
      "minecraft:diamond_axe": {
        "protocol_id": 544
      },
      "minecraft:stick": {
        "protocol_id": 545
      },
      "minecraft:bowl": {
        "protocol_id": 546
      },
      "minecraft:mushroom_stew": {
        "protocol_id": 547
      },
      "minecraft:golden_sword": {
        "protocol_id": 548
      },
      "minecraft:golden_shovel": {
        "protocol_id": 549
      },
      "minecraft:golden_pickaxe": {
        "protocol_id": 550
      },
      "minecraft:golden_axe": {
        "protocol_id": 551
      },
      "minecraft:string": {
        "protocol_id": 552
      },
      "minecraft:feather": {
        "protocol_id": 553
      },
      "minecraft:gunpowder": {
        "protocol_id": 554
      },
      "minecraft:wooden_hoe": {
        "protocol_id": 555
      },
      "minecraft:stone_hoe": {
        "protocol_id": 556
      },
      "minecraft:iron_hoe": {
        "protocol_id": 557
      },
      "minecraft:diamond_hoe": {
        "protocol_id": 558
      },
      "minecraft:golden_hoe": {
        "protocol_id": 559
      },
      "minecraft:wheat_seeds": {
        "protocol_id": 560
      },
      "minecraft:wheat": {
        "protocol_id": 561
      },
      "minecraft:bread": {
        "protocol_id": 562
      },
      "minecraft:leather_helmet": {
        "protocol_id": 563
      },
      "minecraft:leather_chestplate": {
        "protocol_id": 564
      },
      "minecraft:leather_leggings": {
        "protocol_id": 565
      },
      "minecraft:leather_boots": {
        "protocol_id": 566
      },
      "minecraft:chainmail_helmet": {
        "protocol_id": 567
      },
      "minecraft:chainmail_chestplate": {
        "protocol_id": 568
      },
      "minecraft:chainmail_leggings": {
        "protocol_id": 569
      },
      "minecraft:chainmail_boots": {
        "protocol_id": 570
      },
      "minecraft:iron_helmet": {
        "protocol_id": 571
      },
      "minecraft:iron_chestplate": {
        "protocol_id": 572
      },
      "minecraft:iron_leggings": {
        "protocol_id": 573
      },
      "minecraft:iron_boots": {
        "protocol_id": 574
      },
      "minecraft:diamond_helmet": {
        "protocol_id": 575
      },
      "minecraft:diamond_chestplate": {
        "protocol_id": 576
      },
      "minecraft:diamond_leggings": {
        "protocol_id": 577
      },
      "minecraft:diamond_boots": {
        "protocol_id": 578
      },
      "minecraft:golden_helmet": {
        "protocol_id": 579
      },
      "minecraft:golden_chestplate": {
        "protocol_id": 580
      },
      "minecraft:golden_leggings": {
        "protocol_id": 581
      },
      "minecraft:golden_boots": {
        "protocol_id": 582
      },
      "minecraft:flint": {
        "protocol_id": 583
      },
      "minecraft:porkchop": {
        "protocol_id": 584
      },
      "minecraft:cooked_porkchop": {
        "protocol_id": 585
      },
      "minecraft:painting": {
        "protocol_id": 586
      },
      "minecraft:golden_apple": {
        "protocol_id": 587
      },
      "minecraft:enchanted_golden_apple": {
        "protocol_id": 588
      },
      "minecraft:oak_sign": {
        "protocol_id": 589
      },
      "minecraft:spruce_sign": {
        "protocol_id": 590
      },
      "minecraft:birch_sign": {
        "protocol_id": 591
      },
      "minecraft:jungle_sign": {
        "protocol_id": 592
      },
      "minecraft:acacia_sign": {
        "protocol_id": 593
      },
      "minecraft:dark_oak_sign": {
        "protocol_id": 594
      },
      "minecraft:bucket": {
        "protocol_id": 595
      },
      "minecraft:water_bucket": {
        "protocol_id": 596
      },
      "minecraft:lava_bucket": {
        "protocol_id": 597
      },
      "minecraft:minecart": {
        "protocol_id": 598
      },
      "minecraft:saddle": {
        "protocol_id": 599
      },
      "minecraft:redstone": {
        "protocol_id": 600
      },
      "minecraft:snowball": {
        "protocol_id": 601
      },
      "minecraft:oak_boat": {
        "protocol_id": 602
      },
      "minecraft:leather": {
        "protocol_id": 603
      },
      "minecraft:milk_bucket": {
        "protocol_id": 604
      },
      "minecraft:pufferfish_bucket": {
        "protocol_id": 605
      },
      "minecraft:salmon_bucket": {
        "protocol_id": 606
      },
      "minecraft:cod_bucket": {
        "protocol_id": 607
      },
      "minecraft:tropical_fish_bucket": {
        "protocol_id": 608
      },
      "minecraft:brick": {
        "protocol_id": 609
      },
      "minecraft:clay_ball": {
        "protocol_id": 610
      },
      "minecraft:sugar_cane": {
        "protocol_id": 611
      },
      "minecraft:kelp": {
        "protocol_id": 612
      },
      "minecraft:dried_kelp_block": {
        "protocol_id": 613
      },
      "minecraft:bamboo": {
        "protocol_id": 614
      },
      "minecraft:paper": {
        "protocol_id": 615
      },
      "minecraft:book": {
        "protocol_id": 616
      },
      "minecraft:slime_ball": {
        "protocol_id": 617
      },
      "minecraft:chest_minecart": {
        "protocol_id": 618
      },
      "minecraft:furnace_minecart": {
        "protocol_id": 619
      },
      "minecraft:egg": {
        "protocol_id": 620
      },
      "minecraft:compass": {
        "protocol_id": 621
      },
      "minecraft:fishing_rod": {
        "protocol_id": 622
      },
      "minecraft:clock": {
        "protocol_id": 623
      },
      "minecraft:glowstone_dust": {
        "protocol_id": 624
      },
      "minecraft:cod": {
        "protocol_id": 625
      },
      "minecraft:salmon": {
        "protocol_id": 626
      },
      "minecraft:tropical_fish": {
        "protocol_id": 627
      },
      "minecraft:pufferfish": {
        "protocol_id": 628
      },
      "minecraft:cooked_cod": {
        "protocol_id": 629
      },
      "minecraft:cooked_salmon": {
        "protocol_id": 630
      },
      "minecraft:ink_sac": {
        "protocol_id": 631
      },
      "minecraft:red_dye": {
        "protocol_id": 632
      },
      "minecraft:green_dye": {
        "protocol_id": 633
      },
      "minecraft:cocoa_beans": {
        "protocol_id": 634
      },
      "minecraft:lapis_lazuli": {
        "protocol_id": 635
      },
      "minecraft:purple_dye": {
        "protocol_id": 636
      },
      "minecraft:cyan_dye": {
        "protocol_id": 637
      },
      "minecraft:light_gray_dye": {
        "protocol_id": 638
      },
      "minecraft:gray_dye": {
        "protocol_id": 639
      },
      "minecraft:pink_dye": {
        "protocol_id": 640
      },
      "minecraft:lime_dye": {
        "protocol_id": 641
      },
      "minecraft:yellow_dye": {
        "protocol_id": 642
      },
      "minecraft:light_blue_dye": {
        "protocol_id": 643
      },
      "minecraft:magenta_dye": {
        "protocol_id": 644
      },
      "minecraft:orange_dye": {
        "protocol_id": 645
      },
      "minecraft:bone_meal": {
        "protocol_id": 646
      },
      "minecraft:blue_dye": {
        "protocol_id": 647
      },
      "minecraft:brown_dye": {
        "protocol_id": 648
      },
      "minecraft:black_dye": {
        "protocol_id": 649
      },
      "minecraft:white_dye": {
        "protocol_id": 650
      },
      "minecraft:bone": {
        "protocol_id": 651
      },
      "minecraft:sugar": {
        "protocol_id": 652
      },
      "minecraft:cake": {
        "protocol_id": 653
      },
      "minecraft:white_bed": {
        "protocol_id": 654
      },
      "minecraft:orange_bed": {
        "protocol_id": 655
      },
      "minecraft:magenta_bed": {
        "protocol_id": 656
      },
      "minecraft:light_blue_bed": {
        "protocol_id": 657
      },
      "minecraft:yellow_bed": {
        "protocol_id": 658
      },
      "minecraft:lime_bed": {
        "protocol_id": 659
      },
      "minecraft:pink_bed": {
        "protocol_id": 660
      },
      "minecraft:gray_bed": {
        "protocol_id": 661
      },
      "minecraft:light_gray_bed": {
        "protocol_id": 662
      },
      "minecraft:cyan_bed": {
        "protocol_id": 663
      },
      "minecraft:purple_bed": {
        "protocol_id": 664
      },
      "minecraft:blue_bed": {
        "protocol_id": 665
      },
      "minecraft:brown_bed": {
        "protocol_id": 666
      },
      "minecraft:green_bed": {
        "protocol_id": 667
      },
      "minecraft:red_bed": {
        "protocol_id": 668
      },
      "minecraft:black_bed": {
        "protocol_id": 669
      },
      "minecraft:cookie": {
        "protocol_id": 670
      },
      "minecraft:filled_map": {
        "protocol_id": 671
      },
      "minecraft:shears": {
        "protocol_id": 672
      },
      "minecraft:melon_slice": {
        "protocol_id": 673
      },
      "minecraft:dried_kelp": {
        "protocol_id": 674
      },
      "minecraft:pumpkin_seeds": {
        "protocol_id": 675
      },
      "minecraft:melon_seeds": {
        "protocol_id": 676
      },
      "minecraft:beef": {
        "protocol_id": 677
      },
      "minecraft:cooked_beef": {
        "protocol_id": 678
      },
      "minecraft:chicken": {
        "protocol_id": 679
      },
      "minecraft:cooked_chicken": {
        "protocol_id": 680
      },
      "minecraft:rotten_flesh": {
        "protocol_id": 681
      },
      "minecraft:ender_pearl": {
        "protocol_id": 682
      },
      "minecraft:blaze_rod": {
        "protocol_id": 683
      },
      "minecraft:ghast_tear": {
        "protocol_id": 684
      },
      "minecraft:gold_nugget": {
        "protocol_id": 685
      },
      "minecraft:nether_wart": {
        "protocol_id": 686
      },
      "minecraft:potion": {
        "protocol_id": 687
      },
      "minecraft:glass_bottle": {
        "protocol_id": 688
      },
      "minecraft:spider_eye": {
        "protocol_id": 689
      },
      "minecraft:fermented_spider_eye": {
        "protocol_id": 690
      },
      "minecraft:blaze_powder": {
        "protocol_id": 691
      },
      "minecraft:magma_cream": {
        "protocol_id": 692
      },
      "minecraft:brewing_stand": {
        "protocol_id": 693
      },
      "minecraft:cauldron": {
        "protocol_id": 694
      },
      "minecraft:ender_eye": {
        "protocol_id": 695
      },
      "minecraft:glistering_melon_slice": {
        "protocol_id": 696
      },
      "minecraft:bat_spawn_egg": {
        "protocol_id": 697
      },
      "minecraft:bee_spawn_egg": {
        "protocol_id": 698
      },
      "minecraft:blaze_spawn_egg": {
        "protocol_id": 699
      },
      "minecraft:cat_spawn_egg": {
        "protocol_id": 700
      },
      "minecraft:cave_spider_spawn_egg": {
        "protocol_id": 701
      },
      "minecraft:chicken_spawn_egg": {
        "protocol_id": 702
      },
      "minecraft:cod_spawn_egg": {
        "protocol_id": 703
      },
      "minecraft:cow_spawn_egg": {
        "protocol_id": 704
      },
      "minecraft:creeper_spawn_egg": {
        "protocol_id": 705
      },
      "minecraft:dolphin_spawn_egg": {
        "protocol_id": 706
      },
      "minecraft:donkey_spawn_egg": {
        "protocol_id": 707
      },
      "minecraft:drowned_spawn_egg": {
        "protocol_id": 708
      },
      "minecraft:elder_guardian_spawn_egg": {
        "protocol_id": 709
      },
      "minecraft:enderman_spawn_egg": {
        "protocol_id": 710
      },
      "minecraft:endermite_spawn_egg": {
        "protocol_id": 711
      },
      "minecraft:evoker_spawn_egg": {
        "protocol_id": 712
      },
      "minecraft:fox_spawn_egg": {
        "protocol_id": 713
      },
      "minecraft:ghast_spawn_egg": {
        "protocol_id": 714
      },
      "minecraft:guardian_spawn_egg": {
        "protocol_id": 715
      },
      "minecraft:horse_spawn_egg": {
        "protocol_id": 716
      },
      "minecraft:husk_spawn_egg": {
        "protocol_id": 717
      },
      "minecraft:llama_spawn_egg": {
        "protocol_id": 718
      },
      "minecraft:magma_cube_spawn_egg": {
        "protocol_id": 719
      },
      "minecraft:mooshroom_spawn_egg": {
        "protocol_id": 720
      },
      "minecraft:mule_spawn_egg": {
        "protocol_id": 721
      },
      "minecraft:ocelot_spawn_egg": {
        "protocol_id": 722
      },
      "minecraft:panda_spawn_egg": {
        "protocol_id": 723
      },
      "minecraft:parrot_spawn_egg": {
        "protocol_id": 724
      },
      "minecraft:phantom_spawn_egg": {
        "protocol_id": 725
      },
      "minecraft:pig_spawn_egg": {
        "protocol_id": 726
      },
      "minecraft:pillager_spawn_egg": {
        "protocol_id": 727
      },
      "minecraft:polar_bear_spawn_egg": {
        "protocol_id": 728
      },
      "minecraft:pufferfish_spawn_egg": {
        "protocol_id": 729
      },
      "minecraft:rabbit_spawn_egg": {
        "protocol_id": 730
      },
      "minecraft:ravager_spawn_egg": {
        "protocol_id": 731
      },
      "minecraft:salmon_spawn_egg": {
        "protocol_id": 732
      },
      "minecraft:sheep_spawn_egg": {
        "protocol_id": 733
      },
      "minecraft:shulker_spawn_egg": {
        "protocol_id": 734
      },
      "minecraft:silverfish_spawn_egg": {
        "protocol_id": 735
      },
      "minecraft:skeleton_spawn_egg": {
        "protocol_id": 736
      },
      "minecraft:skeleton_horse_spawn_egg": {
        "protocol_id": 737
      },
      "minecraft:slime_spawn_egg": {
        "protocol_id": 738
      },
      "minecraft:spider_spawn_egg": {
        "protocol_id": 739
      },
      "minecraft:squid_spawn_egg": {
        "protocol_id": 740
      },
      "minecraft:stray_spawn_egg": {
        "protocol_id": 741
      },
      "minecraft:trader_llama_spawn_egg": {
        "protocol_id": 742
      },
      "minecraft:tropical_fish_spawn_egg": {
        "protocol_id": 743
      },
      "minecraft:turtle_spawn_egg": {
        "protocol_id": 744
      },
      "minecraft:vex_spawn_egg": {
        "protocol_id": 745
      },
      "minecraft:villager_spawn_egg": {
        "protocol_id": 746
      },
      "minecraft:vindicator_spawn_egg": {
        "protocol_id": 747
      },
      "minecraft:wandering_trader_spawn_egg": {
        "protocol_id": 748
      },
      "minecraft:witch_spawn_egg": {
        "protocol_id": 749
      },
      "minecraft:wither_skeleton_spawn_egg": {
        "protocol_id": 750
      },
      "minecraft:wolf_spawn_egg": {
        "protocol_id": 751
      },
      "minecraft:zombie_spawn_egg": {
        "protocol_id": 752
      },
      "minecraft:zombie_horse_spawn_egg": {
        "protocol_id": 753
      },
      "minecraft:zombie_pigman_spawn_egg": {
        "protocol_id": 754
      },
      "minecraft:zombie_villager_spawn_egg": {
        "protocol_id": 755
      },
      "minecraft:experience_bottle": {
        "protocol_id": 756
      },
      "minecraft:fire_charge": {
        "protocol_id": 757
      },
      "minecraft:writable_book": {
        "protocol_id": 758
      },
      "minecraft:written_book": {
        "protocol_id": 759
      },
      "minecraft:emerald": {
        "protocol_id": 760
      },
      "minecraft:item_frame": {
        "protocol_id": 761
      },
      "minecraft:flower_pot": {
        "protocol_id": 762
      },
      "minecraft:carrot": {
        "protocol_id": 763
      },
      "minecraft:potato": {
        "protocol_id": 764
      },
      "minecraft:baked_potato": {
        "protocol_id": 765
      },
      "minecraft:poisonous_potato": {
        "protocol_id": 766
      },
      "minecraft:map": {
        "protocol_id": 767
      },
      "minecraft:golden_carrot": {
        "protocol_id": 768
      },
      "minecraft:skeleton_skull": {
        "protocol_id": 769
      },
      "minecraft:wither_skeleton_skull": {
        "protocol_id": 770
      },
      "minecraft:player_head": {
        "protocol_id": 771
      },
      "minecraft:zombie_head": {
        "protocol_id": 772
      },
      "minecraft:creeper_head": {
        "protocol_id": 773
      },
      "minecraft:dragon_head": {
        "protocol_id": 774
      },
      "minecraft:carrot_on_a_stick": {
        "protocol_id": 775
      },
      "minecraft:nether_star": {
        "protocol_id": 776
      },
      "minecraft:pumpkin_pie": {
        "protocol_id": 777
      },
      "minecraft:firework_rocket": {
        "protocol_id": 778
      },
      "minecraft:firework_star": {
        "protocol_id": 779
      },
      "minecraft:enchanted_book": {
        "protocol_id": 780
      },
      "minecraft:nether_brick": {
        "protocol_id": 781
      },
      "minecraft:quartz": {
        "protocol_id": 782
      },
      "minecraft:tnt_minecart": {
        "protocol_id": 783
      },
      "minecraft:hopper_minecart": {
        "protocol_id": 784
      },
      "minecraft:prismarine_shard": {
        "protocol_id": 785
      },
      "minecraft:prismarine_crystals": {
        "protocol_id": 786
      },
      "minecraft:rabbit": {
        "protocol_id": 787
      },
      "minecraft:cooked_rabbit": {
        "protocol_id": 788
      },
      "minecraft:rabbit_stew": {
        "protocol_id": 789
      },
      "minecraft:rabbit_foot": {
        "protocol_id": 790
      },
      "minecraft:rabbit_hide": {
        "protocol_id": 791
      },
      "minecraft:armor_stand": {
        "protocol_id": 792
      },
      "minecraft:iron_horse_armor": {
        "protocol_id": 793
      },
      "minecraft:golden_horse_armor": {
        "protocol_id": 794
      },
      "minecraft:diamond_horse_armor": {
        "protocol_id": 795
      },
      "minecraft:leather_horse_armor": {
        "protocol_id": 796
      },
      "minecraft:lead": {
        "protocol_id": 797
      },
      "minecraft:name_tag": {
        "protocol_id": 798
      },
      "minecraft:command_block_minecart": {
        "protocol_id": 799
      },
      "minecraft:mutton": {
        "protocol_id": 800
      },
      "minecraft:cooked_mutton": {
        "protocol_id": 801
      },
      "minecraft:white_banner": {
        "protocol_id": 802
      },
      "minecraft:orange_banner": {
        "protocol_id": 803
      },
      "minecraft:magenta_banner": {
        "protocol_id": 804
      },
      "minecraft:light_blue_banner": {
        "protocol_id": 805
      },
      "minecraft:yellow_banner": {
        "protocol_id": 806
      },
      "minecraft:lime_banner": {
        "protocol_id": 807
      },
      "minecraft:pink_banner": {
        "protocol_id": 808
      },
      "minecraft:gray_banner": {
        "protocol_id": 809
      },
      "minecraft:light_gray_banner": {
        "protocol_id": 810
      },
      "minecraft:cyan_banner": {
        "protocol_id": 811
      },
      "minecraft:purple_banner": {
        "protocol_id": 812
      },
      "minecraft:blue_banner": {
        "protocol_id": 813
      },
      "minecraft:brown_banner": {
        "protocol_id": 814
      },
      "minecraft:green_banner": {
        "protocol_id": 815
      },
      "minecraft:red_banner": {
        "protocol_id": 816
      },
      "minecraft:black_banner": {
        "protocol_id": 817
      },
      "minecraft:end_crystal": {
        "protocol_id": 818
      },
      "minecraft:chorus_fruit": {
        "protocol_id": 819
      },
      "minecraft:popped_chorus_fruit": {
        "protocol_id": 820
      },
      "minecraft:beetroot": {
        "protocol_id": 821
      },
      "minecraft:beetroot_seeds": {
        "protocol_id": 822
      },
      "minecraft:beetroot_soup": {
        "protocol_id": 823
      },
      "minecraft:dragon_breath": {
        "protocol_id": 824
      },
      "minecraft:splash_potion": {
        "protocol_id": 825
      },
      "minecraft:spectral_arrow": {
        "protocol_id": 826
      },
      "minecraft:tipped_arrow": {
        "protocol_id": 827
      },
      "minecraft:lingering_potion": {
        "protocol_id": 828
      },
      "minecraft:shield": {
        "protocol_id": 829
      },
      "minecraft:elytra": {
        "protocol_id": 830
      },
      "minecraft:spruce_boat": {
        "protocol_id": 831
      },
      "minecraft:birch_boat": {
        "protocol_id": 832
      },
      "minecraft:jungle_boat": {
        "protocol_id": 833
      },
      "minecraft:acacia_boat": {
        "protocol_id": 834
      },
      "minecraft:dark_oak_boat": {
        "protocol_id": 835
      },
      "minecraft:totem_of_undying": {
        "protocol_id": 836
      },
      "minecraft:shulker_shell": {
        "protocol_id": 837
      },
      "minecraft:iron_nugget": {
        "protocol_id": 838
      },
      "minecraft:knowledge_book": {
        "protocol_id": 839
      },
      "minecraft:debug_stick": {
        "protocol_id": 840
      },
      "minecraft:music_disc_13": {
        "protocol_id": 841
      },
      "minecraft:music_disc_cat": {
        "protocol_id": 842
      },
      "minecraft:music_disc_blocks": {
        "protocol_id": 843
      },
      "minecraft:music_disc_chirp": {
        "protocol_id": 844
      },
      "minecraft:music_disc_far": {
        "protocol_id": 845
      },
      "minecraft:music_disc_mall": {
        "protocol_id": 846
      },
      "minecraft:music_disc_mellohi": {
        "protocol_id": 847
      },
      "minecraft:music_disc_stal": {
        "protocol_id": 848
      },
      "minecraft:music_disc_strad": {
        "protocol_id": 849
      },
      "minecraft:music_disc_ward": {
        "protocol_id": 850
      },
      "minecraft:music_disc_11": {
        "protocol_id": 851
      },
      "minecraft:music_disc_wait": {
        "protocol_id": 852
      },
      "minecraft:trident": {
        "protocol_id": 853
      },
      "minecraft:phantom_membrane": {
        "protocol_id": 854
      },
      "minecraft:nautilus_shell": {
        "protocol_id": 855
      },
      "minecraft:heart_of_the_sea": {
        "protocol_id": 856
      },
      "minecraft:crossbow": {
        "protocol_id": 857
      },
      "minecraft:suspicious_stew": {
        "protocol_id": 858
      },
      "minecraft:loom": {
        "protocol_id": 859
      },
      "minecraft:flower_banner_pattern": {
        "protocol_id": 860
      },
      "minecraft:creeper_banner_pattern": {
        "protocol_id": 861
      },
      "minecraft:skull_banner_pattern": {
        "protocol_id": 862
      },
      "minecraft:mojang_banner_pattern": {
        "protocol_id": 863
      },
      "minecraft:globe_banner_pattern": {
        "protocol_id": 864
      },
      "minecraft:barrel": {
        "protocol_id": 865
      },
      "minecraft:smoker": {
        "protocol_id": 866
      },
      "minecraft:blast_furnace": {
        "protocol_id": 867
      },
      "minecraft:cartography_table": {
        "protocol_id": 868
      },
      "minecraft:fletching_table": {
        "protocol_id": 869
      },
      "minecraft:grindstone": {
        "protocol_id": 870
      },
      "minecraft:lectern": {
        "protocol_id": 871
      },
      "minecraft:smithing_table": {
        "protocol_id": 872
      },
      "minecraft:stonecutter": {
        "protocol_id": 873
      },
      "minecraft:bell": {
        "protocol_id": 874
      },
      "minecraft:lantern": {
        "protocol_id": 875
      },
      "minecraft:sweet_berries": {
        "protocol_id": 876
      },
      "minecraft:campfire": {
        "protocol_id": 877
      },
      "minecraft:honeycomb": {
        "protocol_id": 878
      },
      "minecraft:bee_nest": {
        "protocol_id": 879
      },
      "minecraft:beehive": {
        "protocol_id": 880
      },
      "minecraft:honey_bottle": {
        "protocol_id": 881
      },
      "minecraft:honey_block": {
        "protocol_id": 882
      },
      "minecraft:honeycomb_block": {
        "protocol_id": 883
      }
}`
