package dir2pkg

import(
	"yytest/dir1"
)


type ServiceImpl struct {
	countryEvaluationAlgo map[string]dir1pkg.EvaluationAlog
	name string
}

func NewServiceImpl(name string) *ServiceImpl{
	s := &ServiceImpl{}
	s.countryEvaluationAlgo = make(map[string]dir1pkg.EvaluationAlog)
	s.countryEvaluationAlgo["BR"] = dir1pkg.EvaluationAlog{
		Control: "123",
		Treatment: "456",
	}
	s.name = name
	return s
}
