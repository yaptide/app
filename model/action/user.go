package action

import (
	"github.com/yaptide/yaptide/errors"
	"github.com/yaptide/yaptide/model"
	"github.com/yaptide/yaptide/model/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserGet ...
func (r *Resolver) UserGet(
	db mongo.DB, userID bson.ObjectId,
) (*model.User, error) {
	user := &model.User{}
	getErr := db.User().FindID(userID).One(user)
	if getErr != mgo.ErrNotFound {
		return nil, errors.ErrNotFound
	}
	if getErr != nil {
		return nil, errors.ErrInternalServerError
	}
	return user, nil
}

// UserRegister ...
func (r *Resolver) UserRegister(
	db mongo.DB, input *model.UserRegisterInput,
) (*model.User, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	user := input.ToUser()

	log.Debugf("Insert user into db %+v", user)
	insertErr := db.User().Insert(user)
	if insertErr != nil {
		return nil, errors.ErrInternalServerError
	}

	return user, nil
}

// UserLogin ...
func (r *Resolver) UserLogin(
	db mongo.DB, input *model.UserLoginInput,
) (string, *model.User, error) {
	user := &model.User{}
	if err := db.User().Find(bson.M{"username": input.Username}).One(user); err != nil {
		log.Warn(err)
		return "", nil, errors.ErrNotFound
	}
	if err := input.Validate(); err != nil {
		log.Warning(err.Error())
		return "", nil, err
	}
	passErr := input.ValidatePassword(user.PasswordHash)
	if passErr != nil {
		return "", nil, passErr
	}

	token, tokenErr := r.GenerateToken(user.ID)
	if tokenErr != nil {
		log.Warning(tokenErr.Error())
		return "", nil, errors.ErrInternalServerError
	}
	return token, user, nil
}
