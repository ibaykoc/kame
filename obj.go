package kame

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Vertex struct {
	position V3f
	uv       V2f
	normal   V3f
}

type V3f struct {
	X, Y, Z float32
}

type V2f struct {
	X, Y float32
}

type Face struct {
	V, VT, VN int
}

func newVertex() Vertex {
	return Vertex{}
}

func LoadOBJ(filePath string, texturePath string) (DrawableModel, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return DrawableModel{}, err
	}

	scanner := bufio.NewScanner(file)

	positions := make([]V3f, 0)
	uvs := make([]V2f, 0)
	normals := make([]V3f, 0)
	faces := make([]Face, 0)

	// Scan position, uv, and normal
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "v ") {
			pos := V3f{}
			fmt.Sscanf(line, "v %f %f %f\n", &pos.X, &pos.Y, &pos.Z)
			positions = append(positions, pos)
		} else if strings.HasPrefix(line, "vt ") {
			uv := V2f{}
			fmt.Sscanf(line, "vt %f %f\n", &uv.X, &uv.Y)
			uvs = append(uvs, uv)
		} else if strings.HasPrefix(line, "vn ") {
			nor := V3f{}
			fmt.Sscanf(line, "vn %f %f %f\n", &nor.X, &nor.Y, &nor.Z)
			normals = append(normals, nor)
		} else if strings.HasPrefix(line, "s ") {
			break
		}
	}

	hasUV := len(uvs) > 0
	hasNormal := len(normals) > 0

	// scan faces
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "f ") {
			line = strings.Replace(line, "f ", "", 1)
			for _, face := range strings.Split(line, " ") {
				f := Face{}
				if !hasUV {
					fmt.Sscanf(face, "%d//%d", &f.V, &f.VN)
				} else {
					fmt.Sscanf(face, "%d/%d/%d", &f.V, &f.VT, &f.VN)
				}
				faces = append(faces, f)
			}
		}
	}
	// Keep track of which face correspond to wich element
	faceToElement := make(map[Face]uint32)

	// Every unique face on obj file should have its own vertex
	faceToVertex := make(map[Face]Vertex)

	for _, face := range faces {
		if _, found := faceToElement[face]; !found {
			faceToElement[face] = uint32(len(faceToElement))
			v := Vertex{}
			v.position = positions[face.V-1]
			if hasUV {
				v.uv = uvs[face.VT-1]
			}
			if hasNormal {
				v.normal = normals[face.VN-1]
			}
			faceToVertex[face] = v
		}
	}

	// Convert faces to elements
	elements := make([]uint32, len(faces))
	for i := 0; i < len(faces); i++ {
		elements[i] = faceToElement[faces[i]]
	}

	// At this point we've got all the data sorted
	// Now convert it so we can create drawable model
	positionData := make([]float32, len(faceToElement)*3)
	var uvData []float32
	var normalData []float32

	if hasUV {
		uvData = make([]float32, len(faceToElement)*2)
	}
	if hasNormal {
		normalData = make([]float32, len(faceToElement)*3)
	}

	for face, element := range faceToElement {
		v := faceToVertex[face]
		positionI := element * 3
		positionData[positionI] = v.position.X
		positionData[positionI+1] = v.position.Y
		positionData[positionI+2] = v.position.Z
		if hasUV {
			uvI := element * 2
			uvData[uvI] = v.uv.X
			uvData[uvI+1] = v.uv.Y
		}
		if hasNormal {
			normalI := element * 3
			normalData[normalI] = v.normal.X
			normalData[normalI+1] = v.normal.Y
			normalData[normalI+2] = v.normal.Z
		}
	}

	var model DrawableModel
	if hasUV && hasNormal {
		model, err = CreateDrawableModel3T(
			positionData,
			uvData,
			normalData,
			elements,
			texturePath,
		)
	} else if hasUV && !hasNormal {
		model, err = CreateDrawableModel1T(positionData, uvData, elements, texturePath)
	} else if !hasUV && hasNormal {
		fmt.Printf("OBJ: %s has no UV data, ignoring texture\n", filePath)
		model, err = CreateDrawableModel2(positionData, normalData, elements)
	} else {
		fmt.Printf("OBJ: %s has no UV data, ignoring texture\n", filePath)
		model, err = CreateDrawableModel0(positionData, elements)
	}
	if err != nil {
		return DrawableModel{}, err
	}
	return model, nil
}

func panciCheck(err error) {
	if err != nil {
		panic(err)
	}
}
