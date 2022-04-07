package sim3core

type Person struct {
	SceneObj

	DNA    *PersonDNA
	Base   *Base
	InBase bool
}

func newPerson(x, y int, scene *Scene, base *Base) *Person {
	person := &Person{
		SceneObj: *newSceneObj(x, y, scene),
		InBase:   true,
	}

	params := scene.MgrPersonDNA.NewPersonDNAParams(base)
	person.DNA = scene.MgrPersonDNA.NewPersonDNA(params)

	return person
}
