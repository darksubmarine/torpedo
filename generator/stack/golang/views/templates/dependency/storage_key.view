package dependency

import (
	"github.com/darksubmarine/torpedo-lib-go/app"
	"github.com/darksubmarine/torpedo-lib-go/conf"
)

type StorageKeyProvider struct {
	app.BaseProvider

	key []byte `torpedo.di:"provide,name=STORAGE_KEY"`

	// private fields initialized by constructor
	cfg conf.Map
}

func NewStorageKeyProvider(config conf.Map) *StorageKeyProvider {
	return &StorageKeyProvider{cfg: config}
}

// Provide provides the storage key instance.
func (p *StorageKeyProvider) Provide(c app.IContainer) error {
	p.key = []byte(
		p.cfg.FetchStringOrElse(
			"the-key-has-to-be-32-bytes-long!",
			"key"))

	return nil
}
