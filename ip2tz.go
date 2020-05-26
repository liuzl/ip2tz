package ip2tz

import (
	"errors"

	"github.com/crawlerclub/phonenumbers"
	"github.com/liuzl/ip2loc"
)

var (
	ErrNoTzFound = errors.New("no timezone found for this ip")
)

func CountryToTz(country string) (string, error) {
	switch country {
	case "US":
		return "America/New_York", nil
	case "CN":
		return "Asia/Shanghai", nil
	default:
		cc := phonenumbers.GetCountryCodeForRegion(country)
		tz, ok := phonenumbers.PrefixToTimezone[cc]
		if !ok || len(tz) == 0 {
			return "", ErrNoTzFound
		}
		return tz[0], nil
	}
}

func Find(ip string) (string, error) {
	country, err := ip2loc.Find(ip)
	if err != nil {
		return "", err
	}
	return CountryToTz(country)
}
