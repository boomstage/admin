package util

import (
	"sort"

	"github.com/hashicorp/go-version"
)

/*
比较版本号
1 (v1>v2)
0 (v1=v2)
-1 (v1<v2)
*/
func CompareV(v1 string, v2 string) int {
	v3, _ := version.NewVersion(v1)
	v4, _ := version.NewVersion(v2)
	return v3.Compare(v4)
}

/*
比较版本号v1是否小于或者等于v2
*/
func LessThanOrEqualV(v1 string, v2 string) bool {
	v3, _ := version.NewVersion(v1)
	v4, _ := version.NewVersion(v2)
	return v3.LessThanOrEqual(v4)
}

/*
比较版本号v1是否大于或者等于v2
*/
func GreaterThanOrEqualV(v1 string, v2 string) bool {
	v3, _ := version.NewVersion(v1)
	v4, _ := version.NewVersion(v2)
	return v3.GreaterThanOrEqual(v4)
}

/*
比较版本号v1是否在某个区间
vv参数示例：">= 1.0, < 1.4"
*/
func ConstraintsVV(v1 string, vv string) bool {
	// 判断范围
	v3, _ := version.NewVersion(v1)
	constraints, _ := version.NewConstraint(vv)
	return constraints.Check(v3)
}

/*
对版本号排序
参数示例：[]string{"1.1", "0.7.1", "1.4-beta", "1.4", "2"}
*/
func SortVV(vv []string) []*version.Version {
	versions := make([]*version.Version, len(vv))
	for i, raw := range vv {
		v, _ := version.NewVersion(raw)
		versions[i] = v
	}
	// After this, the versions are properly sorted
	sort.Sort(version.Collection(versions))
	return versions
}
