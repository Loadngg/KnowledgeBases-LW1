package repository

type DiseaseRepositoryInterface interface {
	GetSymptomsList() (*[]string, error)
	GetRules() (*[]Rule, error)
	GetDiseases() (*[]string, error)
}

type Repository struct {
	DiseaseRepositoryInterface
}

func New(rulesFilepath string, symptomsFilepath string) *Repository {
	return &Repository{
		DiseaseRepositoryInterface: NewDiseaseRepository(rulesFilepath, symptomsFilepath),
	}
}
