package sim3core

import "math/rand"

// 获取目标方向，一般用于没有特定目标的方向，返回 (0/1/-1, 0/1/-1)
type FuncGetTargetDir func(person *Person, scene *Scene) (int, int)

func randGetTargetDir(person *Person, scene *Scene) (int, int) {
	minx := -1
	maxx := 1

	if person.X <= 0 {
		minx = 0
	}

	if person.X >= scene.MapData.Width-1 {
		maxx = 0
	}

	miny := -1
	maxy := 1

	if person.Y <= 0 {
		miny = 0
	}

	if person.Y >= scene.MapData.Height-1 {
		maxy = 0
	}

	if maxx-minx == 0 {
		if maxy-miny == 0 {
			return 0, 0
		}

		cy := rand.Int() % (maxy - miny)
		return 0, cy + miny
	}

	cx := rand.Int() % (maxy - miny)

	if maxy-miny == 0 {
		return cx + minx, 0
	}

	cy := rand.Int() % (maxy - miny)
	return cx + minx, cy + miny
}

func leftGetTargetDir(person *Person, scene *Scene) (int, int) {
	if person.X <= 0 {
		return randGetTargetDir(person, scene)
	}

	return -1, 0
}

func rightGetTargetDir(person *Person, scene *Scene) (int, int) {
	if person.X >= scene.MapData.Width-1 {
		return randGetTargetDir(person, scene)
	}

	return 1, 0
}

func upGetTargetDir(person *Person, scene *Scene) (int, int) {
	if person.Y <= 0 {
		return randGetTargetDir(person, scene)
	}

	return 0, -1
}

func downGetTargetDir(person *Person, scene *Scene) (int, int) {
	if person.Y >= scene.MapData.Height-1 {
		return randGetTargetDir(person, scene)
	}

	return 0, 1
}

func leftupGetTargetDir(person *Person, scene *Scene) (int, int) {
	if person.X <= 0 {
		return upGetTargetDir(person, scene)
	}

	if person.Y <= 0 {
		return leftGetTargetDir(person, scene)
	}

	return -1, -1
}

func leftdownGetTargetDir(person *Person, scene *Scene) (int, int) {
	if person.X <= 0 {
		return downGetTargetDir(person, scene)
	}

	if person.Y >= scene.MapData.Height-1 {
		return leftGetTargetDir(person, scene)
	}

	return -1, 1
}

func rightupGetTargetDir(person *Person, scene *Scene) (int, int) {
	if person.X >= scene.MapData.Width-1 {
		return upGetTargetDir(person, scene)
	}

	if person.Y <= 0 {
		return leftGetTargetDir(person, scene)
	}

	return 1, -1
}

func rightdownGetTargetDir(person *Person, scene *Scene) (int, int) {
	if person.X >= scene.MapData.Width-1 {
		return downGetTargetDir(person, scene)
	}

	if person.Y >= scene.MapData.Height-1 {
		return leftGetTargetDir(person, scene)
	}

	return 1, 1
}
