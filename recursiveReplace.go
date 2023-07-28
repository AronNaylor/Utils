func RecursiveReplace(data interface{}, target string, replacement string) {
	v := reflect.ValueOf(data)

	switch v.Kind() {
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i)
			if item.Kind() == reflect.String {
				fmt.Println("*****************************")
				if strings.Contains(item.String(), target) {
					fmt.Println("Match found:", item.String())
				}
				v.Index(i).SetString(strings.ReplaceAll(item.String(), target, replacement))
			} else {
				RecursiveReplace(item.Interface(), target, replacement)
			}
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			value := v.MapIndex(key)
			if value.Kind() == reflect.String {
				if strings.Contains(value.String(), target) {
					fmt.Println("Match found:", value.String())
				}
				v.SetMapIndex(key, reflect.ValueOf(strings.ReplaceAll(value.String(), target, replacement)))
			} else {
				RecursiveReplace(value.Interface(), target, replacement)
			}
		}
	case reflect.String:
		fmt.Println("=========================")
		if strings.Contains(v.String(), target) {
			fmt.Println("Match found:", v)
		}
	}
}
