package storage

import "github.com/pankratsdarya/linksshorter/links"

// связь с хранилищем.

// здесь должно быть:
// открытие файла, запись и чтение из него (обложить мьютексом!)
// подумать: писать все в мапу, а скидывать ее в файл каждые сколько-то секунд.

var Storage map[string]links.Link
