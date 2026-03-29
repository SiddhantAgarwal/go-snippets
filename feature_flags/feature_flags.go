package feature_flags

type FeatureFlag uint64

const (
	FeatureA FeatureFlag = 1 << iota
	FeatureB
	FeatureC
	FeatureD
	FeatureE
	FeatureF
)

type FeatureSet struct {
	mask FeatureFlag
}

func (f *FeatureSet) Add(flag FeatureFlag) {
	f.mask |= flag
}

func (f *FeatureSet) Remove(flag FeatureFlag) {
	f.mask &^= flag
}

func (f *FeatureSet) IsEnabled(flag FeatureFlag) bool {
	return f.mask&flag != 0
}
