package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"

	"github.com/pushthat/nboard-engine/internal/kv"
)

type API struct {
	table dynamo.Table
}

type SingleTableElem struct {
	Pk    string `dynamo:"pk"`
	Sk    string `dynamo:"sk"`
	Value []byte `dynamo:"value"`
}

func New(tableName, region string) *API {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String(region)})
	table := db.Table(tableName)

	return &API{
		table: table,
	}
}

func (d *API) Get(pk kv.PK, sk kv.SK) ([]byte, error) {
	var value SingleTableElem
	err := d.table.Get("pk", string(pk)).Range("sk", dynamo.Equal, string(sk)).One(&value)
	return value.Value, err
}

func (d *API) Set(pk kv.PK, sk kv.SK, content []byte) error {
	s := SingleTableElem{Pk: string(pk), Sk: string(sk), Value: content}
	return d.table.Put(s).Run()
}

func (d *API) Delete(pk kv.PK, sk kv.SK) error {
	return d.table.Delete("pk", string(pk)).Range("sk", string(sk)).Run()
}
