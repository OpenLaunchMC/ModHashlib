package mod

type mod struct {
	name string
	authors []string
  latestVersion string
	allVersions []string
	repoURL []string  //Code Repository URL
	downloadURL []string //Where to download
	documentURL string  //Document
	virustotalURL string //Virustotal Scanning History
	sha256hash string
	sha512hash string
	md5hash string
	signingKey string
}

