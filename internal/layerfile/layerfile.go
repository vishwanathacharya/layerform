package layerfile

import (
	"encoding/json"
	"os"
	"path"

	"github.com/pkg/errors"

	"github.com/ergomake/layerform/internal/data/model"
)

type layerfile struct {
	sourceFilepath string           `json:"-"`
	Layers         []layerfileLayer `json:"layers"`
}

type layerfileLayer struct {
	Name         string   `json:"name"`
	Files        []string `json:"files"`
	Dependencies []string `json:"dependencies"`
}

func FromFile(sourceFilepath string) (*layerfile, error) {
	bs, err := os.ReadFile(sourceFilepath)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to read %s", sourceFilepath)
	}

	lf := &layerfile{sourceFilepath: sourceFilepath}
	err = json.Unmarshal(bs, lf)

	return lf, errors.Wrapf(err, "fail to decode %s into layerfile", lf)
}

func (lf *layerfile) ToLayers() ([]*model.Layer, error) {
	dir := path.Dir(lf.sourceFilepath)

	modelLayers := make([]*model.Layer, len(lf.Layers))
	for i, l := range lf.Layers {
		files := make([]model.LayerFile, len(l.Files))
		for j, f := range l.Files {
			fpath := path.Join(dir, f)
			content, err := os.ReadFile(fpath)
			if err != nil {
				return nil, errors.Wrapf(err, "could not read %s", fpath)
			}

			files[j] = model.LayerFile{
				Path:    f,
				Content: content,
			}
		}

		modelLayers[i] = &model.Layer{
			Name:         l.Name,
			Files:        files,
			Dependencies: l.Dependencies,
		}
	}

	return modelLayers, nil
}