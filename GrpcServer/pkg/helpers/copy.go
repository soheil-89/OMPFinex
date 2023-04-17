package helpers

import "github.com/jinzhu/copier"

func Inject(p interface{}, c interface{}) {
	_ = copier.Copy(p, c)
}
