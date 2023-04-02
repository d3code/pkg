package common_util

func ArrayContainsString(array []string, findString string) bool {
    for _, a := range array {
        if a == findString {
            return true
        }
    }
    return false
}
