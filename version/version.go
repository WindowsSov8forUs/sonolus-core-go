package version

type Version struct {
	Package string
	Sonolus string
}

func CurrentVersion() Version {
	return Version{
		Package: "0.2.0",
		Sonolus: "1.1.1",
	}
}
