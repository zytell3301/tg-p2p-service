package core

type Service struct {
	repository Repository
}

type ServiceConfig struct {
}

type Dependencies struct {
	Repository Repository
}

func NewMessagesCore(config ServiceConfig, dependencies Dependencies) Service {
	return Service{
		repository: dependencies.Repository,
	}
}
