package kame

type KGameWindow struct {
	*KwindowController
	scenes            []Scene
	currentSceneIndex int
	entityCount       int
}

func (gw *KGameWindow) initialize(scenes []Scene) {
	gw.currentSceneIndex = 0
	gw.scenes = scenes
}

func (gw *KGameWindow) Start() {
	currentScene := &gw.scenes[gw.currentSceneIndex]
	(*currentScene).CreateProcessorSystems()
	(*currentScene).CreateDrawerSystems()
	(*currentScene).CreateInputProcessorSystems()
	sceneProcessorSystems := (*currentScene).GetProcessorSystemPointers()
	for _, p := range sceneProcessorSystems {
		(*p).OnCreate()
	}
	sceneDrawerSystems := (*currentScene).GetDrawerSystemPointers()
	for _, d := range sceneDrawerSystems {
		(*d).OnCreate()
	}
	sceneInputProcessorSystems := (*currentScene).GetInputProcessorSystemPointers()
	for _, i := range sceneInputProcessorSystems {
		(*i).OnCreate()
	}
	(*currentScene).CreateEntities()
	sceneEntities := (*currentScene).GetEntityPointers()
	for _, e := range sceneEntities {
		(*e).ReceiveID(gw.entityCount)
		gw.entityCount++
		(*e).CreateComponents()
	}
	for _, s := range sceneInputProcessorSystems {
		(*s).OnEntitiesAdded(sceneEntities)
	}
	for _, s := range sceneProcessorSystems {
		(*s).OnEntitiesAdded(sceneEntities)
	}
	for _, s := range sceneDrawerSystems {
		(*s).OnEntitiesAdded(sceneEntities)
	}
}

func (gw *KGameWindow) processInput(windowInput KwindowInput) {
	iSystems := gw.scenes[gw.currentSceneIndex].GetInputProcessorSystemPointers()
	for _, i := range iSystems {
		(*i).ProcessInput(windowInput)
	}
}

func (gw *KGameWindow) update(timeSinceLastFrame float32) {
	pSystems := gw.scenes[gw.currentSceneIndex].GetProcessorSystemPointers()
	for _, s := range pSystems {
		(*s).Process(timeSinceLastFrame)
	}
}

func (gw *KGameWindow) draw(kdrawer *KwindowDrawer) {
	dSystems := gw.scenes[gw.currentSceneIndex].GetDrawerSystemPointers()
	for _, s := range dSystems {
		(*s).Draw(kdrawer)
	}
}

func (gw *KGameWindow) RemoveEntities(entitieIDs []int) {
	currentScene := &gw.scenes[gw.currentSceneIndex]
	sceneProcessorSystems := (*currentScene).GetProcessorSystemPointers()
	for i := 0; i < len(sceneProcessorSystems); i++ {
		(*sceneProcessorSystems[i]).OnRemoveEntities(entitieIDs)
	}
	sceneDrawerSystems := (*currentScene).GetDrawerSystemPointers()
	for i := 0; i < len(sceneDrawerSystems); i++ {
		(*sceneDrawerSystems[i]).OnRemoveEntities(entitieIDs)
	}
	sceneInputProcessorSystems := (*currentScene).GetInputProcessorSystemPointers()
	for i := 0; i < len(sceneInputProcessorSystems); i++ {
		(*sceneInputProcessorSystems[i]).OnRemoveEntities(entitieIDs)
	}
	(*currentScene).OnRemoveEntities(entitieIDs)
}
