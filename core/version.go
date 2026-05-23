package core

type Version struct {
	Package string
	Sonolus string
}

func CurrentVersion() Version {
	return Version{
		Package: "7.15.1",
		Sonolus: "1.1.1",
	}
}
