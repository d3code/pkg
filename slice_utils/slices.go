package slice_utils

func RemoveString(slice []string, entry string) []string {
    var result []string
    for _, value := range slice {
        if value != entry {
            result = append(result, value)
        }
    }
    return result
}
