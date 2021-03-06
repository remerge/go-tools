package normalize

%%{
	machine os_ios;
	write data;
}%%

func MatchOsiOS(data string) bool {
	cs, p, pe := 0, 0, len(data)
	%%{
		main := any* ('ipod'i | 'iphone'i | 'ipad'i | 'ios'i) @{ return true } ;
		write init;
		write exec;
	}%%
	return false
}

%%{
	machine os_android;
	write data;
}%%

func MatchOsAndroid(data string) bool {
	cs, p, pe := 0, 0, len(data)
	%%{
		main := any* 'android'i @{ return true } ;
		write init;
		write exec;
	}%%
	return false
}

// vim: ft=go ts=4 sw=4 noet
