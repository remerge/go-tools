//line uuid_fsm.rl:1
package uuid

//line uuid_fsm.go:7
const uuidRegex_start int = 1
const uuidRegex_first_final int = 37
const uuidRegex_error int = 0

const uuidRegex_en_main int = 1

//line uuid_fsm.rl:6

func matchUuidRegex(data string) bool {
	cs, p, pe := 0, 0, len(data)
	eof := pe

//line uuid_fsm.go:22
	{
		cs = uuidRegex_start
	}

//line uuid_fsm.go:27
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
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
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
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
		}
		goto st_out
	st_case_1:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st2
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st3
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st3
			}
		default:
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st4
			}
		default:
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st5
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st5
			}
		default:
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st6
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st6
			}
		default:
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st7
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st8
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st8
			}
		default:
			goto st8
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st9
			}
		default:
			goto st9
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 45 {
			goto st10
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st11
			}
		default:
			goto st11
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st12
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st12
			}
		default:
			goto st12
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st13
			}
		default:
			goto st13
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st14
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st14
			}
		default:
			goto st14
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 45 {
			goto st15
		}
		goto st0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st16
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st16
			}
		default:
			goto st16
		}
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st17
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st17
			}
		default:
			goto st17
		}
		goto st0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st18
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st18
			}
		default:
			goto st18
		}
		goto st0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st19
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st19
			}
		default:
			goto st19
		}
		goto st0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if data[p] == 45 {
			goto st20
		}
		goto st0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st21
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st21
			}
		default:
			goto st21
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st22
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st22
			}
		default:
			goto st22
		}
		goto st0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st23
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st24
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st24
			}
		default:
			goto st24
		}
		goto st0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if data[p] == 45 {
			goto st25
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st26
			}
		default:
			goto st26
		}
		goto st0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st27
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st27
			}
		default:
			goto st27
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st28
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st28
			}
		default:
			goto st28
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st29
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st29
			}
		default:
			goto st29
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st30
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st30
			}
		default:
			goto st30
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st31
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st31
			}
		default:
			goto st31
		}
		goto st0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st32
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st32
			}
		default:
			goto st32
		}
		goto st0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st33
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st33
			}
		default:
			goto st33
		}
		goto st0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st34
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st34
			}
		default:
			goto st34
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st35
			}
		default:
			goto st35
		}
		goto st0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st36
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st36
			}
		default:
			goto st36
		}
		goto st0
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st37
			}
		default:
			goto st37
		}
		goto st0
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		goto st0
	st_out:
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
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof30:
		cs = 30
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

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 37:
//line uuid_fsm.rl:12
				return true
//line uuid_fsm.go:773
			}
		}

	_out:
		{
		}
	}

//line uuid_fsm.rl:15

	return false
}

//line uuid_fsm.go:786
const uuidRegexiOS_start int = 1
const uuidRegexiOS_first_final int = 37
const uuidRegexiOS_error int = 0

const uuidRegexiOS_en_main int = 1

//line uuid_fsm.rl:22

func matchUuidRegexiOS(data string) bool {
	cs, p, pe := 0, 0, len(data)
	eof := pe

//line uuid_fsm.go:801
	{
		cs = uuidRegexiOS_start
	}

//line uuid_fsm.go:806
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
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
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
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
		}
		goto st_out
	st_case_1:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st2
			}
		case data[p] >= 48:
			goto st2
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st3
			}
		case data[p] >= 48:
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st4
			}
		case data[p] >= 48:
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st5
			}
		case data[p] >= 48:
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st6
			}
		case data[p] >= 48:
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st7
			}
		case data[p] >= 48:
			goto st7
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st8
			}
		case data[p] >= 48:
			goto st8
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st9
			}
		case data[p] >= 48:
			goto st9
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 45 {
			goto st10
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st11
			}
		case data[p] >= 48:
			goto st11
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st12
			}
		case data[p] >= 48:
			goto st12
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st13
			}
		case data[p] >= 48:
			goto st13
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st14
			}
		case data[p] >= 48:
			goto st14
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 45 {
			goto st15
		}
		goto st0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st16
			}
		case data[p] >= 48:
			goto st16
		}
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st17
			}
		case data[p] >= 48:
			goto st17
		}
		goto st0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st18
			}
		case data[p] >= 48:
			goto st18
		}
		goto st0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st19
			}
		case data[p] >= 48:
			goto st19
		}
		goto st0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if data[p] == 45 {
			goto st20
		}
		goto st0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		if data[p] == 124 {
			goto st21
		}
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 66 {
				goto st21
			}
		case data[p] >= 56:
			goto st21
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st22
			}
		case data[p] >= 48:
			goto st22
		}
		goto st0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st23
			}
		case data[p] >= 48:
			goto st23
		}
		goto st0
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st24
			}
		case data[p] >= 48:
			goto st24
		}
		goto st0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if data[p] == 45 {
			goto st25
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st26
			}
		case data[p] >= 48:
			goto st26
		}
		goto st0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st27
			}
		case data[p] >= 48:
			goto st27
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st28
			}
		case data[p] >= 48:
			goto st28
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st29
			}
		case data[p] >= 48:
			goto st29
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st30
			}
		case data[p] >= 48:
			goto st30
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st31
			}
		case data[p] >= 48:
			goto st31
		}
		goto st0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st32
			}
		case data[p] >= 48:
			goto st32
		}
		goto st0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st33
			}
		case data[p] >= 48:
			goto st33
		}
		goto st0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st34
			}
		case data[p] >= 48:
			goto st34
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st35
			}
		case data[p] >= 48:
			goto st35
		}
		goto st0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st36
			}
		case data[p] >= 48:
			goto st36
		}
		goto st0
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		switch {
		case data[p] > 57:
			if 65 <= data[p] && data[p] <= 70 {
				goto st37
			}
		case data[p] >= 48:
			goto st37
		}
		goto st0
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		goto st0
	st_out:
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
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof30:
		cs = 30
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

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 37:
//line uuid_fsm.rl:28
				return true
