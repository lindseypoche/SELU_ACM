package elasticsearch

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"github.com/cmd-ctrl-q/SELU_HEX_TEST/internal/logger"
// 	"github.com/olivere/elastic"
// )

// var (
// 	Client esClientInterface = &esClient{}
// )

// type esClientInterface interface {
// 	setClient(c *elastic.Client)
// 	Index(string, interface{}) (*elastic.IndexResponse, error) // for indexing a new document in ES
// }

// type esClient struct {
// 	client *elastic.Client
// }

// func Init() {
// 	log := logger.GetLogger()
// 	// create a new client
// 	client, err := elastic.NewClient(
// 		elastic.SetURL("http://127.0.0.1:9200"),
// 		elastic.SetHealthcheckInterval(10*time.Second),
// 		elastic.SetErrorLog(log),
// 		elastic.SetInfoLog(log),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// set global var Client to the created client
// 	Client.setClient(client)

// 	// Create index (ie es database) if it does not exist.
// }

// func (c *esClient) setClient(client *elastic.Client) {
// 	c.client = client
// }

// func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
// 	ctx := context.Background()
// 	// return c.client.Index().Do(ctx)
// 	result, err := c.client.Index().
// 		Index(index).
// 		BodyJson(doc).
// 		Do(ctx)

// 	if err != nil {
// 		// put logger where error occurs
// 		logger.Error(
// 			fmt.Sprintf("error when trying to index document in es: %s", index), err)
// 		return nil, err
// 	}
// 	return result, nil
// }
