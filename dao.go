package dao

import (
	"labix.org/v2/mgo"
)

//------------------------------------------------------------
// Open/close methods
//------------------------------------------------------------

// Opens new session with collection. Don't forget to close.
func (dao *DAO) Open(servername, dbname, collname string, idxs *CollIndexes) {

	// Validate

	if servername == "" {
		panic("DAO received empty string as server name!")
	}

	if dbname == "" {
		panic("DAO received empty string as database name!")
	}

	if collname == "" {
		panic("DAO received empty string as collection name!")
	}

	if idxs == nil {
		panic("DAO received nil as index!")
	}

	if idxs.indexes == nil {
		panic("DAO received indexes with nil array!")
	}

	// Create session
	dao.server = servername
	dao.dbname = dbname
	dao.collname = collname
	dao.indexes = idxs
	dao.sess, dao.Coll = getSession(dao.server, dao.dbname, dao.collname)

	if dao.sess == nil || dao.Coll == nil {
		panic("DAO cannot acquire DB session. Perhaps connection was closed or DAO wasn't configured?")
	}

	// Ensure indexes
	dao.EnsureIndexes()
}

// Changes consistency to strong.
func (dao *DAO) SetStrongMode() {
	if dao.sess != nil {
		dao.sess.SetMode(mgo.Strong, true)
	}
}

// Closes session.
func (dao *DAO) Close() {
	if dao.sess != nil {
		dao.sess.Close()
	}
}
