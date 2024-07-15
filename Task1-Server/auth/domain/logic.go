package domain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/golang-jwt/jwt"
)

const (
	salt       = "euv45675kdfjd458dhg43"
	signingKey = "dfjh47ty34hfd89wofdhf"
	tokenTTL   = 4 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
	Role   int    `json:"role"`
}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, l log.Logger) Service {
	return &service{
		repo:   repo,
		logger: l,
	}
}

func (s *service) CreateUser(user User) error {
	logger := log.With(s.logger, "method", "CreateUser")

	user.Password = s.generatePassHash(user.Password)
	user.Role = RoleUser
	err := s.repo.CreateUser(user)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}

	logger.Log("successfully created user", user.Id.Hex())

	return nil
}

func (s *service) UserById(id string) (User, error) {
	logger := log.With(s.logger, "method", "UserById")

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return User{}, err
	}

	user, err := s.repo.UserById(objId)
	if err != nil {
		level.Error(logger).Log("err", err)
		return User{}, err
	}

	logger.Log("successfully got user by id", user)

	return user, nil
}

func (s *service) UpdateUser(user UserUpdate) error {
	logger := log.With(s.logger, "method", "UpdateUser")

	objId, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}
	if user.Password != nil {
		*user.Password = s.generatePassHash(*user.Password)
	}

	err = s.repo.UpdateUser(objId, user)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}

	logger.Log("successfully updated user", user)

	return nil
}

func (s *service) DeleteUser(id string) error {
	logger := log.With(s.logger, "method", "DeleteUser")

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}

	err = s.repo.DeleteUser(objId)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}

	logger.Log("successfully deleted user", id)

	return nil
}

func (s *service) BecomeVolunteer(userID string) error {
	logger := log.With(s.logger, "method", "BecomeVolunteer")

	userObjId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}

	err = s.repo.BecomeVolunteer(userObjId)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}

	logger.Log("successfully became volunteer", userID)

	return nil
}

func (s *service) CreateOrganization(userId string, org Organization) error {
	logger := log.With(s.logger, "method", "CreateOrganization")

	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}
	org.UserID = userObjId

	err = s.repo.CreateOrganization(org)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}
	orgRole := RoleOrganization
	err = s.repo.UpdateUser(userObjId, UserUpdate{Role: &orgRole})
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}

	logger.Log("successfully created organization", org.ID.Hex())

	return nil

}

func (s *service) JoinOrganization(userID, orgID string) error {
	logger := log.With(s.logger, "method", "JoinOrganization")

	userObjId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}

	orgObjId, err := primitive.ObjectIDFromHex(orgID)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}

	err = s.repo.JoinOrganization(userObjId, orgObjId)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}

	logger.Log("successfully joined organization", orgID)

	return nil
}

func (s *service) LeaveOrganization(userID, orgID string) error {
	logger := log.With(s.logger, "method", "LeaveOrganization")

	userObjId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}

	orgObjId, err := primitive.ObjectIDFromHex(orgID)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}

	err = s.repo.LeaveOrganization(userObjId, orgObjId)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}

	logger.Log("successfully left organization", orgID)

	return nil

}

func (s *service) GetOrganizationById(id string) (Organization, error) {
	logger := log.With(s.logger, "method", "GetOrganizationById")

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return Organization{}, err
	}

	org, err := s.repo.GetOrganizationById(objId)
	if err != nil {
		level.Error(logger).Log("err", err)
		return Organization{}, err
	}

	logger.Log("successfully got organization by id", org)

	return org, nil

}

func (s *service) GetAllUsers() ([]User, error) {
	logger := log.With(s.logger, "method", "GetAllUsers")

	users, err := s.repo.GetAllUsers()
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	logger.Log("successfully got all users")

	return users, nil
}

func (s *service) GetAllOrganizations() ([]Organization, error) {
	logger := log.With(s.logger, "method", "GetAllOrganizations")

	orgs, err := s.repo.GetAllOrganizations()
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	logger.Log("successfully got all organizations")

	return orgs, nil
}

func (s *service) GenerateToken(username, pass string) (string, error) {
	logger := log.With(s.logger, "method", "GenerateToken")

	user, err := s.repo.UserByLogin(username, s.generatePassHash(pass))
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id.Hex(),
		user.Role,
	})

	logger.Log("successfully generated token")

	return token.SignedString([]byte(signingKey))
}

func (s *service) ParseToken(token string) (Identity, error) {
	logger := log.With(s.logger, "method", "ParseToken")

	var u Identity

	userToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid jwt signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		level.Error(logger).Log("err", err)
		return u, err
	}

	claims, ok := userToken.Claims.(*tokenClaims)
	if !ok {
		level.Error(logger).Log("err", errors.New("wrong claims type"))
		return u, errors.New("wrong claims type")
	}

	u.Id = claims.UserId
	u.Role = claims.Role

	logger.Log("successfully parsed token", u)

	return u, nil
}

func (s *service) generatePassHash(pass string) string {
	hash := sha256.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
