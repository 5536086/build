// Code generated by "stringer -type=Record"; DO NOT EDIT.

package egoscale

import "strconv"

const _Record_name = "AAAAAALIASCNAMEHINFOMXNAPTRNSPOOLSPFSRVSSHFPTXTURL"

var _Record_index = [...]uint8{0, 1, 5, 10, 15, 20, 22, 27, 29, 33, 36, 39, 44, 47, 50}

func (i Record) String() string {
	if i < 0 || i >= Record(len(_Record_index)-1) {
		return "Record(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Record_name[_Record_index[i]:_Record_index[i+1]]
}
