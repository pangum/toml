package yaml

import (
	"github.com/pangum/pangu"
	"github.com/pangum/toml/internal"
)

func init() {
	pangu.New().Config().Loader(internal.NewLoader()).Build()
}
