package persistence

import (
	"log"
	"sort"

	"github.com/jeromelesaux/fizzbuzz/configuration"
	"github.com/jeromelesaux/fizzbuzz/model"
)

var (
	RequestsStored    = make(map[model.Parameters]int)
	AddParametersChan chan model.Parameters
)

func init() {
	AddParametersChan = make(chan model.Parameters)
	go runChannel()
}

func runChannel() {
	for p := range AddParametersChan {
		err := Add(p)
		if err != nil {
			panic(err)
		}
	}
}

func Add(p model.Parameters) error {
	if configuration.StaticConfiguration.Persistence == configuration.MemoryType {
		AddInMemory(p)
	} else {
		err := AddInDB(p)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetMostFrequent() (model.Parameters, int) {
	if configuration.StaticConfiguration.Persistence == configuration.MemoryType {
		return GetMostFrequentMemory()
	} else {
		p, err := GetMostFrequentDB()
		if err != nil {
			log.Printf("%s\n", err.Error())
		}
		return p, int(p.Hits)
	}
}

func AddInMemory(p model.Parameters) {
	log.Printf("adding new model.Parameters\n")
	RequestsStored[p]++
}

type Pair struct {
	P model.Parameters
	V int
}

type Pairlist []Pair

func (p Pairlist) Len() int           { return len(p) }
func (p Pairlist) Less(i, j int) bool { return p[i].V < p[j].V }
func (p Pairlist) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func GetMostFrequentMemory() (model.Parameters, int) {
	if len(RequestsStored) == 0 {
		return model.Parameters{}, 0
	}

	pairs := make(Pairlist, len(RequestsStored))
	i := 0
	for k, v := range RequestsStored {
		pairs[i] = Pair{P: k, V: v}
		i++
	}
	sort.Sort(sort.Reverse(pairs))
	return pairs[0].P, pairs[0].V
}
