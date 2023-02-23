package stat

import (
	"github.com/zonewave/pkgs/util/expr"
	"github.com/zonewave/pkgs/util/maputil"
	"github.com/zonewave/pkgs/util/sliceutil"
	"math/rand"
	"sort"
	"time"
)

func ShuffleSequences[T comparable](sequences []T, r *rand.Rand) []T {
	r.Shuffle(len(sequences), func(i, j int) {
		sequences[i], sequences[j] = sequences[j], sequences[i]
	})
	return sequences
}

type RandomByWeightsInMap[Key comparable] struct {
	random            *RandomByWeights
	weightIndexMapKey map[int]Key
}

func NewRandomByWeightsInMap[Key comparable](weights map[Key]int32, r ...*rand.Rand) *RandomByWeightsInMap[Key] {
	weightSlice := make([]int32, 0, len(weights))
	weightIndexMapKey := make(map[int]Key, len(weights))
	curIndex := 0
	maputil.IterFn(weights, func(k Key, v int32) bool {
		weightSlice = append(weightSlice, v)
		weightIndexMapKey[curIndex] = k
		curIndex += 1
		return true
	})

	ret := &RandomByWeightsInMap[Key]{
		random:            NewRandomByWeights(weightSlice, r...),
		weightIndexMapKey: weightIndexMapKey,
	}
	return ret
}

func (w *RandomByWeightsInMap[T]) Rand() T {
	return w.weightIndexMapKey[w.random.Rand()]
}

type RandomByWeights struct {
	cumWeights []int32
	r          *rand.Rand
}

func NewDefaultRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
func NewRandomByWeights(weights []int32, rOpt ...*rand.Rand) *RandomByWeights {
	if len(weights) == 0 {
		return &RandomByWeights{}
	}
	return &RandomByWeights{
		cumWeights: sliceutil.Reduce(weights, expr.Add[int32]),
		r:          expr.CondExpr(len(rOpt) > 0, rOpt[0], NewDefaultRand()),
	}
}

func (w *RandomByWeights) Rand() int {
	maxWeight := w.cumWeights[len(w.cumWeights)-1]
	x := w.r.Int31n(maxWeight) + 1
	ret := sort.Search(len(w.cumWeights), func(i int) bool { return w.cumWeights[i] >= x })
	return ret
}

type RandomBox interface {
	GetWeight(int32) int32
}

func RandByWeightNoReplace(randBox RandomBox, srcIds []int32, r *rand.Rand) []int32 {
	ids := make([]int32, len(srcIds))
	copy(ids, srcIds)
	size := len(ids)

	totalWeight := sliceutil.Sum(ids, randBox.GetWeight)
	randIdsList := make([]int32, 0, size)
	for i := 0; i < size; i++ {
		randWeight := r.Int31n(totalWeight) + 1
		var currWeight int32
		for index, id := range ids {
			idWeight := randBox.GetWeight(id)
			currWeight += idWeight

			if randWeight <= currWeight {
				ids = append(ids[:index], ids[index+1:]...)
				totalWeight -= idWeight
				randIdsList = append(randIdsList, id)
				break
			}

		}

	}
	return randIdsList

}
