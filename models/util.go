package models

func getBool(i interface{}) bool {
	switch v := i.(type) {
	case float64:
		if v > 0 {
			return true
		} else {
			return false
		}
	case int:
		if v > 0 {
			return true
		} else {
			return false
		}
	case bool:
		return v
	case string:
		switch v {
		case "0", "false", "Flase":
			return false
		default:
			return true
		}
	default:
		return false
	}

}
