// Code generated by "stringer -type TileFlags"; DO NOT EDIT.

package mt

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TileBackfaceCull-1]
	_ = x[TileAbleH-2]
	_ = x[TileAbleV-4]
	_ = x[TileColor-8]
	_ = x[TileScale-16]
	_ = x[TileAlign-32]
}

const (
	_TileFlags_name_0 = "TileBackfaceCullTileAbleH"
	_TileFlags_name_1 = "TileAbleV"
	_TileFlags_name_2 = "TileColor"
	_TileFlags_name_3 = "TileScale"
	_TileFlags_name_4 = "TileAlign"
)

var (
	_TileFlags_index_0 = [...]uint8{0, 16, 25}
)

func (i TileFlags) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _TileFlags_name_0[_TileFlags_index_0[i]:_TileFlags_index_0[i+1]]
	case i == 4:
		return _TileFlags_name_1
	case i == 8:
		return _TileFlags_name_2
	case i == 16:
		return _TileFlags_name_3
	case i == 32:
		return _TileFlags_name_4
	default:
		return "TileFlags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
