// ReplaceInMap is a helper function to replace complete strings and substrings in a map
func ReplaceInMap(inputMap map[interface{}]interface{}, oldString, newString string) map[interface{}]interface{} {
	for key, value := range inputMap {
		switch concreteVal := value.(type) {
		case map[interface{}]interface{}:
			inputMap[key] = ReplaceInMap(concreteVal, oldString, newString)
		case []interface{}:
			for i, val := range concreteVal {
				valMap, ok := val.(map[interface{}]interface{})
				if ok {
					concreteVal[i] = ReplaceInMap(valMap, oldString, newString)
				} else {
					strVal, ok := val.(string)
					if ok {
						concreteVal[i] = strings.ReplaceAll(strVal, oldString, newString)
					}
				}
			}
		case string:
			inputMap[key] = strings.ReplaceAll(concreteVal, oldString, newString)
		}
	}
	return inputMap
}

// RecursiveRepalce is the entrypoint for a recursive walk of an unmarshalld YAML document, if the type is not a string it recurses
// using a helper function to find and replace strings
func RecursiveReplace(yamlFile []byte, oldString, newString string) ([]byte, error) {
	m := make(map[interface{}]interface{})

	err := yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling YAML: %w", err)
	}

	m = ReplaceInMap(m, oldString, newString)

	newYaml, err := yaml.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("error marshalling YAML: %w", err)
	}

	return newYaml, nil
}
