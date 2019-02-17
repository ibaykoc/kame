package kame

type EntitiesCreator interface {
	CreateEntities()
}

type EntitiesOwner interface {
	GetEntities() *[]Entity
}
type ProcessorSystemsCreator interface {
	CreateProcessorSystems()
}

type ProcessorSystemsOwner interface {
	GetProcessorSystems() *[]ProcessorSystem
}

type DrawerSystemsCreator interface {
	CreateDrawerSystems()
}

type DrawerSystemsOwner interface {
	GetDrawerSystems() *[]DrawerSystem
}

type Scene interface {
	EntitiesCreator
	EntitiesOwner
	DrawerSystemsCreator
	DrawerSystemsOwner
	ProcessorSystemsCreator
	ProcessorSystemsOwner
}
