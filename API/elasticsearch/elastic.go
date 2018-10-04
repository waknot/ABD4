package elasticsearch

import (
	"ABD4/API/utils"
	"fmt"

	elastic "gopkg.in/olivere/elastic.v5"
)

// Instanciate elasticsearch client for context
func Instanciate() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		msg := fmt.Errorf("%s elastic client Init failed: %s", utils.Use().GetStack(Instanciate), err.Error())
		return client, msg
	}

	return client, nil
}
