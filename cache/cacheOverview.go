package cache

type Overview struct {
	totalKeysNum   int
	totalKeysSize  int
	totalValueSize int
}

func (o *Overview) Add(value *Frame) {
	o.totalKeysNum += 1
	o.totalKeysSize += value.GetKeySize()
	o.totalValueSize += value.GetValueSize()
}

func (o *Overview) Del(value *Frame) {
	o.totalKeysNum -= 1
	o.totalKeysSize -= value.GetKeySize()
	o.totalValueSize -= value.GetValueSize()
}

func (o *Overview) GetKeysNums() int {
	return o.totalKeysNum
}

func (o *Overview) GetKeysSize() int {
	return o.totalKeysSize
}

func (o *Overview) GetValueSize() int {
	return o.totalValueSize
}

func NewOverview() *Overview {
	return &Overview{0, 0, 0}
}
