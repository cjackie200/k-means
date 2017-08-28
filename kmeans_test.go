package k-means

import (
	"fmt"
	"sort"
	"testing"
)

type KPoint interface {
	KValue() int
}

type point struct {
	index int
	KPoint
}

func newPoint(i int, p KPoint) *point {
	return &point{
		index:  i,
		KPoint: p,
	}
}

type KMeansManager struct {
	index int
	data  []*point
	means []([]*point)
	key   []int
}

func (k *KMeansManager) Sort() *KMeansManager {
	sort.Sort(k)
	return k
}

func (k *KMeansManager) Len() int {
	return len(k.data)
}

func (k *KMeansManager) Less(i int, j int) bool {
	return k.data[i].KValue() < k.data[j].KValue()
}

func (k *KMeansManager) Swap(i int, j int) {
	k.data[i], k.data[j] = k.data[j], k.data[i]
	k.data[i].index = i
	k.data[j].index = j
}

func NewKMeansManager() *KMeansManager {
	return &KMeansManager{}
}

func (k *KMeansManager) Load(v []KPoint) *KMeansManager {
	l := len(v)
	k.data = make([]*point, l)
	for i := 0; i < l; i++ {
		k.data[i] = newPoint(i, v[i])
	}
	return k
}

func (k *KMeansManager) GetMeans() []([]KPoint) {
	l := len(k.means)
	res := make([]([]KPoint), l)
	for i := 0; i < l; i++ {
		n := len(k.means[i])
		temp := make([]KPoint, n)
		for j := 0; j < n; j++ {
			temp[j] = k.means[i][j].KPoint
		}
		res[i] = temp
	}
	return res
}

func (k *KMeansManager) KMeans(n int) *KMeansManager {
	k.firstKey(n)
	for {
		k.kMeansAssign()
		nextKey := k.nextKey()
		//if i > 1 {
		//	return k
		//}
		if equalKey(k.key, nextKey) {
			return k
		} else {
			k.key = nextKey
		}
	}
}

func (k *KMeansManager) kMeansAssign() *KMeansManager {
	n := len(k.key)
	k.means = makeKMeanList(n, len(k.data))
	k.means[0] = append(k.means[0], k.data[:(k.key[0])]...)
	k.means[n-1] = append(k.means[n-1], k.data[(k.key[n-1]):]...)
	for i := 0; i < n-1; i++ {
		for j := k.key[i]; j < k.key[i+1]; j++ {
			if abs(k.data[j], k.data[k.key[i]]) > abs(k.data[j], k.data[k.key[i+1]]) {
				k.means[i+1] = append(k.means[i+1], k.data[j:k.key[i+1]]...)
				break
			} else {
				k.means[i] = append(k.means[i], k.data[j])
			}
		}
	}
	return k
}

func (k *KMeansManager) nextKey() []int {
	sortKMeansPoint(k.means)
	l := len(k.means)
	key := make([]int, l)
	for i := 0; i < l; i++ {
		key[i] = k.means[i][len(k.means[i])/2].index
	}
	return key
}

func equalKey(a []int, b []int) bool {
	l := len(a)
	if l != len(b) {
		return false
	} else {
		for i := 0; i < l; i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
}

func sortKMeansPoint(p [][]*point) {
	n := len(p)
	for m := 0; m < n; m++ {
		sort.Slice(p[m], func(i, j int) bool {
			return p[m][i].index < p[m][j].index
		})
	}
}

func (k *KMeansManager) firstKey(n int) *KMeansManager {
	if 0 == len(k.key) {
		k.key = make([]int, n)
		oneKey := len(k.data) / n
		for i := 0; i < n; i++ {
			k.key[i] = oneKey*i + oneKey/2
		}
	}
	return k
}

func makeKMeanList(n int, m int) [][]*point {
	res := make([][]*point, n)
	for i := 0; i < n; i++ {
		res[i] = make([]*point, 0, m)
	}
	return res
}

func abs(a *point, b *point) int {
	c := a.KValue() - b.KValue()
	if c < 0 {
		return -c
	} else {
		return c
	}
}

type TPoint struct {
	a int
}

func (t *TPoint) KValue() int {
	return t.a
}

func TestKm(*testing.T) {
	a := []KPoint{
		&TPoint{1},
		&TPoint{2},
		&TPoint{3},
		&TPoint{3}, //3
		&TPoint{3},
		&TPoint{3},
		&TPoint{3},
		&TPoint{21},
		&TPoint{22},
		&TPoint{23},
		&TPoint{23},
		&TPoint{23},
		&TPoint{23}, //12
		&TPoint{23},
		&TPoint{23},
		&TPoint{23},
		&TPoint{23},
		&TPoint{23},
		&TPoint{23},
		&TPoint{23},
		&TPoint{23},
		&TPoint{23},
		&TPoint{23},
		&TPoint{23},
		&TPoint{23},
		&TPoint{23},
		&TPoint{23},
		&TPoint{23},
	}
	n := 2
	km := NewKMeansManager().Load(a).Sort().KMeans(n).GetMeans()
	for i := 0; i < n; i++ {
		fmt.Printf("第 %00d 组", i)
		for j := 0; j < len(km[i]); j++ {
			fmt.Printf("%00d => %v \n", j, km[i][j].KValue())
		}
	}
}
