package sim3core

type Base struct {
	SceneObj

	ListPerson []*Person
}

func (base *Base) buildPersonDNAParams(params *PersonDNAParams) {
	for _, v := range base.ListPerson {
		params.WeightsGetTargetDir.AddWeight(v.DNA.NameGetTargetDir, 1, false)
	}
}

func (base *Base) NewPerson() *Person {
	person := newPerson(base.X, base.Y, base.Scene, base)

	base.ListPerson = append(base.ListPerson, person)

	return person
}
