package parsing
import(
	"testing"
	"car-listings/models"
)

func TestParseCarFromCSV(t *testing.T) {
	tt := []struct {
		Name	string
		Line	string
		Expected models.Car
	}{
		{Name: "valid line", Line: "1,mercedes,a180,123,2014,black,15950", Expected: models.Car{"1", "mercedes", "a180", 123, 2014, "black", 15950},},
		{Name: "more coulmns", Line: "2,bmw,a180,123,2014,black,555555,3333,asdf", Expected: models.Car{"2", "bmw", "a180", 123, 2014, "black", 555555},},
	}

	for _, tc := range tt {
		car, err := parseCarFromCSV(tc.Line)
		if err != nil {
			t.Fatalf("Test '%s' failed with %v", tc.Name, err)
		}
		if car != tc.Expected {
			t.Fatalf("Test '%s' failed, Expected %v Found %v", tc.Name, tc.Expected,car)
		}
	}
}
// Note! every panic test should be tested seperately
func TestParseCarFromCSVPanicEmptyLine(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code didn't panic")
        }
	}()
	car, _ := parseCarFromCSV("")
	t.Fatalf("Car %v shoulden't be reached here, this must panic",car)
}


func TestParseCarFromCSVPanicMissingColumns(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code didn't panic")
        }
	}()
	car, _ := parseCarFromCSV("a180,123,2014,black")
	t.Fatalf("Car %v shoulden't be reached here, this must panic",car)
}
