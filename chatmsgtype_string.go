// Code generated by "stringer -linecomment -type ChatMsgType"; DO NOT EDIT.

package mt

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RawMsg-0]
	_ = x[NormalMsg-1]
	_ = x[AnnounceMsg-2]
	_ = x[SysMsg-3]
	_ = x[maxMsg-4]
}

const _ChatMsgType_name = "rawnormalannouncesysmaxMsg"

var _ChatMsgType_index = [...]uint8{0, 3, 9, 17, 20, 26}

func (i ChatMsgType) String() string {
	if i >= ChatMsgType(len(_ChatMsgType_index)-1) {
		return "ChatMsgType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ChatMsgType_name[_ChatMsgType_index[i]:_ChatMsgType_index[i+1]]
}
