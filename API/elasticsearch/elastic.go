package elasticsearch

import (
	"ABD4/API/utils"
	"context"
	"fmt"
	"io/ioutil"

	elastic "gopkg.in/olivere/elastic.v5"
)

// Instanciate elasticsearch client for context
func Instanciate(serv string) (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://" + serv + ":9200"))
	if err != nil {
		msg := fmt.Errorf("%s elastic client Init failed: %s", utils.Use().GetStack(Instanciate), err.Error())
		return client, msg
	}

	return client, nil
}

//IndexAll launch indexation or re-indexation
func IndexAll(es *elastic.Client, reindex bool) error {
	files, err := ioutil.ReadDir("ABD4/API/elasticsearch/esmapping/")
	if err != nil {
		return err
	}

	if reindex {
		for _, file := range files {
			ReIndex(es, utils.NoFileExtension(file.Name()))
		}
	} else {
		for _, file := range files {
			Index(es, utils.NoFileExtension(file.Name()))
		}
	}

	return nil

}

// Index all the entities mapped in elasticsearch/esmapping/
// TODO ajouter une boucle for sur les fichiers pour indexer.
func Index(es *elastic.Client, index string) error {
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
func ReIndex(es *elastic.Client, index string) error {
	exists, err := es.IndexExists(index).Do(context.Background())
	if err != nil {
		return err
	}

	if exists {
		deleteIndex, err := es.DeleteIndex(index).Do(context.Background())
		if err != nil {
			// Handle error
			return err
		}
		if !deleteIndex.Acknowledged {
			return err
		}
	}

	err = Index(es, index)
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

	//Casts to string for BodyString method
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
