package models

import (
	"errors"
	"time"
	"gopkg.in/mgo.v2/bson"
	"ritefruit-api/models/db"
)

func init() {

}

func newUsersCollection() *db.Collection {
	return db.NewCollectionSession("users")
}

type User struct {
	ID bson.ObjectId `json:"id" bson:"_id"`
	UserName string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	FirstName string `json:"firstname" bson:"firstname"`
	LasttName string `json:"lastname" bson:"lastname"`
	RefCount  int    `json:"refcount" bson:"refcount"`
	EmailId  int    `json:"email-id" bson:"email-id"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`

}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

func AddUser(u User) string {
	u.ID = bson.NewObjectId()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.RefCount = 0

	// Get post collection connection
	c := newUsersCollection()
	defer c.Close()

	err := c.Session.Insert(&u)

	if err != nil {
		return "not created"
	}
	return u.ID.String()
}

func GetUser(uid string) (User, error) {
	var (
		err error
		u User
	)
	uId := bson.ObjectIdHex(uid)
	// Get post collection connection
	c := newUsersCollection()
	defer c.Close()
	// get post
	err = c.Session.FindId(uId).One(&u)
	if err != nil {
		return u,err
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() ([]User) {
	var (
		//err error
		u []User
	)
	// Get post collection connection
	c := newUsersCollection()
	defer c.Close()
	//err =
		c.Session.Find(bson.M{}).All(&u)
	return u
}

func UpdateUser(uID bson.ObjectId, uu *User) (*User,  error) {
	var (
		err error
		u *User
	)
	// Get post collection connection
	c := newUsersCollection()
	defer c.Close()
	// get post
	err = c.Session.FindId(uID).One(&u)
	if err != nil {
		//FIXME do update properly
		return u,err
	}
	return nil, errors.New("User not exists")
}

func Login(username, password string) bool {
	return true
}

func DeleteUser(uID bson.ObjectId) (error){
	var (
		err error
	)
	// Get post collection connection
	c := newUsersCollection()
	defer c.Close()
	// remove post
	err = c.Session.Remove(bson.M{"_id": uID})
	if err != nil {
		return err
	}
	return err
}
