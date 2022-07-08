package amarisapi

import (
	"context"
	"fmt"
	amaris "testamaris/gen/amaris"
)

// Receives two words and returns if these are friendly
func (s *amarissrvc) FriendlyChains(ctx context.Context, p *amaris.FriendlyChainsPayload) (res *amaris.Friendlychainsresponse, err error) {
	res = &amaris.Friendlychainsresponse{}

	var result string
	if IsFriendlyChains(p.Request.A, p.Request.B) {
		result = fmt.Sprintf("Las cadenas de texto %s y %s son amigas", p.Request.A, p.Request.B)
	} else {
		result = fmt.Sprintf("Las cadenas de texto %s y %s no son amigas", p.Request.A, p.Request.B)
	}

	res.Result = &result

	return
}

func IsFriendlyChains(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for index := range a {
		u := a[0 : index+1]
		v := a[index+1:]

		if b == v+u {
			return true
		}
	}

	return false
}
