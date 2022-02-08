package data

//Global conf
const (
	LocalPort          string = "53"
	SinkholeAddr       string = "1.2.3.4"
	DefaultFallbackDNS string = "8.8.8.8:53"
)

//DNS data
var DomainsBlock map[string]bool = map[string]bool{}

func AddDomainToBlock(url string) {
	DomainsBlock[url] = true
}

func IsInBlockList(url string) bool {
	_, b := DomainsBlock[url]
	return b
}
