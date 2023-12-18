package main

import (
	_ "embed"
	"os"
)

var (
	//go:embed delete.tpl
	deleteFile []byte
	//go:embed err.tpl
	errFile []byte
	//go:embed field.tpl
	fieldFile []byte
	//go:embed find-one.tpl
	findOneFile []byte
	//go:embed find-one-by-field.tpl
	findOneByField []byte
	//go:embed find-one-by-field-extra-method.tpl
	findOneByFieldExtraMethod []byte
	//go:embed import.tpl
	importFile []byte
	//go:embed import-no-cache.tpl
	importNoCacheFile []byte
	//go:embed insert.tpl
	insertFile []byte
	//go:embed interface-delete.tpl
	interfaceDeleteFile []byte
	//go:embed interface-find-one.tpl
	interfaceFindOneFile []byte
	//go:embed interface-find-one-by-field.tpl
	interfaceFindOneByFieldFile []byte
	//go:embed interface-insert.tpl
	interfaceInsertFile []byte
	//go:embed interface-update.tpl
	interfaceUpdateFile []byte
	//go:embed model.tpl
	modelFile []byte
	//go:embed model-gen.tpl
	modelGenFile []byte
	//go:embed model-new.tpl
	modelNewFile []byte
	//go:embed table-name.tpl
	tableNameFile []byte
	//go:embed tag.tpl
	tagName []byte
	//go:embed types.tpl
	typesFile []byte
	//go:embed update.tpl
	updateFile []byte
	//go:embed var.tpl
	varFile []byte
)

var fileMap = map[string][]byte{
	"delete.tpl":                         deleteFile,
	"err.tpl":                            errFile,
	"field.tlp":                          fieldFile,
	"find-one.tpl":                       findOneFile,
	"find-one-by-field.tpl":              findOneByField,
	"find-one-by-field-extra-method.tpl": findOneByFieldExtraMethod,
	"import.tpl":                         importFile,
	"import-no-cache.tpl":                importNoCacheFile,
	"insert.tpl":                         insertFile,
	"interface-delete.tpl":               interfaceDeleteFile,
	"interface-find-one.tpl":             interfaceFindOneFile,
	"interface-find-one-by-field.tpl":    interfaceFindOneByFieldFile,
	"interface-insert.tpl":               interfaceInsertFile,
	"interface-update.tpl":               interfaceUpdateFile,
	"model.tpl":                          modelFile,
	"model-gen.tpl":                      modelGenFile,
	"model-new.tpl":                      modelNewFile,
	"table-name.tpl":                     tableNameFile,
	"tag.tpl":                            tagName,
	"types.tpl":                          typesFile,
	"update.tpl":                         updateFile,
	"var.tpl":                            varFile,
}

func main() {
	for filename, data := range fileMap {
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		if _, err = file.Write(data); err != nil {
			panic(err)
		}

		file.Close()
	}

}
