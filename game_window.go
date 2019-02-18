package kame

import (
	"reflect"
)

type GameWindow struct {
	*Window
	scenes            []Scene
	currentSceneIndex int
	entityCount       int
}

func (gw *GameWindow) initialize(scenes []Scene) {
	gw.currentSceneIndex = 0
	gw.scenes = scenes
}

func (gw *GameWindow) Start() {
	currentScene := gw.scenes[gw.currentSceneIndex]
	currentScene.CreateProcessorSystems()
	currentScene.CreateDrawerSystems()
	sceneProcessorSystems := currentScene.GetProcessorSystems()
	for _, p := range *sceneProcessorSystems {
		p.OnCreate()
		p.CreateEntityFilters()
	}
	sceneDrawerSystems := currentScene.GetDrawerSystems()
	for _, d := range *sceneDrawerSystems {
		d.OnCreate()
		d.CreateEntityFilters()
	}

	currentScene.CreateEntities()
	sceneEntities := currentScene.GetEntities()
	for _, e := range *sceneEntities {
		e.ReceiveID(gw.entityCount)
		gw.entityCount++
		e.CreateComponents()
		for _, s := range *sceneProcessorSystems {
			if comps, ok := filerEntityForSystem(&e, s.GetEntityFilters()); ok {
				s.OnEntityMatch(&e, comps)
			}
		}
		for _, s := range *sceneDrawerSystems {
			if comps, ok := filerEntityForSystem(&e, s.GetEntityFilters()); ok {
				s.OnEntityMatch(&e, comps)
			}
		}
	}
}

func (gw *GameWindow) update(timeSinceLastFrame float32) {
	pSystems := gw.scenes[gw.currentSceneIndex].GetProcessorSystems()
	for _, s := range *pSystems {
		s.Process(timeSinceLastFrame)
	}
}

func (gw *GameWindow) draw(kdrawer *KDrawer) {
	dSystems := gw.scenes[gw.currentSceneIndex].GetDrawerSystems()
	for _, s := range *dSystems {
		s.Draw(kdrawer)
	}
}

func filerEntityForSystem(entity *Entity, filterSets []EntityFilterSet) ([]*Component, bool) {
	shc := []Component{}
	snhc := []Component{}
	for _, filterSet := range filterSets {
		if filterSet.Need {
			shc = append(shc, filterSet.ComponentType)
		} else {
			snhc = append(snhc, filterSet.ComponentType)
		}
	}
	components := (*entity).GetComponents()
	eComToFiltering := make(map[Component]bool)
	for _, c := range *components {
		eComToFiltering[c] = true
		for _, snh := range snhc {
			if reflect.TypeOf(c) == reflect.TypeOf(snh) {
				// Entity have component that should not have, bail out
				return nil, false
			}
		}
	}
	matchedComponents := []*Component{}
	for _, sh := range shc {
		for eCom, filtering := range eComToFiltering {
			if !filtering {
				continue
			}
			if reflect.TypeOf((eCom)) == reflect.TypeOf(sh) {
				matchedComponents = append(matchedComponents, &eCom)
				if len(matchedComponents) == len(shc) {
					return matchedComponents, true
				}
				eComToFiltering[eCom] = false
				break

			}
		}
	}

	return nil, false
}
