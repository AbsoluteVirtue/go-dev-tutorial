// https://go.dev/doc/effective_go

// the "comma-ok" idiom
func offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok { // if tz isn't in the map, seconds is zeroed out, ok is false
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}

// min function that chooses the least of a list of integers
func Min(a ...int) int {
    min := int(^uint(0) >> 1)  // largest int
    for _, i := range a {
        if i < min {
            min = i
        }
    }
    return min
}
