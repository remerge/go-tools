package normalize

%%{
	machine device_ipod;
	write data;
}%%

func MatchDeviceiPod(data string) bool {
	cs, p, pe := 0, 0, len(data)
	%%{
		main := any* 'ipod'i @{ return true } ;
		write init;
		write exec;
	}%%
	return false
}

%%{
	machine device_ipad;
	write data;
}%%

func MatchDeviceiPad(data string) bool {
	cs, p, pe := 0, 0, len(data)
	%%{
		main := any* 'ipad'i @{ return true } ;
		write init;
		write exec;
	}%%
	return false
}

%%{
	machine device_iphone;
	write data;
}%%

func MatchDeviceiPhone(data string) bool {
	cs, p, pe := 0, 0, len(data)
	%%{
		main := any* 'iphone'i @{ return true } ;
		write init;
		write exec;
	}%%
	return false
}

// vim: ft=go ts=4 sw=4 noet
