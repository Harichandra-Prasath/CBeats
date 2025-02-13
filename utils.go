package main

func isAllowed(directive string) bool {

	for _, _directive := range ALLOWED_DIRECTIVES {
		if directive == _directive {
			return true
		}

	}

	return false

}
