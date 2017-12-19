package uuid

%%{
	machine uuidRegex;
	write data;
}%%

func matchUuidRegex(data []byte) bool {
	cs, p, pe := 0, 0, len(data)
	eof := pe
	%%{
		main := [a-fA-F0-9]{8} '-' [a-fA-F0-9]{4} '-4' [a-fA-F0-9]{3} '-' [8|9|aA|bB] [a-fA-F0-9]{3} '-' [a-fA-F0-9]{12} %{ return true } ;
		write init;
		write exec;
	}%%
	return false
}

%%{
	machine uuidRegexiOS;
	write data;
}%%

func matchUuidRegexiOS(data []byte) bool {
	cs, p, pe := 0, 0, len(data)
	eof := pe
	%%{
		main := [A-F0-9]{8} '-' [A-F0-9]{4} '-4' [A-F0-9]{3} '-' [8|9|A|B] [A-F0-9]{3} '-' [A-F0-9]{12} %{ return true } ;
		write init;
		write exec;
	}%%
	return false
}

%%{
	machine uuidRegexAndroid;
	write data;
}%%

func matchUuidRegexAndroid(data []byte) bool {
	cs, p, pe := 0, 0, len(data)
	eof := pe
	%%{
		main := [a-f0-9]{8} '-' [a-f0-9]{4} '-4' [a-f0-9]{3} '-' [8|9|a|b] [a-f0-9]{3} '-' [a-f0-9]{12} %{ return true } ;
		write init;
		write exec;
	}%%
	return false
}

// vim: ft=go ts=4 sw=4 noet
