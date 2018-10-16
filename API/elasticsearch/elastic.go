package elasticsearch

import (
	"ABD4/API/utils"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

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
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	mappingDir := filepath.Join(dir, "API/elasticsearch/esmapping/")
	fmt.Printf("%s mapping dir: %s", utils.Use().GetStack(IndexAll), mappingDir)
	files, err := ioutil.ReadDir(mappingDir)
	if err != nil {
		return fmt.Errorf("%s %s", utils.Use().GetStack(Instanciate), err.Error())
	}

	if reindex {
		for _, file := range files {
			err = ReIndex(es, utils.NoFileExtension(file.Name()))
		}
	} else {
		for _, file := range files {
			err = Index(es, utils.NoFileExtension(file.Name()))
		}
	}
	return err
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

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	mappingDir := filepath.Join(dir, "API/elasticsearch/esmapping/")
	fmt.Printf("%s mapping dir: %s", utils.Use().GetStack(IndexAll), mappingDir)
	mapping, err := ioutil.ReadFile(filepath.Join(mappingDir, index+".json"))
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
