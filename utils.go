package main

func isIncluded(target string, list []string) bool {

	for _, el := range list {
		if target == el {
			return true
		}

	}

	return false

}
