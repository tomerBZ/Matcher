package match

import (
	"index/suffixarray"
	"strings"

	"BigID/aggregator"
	"BigID/domain"
)

var Keys = [50]string{
	"James",
	"John",
	"Robert",
	"Michael",
	"William",
	"David",
	"Richard",
	"Charles",
	"Joseph",
	"Thomas",
	"Christopher",
	"Daniel",
	"Paul",
	"Mark",
	"Donal d",
	"George",
	"Kenneth",
	"Steven",
	"Edward",
	"Brian",
	"Ronald",
	"Anthony",
	"Kevin",
	"Jason",
	"Matthew",
	"Gary",
	"Timothy",
	"Jose",
	"Larry",
	"Jeffrey",
	"Frank",
	"Scott",
	"Eric",
	"Stephen",
	"Andrew",
	"Raymond",
	"Gregory",
	"Joshua",
	"Jerry",
	"Dennis",
	"Walter",
	"Patrick",
	"Peter",
	"Harold",
	"Douglas",
	"H enry",
	"Carl",
	"Arthur",
	"Ryan",
	"Roger",
}

type Matcher interface {
	Find(chunk []byte)
}

type matcher struct {
	aggregator aggregator.Aggregator
}

func NewMatcher(aggregator aggregator.Aggregator) *matcher {
	return &matcher{aggregator: aggregator}
}

func (m *matcher) Find(chunk []byte) {
	index := suffixarray.New(chunk)
	chunkString := string(chunk)
	chunkLines := strings.Split(chunkString, "\n")
	for _, element := range Keys {
		r := index.Lookup([]byte(element), -1)
		if len(r) > 0 {
			for _, in := range r {
				lineCounter := 0
				lineOffset := 0
				for _, line := range chunkLines {
					if strings.Contains(line, element) {
						lineOffset = lineCounter
					}
					lineCounter++
				}
				m.aggregator.Aggregate(
					element, domain.Position{
						LineOffset: lineOffset,
						CharOffset: in,
					},
				)
			}
		}
	}
}
