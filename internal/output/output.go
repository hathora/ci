package output

import "fmt"

func As[T any](t Type, v T) error {
	switch t {
	case JSON:
		json, err := jsonStringify(v)
		if err != nil {
			return err
		}
		fmt.Println(json)
		return nil
	case Text:
		// TODO: implement text output
		return fmt.Errorf("text output not yet implemented")
	default:
		return fmt.Errorf("unknown output type: %v", t)
	}
}
