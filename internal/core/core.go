package core

import uuid_generator "github.com/zytell3301/uuid-generator"

type Service struct {
	repository    Repository
	uuidGenerator uuid_generator.Generator
}

type ServiceConfig struct {
}

type Dependencies struct {
	Repository    Repository
	UuidGenerator uuid_generator.Generator
}

func NewMessagesCore(config ServiceConfig, dependencies Dependencies) Service {
	return Service{
		repository:    dependencies.Repository,
		uuidGenerator: dependencies.UuidGenerator,
	}
}
