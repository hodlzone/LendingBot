package mongo

import (
	"gopkg.in/mgo.v2"
)

const (
	USER_DB        = "userdb"
	USER_DB_TEST   = "userdb_test"
	USER_DB_C_USER = "user_c"

	STAT_DB = "statdb"

	C_UserStat_POL = "poloniexUserStat"
	C_LendHist_POL = "poloniexLendingHist"
	C_Exchange_POL = "poloniexExchange"

	C_UserStat_BIT = "bitfinexUserStat"
	C_LendHist_BIT = "bitfinexLendingHist"
	C_Exchange_BIT = "bitfinexExchange"
)

type MongoDB struct {
	uri         string
	DbName      string
	baseSession *mgo.Session
}

func (c *MongoDB) GetURI() string {
	return c.uri
}

func createMongoDB(uri string, dbname string) *MongoDB {
	mongoDB := &MongoDB{
		uri,
		dbname,
		nil,
	}
	return mongoDB
}

//DONT USE THIS ONE
func (c *MongoDB) CreateSession() (session *mgo.Session, err error) {

	if c.baseSession == nil {
		session, err = mgo.Dial(c.uri)
		if err != nil {
			return nil, err
		}
		c.baseSession = session
	} else {
		session = c.baseSession.Clone()
	}

	// See https://godoc.org/labix.org/v2/mgo#Session.SetMode
	session.SetMode(mgo.Monotonic, true)

	return
}

func (c *MongoDB) GetCollection(collectionName string) (*mgo.Session, *mgo.Collection, error) {
	session, err := c.CreateSession()
	if err != nil {
		return nil, nil, err
	}

	return session, session.DB(c.DbName).C(collectionName), nil
}

func CreateUserDB(uri string) (*MongoDB, error) {
	db := createMongoDB(uri, USER_DB)

	session, err := db.CreateSession()
	if err != nil {
		return nil, err
	}

	c := session.DB(USER_DB).C(USER_DB_C_USER)

	index := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateStatDB(uri string) (*MongoDB, error) {
	db := createMongoDB(uri, STAT_DB)

	session, err := db.CreateSession()
	if err != nil {
		return nil, err
	}

	c := session.DB(STAT_DB).C(C_UserStat_POL)

	index := mgo.Index{
		Key:        []string{"time"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		return nil, err
	}

	c = session.DB(STAT_DB).C(C_LendHist_POL)
	index = mgo.Index{
		Key:        []string{},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		return nil, err
	}

	c = session.DB(STAT_DB).C(C_Exchange_POL)
	index = mgo.Index{
		Key:        []string{},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTestUserDB(uri string) (*MongoDB, error) {
	db := createMongoDB(uri, USER_DB_TEST)

	session, err := db.CreateSession()
	if err != nil {
		return nil, err
	}

	c := session.DB(USER_DB_TEST).C(USER_DB_C_USER)

	index := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		return nil, err
	}
	return db, nil
}
