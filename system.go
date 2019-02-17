package kame

type EntityFilterSet struct {
	ComponentType Component
	Need          bool
}

type EntityFilterSetsCreator interface {
	CreateEntityFilters()
}

type EntityFilterSetsOwner interface {
	GetEntityFilters() []EntityFilterSet
}

// MatchEntityListener is an interface to receive event when newly added entity matched
// with the systems entity filter set.
// Interface will receive the matched entity and the components that current system needs.
// Received components are in the same order with provided filter set .
type MatchEntityListener interface {
	OnEntityMatch(entity *Entity, components []*Component)
}

type Processor interface {
	Process(timeSinceLastFrame float32)
}

type Drawer interface {
	Draw(kdrawer *KDrawer)
}

type OnCreateListener interface {
	OnCreate()
}

type System interface {
	OnCreateListener
	EntityFilterSetsCreator
	EntityFilterSetsOwner
	MatchEntityListener
}

type ProcessorSystem interface {
	Processor
	System
}

type DrawerSystem interface {
	Drawer
	System
}
