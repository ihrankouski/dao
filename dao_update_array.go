package dao

import "labix.org/v2/mgo/bson"

//------------------------------------------------------------
// DAO update array methods
//------------------------------------------------------------

// Adds pushObj element to array pushTo.
func (dao *DAO) Update_ArrayPush(id bson.ObjectId, pushTo string, pushObj interface{}) (err error) {

	q := M{
		"$push": M{pushTo: pushObj},
	}
	err = dao.Coll.UpdateId(id, q)
	return
}

// Adds pushObjs elements to array pushTo.
func (dao *DAO) Update_ArrayPushMany(id bson.ObjectId, pushTo string, pushObjs []interface{}) (err error) {

	q := M{
		"$pushAll": M{pushTo: pushObjs},
	}
	err = dao.Coll.UpdateId(id, q)
	return
}

// Adds addObj element to array addTo only if it is not already there.
func (dao *DAO) Update_ArrayAdd(id bson.ObjectId, addTo string, addObj interface{}) (err error) {

	q := M{
		"$addToSet": M{addTo: addObj},
	}
	err = dao.Coll.UpdateId(id, q)
	return
}

// Adds addObjs elements to addTo array only if they are not already there.
func (dao *DAO) Update_ArrayAddMany(id bson.ObjectId, addTo string, addObjs []interface{}) (err error) {

	q := M{
		"$addToSet": M{addTo: M{"$each": addObjs}},
	}
	err = dao.Coll.UpdateId(id, q)
	return
}

// Removes pullObj element from pullFrom array.
// Adds pushObj element to pushTo array.
func (dao *DAO) Update_ArraysPullPush(id bson.ObjectId, pullFrom string, pullObj interface{}, pushTo string, pushObj interface{}) (err error) {

	q := M{
		"$pull": M{pullFrom: pullObj},
		"$push": M{pushTo: pushObj},
	}
	err = dao.Coll.UpdateId(id, q)
	return
}

// Removes pullObjs elements from pullFrom array.
// Adds pushObjs elements to pushTo array.
func (dao *DAO) Update_ArraysPullPushMany(id bson.ObjectId, pullFrom string, pullObjs []interface{}, pushTo string, pushObjs []interface{}) (err error) {

	q := M{
		"$pullAll": M{pullFrom: pullObjs},
		"$pushAll": M{pushTo: pushObjs},
	}
	err = dao.Coll.UpdateId(id, q)
	return
}

// Removes pullObjs elements from pullFrom array.
// Adds addObjs elements to addTo array only if they are not already there.
func (dao *DAO) Update_ArraysPullAddMany(id bson.ObjectId, pullFrom string, pullObjs []interface{}, addTo string, addObjs []interface{}) (err error) {

	q := M{
		"$pullAll": M{pullFrom: pullObjs},
		"$addToSet": M{addTo: M{"$each": addObjs}},
	}
	err = dao.Coll.UpdateId(id, q)
	return
}

// Removes pullFrom array element that matches pullObj.
func (dao *DAO) Update_ArrayPull(id bson.ObjectId, pullFrom string, pullObj interface{}) (err error) {

	q := M{
		"$pull": M{pullFrom: pullObj},
	}
	err = dao.Coll.UpdateId(id, q)
	return
}

// Removes pullFrom array elements that matches pullObjs.
func (dao *DAO) Update_ArrayPullMany(id bson.ObjectId, pullFrom string, pullObjs []interface{}) (err error) {

	q := M{
		"$pullAll": M{pullFrom: pullObjs},
	}
	err = dao.Coll.UpdateId(id, q)
	return
}
