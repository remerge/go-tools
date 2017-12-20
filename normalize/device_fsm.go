// line 1 "device_fsm.rl"
package normalize

// line 7 "device_fsm.go"
const device_ipod_start int = 0
const device_ipod_first_final int = 4
const device_ipod_error int = -1

const device_ipod_en_main int = 0

// line 6 "device_fsm.rl"

func MatchDeviceiPod(data string) bool {
	cs, p, pe := 0, 0, len(data)

	// line 21 "device_fsm.go"
	{
		cs = device_ipod_start
	}

	// line 26 "device_fsm.go"
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 0:
			goto st_case_0
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		}
		goto st_out
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 73:
			goto st1
		case 105:
			goto st1
		}
		goto st0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 73:
			goto st1
		case 80:
			goto st2
		case 105:
			goto st1
		case 112:
			goto st2
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 73:
			goto st1
		case 79:
			goto st3
		case 105:
			goto st1
		case 111:
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 68:
			goto tr4
		case 73:
			goto st1
		case 100:
			goto tr4
		case 105:
			goto st1
		}
		goto st0
	tr4:
		// line 11 "device_fsm.rl"

		return true
		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		// line 114 "device_fsm.go"
		switch data[p] {
		case 73:
			goto st1
		case 105:
			goto st1
		}
		goto st0
	st_out:
	_test_eof0:
		cs = 0
		goto _test_eof
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof

	_test_eof:
		{
		}
	}

	// line 14 "device_fsm.rl"

	return false
}

// line 138 "device_fsm.go"
const device_ipad_start int = 0
const device_ipad_first_final int = 4
const device_ipad_error int = -1

const device_ipad_en_main int = 0

// line 21 "device_fsm.rl"

func MatchDeviceiPad(data string) bool {
	cs, p, pe := 0, 0, len(data)

	// line 152 "device_fsm.go"
	{
		cs = device_ipad_start
	}

	// line 157 "device_fsm.go"
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 0:
			goto st_case_0
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		}
		goto st_out
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 73:
			goto st1
		case 105:
			goto st1
		}
		goto st0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 73:
			goto st1
		case 80:
			goto st2
		case 105:
			goto st1
		case 112:
			goto st2
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 65:
			goto st3
		case 73:
			goto st1
		case 97:
			goto st3
		case 105:
			goto st1
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 68:
			goto tr4
		case 73:
			goto st1
		case 100:
			goto tr4
		case 105:
			goto st1
		}
		goto st0
	tr4:
		// line 26 "device_fsm.rl"

		return true
		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		// line 245 "device_fsm.go"
		switch data[p] {
		case 73:
			goto st1
		case 105:
			goto st1
		}
		goto st0
	st_out:
	_test_eof0:
		cs = 0
		goto _test_eof
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof

	_test_eof:
		{
		}
	}

	// line 29 "device_fsm.rl"

	return false
}

// line 269 "device_fsm.go"
const device_iphone_start int = 0
const device_iphone_first_final int = 6
const device_iphone_error int = -1

const device_iphone_en_main int = 0

// line 36 "device_fsm.rl"

func MatchDeviceiPhone(data string) bool {
	cs, p, pe := 0, 0, len(data)

	// line 283 "device_fsm.go"
	{
		cs = device_iphone_start
	}

	// line 288 "device_fsm.go"
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 0:
			goto st_case_0
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		}
		goto st_out
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 73:
			goto st1
		case 105:
			goto st1
		}
		goto st0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 73:
			goto st1
		case 80:
			goto st2
		case 105:
			goto st1
		case 112:
			goto st2
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 72:
			goto st3
		case 73:
			goto st1
		case 104:
			goto st3
		case 105:
			goto st1
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 73:
			goto st1
		case 79:
			goto st4
		case 105:
			goto st1
		case 111:
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 73:
			goto st1
		case 78:
			goto st5
		case 105:
			goto st1
		case 110:
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 69:
			goto tr6
		case 73:
			goto st1
		case 101:
			goto tr6
		case 105:
			goto st1
		}
		goto st0
	tr6:
		// line 41 "device_fsm.rl"

		return true
		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		// line 412 "device_fsm.go"
		switch data[p] {
		case 73:
			goto st1
		case 105:
			goto st1
		}
		goto st0
	st_out:
	_test_eof0:
		cs = 0
		goto _test_eof
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof

	_test_eof:
		{
		}
	}

	// line 44 "device_fsm.rl"

	return false
}

// vim: ft=go ts=4 sw=4 noet
