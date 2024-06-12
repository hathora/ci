package altsrc

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v3"
)

func ConfigFile(fileFlag string, valuePath string) cli.ValueSource {
	return &fileValueSource{
		fileFlag:  fileFlag,
		valuePath: valuePath,
	}
}

type fileValueSource struct {
	fileFlag  string
	fileName  string
	fileType  string
	valuePath string

	cache *fileValueCache
}

type fileValueCache struct {
	fileName string
	values   map[string]any
}

var _ FlagInitializedValueSource = (*fileValueSource)(nil)

func (fvs *fileValueSource) Lookup() (string, bool) {
	if err := fvs.populateFileCache(); err != nil {
		return "", false
	}

	if fvs.cache == nil {
		return "", false
	}

	if val, ok := fvs.resolvePath(fvs.cache.values); ok {
		return fmt.Sprintf("%[1]v", val), true
	}

	return "", false

}

func (fvs *fileValueSource) String() string {
	return fmt.Sprintf("file %[1]q at path %[2]q", fvs.fileName, fvs.valuePath)
}

func (fvs *fileValueSource) GoString() string {
	return fmt.Sprintf(
		"&fileValueSource{fileFlag:%[1]q,fileName:%[2]q,fileType:%[3]q,valuePath:%[4]q}",
		fvs.fileFlag,
		fvs.fileName,
		fvs.fileType,
		fvs.valuePath,
	)
}

func (fvs *fileValueSource) Initialize(cmd *cli.Command) error {
	if fvs.fileName == "" && fvs.fileFlag != "" {
		fvs.fileName = cmd.String(fvs.fileFlag)
	}

	if fvs.fileName == "" {
		fvs.cache = &fileValueCache{values: map[string]any{}}
		return nil
	}

	cacheMetadataKey := fmt.Sprintf("fileValueCache#%s", fvs.fileName)
	cache, hasKey := cmd.Metadata[cacheMetadataKey]
	if !hasKey {
		cache = &fileValueCache{fileName: fvs.fileName}
		cmd.Metadata[cacheMetadataKey] = cache
	}

	typedCache, ok := cache.(*fileValueCache)
	if !ok {
		return fmt.Errorf("metadata key %q has unexpected type", cacheMetadataKey)
	}

	fvs.cache = typedCache
	return nil
}

func (fvs *fileValueSource) resolvePath(values map[string]any) (any, bool) {
	if fvs.valuePath == "" {
		return values, true
	}

	sections := strings.Split(fvs.valuePath, ".")
	switch len(sections) {
	case 1:
		if val, ok := values[sections[0]]; ok {
			return val, true
		}
	default:
		node := values
		for _, section := range sections[:len(sections)-1] {
			child, ok := node[section]
			if !ok {
				return nil, false
			}

			switch child := child.(type) {
			case map[string]any:
				node = make(map[string]any, len(child))
				for k, v := range child {
					node[k] = v
				}
			default:
				return nil, false
			}
		}
		if val, ok := node[sections[len(sections)-1]]; ok {
			return val, true
		}
	}

	return nil, false
}

func (fvs *fileValueSource) determineFileType() error {
	if fvs.fileType != "" {
		return nil
	}

	if fvs.fileName == "" {
		return fmt.Errorf("file name is required in order to determine file type")
	}

	ext := filepath.Ext(fvs.fileName)
	if ext == "" {
		return fmt.Errorf("file name %q does not have an extension", fvs.fileName)
	}

	switch ext {
	case ".json":
		fvs.fileType = "json"
	case ".yaml", ".yml":
		fvs.fileType = "yaml"
	default:
		return fmt.Errorf("unsupported file type %q", ext)
	}

	return nil
}

func (fvs *fileValueSource) populateFileCache() error {
	if fvs.cache == nil {
		fvs.cache = &fileValueCache{fileName: fvs.fileName}
	}

	if fvs.cache.values != nil {
		return nil
	}

	if fvs.fileType == "" {
		if err := fvs.determineFileType(); err != nil {
			return err
		}
	}

	switch fvs.fileType {
	case "json": // json is a subset of yaml
		err := fvs.cache.populateYAML()
		if err != nil {
			return err
		}
	case "yaml":
		err := fvs.cache.populateYAML()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported file type %q", fvs.fileType)
	}

	return nil
}

func (fvc *fileValueCache) populateYAML() error {
	if fvc.fileName == "" {
		return fmt.Errorf("file name is required in order to populate values from YAML")
	}
	if fvc.values != nil {
		return nil
	}

	values, err := yamlUnmarshalFile(fvc.fileName)
	if err != nil {
		return err
	}

	fvc.values = values
	return nil
}

func yamlUnmarshalFile(filePath string) (map[string]any, error) {
	b, err := readURI(filePath)
	if err != nil {
		return nil, err
	}

	values := map[string]any{}
	if err := yaml.Unmarshal(b, &values); err != nil {
		return nil, err
	}

	return values, nil
}
