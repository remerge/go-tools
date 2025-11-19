//line os_fsm.rl:1
package normalize

//line os_fsm.go:7
const os_ios_start int = 0
const os_ios_first_final int = 8
const os_ios_error int = -1

const os_ios_en_main int = 0

//line os_fsm.rl:6

func MatchOsiOS(data string) bool {
	cs, p, pe := 0, 0, len(data)

//line os_fsm.go:21
	{
		cs = os_ios_start
	}

//line os_fsm.go:26
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
//line os_fsm.rl:11
		return true
		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line os_fsm.go:109
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

//line os_fsm.rl:14

	return false
}

//line os_fsm.go:225
const os_android_start int = 0
const os_android_first_final int = 7
const os_android_error int = -1

const os_android_en_main int = 0

//line os_fsm.rl:21

func MatchOsAndroid(data string) bool {
	cs, p, pe := 0, 0, len(data)

//line os_fsm.go:239
	{
		cs = os_android_start
	}

//line os_fsm.go:244
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
//line os_fsm.rl:26
		return true
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line os_fsm.go:385
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

//line os_fsm.rl:29

	return false
}

//line os_fsm.go:412
const os_tvos_start int = 0
const os_tvos_first_final int = 4
const os_tvos_error int = -1

const os_tvos_en_main int = 0

//line os_fsm.rl:36

func MatchOsTvOS(data string) bool {
	cs, p, pe := 0, 0, len(data)

//line os_fsm.go:426
	{
		cs = os_tvos_start
	}

//line os_fsm.go:431
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
		case 84:
			goto st1
		case 116:
			goto st1
		}
		goto st0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 84:
			goto st1
		case 86:
			goto st2
		case 116:
			goto st1
		case 118:
			goto st2
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 79:
			goto st3
		case 84:
			goto st1
		case 111:
			goto st3
		case 116:
			goto st1
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 83:
			goto tr4
		case 84:
			goto st1
		case 115:
			goto tr4
		case 116:
			goto st1
		}
		goto st0
	tr4:
//line os_fsm.rl:41
		return true
		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line os_fsm.go:518
		switch data[p] {
		case 84:
			goto st1
		case 116:
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

//line os_fsm.rl:44

	return false
}

//line os_fsm.go:542
const os_roku_start int = 0
const os_roku_first_final int = 6
const os_roku_error int = -1

const os_roku_en_main int = 0

//line os_fsm.rl:51

func MatchOsRoku(data string) bool {
	cs, p, pe := 0, 0, len(data)

//line os_fsm.go:556
	{
		cs = os_roku_start
	}

//line os_fsm.go:561
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
		case 6:
			goto st_case_6
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
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
		case 82:
			goto st1
		case 114:
			goto st1
		}
		goto st0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 79:
			goto st2
		case 82:
			goto st1
		case 111:
			goto st2
		case 114:
			goto st1
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 75:
			goto st3
		case 82:
			goto st1
		case 107:
			goto st3
		case 114:
			goto st1
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 82:
			goto st1
		case 85:
			goto tr4
		case 114:
			goto st1
		case 117:
			goto tr4
		}
		goto st0
	tr4:
//line os_fsm.rl:56
		return true
		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line os_fsm.go:654
		switch data[p] {
		case 32:
			goto st4
		case 43:
			goto st4
		case 79:
			goto st5
		case 82:
			goto st1
		case 111:
			goto st5
		case 114:
			goto st1
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 79:
			goto st5
		case 82:
			goto st1
		case 111:
			goto st5
		case 114:
			goto st1
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 82:
			goto st1
		case 83:
			goto tr6
		case 114:
			goto st1
		case 115:
			goto tr6
		}
		goto st0
	tr6:
//line os_fsm.rl:56
		return true
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line os_fsm.go:714
		switch data[p] {
		case 82:
			goto st1
		case 114:
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
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof

	_test_eof:
		{
		}
	}

//line os_fsm.rl:59

	return false
}

//line os_fsm.go:741
const os_fireos_start int = 0
const os_fireos_first_final int = 7
const os_fireos_error int = -1

const os_fireos_en_main int = 0

//line os_fsm.rl:66

