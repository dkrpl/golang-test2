package pasien

import (
	"golang-test2/models"
	schemas_pasien "golang-test2/schemas/pasien"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewPasienRepository(db *mgo.Database) PasienRepository {
	return &PasienRepositoryImpl{
		Collection: db.C("pasiens"),
		KamarPasienRepositoryImpl: &KamarPasienRepositoryImpl{
			Collection: db.C("pasien"),
		},
	}
}

type PasienRepositoryImpl struct {
	Collection                *mgo.Collection
	KamarPasienRepositoryImpl *KamarPasienRepositoryImpl
}

type KamarPasienRepositoryImpl struct {
	Collection *mgo.Collection
}

func (this_ *PasienRepositoryImpl) Create(schemas *schemas_pasien.Input) error {
	id := bson.NewObjectId()
	return this_.Collection.Insert(bson.M{
		"_id":         id,
		"nama_pasien": schemas.Nama_pasien,
		"alamat":      schemas.Alamat,
		"no_rumah":    schemas.No_rumah,
		"kecamatan":   schemas.Kecamatan,
		"kabupaten":   schemas.Kabupaten,
		"kamars":      []interface{}{},
	})
}

func (this_ *PasienRepositoryImpl) Edit(id bson.ObjectId, schemas schemas_pasien.Edit) error {
	err := this_.Collection.Update(bson.M{"_id": id},
		bson.M{"$set": bson.M{
			"nama_pasien": schemas.Nama_pasien,
			"alamat":      schemas.Alamat,
			"no_rumah":    schemas.No_rumah,
			"kecamatan":   schemas.Kecamatan,
			"kabupaten":   schemas.Kabupaten,
		}})
	return err
}

func (this_ *PasienRepositoryImpl) List(query map[string]interface{}, per_page, page_no int) ([]models.Pasien, int) {
	var pasiens []models.Pasien
	this_.Collection.Find(query).Skip((page_no - 1) * per_page).Limit(per_page).All(&pasiens)
	count, _ := this_.Collection.Find(query).Count()
	return pasiens, count
}

func (this_ *PasienRepositoryImpl) Get(id bson.ObjectId) models.Pasien {
	var pasien models.Pasien
	this_.Collection.Find(bson.M{"_id": id}).One(&pasien)
	return pasien
}

func (this_ *PasienRepositoryImpl) Delete(id bson.ObjectId) error {
	return this_.Collection.Remove(bson.M{"_id": id})
}

func (this_ *PasienRepositoryImpl) GetByNo_rumah(no_rumah string) (pasien models.Pasien) {
	this_.Collection.Find(bson.M{"no_rumah": no_rumah}).One(&pasien)
	return pasien
}

func (this_ *PasienRepositoryImpl) GetByKamar(kamar string) (pasien []models.Pasien, err error) {
	err = this_.Collection.Find(bson.M{
		"kamars." + kamar: bson.M{
			"$exists": true,
		},
	}).All(&pasien)
	return
}
