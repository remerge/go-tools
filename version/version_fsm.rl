package version

%%{
	machine version_parser;
	write data;
}%%

func Parse(data []byte) *Version {
	version := &Version{}

	if len(data) < 1 {
		return version
	}

	cs, p, pe := 0, 0, len(data)

	%%{
		action add_major { version.Major = version.Major * 10 + (int(fc) - '0') }
		action add_minor { version.Minor = version.Minor * 10 + (int(fc) - '0') }
		action add_patch { version.Patch = version.Patch * 10 + (int(fc) - '0') }
		main := ^digit* ( digit @add_major )+ [\.\_] ( digit @add_minor )+ [\.\_] ( digit @add_patch )+;
		write init;
		write exec;
	}%%

	// remove some insane versions
	if version.Major > 1000000 || version.Major < 0{
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