func MatchOsFireOS(data string) bool {
	cs, p, pe := 0, 0, len(data)

//line os_fsm.go:755
	{
		cs = os_fireos_start
	}

//line os_fsm.go:760
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
		case 7:
			goto st_case_7
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 8:
			goto st_case_8
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
		case 70:
			goto st1
		case 102:
			goto st1
		}
		goto st0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 70:
			goto st1
		case 73:
			goto st2
		case 102:
			goto st1
		case 105:
			goto st2
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 70:
			goto st1
		case 82:
			goto st3
		case 102:
			goto st1
		case 114:
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 69:
			goto tr4
		case 70:
			goto st1
		case 101:
			goto tr4
		case 102:
			goto st1
		}
		goto st0
	tr4:
//line os_fsm.rl:71
		return true
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line os_fsm.go:855
		switch data[p] {
		case 43:
			goto st4
		case 45:
			goto st4
		case 70:
			goto st1
		case 79:
			goto st5
		case 84:
			goto st6
		case 102:
			goto st1
		case 111:
			goto st5
		case 116:
			goto st6
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 70:
			goto st1
		case 79:
			goto st5
		case 102:
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
		case 70:
			goto st1
		case 83:
			goto tr6
		case 102:
			goto st1
		case 115:
			goto tr6
		}
		goto st0
	tr6:
//line os_fsm.rl:71
		return true
		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line os_fsm.go:916
		switch data[p] {
		case 70:
			goto st1
		case 102:
			goto st1
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 70:
			goto st1
		case 86:
			goto st4
		case 102:
			goto st1
		case 118:
			goto st4
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
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof

	_test_eof:
		{
		}
	}

//line os_fsm.rl:74

	return false
}

//line os_fsm.go:960
const os_smartcast_start int = 0
const os_smartcast_first_final int = 13
const os_smartcast_error int = -1

const os_smartcast_en_main int = 0

//line os_fsm.rl:81

func MatchOsSmartCast(data string) bool {
	cs, p, pe := 0, 0, len(data)

//line os_fsm.go:974
	{
		cs = os_smartcast_start
	}

//line os_fsm.go:979
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
		case 8:
			goto st_case_8
		case 13:
			goto st_case_13
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 14:
			goto st_case_14
		}
		goto st_out
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 83:
			goto st1
		case 115:
			goto st1
		}
		goto st0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 77:
			goto st2
		case 83:
			goto st1
		case 109:
			goto st2
		case 115:
			goto st1
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
		case 83:
			goto st1
		case 97:
			goto st3
		case 115:
			goto st1
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 82:
			goto st4
		case 83:
			goto st1
		case 114:
			goto st4
		case 115:
			goto st1
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 83:
			goto st1
		case 84:
			goto st5
		case 115:
			goto st1
		case 116:
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 67:
			goto st6
		case 83:
			goto st1
		case 99:
			goto st6
		case 115:
			goto st1
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 65:
			goto st7
		case 83:
			goto st1
		case 97:
			goto st7
		case 115:
			goto st1
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 83:
			goto st8
		case 115:
			goto st8
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 77:
			goto st2
		case 83:
			goto st1
		case 84:
			goto tr9
		case 109:
			goto st2
		case 115:
			goto st1
		case 116:
			goto tr9
		}
		goto st0
	tr9:
//line os_fsm.rl:86
		return true
		goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line os_fsm.go:1166
		switch data[p] {
		case 37:
			goto st9
		case 79:
			goto st12
		case 83:
			goto st1
		case 111:
			goto st12
		case 115:
			goto st1
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 50:
			goto st10
		case 83:
			goto st1
		case 115:
			goto st1
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 48:
			goto st11
		case 83:
			goto st1
		case 115:
			goto st1
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 79:
			goto st12
		case 83:
			goto st1
		case 111:
			goto st12
		case 115:
			goto st1
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 83:
			goto tr13
		case 115:
			goto tr13
		}
		goto st0
	tr13:
//line os_fsm.rl:86
		return true
		goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line os_fsm.go:1245
		switch data[p] {
		case 77:
			goto st2
		case 83:
			goto st1
		case 109:
			goto st2
		case 115:
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
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof

	_test_eof:
		{
		}
	}

//line os_fsm.rl:89

	return false
}

//line os_fsm.go:1283
const os_webos_start int = 0
const os_webos_first_final int = 7
const os_webos_error int = -1

const os_webos_en_main int = 0

//line os_fsm.rl:96

