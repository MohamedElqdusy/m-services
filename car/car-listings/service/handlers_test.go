package service

import (
	"encoding/json"
	"net/http"
	"reflect"
	"net/http/httptest"
	"car-listings/db"
	"car-listings/models"
	"testing"
	"github.com/julienschmidt/httprouter"
)

func TestAllCars(t *testing.T) {

	mockStore := db.InitMockStore()
	cars:= []models.Car{{"a","renault","megane",132,2015,"red",2584},{"a2","renault","megane",132,2015,"re",2584},{"1","mercedes","a180",123,2014,"black",15950},{"2","audi","a3",111,2016,"white",17210},{"3","vw","golf",86,2018,"green",14980},{"4","skoda","octavia",86,2018,"black",16990}}
	mockStore.On("FindAllCars").Return(cars, nil).Once()

	req, err := http.NewRequest("GET", "/cars/search/", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router := httprouter.New()
	router.Handle("GET", "/cars/search/", AllCars)
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := []models.Car{{"a","renault","megane",132,2015,"red",2584},{"a2","renault","megane",132,2015,"re",2584},{"1","mercedes","a180",123,2014,"black",15950},{"2","audi","a3",111,2016,"white",17210},{"3","vw","golf",86,2018,"green",14980},{"4","skoda","octavia",86,2018,"black",16990}}
	b := []models.Car{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b

	if !reflect.DeepEqual(actual,expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}

	// the expectations that we defined in the `On` method are asserted here
	mockStore.AssertExpectations(t)
}

func TestSearchCarsByColor(t *testing.T) {

	mockStore := db.InitMockStore()
	cars:= []models.Car{{"a","renault","megane",132,2015,"red",2584}}
	mockStore.On("FindCarsByColor").Return(cars, nil).Once()

	req, err := http.NewRequest("GET", "/cars/color/red/", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router := httprouter.New()
	router.Handle("GET", "/cars/color/:color/", SearchCarsByColor)
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := []models.Car{{"a","renault","megane",132,2015,"red",2584}}
	b := []models.Car{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b

	if !reflect.DeepEqual(actual,expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
	mockStore.AssertExpectations(t)
}

func TestSearchCarsByYear(t *testing.T) {

	mockStore := db.InitMockStore()
	cars:= []models.Car{{"3","vw","golf",86,2018,"green",14980},{"4","skoda","octavia",86,2018,"black",16990}}
	mockStore.On("FindCarsByYear").Return(cars, nil).Once()

	req, err := http.NewRequest("GET", "/cars/year/2018/", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router := httprouter.New()
	router.Handle("GET", "/cars/year/:year/", SearchCarsByYear)
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := []models.Car{{"3","vw","golf",86,2018,"green",14980},{"4","skoda","octavia",86,2018,"black",16990}}
	b := []models.Car{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b

	if !reflect.DeepEqual(actual,expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
	mockStore.AssertExpectations(t)
}

func TestSearchCarsByMake(t *testing.T) {

	mockStore := db.InitMockStore()
	cars:= []models.Car{{"a2","renault","megane",132,2015,"re",2584}}
	mockStore.On("FindCarsByMake").Return(cars, nil).Once()

	req, err := http.NewRequest("GET", "/cars/make/megan/", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router := httprouter.New()
	router.Handle("GET", "/cars/make/:make/", SearchCarsByMake)
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := []models.Car{{"a2","renault","megane",132,2015,"re",2584}}
	b := []models.Car{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b

	if !reflect.DeepEqual(actual,expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
	mockStore.AssertExpectations(t)
}

func TestSearchCarsByModel(t *testing.T) {

	mockStore := db.InitMockStore()
	cars:= []models.Car{{"1","mercedes","a180",123,2014,"black",15950}}
	mockStore.On("FindCarsByModel").Return(cars, nil).Once()

	req, err := http.NewRequest("GET", "/cars/model/mercedes/", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router := httprouter.New()
	router.Handle("GET", "/cars/model/:model/", SearchCarsByModel)
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := []models.Car{{"1","mercedes","a180",123,2014,"black",15950}}
	b := []models.Car{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b

	if !reflect.DeepEqual(actual,expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
	mockStore.AssertExpectations(t)
}