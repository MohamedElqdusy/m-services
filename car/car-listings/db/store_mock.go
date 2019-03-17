package db

import (
	"context"
	"car-listings/models"

	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) StoreCar(ctx context.Context, car models.Car, dealer_id uint64) error {
	return repositoryImpl.StoreCar(ctx, car, dealer_id)
}

func (m *MockStore) FindCarsByMake(ctx context.Context, make string) ([]models.Car, error) {
	rets := m.Called()
	return rets.Get(0).([]models.Car), rets.Error(1)
}

func (m *MockStore) FindCarsByModel(ctx context.Context, model string)([]models.Car, error){
	rets := m.Called()
	return rets.Get(0).([]models.Car), rets.Error(1)
}

func (m *MockStore) FindCarsByYear(ctx context.Context, year uint64) ([]models.Car, error){
	rets := m.Called()
	return rets.Get(0).([]models.Car), rets.Error(1)
}

func (m *MockStore) FindCarsByColor(ctx context.Context, color string) ([]models.Car, error){
	rets := m.Called()
	return rets.Get(0).([]models.Car), rets.Error(1)
}

func (m *MockStore) FindAllCars(ctx context.Context) ([]models.Car, error) {
	rets := m.Called()
	return rets.Get(0).([]models.Car), rets.Error(1)
}

func (m *MockStore) Close() {
}

func InitMockStore() *MockStore {
	store := new(MockStore)
	SetRepository(store)
	return store
}
