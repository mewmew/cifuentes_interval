// Code generated by "stringer -linecomment -type LoopType"; DO NOT EDIT.

package interval

import "strconv"

const _LoopType_name = "LoopTypeNonepre_looppost_loopinf_loop"

var _LoopType_index = [...]uint8{0, 12, 20, 29, 37}

func (i LoopType) String() string {
	if i >= LoopType(len(_LoopType_index)-1) {
		return "LoopType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LoopType_name[_LoopType_index[i]:_LoopType_index[i+1]]
}
