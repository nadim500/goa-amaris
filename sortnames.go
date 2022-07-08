package amarisapi

import (
	"context"
	"sort"
	"strings"
	amaris "testamaris/gen/amaris"
)

// Receives text string and returns two parameters: an array string and the
// mount of array
func (s *amarissrvc) SortNames(ctx context.Context, p *amaris.SortNamesPayload) (res *amaris.Sortnamesresponse, err error) {
	res = &amaris.Sortnamesresponse{}

	var quantity int
	res.Data, quantity = sortByName(p.Request.Text)
	res.Quantity = &quantity

	return
}

func sortByName(data string) ([]string, int) {
	buffArr := strings.Split(data, ",")

	sort.Strings(buffArr)

	return buffArr, len(buffArr)
}
