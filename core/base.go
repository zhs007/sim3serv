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
