package kame

type EntitiesCreator interface {
	CreateEntities()
}

type EntitiesOwner interface {
	GetEntityPointers() []*Entity
}

type EntitiesRemoveListener interface {
	OnRemoveEntities(entityIDs []int)
}

type ProcessorSystemsCreator interface {
	CreateProcessorSystems()
}

type ProcessorSystemsOwner interface {
	GetProcessorSystemPointers() []*ProcessorSystem
}

type DrawerSystemsCreator interface {
	CreateDrawerSystems()
}

type DrawerSystemsOwner interface {
	GetDrawerSystemPointers() []*DrawerSystem
}

type Scene interface {
	EntitiesCreator
	EntitiesOwner
	EntitiesRemoveListener
	DrawerSystemsCreator
	DrawerSystemsOwner
	ProcessorSystemsCreator
	ProcessorSystemsOwner
}
