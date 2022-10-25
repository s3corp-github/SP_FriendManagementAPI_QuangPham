package relation

import (
	"context"
	models "github.com/quangpham789/golang-assessment/api/internal/repository/orm/models"
	"github.com/stretchr/testify/mock"
)

type MockRelationRepo struct {
	mock.Mock
}

func (m MockRelationRepo) CreateRelation(ctx context.Context, input models.Relation) (bool, error) {
	args := m.Called(ctx, input)
	return true, args.Error(1)
}

func (m MockRelationRepo) GetRelationByIdsAndType(ctx context.Context, requesterId int, addresseeId int, relationType int) (models.Relation, error) {
	args := m.Called(ctx, requesterId, addresseeId, relationType)
	return args.Get(0).(models.Relation), args.Error(1)
}

func (m MockRelationRepo) GetRelationFriendIDsOfUser(ctx context.Context, requesterId int) ([]int, error) {
	//args := m.Called(ctx, requesterId)
	//return args.Get(0).(models.UserSlice.), args.Error(1)
	panic("To implement")
}
