package Customer

import (
	"HepsiGonulden/Customer/types"
	"HepsiGonulden/config"
	"context"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

/*
Unit Test: Birim testidir. Kodun butun dis bagimliliklari (mongodb, redis, kafka ...) calistirilmaksizin, MOCK'lanarak yapilan testlerdir.
Integration Test: Gercekten mongodb, redis, kafka gibi fiziksel makinelere baglanip yapilan testlerdir.
*/

func TestFindByID(t *testing.T) {
	/*
	   - DONE / DB'de kayit varken kaydin donmesi
	   - TODO(Umut) / DB'de kayit yokken nil donmesi ama err donmemesi
	   - TODO(Umut) / DB baglantisinda hata olusursa err'in nil olmamasi
	*/

	config.Init() // TODO(Umut): repository kodum su anda test edilmeye uygun olmadigi icin, config.Init cagiriliyor

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).DatabaseName("customer_db"))

	mt.Run("find customer by id", func(mt *mtest.T) {
		dbCustomer := types.Customer{
			Id:        uuid.NewString(),
			FirstName: "Umut",
			LastName:  "Com",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(0, "foo.bar", mtest.FirstBatch, primitive.D{
			{Key: "_id", Value: dbCustomer.Id},
			{Key: "first_name", Value: dbCustomer.FirstName},
			{Key: "last_name", Value: dbCustomer.LastName},
		}))

		repo, err := NewRepository(mt.Client)
		assert.Nil(t, err)

		customer, err := repo.FindByID(context.Background(), dbCustomer.Id)
		assert.Nil(t, err)
		assert.NotNil(t, customer)
		assert.Equal(t, dbCustomer.Id, customer.Id)
		assert.Equal(t, dbCustomer.FirstName, customer.FirstName)
		assert.Equal(t, dbCustomer.LastName, customer.LastName)
	})
}
