// line 1 "os_fsm.rl"
package normalize

// line 7 "os_fsm.go"
const os_ios_start int = 0
const os_ios_first_final int = 8
const os_ios_error int = -1

const os_ios_en_main int = 0

// line 6 "os_fsm.rl"

func MatchOsiOS(data string) bool {
	cs, p, pe := 0, 0, len(data)

	// line 21 "os_fsm.go"
	{
		cs = os_ios_start
	}

	// line 26 "os_fsm.go"
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
		case 8:
			goto st_case_8
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
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
		case 79:
			goto st2
		case 80:
			goto st3
		case 105:
			goto st1
		case 111:
			goto st2
		case 112:
			goto st3
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
		case 83:
			goto tr4
		case 105:
			goto st1
		case 115:
			goto tr4
		}
		goto st0
	tr4:
		// line 11 "os_fsm.rl"

		return true
		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		// line 110 "os_fsm.go"
		switch data[p] {
		case 73:
			goto st1
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
		case 65:
			goto st4
		case 72:
			goto st5
		case 73:
			goto st1
		case 79:
			goto st4
		case 97:
			goto st4
		case 104:
			goto st5
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
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 73:
			goto st1
		case 79:
			goto st6
		case 105:
			goto st1
		case 111:
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 73:
			goto st1
		case 78:
			goto st7
		case 105:
			goto st1
		case 110:
			goto st7
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 69:
			goto tr4
		case 73:
			goto st1
		case 101:
			goto tr4
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
	_test_eof8:
		cs = 8
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
	_test_eof7:
		cs = 7
		goto _test_eof

	_test_eof:
		{
		}
	}

	// line 14 "os_fsm.rl"

	return false
}

// line 226 "os_fsm.go"
const os_android_start int = 0
const os_android_first_final int = 7
const os_android_error int = -1

const os_android_en_main int = 0

// line 21 "os_fsm.rl"

func MatchOsAndroid(data string) bool {
	cs, p, pe := 0, 0, len(data)

	// line 240 "os_fsm.go"
	{
		cs = os_android_start
	}

	// line 245 "os_fsm.go"
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
		case 7:
			goto st_case_7
		}
		goto st_out
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 65:
			goto st1
		case 97:
			goto st1
		}
		goto st0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 65:
			goto st1
		case 78:
			goto st2
		case 97:
			goto st1
		case 110:
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
			goto st1
		case 68:
			goto st3
		case 97:
			goto st1
		case 100:
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 65:
			goto st1
		case 82:
			goto st4
		case 97:
			goto st1
		case 114:
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 65:
			goto st1
		case 79:
			goto st5
		case 97:
			goto st1
		case 111:
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 65:
			goto st1
		case 73:
			goto st6
		case 97:
			goto st1
		case 105:
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 65:
			goto st1
		case 68:
			goto tr7
		case 97:
			goto st1
		case 100:
			goto tr7
		}
		goto st0
	tr7:
		// line 26 "os_fsm.rl"

		return true
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		// line 387 "os_fsm.go"
		switch data[p] {
		case 65:
			goto st1
		case 97:
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
	_test_eof7:
		cs = 7
		goto _test_eof

	_test_eof:
		{
		}
	}

	// line 29 "os_fsm.rl"

	return false
}

// vim: ft=go ts=4 sw=4 noet
