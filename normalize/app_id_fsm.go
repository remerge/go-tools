// line 1 "app_id_fsm.rl"
package normalize

// line 7 "app_id_fsm.go"
const app_id_start int = 1
const app_id_first_final int = 4
const app_id_error int = 0

const app_id_en_main int = 1

// line 6 "app_id_fsm.rl"

func AppId(data string) string {
	cs, p, pe := 0, 0, len(data)
	mark, eof := p, pe

	// if we can't match an iOS ID we return the ref as-is
	ref := data

	// line 26 "app_id_fsm.go"
	{
		cs = app_id_start
	}

	// line 31 "app_id_fsm.go"
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 4:
			goto st_case_4
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		}
		goto st_out
	st_case_1:
		if data[p] == 105 {
			goto st2
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr0
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	tr0:
		// line 16 "app_id_fsm.rl"

		mark = p
		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		// line 71 "app_id_fsm.go"
		if 48 <= data[p] && data[p] <= 57 {
			goto st4
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 100 {
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr0
		}
		goto st0
	st_out:
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 4:
				// line 17 "app_id_fsm.rl"

				ref = data[mark:p]
				// line 106 "app_id_fsm.go"
			}
		}

	_out:
		{
		}
	}

	// line 21 "app_id_fsm.rl"

	return ref
}

// vim: ft=go ts=4 sw=4 noet
