package slice_utils

import "reflect"

func RemoveString(slice []string, entry string) []string {
    var result []string
    for _, value := range slice {
        if value != entry {
            result = append(result, value)
        }
    }
    return result
}

func ContainsString(slice []string, entry string) bool {
    for _, value := range slice {
        if value == entry {
            return true
        }
    }
    return false
}

func Keys(m interface{}) []string {
    v := reflect.ValueOf(m)
    if v.Kind() != reflect.Map {
        return nil
    }
    keys := v.MapKeys()
    result := make([]string, len(keys))
    for i, key := range keys {
        result[i] = key.String()
    }
    return result
}
