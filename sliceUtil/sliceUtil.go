package sliceUtil

func removeWithOrder[T any](slice []T, s int) []T {
    return append(slice[:s], slice[s+1:])    
}

func removeWithoutOrder[T any](slice []T, s int) []T {
    slice[s] = slice[len(slice)-1]
    return slice[:len(slice)-1]
}
