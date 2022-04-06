package sim3core

// 获取目标方向，一般用于没有特定目标的方向，返回 (0/1/-1, 0/1/-1)
type FuncGetTargetDir func(person *Person, scene *Scene) (int, int)

type PersonDNA struct {
	GetTargetDir FuncGetTargetDir
}