func MatchOsWebOS(data string) bool {
	cs, p, pe := 0, 0, len(data)

//line os_fsm.go:1297
	{
		cs = os_webos_start
	}

//line os_fsm.go:1302
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
		case 7:
			goto st_case_7
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 8:
			goto st_case_8
		}
		goto st_out
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 87:
			goto st1
		case 119:
			goto st1
		}
		goto st0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 69:
			goto st2
		case 87:
			goto st1
		case 101:
			goto st2
		case 119:
			goto st1
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 66:
			goto st3
		case 87:
			goto st1
		case 98:
			goto st3
		case 119:
			goto st1
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 79:
			goto st4
		case 87:
			goto st1
		case 111:
			goto st4
		case 119:
			goto st1
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 83:
			goto tr5
		case 87:
			goto st1
		case 115:
			goto tr5
		case 119:
			goto st1
		}
		goto st0
	tr5:
//line os_fsm.rl:101
		return true
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line os_fsm.go:1413
		switch data[p] {
		case 32:
			goto st5
		case 43:
			goto st5
		case 84:
			goto st6
		case 87:
			goto st1
		case 116:
			goto st6
		case 119:
			goto st1
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 84:
			goto st6
		case 87:
			goto st1
		case 116:
			goto st6
		case 119:
			goto st1
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 86:
			goto tr7
		case 87:
			goto st1
		case 118:
			goto tr7
		case 119:
			goto st1
		}
		goto st0
	tr7:
//line os_fsm.rl:101
		return true
		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line os_fsm.go:1473
		switch data[p] {
		case 87:
			goto st1
		case 119:
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
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof

	_test_eof:
		{
		}
	}

//line os_fsm.rl:104

	return false
}

//line os_fsm.go:1501
const os_tizen_start int = 0
const os_tizen_first_final int = 6
const os_tizen_error int = -1

const os_tizen_en_main int = 0

//line os_fsm.rl:111

func MatchOsTizen(data string) bool {
	cs, p, pe := 0, 0, len(data)

//line os_fsm.go:1515
	{
		cs = os_tizen_start
	}

//line os_fsm.go:1520
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
		case 6:
			goto st_case_6
		case 5:
			goto st_case_5
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
		case 84:
			goto st1
		case 116:
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
			goto st2
		case 84:
			goto st1
		case 105:
			goto st2
		case 116:
			goto st1
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 84:
			goto st1
		case 90:
			goto st3
		case 116:
			goto st1
		case 122:
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 69:
			goto st4
		case 84:
			goto st1
		case 101:
			goto st4
		case 116:
			goto st1
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 78:
			goto tr5
		case 84:
			goto st1
		case 110:
			goto tr5
		case 116:
			goto st1
		}
		goto st0
	tr5:
//line os_fsm.rl:116
		return true
		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line os_fsm.go:1629
		switch data[p] {
		case 84:
			goto st5
		case 116:
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 73:
			goto st2
		case 84:
			goto st1
		case 86:
			goto tr6
		case 105:
			goto st2
		case 116:
			goto st1
		case 118:
			goto tr6
		}
		goto st0
	tr6:
//line os_fsm.rl:116
		return true
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line os_fsm.go:1666
		switch data[p] {
		case 84:
			goto st1
		case 116:
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
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof

	_test_eof:
		{
		}
	}

//line os_fsm.rl:119

	return false
}

//line os_fsm.go:1693
const os_linux_start int = 0
const os_linux_first_final int = 5
const os_linux_error int = -1

const os_linux_en_main int = 0

//line os_fsm.rl:126

func MatchOsLinux(data string) bool {
	cs, p, pe := 0, 0, len(data)

//line os_fsm.go:1707
	{
		cs = os_linux_start
	}

//line os_fsm.go:1712
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
		}
		goto st_out
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 76:
			goto st1
		case 108:
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
			goto st2
		case 76:
			goto st1
		case 105:
			goto st2
		case 108:
			goto st1
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 76:
			goto st1
		case 78:
			goto st3
		case 108:
			goto st1
		case 110:
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 76:
			goto st1
		case 85:
			goto st4
		case 108:
			goto st1
		case 117:
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 76:
			goto st1
		case 88:
			goto tr5
		case 108:
			goto st1
		case 120:
			goto tr5
		}
		goto st0
	tr5:
//line os_fsm.rl:131
		return true
		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line os_fsm.go:1817
		switch data[p] {
		case 76:
			goto st1
		case 108:
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

	_test_eof:
		{
		}
	}

//line os_fsm.rl:134

	return false
}

