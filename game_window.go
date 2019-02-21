package kame

// import (
// 	"reflect"
// )

// type GameWindow struct {
// 	*Kwindow
// 	scenes            []Scene
// 	currentSceneIndex int
// 	entityCount       int
// }

// func (gw *GameWindow) initialize(scenes []Scene) {
// 	gw.currentSceneIndex = 0
// 	gw.scenes = scenes
// }

// func (gw *GameWindow) RemoveEntities(entitieIDs []int) {
// 	currentScene := &gw.scenes[gw.currentSceneIndex]
// 	sceneProcessorSystems := (*currentScene).GetProcessorSystemPointers()
// 	for i := 0; i < len(sceneProcessorSystems); i++ {
// 		(*sceneProcessorSystems[i]).OnRemoveEntities(entitieIDs)

// 	}
// 	sceneDrawerSystems := (*currentScene).GetDrawerSystemPointers()
// 	for i := 0; i < len(sceneDrawerSystems); i++ {
// 		(*sceneDrawerSystems[i]).OnRemoveEntities(entitieIDs)
// 	}
// 	(*currentScene).OnRemoveEntities(entitieIDs)
// }

// func (gw *GameWindow) Start() {
// 	gw.Window.Start()
// 	currentScene := &gw.scenes[gw.currentSceneIndex]
// 	(*currentScene).CreateProcessorSystems()
// 	(*currentScene).CreateDrawerSystems()
// 	sceneProcessorSystems := (*currentScene).GetProcessorSystemPointers()
// 	for _, p := range sceneProcessorSystems {
// 		(*p).OnCreate()
// 		(*p).CreateEntityFilters()
// 	}
// 	sceneDrawerSystems := (*currentScene).GetDrawerSystemPointers()
// 	for _, d := range sceneDrawerSystems {
// 		(*d).OnCreate()
// 		(*d).CreateEntityFilters()
// 	}

// 	(*currentScene).CreateEntities()
// 	sceneEntities := (*currentScene).GetEntityPointers()
// 	for _, e := range sceneEntities {
// 		(*e).ReceiveID(gw.entityCount)
// 		gw.entityCount++
// 		(*e).CreateComponents()
// 		for _, s := range sceneProcessorSystems {
// 			if comps, ok := filerEntityForSystem(e, (*s).GetEntityFilters()); ok {
// 				(*s).OnEntityMatch(e, comps)
// 			}
// 		}
// 		for _, s := range sceneDrawerSystems {
// 			if comps, ok := filerEntityForSystem(e, (*s).GetEntityFilters()); ok {
// 				(*s).OnEntityMatch(e, comps)
// 			}
// 		}
// 	}
// }

// func (gw *GameWindow) update(timeSinceLastFrame float32) {
// 	pSystems := gw.scenes[gw.currentSceneIndex].GetProcessorSystemPointers()
// 	for _, s := range pSystems {
// 		(*s).Process(timeSinceLastFrame)
// 	}
// }

// func (gw *GameWindow) draw(kdrawer *KDrawer) {
// 	dSystems := gw.scenes[gw.currentSceneIndex].GetDrawerSystemPointers()
// 	for _, s := range dSystems {
// 		(*s).Draw(kdrawer)
// 	}
// }

// func filerEntityForSystem(entity *Entity, filterSets []EntityFilterSet) ([]*Component, bool) {
// 	shc := []Component{}
// 	snhc := []Component{}
// 	for _, filterSet := range filterSets {
// 		if filterSet.Need {
// 			shc = append(shc, filterSet.ComponentType)
// 		} else {
// 			snhc = append(snhc, filterSet.ComponentType)
// 		}
// 	}
// 	components := (*entity).GetComponentPointers()
// 	eComToFiltering := make(map[*Component]bool)
// 	for i := 0; i < len(components); i++ {
// 		c := *components[i]
// 		eComToFiltering[components[i]] = true
// 		for _, snh := range snhc {
// 			if reflect.TypeOf(c) == reflect.TypeOf(snh) {
// 				// Entity have component that should not have, bail out
// 				return nil, false
// 			}
// 		}
// 	}
// 	matchedComponents := []*Component{}
// 	for _, sh := range shc {
// 		for eCom, filtering := range eComToFiltering {
// 			if !filtering {
// 				continue
// 			}
// 			if reflect.TypeOf((*eCom)) == reflect.TypeOf(sh) {
// 				matchedComponents = append(matchedComponents, eCom)
// 				if len(matchedComponents) == len(shc) {
// 					return matchedComponents, true
// 				}
// 				eComToFiltering[eCom] = false
// 				break

// 			}
// 		}
// 	}

// 	return nil, false
// }
