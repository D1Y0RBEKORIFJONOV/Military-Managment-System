package user_service

import (
	"storehouse-service/logger"
)

type Storehouse struct {
	log      logger.ILogger
	creater  StorehouseCreater
	provider StorehouseProvider
	updater  StorehouseUpdater
	deleter  StorehouseDeleter
}

func NewUser(log logger.ILogger,
	provider StorehouseProvider,
	creater StorehouseCreater,
	deleter StorehouseDeleter,
	updater StorehouseUpdater) *Storehouse {
	return &Storehouse{
		log:      log,
		provider: provider,
		creater:  creater,
		deleter:  deleter,
		updater:  updater,
	}
}
