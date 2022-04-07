package sim3core

type SceneObj struct {
	X, Y  int
	Scene *Scene
}

func newSceneObj(x, y int, scene *Scene) *SceneObj {
	return &SceneObj{
		X:     x,
		Y:     y,
		Scene: scene,
	}
}
