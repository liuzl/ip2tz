package ip2tz

import (
	"errors"
	"github.com/liuzl/ip2loc"
	"github.com/ttacon/libphonenumber"
)

var (
	ErrNoTzFound = errors.New("no timezone found for this ip")
)

func Find(ip string) (string, error) {
	country, err := ip2loc.Find(ip)
	if err != nil {
		return "", err
	}
	switch country {
	case "US":
		return "America/New_York", nil
	case "CN":
		return "Asia/Shanghai", nil
	default:
		cc := libphonenumber.GetCountryCodeForRegion(country)
		tz, ok := libphonenumber.CountryCodeToTimeZones[cc]
		if !ok || len(tz) == 0 {
			return "", ErrNoTzFound
		}
		return tz[0], nil
	}
}