//line uuid_fsm.go:1427
			}
		}

	_out:
		{
		}
	}

//line uuid_fsm.rl:31

	return false
}

//line uuid_fsm.go:1440
const uuidRegexAndroid_start int = 1
const uuidRegexAndroid_first_final int = 37
const uuidRegexAndroid_error int = 0

const uuidRegexAndroid_en_main int = 1

//line uuid_fsm.rl:38

func matchUuidRegexAndroid(data string) bool {
	cs, p, pe := 0, 0, len(data)
	eof := pe

//line uuid_fsm.go:1455
	{
		cs = uuidRegexAndroid_start
	}

//line uuid_fsm.go:1460
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
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
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
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
		}
		goto st_out
	st_case_1:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st2
			}
		case data[p] >= 48:
			goto st2
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st3
			}
		case data[p] >= 48:
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st4
			}
		case data[p] >= 48:
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st5
			}
		case data[p] >= 48:
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st6
			}
		case data[p] >= 48:
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st7
			}
		case data[p] >= 48:
			goto st7
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st8
			}
		case data[p] >= 48:
			goto st8
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st9
			}
		case data[p] >= 48:
			goto st9
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 45 {
			goto st10
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st11
			}
		case data[p] >= 48:
			goto st11
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st12
			}
		case data[p] >= 48:
			goto st12
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st13
			}
		case data[p] >= 48:
			goto st13
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st14
			}
		case data[p] >= 48:
			goto st14
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 45 {
			goto st15
		}
		goto st0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		if data[p] == 52 {
			goto st16
		}
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st17
			}
		case data[p] >= 48:
			goto st17
		}
		goto st0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st18
			}
		case data[p] >= 48:
			goto st18
		}
		goto st0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st19
			}
		case data[p] >= 48:
			goto st19
		}
		goto st0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if data[p] == 45 {
			goto st20
		}
		goto st0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		if data[p] == 124 {
			goto st21
		}
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 98 {
				goto st21
			}
		case data[p] >= 56:
			goto st21
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st22
			}
		case data[p] >= 48:
			goto st22
		}
		goto st0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st23
			}
		case data[p] >= 48:
			goto st23
		}
		goto st0
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st24
			}
		case data[p] >= 48:
			goto st24
		}
		goto st0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if data[p] == 45 {
			goto st25
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st26
			}
		case data[p] >= 48:
			goto st26
		}
		goto st0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st27
			}
		case data[p] >= 48:
			goto st27
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st28
			}
		case data[p] >= 48:
			goto st28
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st29
			}
		case data[p] >= 48:
			goto st29
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st30
			}
		case data[p] >= 48:
			goto st30
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st31
			}
		case data[p] >= 48:
			goto st31
		}
		goto st0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st32
			}
		case data[p] >= 48:
			goto st32
		}
		goto st0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st33
			}
		case data[p] >= 48:
			goto st33
		}
		goto st0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st34
			}
		case data[p] >= 48:
			goto st34
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st35
			}
		case data[p] >= 48:
			goto st35
		}
		goto st0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st36
			}
		case data[p] >= 48:
			goto st36
		}
		goto st0
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		switch {
		case data[p] > 57:
			if 97 <= data[p] && data[p] <= 102 {
				goto st37
			}
		case data[p] >= 48:
			goto st37
		}
		goto st0
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		goto st0
	st_out:
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
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof30:
		cs = 30
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

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 37:
//line uuid_fsm.rl:44
				return true
//line uuid_fsm.go:2076
			}
		}

	_out:
		{
		}
	}

//line uuid_fsm.rl:47

	return false
}

// vim: ft=go ts=4 sw=4 noet
