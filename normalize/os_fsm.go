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
const os_tv_start int = 0
const os_tv_first_final int = 40
const os_tv_error int = -1

const os_tv_en_main int = 0

//line os_fsm.rl:36

func MatchOsCTV(data string) bool {
	cs, p, pe := 0, 0, len(data)

//line os_fsm.go:426
	{
		cs = os_tv_start
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
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 40:
			goto st_case_40
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 41:
			goto st_case_41
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 42:
			goto st_case_42
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 43:
			goto st_case_43
		case 39:
			goto st_case_39
		case 44:
			goto st_case_44
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
		case 70:
			goto st2
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
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
		case 70:
			goto st2
		case 78:
			goto st13
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 110:
			goto st13
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
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
		case 70:
			goto st2
		case 73:
			goto st3
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 105:
			goto st3
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
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
		case 70:
			goto st2
		case 82:
			goto st4
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 114:
			goto st4
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
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
		case 69:
			goto st5
		case 70:
			goto st2
		case 79:
			goto st8
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 101:
			goto st5
		case 102:
			goto st2
		case 111:
			goto st8
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 13:
			goto st6
		case 32:
			goto st6
		case 65:
			goto st1
		case 70:
			goto st2
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st39
		case 86:
			goto st20
		case 87:
			goto st23
		case 95:
			goto st6
		case 97:
			goto st1
		case 102:
			goto st2
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st39
		case 118:
			goto st20
		case 119:
			goto st23
		}
		if 9 <= data[p] && data[p] <= 10 {
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
		case 70:
			goto st2
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st39
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st39
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 79:
			goto st8
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 111:
			goto st8
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 75:
			goto st9
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 107:
			goto st9
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 85:
			goto tr16
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 117:
			goto tr16
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 77:
			goto st11
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 109:
			goto st11
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 65:
			goto st12
		case 70:
			goto st2
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st12
		case 102:
			goto st2
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 78:
			goto st13
		case 82:
			goto st34
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 110:
			goto st13
		case 114:
			goto st34
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch data[p] {
		case 65:
			goto st1
		case 68:
			goto st14
		case 70:
			goto st2
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 100:
			goto st14
		case 102:
			goto st2
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 82:
			goto st15
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 114:
			goto st15
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 79:
			goto st16
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 111:
			goto st16
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 73:
			goto st17
		case 75:
			goto st9
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 105:
			goto st17
		case 107:
			goto st9
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch data[p] {
		case 65:
			goto st1
		case 68:
			goto st5
		case 70:
			goto st2
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 100:
			goto st5
		case 102:
			goto st2
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 73:
			goto st19
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st33
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 105:
			goto st19
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st33
		case 119:
			goto st23
		}
		goto st0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 90:
			goto st31
		case 97:
			goto st1
		case 102:
			goto st2
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		case 122:
			goto st31
		}
		goto st0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch data[p] {
		case 65:
			goto st1
		case 69:
			goto st21
		case 70:
			goto st2
		case 73:
			goto st28
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 101:
			goto st21
		case 102:
			goto st2
		case 105:
			goto st28
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st22
		case 97:
			goto st1
		case 102:
			goto st2
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st22
		}
		goto st0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch data[p] {
		case 65:
			goto st1
		case 68:
			goto tr16
		case 69:
			goto st24
		case 70:
			goto st2
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 100:
			goto tr16
		case 101:
			goto st24
		case 102:
			goto st2
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	tr16:
//line os_fsm.rl:52
		return true
		goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line os_fsm.go:1455
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		switch data[p] {
		case 65:
			goto st1
		case 69:
			goto st24
		case 70:
			goto st2
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 101:
			goto st24
		case 102:
			goto st2
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		switch data[p] {
		case 65:
			goto st1
		case 66:
			goto st25
		case 70:
			goto st2
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 98:
			goto st25
		case 102:
			goto st2
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch data[p] {
		case 13:
			goto st26
		case 32:
			goto st26
		case 65:
			goto st1
		case 70:
			goto st2
		case 79:
			goto st27
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 95:
			goto st26
		case 97:
			goto st1
		case 102:
			goto st2
		case 111:
			goto st27
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		if 9 <= data[p] && data[p] <= 10 {
			goto st26
		}
		goto st0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 79:
			goto st27
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 111:
			goto st27
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 82:
			goto st7
		case 83:
			goto tr34
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 114:
			goto st7
		case 115:
			goto tr34
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	tr34:
//line os_fsm.rl:52
		return true
		goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line os_fsm.go:1701
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 77:
			goto st11
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 109:
			goto st11
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 65:
			goto st1
		case 68:
			goto st29
		case 70:
			goto st2
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 100:
			goto st29
		case 102:
			goto st2
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch data[p] {
		case 65:
			goto st30
		case 70:
			goto st2
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st30
		case 102:
			goto st2
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 65:
			goto tr37
		case 70:
			goto st2
		case 78:
			goto st13
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto tr37
		case 102:
			goto st2
		case 110:
			goto st13
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	tr37:
//line os_fsm.rl:52
		return true
		goto st42
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
//line os_fsm.go:1862
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 78:
			goto st13
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 110:
			goto st13
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 65:
			goto st1
		case 69:
			goto st32
		case 70:
			goto st2
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 101:
			goto st32
		case 102:
			goto st2
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 78:
			goto tr16
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 110:
			goto tr16
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 13:
			goto st26
		case 32:
			goto st26
		case 65:
			goto st1
		case 69:
			goto st21
		case 70:
			goto st2
		case 73:
			goto st28
		case 79:
			goto st27
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 95:
			goto st26
		case 97:
			goto st1
		case 101:
			goto st21
		case 102:
			goto st2
		case 105:
			goto st28
		case 111:
			goto st27
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		if 9 <= data[p] && data[p] <= 10 {
			goto st26
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 79:
			goto st8
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st35
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 111:
			goto st8
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st35
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 65:
			goto st1
		case 67:
			goto st36
		case 70:
			goto st2
		case 73:
			goto st19
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st33
		case 87:
			goto st23
		case 97:
			goto st1
		case 99:
			goto st36
		case 102:
			goto st2
		case 105:
			goto st19
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st33
		case 119:
			goto st23
		}
		goto st0
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		switch data[p] {
		case 65:
			goto st37
		case 70:
			goto st2
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st37
		case 102:
			goto st2
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 78:
			goto st13
		case 82:
			goto st7
		case 83:
			goto st38
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 110:
			goto st13
		case 114:
			goto st7
		case 115:
			goto st38
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 77:
			goto st11
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto tr43
		case 86:
			goto st20
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 109:
			goto st11
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto tr43
		case 118:
			goto st20
		case 119:
			goto st23
		}
		goto st0
	tr43:
//line os_fsm.rl:52
		return true
		goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line os_fsm.go:2244
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 73:
			goto st19
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st33
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 105:
			goto st19
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st33
		case 119:
			goto st23
		}
		goto st0
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		switch data[p] {
		case 65:
			goto st1
		case 70:
			goto st2
		case 73:
			goto st19
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto tr44
		case 87:
			goto st23
		case 97:
			goto st1
		case 102:
			goto st2
		case 105:
			goto st19
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto tr44
		case 119:
			goto st23
		}
		goto st0
	tr44:
//line os_fsm.rl:52
		return true
		goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line os_fsm.go:2329
		switch data[p] {
		case 13:
			goto st26
		case 32:
			goto st26
		case 65:
			goto st1
		case 69:
			goto st21
		case 70:
			goto st2
		case 73:
			goto st28
		case 79:
			goto st27
		case 82:
			goto st7
		case 83:
			goto st10
		case 84:
			goto st18
		case 86:
			goto st20
		case 87:
			goto st23
		case 95:
			goto st26
		case 97:
			goto st1
		case 101:
			goto st21
		case 102:
			goto st2
		case 105:
			goto st28
		case 111:
			goto st27
		case 114:
			goto st7
		case 115:
			goto st10
		case 116:
			goto st18
		case 118:
			goto st20
		case 119:
			goto st23
		}
		if 9 <= data[p] && data[p] <= 10 {
			goto st26
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
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof19:
		cs = 19
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof40:
		cs = 40
		goto _test_eof
	_test_eof23:
		cs = 23
		goto _test_eof
	_test_eof24:
		cs = 24
		goto _test_eof
	_test_eof25:
		cs = 25
		goto _test_eof
	_test_eof26:
		cs = 26
		goto _test_eof
	_test_eof27:
		cs = 27
		goto _test_eof
	_test_eof41:
		cs = 41
		goto _test_eof
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof30:
		cs = 30
		goto _test_eof
	_test_eof42:
		cs = 42
		goto _test_eof
	_test_eof31:
		cs = 31
		goto _test_eof
	_test_eof32:
		cs = 32
		goto _test_eof
	_test_eof33:
		cs = 33
		goto _test_eof
	_test_eof34:
		cs = 34
		goto _test_eof
	_test_eof35:
		cs = 35
		goto _test_eof
	_test_eof36:
		cs = 36
		goto _test_eof
	_test_eof37:
		cs = 37
		goto _test_eof
	_test_eof38:
		cs = 38
		goto _test_eof
	_test_eof43:
		cs = 43
		goto _test_eof
	_test_eof39:
		cs = 39
		goto _test_eof
	_test_eof44:
		cs = 44
		goto _test_eof

	_test_eof:
		{
		}
	}

//line os_fsm.rl:55

	return false
}

// vim: ft=go ts=4 sw=4 noet
