//line version_fsm.rl:1
package version

//line version_fsm.go:7
const version_parser_start int = 1
const version_parser_first_final int = 6
const version_parser_error int = 0

const version_parser_en_main int = 1

//line version_fsm.rl:6

func Parse(data string) *Version {
	version := &Version{}

	if len(data) < 1 {
		return version
	}

	cs, p, pe := 0, 0, len(data)

//line version_fsm.go:28
	{
		cs = version_parser_start
	}

//line version_fsm.go:33
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 0:
			goto st_case_0
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
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st1
	tr1:
//line version_fsm.rl:18
		version.Major = version.Major*10 + (int(data[p]) - '0')
		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line version_fsm.go:73
		switch data[p] {
		case 46:
			goto st3
		case 95:
			goto st3
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr4
		}
		goto st0
	tr4:
//line version_fsm.rl:19
		version.Minor = version.Minor*10 + (int(data[p]) - '0')
		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line version_fsm.go:106
		switch data[p] {
		case 46:
			goto st5
		case 95:
			goto st5
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr4
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr6
		}
		goto st0
	tr6:
//line version_fsm.rl:20
		version.Patch = version.Patch*10 + (int(data[p]) - '0')
		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line version_fsm.go:135
		if 48 <= data[p] && data[p] <= 57 {
			goto tr6
		}
		goto st0
	st_out:
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
	_out:
		{
		}
	}

//line version_fsm.rl:24

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
