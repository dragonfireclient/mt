// Code generated by "stringer -type WaveType"; DO NOT EDIT.

package mt

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NotWaving-0]
	_ = x[PlantWaving-1]
	_ = x[LeafWaving-2]
	_ = x[LiquidWaving-3]
}

const _WaveType_name = "NotWavingPlantWavingLeafWavingLiquidWaving"

var _WaveType_index = [...]uint8{0, 9, 20, 30, 42}

func (i WaveType) String() string {
	if i >= WaveType(len(_WaveType_index)-1) {
		return "WaveType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _WaveType_name[_WaveType_index[i]:_WaveType_index[i+1]]
}