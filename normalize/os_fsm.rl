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

%%{
	machine os_tvos;
	write data;
}%%

func MatchOsTvOS(data string) bool {
	cs, p, pe := 0, 0, len(data)
	%%{
		main := any* 'tvos'i @{ return true } ;
		write init;
		write exec;
	}%%
	return false
}

%%{
	machine os_roku;
	write data;
}%%

func MatchOsRoku(data string) bool {
	cs, p, pe := 0, 0, len(data)
	%%{
		main := any* ('roku'i ('os'i | ('+'|space) 'os'i)? ) @{ return true } ;
		write init;
		write exec;
	}%%
	return false
}

%%{
	machine os_fireos;
	write data;
}%%

func MatchOsFireOS(data string) bool {
	cs, p, pe := 0, 0, len(data)
	%%{
		main := any* ('fire'i (('tv'i)? 'os'i | ('-'|'+') 'os'i)? ) @{ return true } ;
		write init;
		write exec;
	}%%
	return false
}

%%{
	machine os_smartcast;
	write data;
}%%

func MatchOsSmartCast(data string) bool {
	cs, p, pe := 0, 0, len(data)
	%%{
		main := any* (('vizio'i)? 'smartcast'i ('os'i | '%20os'i)?) @{ return true } ;
		write init;
		write exec;
	}%%
	return false
}

%%{
	machine os_webos;
	write data;
}%%

func MatchOsWebOS(data string) bool {
	cs, p, pe := 0, 0, len(data)
	%%{
		main := any* ('webos'i ('tv'i | ('+'|space) 'tv'i)? ) @{ return true } ;
		write init;
		write exec;
	}%%
	return false
}

%%{
	machine os_tizen;
	write data;
}%%

func MatchOsTizen(data string) bool {
	cs, p, pe := 0, 0, len(data)
	%%{
		main := any* ('tizen'i ('tv'i)? ) @{ return true } ;
		write init;
		write exec;
	}%%
	return false
}

%%{
	machine os_linux;
	write data;
}%%

func MatchOsLinux(data string) bool {
	cs, p, pe := 0, 0, len(data)
	%%{
		main := any* 'linux'i @{ return true } ;
		write init;
		write exec;
	}%%
	return false
}

%%{
	machine os_vidaa;
	write data;
}%%

func MatchOsVidaa(data string) bool {
	cs, p, pe := 0, 0, len(data)
	%%{
		main := any* 'vidaa'i @{ return true } ;
		write init;
		write exec;
	}%%
	return false
}

%%{
	machine os_reachtv;
	write data;
}%%

func MatchOsReachTV(data string) bool {
	cs, p, pe := 0, 0, len(data)
	%%{
		main := any* 'reachtv'i @{ return true } ;
		write init;
		write exec;
	}%%
	return false
}

%%{
	machine os_chromeos;
	write data;
}%%

func MatchOsChromeOS(data string) bool {
	cs, p, pe := 0, 0, len(data)
	%%{
		main := any* 'chromeos'i @{ return true } ;
		write init;
		write exec;
	}%%
	return false
}

// vim: ft=go ts=4 sw=4 noet
