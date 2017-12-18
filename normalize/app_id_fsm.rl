package normalize

%%{
	machine app_id;
	write data;
}%%

func AppId(data string) string {
	cs, p, pe := 0, 0, len(data)
	mark, eof := p, pe

	// if we can't match an iOS ID we return the ref as-is
	ref := data

	%%{
		action mark { mark = p }
		action push { ref = data[mark:p] }
		main := 'id'? digit+ >mark %push;
		write init;
		write exec;
	}%%

	return ref
}

// vim: ft=go ts=4 sw=4 noet
