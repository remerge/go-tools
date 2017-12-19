// line 1 "version_fsm.rl"
package version

// line 7 "version_fsm.go"
var _version_parser_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2,
}

var _version_parser_key_offsets []byte = []byte{
	0, 0, 2, 6, 8, 12, 14,
}

var _version_parser_trans_keys []byte = []byte{
	48, 57, 46, 95, 48, 57, 48, 57,
	46, 95, 48, 57, 48, 57, 48, 57,
}

var _version_parser_single_lengths []byte = []byte{
	0, 0, 2, 0, 2, 0, 0,
}

var _version_parser_range_lengths []byte = []byte{
	0, 1, 1, 1, 1, 1, 1,
}

var _version_parser_index_offsets []byte = []byte{
	0, 0, 2, 6, 8, 12, 14,
}

var _version_parser_indicies []byte = []byte{
	1, 0, 2, 2, 1, 3, 4, 3,
	5, 5, 4, 3, 6, 3, 6, 3,
}

var _version_parser_trans_targs []byte = []byte{
	1, 2, 3, 0, 4, 5, 6,
}

var _version_parser_trans_actions []byte = []byte{
	0, 1, 0, 0, 3, 0, 5,
}

const version_parser_start int = 1
const version_parser_first_final int = 6
const version_parser_error int = 0

const version_parser_en_main int = 1

// line 6 "version_fsm.rl"

func Parse(data []byte) *Version {
	version := &Version{}

	if len(data) < 1 {
		return version
	}

	cs, p, pe := 0, 0, len(data)

	// line 68 "version_fsm.go"
	{
		cs = version_parser_start
	}

	// line 73 "version_fsm.go"
	{
		var _klen int
		var _trans int
		var _acts int
		var _nacts uint
		var _keys int
		if p == pe {
			goto _test_eof
		}
		if cs == 0 {
			goto _out
		}
	_resume:
		_keys = int(_version_parser_key_offsets[cs])
		_trans = int(_version_parser_index_offsets[cs])

		_klen = int(_version_parser_single_lengths[cs])
		if _klen > 0 {
			_lower := int(_keys)
			var _mid int
			_upper := int(_keys + _klen - 1)
			for {
				if _upper < _lower {
					break
				}

				_mid = _lower + ((_upper - _lower) >> 1)
				switch {
				case data[p] < _version_parser_trans_keys[_mid]:
					_upper = _mid - 1
				case data[p] > _version_parser_trans_keys[_mid]:
					_lower = _mid + 1
				default:
					_trans += int(_mid - int(_keys))
					goto _match
				}
			}
			_keys += _klen
			_trans += _klen
		}

		_klen = int(_version_parser_range_lengths[cs])
		if _klen > 0 {
			_lower := int(_keys)
			var _mid int
			_upper := int(_keys + (_klen << 1) - 2)
			for {
				if _upper < _lower {
					break
				}

				_mid = _lower + (((_upper - _lower) >> 1) & ^1)
				switch {
				case data[p] < _version_parser_trans_keys[_mid]:
					_upper = _mid - 2
				case data[p] > _version_parser_trans_keys[_mid+1]:
					_lower = _mid + 2
				default:
					_trans += int((_mid - int(_keys)) >> 1)
					goto _match
				}
			}
			_trans += _klen
		}

	_match:
		_trans = int(_version_parser_indicies[_trans])
		cs = int(_version_parser_trans_targs[_trans])

		if _version_parser_trans_actions[_trans] == 0 {
			goto _again
		}

		_acts = int(_version_parser_trans_actions[_trans])
		_nacts = uint(_version_parser_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _version_parser_actions[_acts-1] {
			case 0:
				// line 18 "version_fsm.rl"

				version.Major = version.Major*10 + (int(data[p]) - '0')
			case 1:
				// line 19 "version_fsm.rl"

				version.Minor = version.Minor*10 + (int(data[p]) - '0')
			case 2:
				// line 20 "version_fsm.rl"

				version.Patch = version.Patch*10 + (int(data[p]) - '0')
				// line 164 "version_fsm.go"
			}
		}

	_again:
		if cs == 0 {
			goto _out
		}
		p++
		if p != pe {
			goto _resume
		}
	_test_eof:
		{
		}
	_out:
		{
		}
	}

	// line 24 "version_fsm.rl"

	// remove some insane versions
	if version.Major > 1000000 || version.Major < 0 {
		version.Major = 1000000
	}
	if version.Minor > 1000000 || version.Minor < 0 {
		version.Minor = 1000000
	}
	if version.Patch > 1000000 || version.Patch < 0 {
		version.Patch = 1000000
	}

	return version
}

// vim: ft=go ts=4 sw=4 noet
