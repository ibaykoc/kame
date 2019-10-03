package kame

type EntityFilterSet struct {
	ComponentType Component
	Need          bool
}

type Processor interface {
	Process(timeSinceLastFrame float32)
}

type EntitiesAddedListener interface {
	OnEntitiesAdded(entities []*Entity)
}

type Drawer interface {
	Draw(kdrawer *KwindowDrawer)
}

type InputProcessor interface {
	ProcessInput(windowInput KwindowInput)
}

type OnCreateListener interface {
	OnCreate()
}

type System interface {
	OnCreateListener
	EntitiesAddedListener
	EntitiesRemoveListener
}

type ProcessorSystem interface {
	Processor
	System
}

type DrawerSystem interface {
	Drawer
	System
}

type InputProcessorSystem interface {
	InputProcessor
	System
}
