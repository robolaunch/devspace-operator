package hybrid

func ContainsInstance(instances []string, instance string) bool {

	for _, v := range instances {
		if v == instance {
			return true
		}
	}

	return false
}
