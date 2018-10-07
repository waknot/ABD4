package elasticsearch

import (
	"ABD4/API/utils"
	"context"
	"fmt"
	"io/ioutil"

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

// Index all the entities mapped in elasticsearch/esmapping/
// TODO ajouter une boucle for sur les fichiers pour indexer.
func Index(es *elastic.Client) error {
	var index = "users"
	exists, err := es.IndexExists(index).Do(context.Background())
	if err != nil {
		return err
	}

	if !exists {
		err = createIndex(es, index)
		if err != nil {
			return err
		}
	}

	//TODO penser à indexer depui la db pour chaque entité.

	return nil
}

// ReIndex delete and create index
// TODO: Quand on aura toutes les db il faudra faire en sorte qu'elle get tous les mappings :)
func ReIndex(es *elastic.Client, reindex bool) error {
	var index = "users"
	exists, err := es.IndexExists(index).Do(context.Background())
	if err != nil {
		return err
	}

	if exists {
		deleteIndex, err := es.DeleteIndex("users").Do(context.Background())
		if err != nil {
			// Handle error
			return err
		}
		if !deleteIndex.Acknowledged {
			return err
		}
	}

	err = Index(es)
	if err != nil {
		return err
	}

	return nil
}

//creatIndex create ES index
func createIndex(es *elastic.Client, index string) error {

	mapping, err := ioutil.ReadFile("ABD4/API/elasticsearch/esmapping/" + index + ".json")
	if err != nil {
		return err
	}

	mappingstr := string(mapping)
	createIndex, err := es.CreateIndex(index).BodyString(mappingstr).Do(context.Background())
	if err != nil {
		return err
	}
	if !createIndex.Acknowledged {
		return err
	}

	return nil
}
