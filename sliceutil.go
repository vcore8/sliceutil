package sliceutil

// Check if a slice contains a string element and return bool
func sliceContainsElement(slice []string, element string) bool {
	retval := false
	for _, e := range slice {
		if e == element {
			retval = true
		}
	}
	return retval
}

// Remove duplicated items on slice of strings
func sliceUniqueElements(slice []string) []string {
    keys := make(map[string]bool)
    list := []string{}
    for _, entry := range slice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            if len(entry) > 2 {
                list = append(list, entry)
            }
        }
    }
    return list
}
// Return the difererence between 2 slices of strings
func sliceDifference(slice1 []string, slice2 []string) []string {
    var diff []string

    for i := 0; i < 2; i++ {
        for _, s1 := range slice1 {
            found := false
            for _, s2 := range slice2 {
                if s1 == s2 {
                    found = true
                    break
                }
            }
            if !found {
                diff = append(diff, s1)
            }
        }
        if i == 0 {
            slice1, slice2 = slice2, slice1
        }
    }
    return diff
}


// merge 2 slices of strings
func Merge(x, y []string]) interface{} {
	return append(x, y)
}


func Count(x []string]) int {
	return x.Len()
}

func Len(x []string]) int{
	return len(x)
}


