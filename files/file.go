package files

import (
    "fmt"
    "os"
)

func Exist(file string) bool {
    _, err := os.Stat(file)
    return !os.IsNotExist(err)
}

func Save(path string, fileName string, byteArray []byte, overwrite bool) error {
    err := os.MkdirAll(path, 0755)
    if err != nil {
        return err
    }

    file := fmt.Sprintf("%s/%s", path, fileName)

    if !overwrite && Exist(file) {
        return fmt.Errorf("file already exists")
    }

    return os.WriteFile(file, byteArray, 0666)
}