//line os_fsm.go:1842
const os_vidaa_start int = 0
const os_vidaa_first_final int = 5
const os_vidaa_error int = -1

const os_vidaa_en_main int = 0

//line os_fsm.rl:141

func MatchOsVidaa(data string) bool {
	cs, p, pe := 0, 0, len(data)

//line os_fsm.go:1856
	{
		cs = os_vidaa_start
	}

//line os_fsm.go:1861
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
		}
		goto st_out
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 86:
			goto st1
		case 118:
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
			goto st2
		case 86:
			goto st1
		case 105:
			goto st2
		case 118:
			goto st1
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 68:
			goto st3
		case 86:
			goto st1
		case 100:
			goto st3
		case 118:
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
		case 86:
			goto st1
		case 97:
			goto st4
		case 118:
			goto st1
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 65:
			goto tr5
		case 86:
			goto st1
		case 97:
			goto tr5
		case 118:
			goto st1
		}
		goto st0
	tr5:
//line os_fsm.rl:146
		return true
		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line os_fsm.go:1966
		switch data[p] {
		case 86:
			goto st1
		case 118:
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

	_test_eof:
		{
		}
	}

//line os_fsm.rl:149

	return false
}

//line os_fsm.go:1991
const os_reachtv_start int = 0
const os_reachtv_first_final int = 7
const os_reachtv_error int = -1

const os_reachtv_en_main int = 0

//line os_fsm.rl:156

func MatchOsReachTV(data string) bool {
	cs, p, pe := 0, 0, len(data)

//line os_fsm.go:2005
	{
		cs = os_reachtv_start
	}

//line os_fsm.go:2010
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
		case 82:
			goto st1
		case 114:
			goto st1
		}
		goto st0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 69:
			goto st2
		case 82:
			goto st1
		case 101:
			goto st2
		case 114:
			goto st1
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
		case 82:
			goto st1
		case 97:
			goto st3
		case 114:
			goto st1
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 67:
			goto st4
		case 82:
			goto st1
		case 99:
			goto st4
		case 114:
			goto st1
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 72:
			goto st5
		case 82:
			goto st1
		case 104:
			goto st5
		case 114:
			goto st1
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 82:
			goto st1
		case 84:
			goto st6
		case 114:
			goto st1
		case 116:
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 82:
			goto st1
		case 86:
			goto tr7
		case 114:
			goto st1
		case 118:
			goto tr7
		}
		goto st0
	tr7:
//line os_fsm.rl:161
		return true
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line os_fsm.go:2151
		switch data[p] {
		case 82:
			goto st1
		case 114:
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

//line os_fsm.rl:164

	return false
}

//line os_fsm.go:2178
const os_chromeos_start int = 0
const os_chromeos_first_final int = 8
const os_chromeos_error int = -1

const os_chromeos_en_main int = 0

//line os_fsm.rl:171

func MatchOsChromeOS(data string) bool {
	cs, p, pe := 0, 0, len(data)

//line os_fsm.go:2192
	{
		cs = os_chromeos_start
	}

//line os_fsm.go:2197
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
		case 8:
			goto st_case_8
		}
		goto st_out
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		switch data[p] {
		case 67:
			goto st1
		case 99:
			goto st1
		}
		goto st0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 67:
			goto st1
		case 72:
			goto st2
		case 99:
			goto st1
		case 104:
			goto st2
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 67:
			goto st1
		case 82:
			goto st3
		case 99:
			goto st1
		case 114:
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 67:
			goto st1
		case 79:
			goto st4
		case 99:
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
		case 67:
			goto st1
		case 77:
			goto st5
		case 99:
			goto st1
		case 109:
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 67:
			goto st1
		case 69:
			goto st6
		case 99:
			goto st1
		case 101:
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 67:
			goto st1
		case 79:
			goto st7
		case 99:
			goto st1
		case 111:
			goto st7
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 67:
			goto st1
		case 83:
			goto tr8
		case 99:
			goto st1
		case 115:
			goto tr8
		}
		goto st0
	tr8:
//line os_fsm.rl:176
		return true
		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line os_fsm.go:2356
		switch data[p] {
		case 67:
			goto st1
		case 99:
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
	_test_eof8:
		cs = 8
		goto _test_eof

	_test_eof:
		{
		}
	}

//line os_fsm.rl:179

	return false
}

// vim: ft=go ts=4 sw=4 noet
