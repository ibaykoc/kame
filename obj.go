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
		} else if strings.HasPrefix(line, "f ") {
			line = strings.Replace(line, "f ", "", 1)
			for _, face := range strings.Split(line, " ") {
				f := Face{}
				fmt.Sscanf(face, "%d/%d/%d", &f.V, &f.VT, &f.VN)
				faces = append(faces, f)
			}
		}
	}
	facesToIndices := make(map[Face]uint32)

	verts := make(map[Face]Vertex)
	for _, face := range faces {
		if _, found := facesToIndices[face]; !found {
			facesToIndices[face] = uint32(len(facesToIndices))
		}
		verts[face] = Vertex{
			position: positions[face.V-1],
			uv:       uvs[face.VT-1],
			normal:   normals[face.VN-1],
		}
	}

	indices := make([]uint32, len(faces))
	for i := 0; i < len(faces); i++ {
		indices[i] = facesToIndices[faces[i]]
	}

	positionData := make([]float32, len(facesToIndices)*3)
	uvData := make([]float32, len(facesToIndices)*2)
	normalData := make([]float32, len(facesToIndices)*3)
	for face, indice := range facesToIndices {
		v := verts[face]
		positionI := indice * 3
		uvI := indice * 2
		normalI := indice * 3
		positionData[positionI] = v.position.X
		positionData[positionI+1] = v.position.Y
		positionData[positionI+2] = v.position.Z
		uvData[uvI] = v.uv.X
		uvData[uvI+1] = v.uv.Y
		normalData[normalI] = v.normal.X
		normalData[normalI+1] = v.normal.Y
		normalData[normalI+2] = v.normal.Z
	}

	// fmt.Printf("verts: %v\n\n", verts)
	// fmt.Println("")
	// fmt.Println(facesToIndices)
	// fmt.Println(indices)

	model, err := CreateDrawableModel2(
		positionData,
		uvData,
		normalData,
		indices,
		texturePath,
	)
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
