package sqlstring

import (
	"reflect"
	"testing"
	"time"
)

func TestNULLEscape(t *testing.T) {
	result := Escape(nil)
	if result != "NULL" {
		t.Fatalf("escape error")
	}
}

func TestEmptyStringEscape(t *testing.T) {
	result := Escape("")
	t.Logf("result :%s", result)
	if result != "''" {
		t.Fatalf("escape empty string error")
	}
}
func TestBoolEscape(t *testing.T) {

	result := Escape(true)
	if result != "true" {
		t.Fatalf("escape error")
	}

	result = Escape(false)
	if result != "false" {
		t.Fatalf("escape error")
	}
}

func TestTimeToString(t *testing.T) {
	bt, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-01-01 15:00:09", time.Local)

	result := Escape(bt)
	t.Logf("result time %s", result)
	if result != "'2021-01-01 15:00:09'" {
		t.Fatalf("escape time error")
	}

	result = Escape(&bt)
	t.Logf("result time2 %s", result)
	if result != "'2021-01-01 15:00:09'" {
		t.Fatalf("escape time error")
	}
}

func TestArrayToString(t *testing.T) {

	var a = []int{1, 2, 3, 4}
	val := reflect.ValueOf(a)
	result := arrayToString(val, time.Local)

	if result != "1,2,3,4" {
		t.Fatalf("escape slice error")

	}

	b := [3]string{"1", "2", "3"}
	val = reflect.ValueOf(b)
	result = arrayToString(val, time.Local)

	if result != "'1','2','3'" {
		t.Fatalf("escape arr error")

	}

}

func TestStringEscape(t *testing.T) {
	s := "hello world"
	result := Escape(s)
	if result != "'hello world'" {
		t.Fatalf("escape string error")

	}

	s = "hello ' world"
	result = Escape(s)
	if result != "'hello \\' world'" {
		t.Fatalf("escape string error")

	}
}

func TestBytesEscape(t *testing.T) {
	s := []byte{0, 1, 254, 255}
	result := Escape(s)
	t.Logf("TestBytesEscape result: %s", result)
	if result != "X'0001feff'" {
		t.Fatalf("escape bytes error")

	}

}

func TestIntEscape(t *testing.T) {
	var i int = 10
	result := Escape(i)
	t.Logf("TestBytesEscape result: %s", result)
	if result != "10" {
		t.Fatalf("escape int error")

	}

	var i2 int8 = 7
	result = Escape(i2)
	t.Logf("TestBytesEscape result: %s", result)
	if result != "7" {
		t.Fatalf("escape int8 error")

	}

	var i3 int16 = 12
	result = Escape(i3)
	t.Logf("TestBytesEscape result: %s", result)
	if result != "12" {
		t.Fatalf("escape int16 error")

	}

	var i4 int32 = 13
	result = Escape(i4)
	t.Logf("TestBytesEscape result: %s", result)
	if result != "13" {
		t.Fatalf("escape int32 error")

	}

	var i5 int64 = 14
	result = Escape(i5)
	t.Logf("TestBytesEscape result: %s", result)
	if result != "14" {
		t.Fatalf("escape int32 error")

	}

}

func TestUIntEscape(t *testing.T) {
	var i uint = 10
	result := Escape(i)
	t.Logf("TestBytesEscape result: %s", result)
	if result != "10" {
		t.Fatalf("escape int error")

	}

	var i2 uint8 = 7
	result = Escape(i2)
	t.Logf("TestBytesEscape result: %s", result)
	if result != "7" {
		t.Fatalf("escape int8 error")

	}

	var i3 uint16 = 12
	result = Escape(i3)
	t.Logf("TestBytesEscape result: %s", result)
	if result != "12" {
		t.Fatalf("escape int16 error")

	}

	var i4 uint32 = 13
	result = Escape(i4)
	t.Logf("TestBytesEscape result: %s", result)
	if result != "13" {
		t.Fatalf("escape int32 error")

	}

	var i5 uint64 = 14
	result = Escape(i5)
	t.Logf("TestBytesEscape result: %s", result)
	if result != "14" {
		t.Fatalf("escape int32 error")

	}

}

func TestOtherEscape(t *testing.T) {
	x := map[string]string{
		"name": "asd'fsadf",
		"key":  "test",
	}
	result := Escape(x)
	t.Logf("escape reuslt %s", result)

	if result != "'{\"key\":\"test\",\"name\":\"asd\\'fsadf\"}'" {
		t.Fatalf("escape map error")

	}

}

func TestFormatSql(t *testing.T) {

	sql := Format("select * from users where name=? and age=? limit ?,?", "t'est", 10, 10, 10)
	t.Logf("sql %s", sql)

	if sql != "select * from users where name='t\\'est' and age=10 limit 10,10" {
		t.Fatalf("escape format error")
	}

	sql = Format("? and ?", "a", "b")
	t.Logf("sql %s", sql)

	if sql != "'a' and 'b'" {
		t.Fatalf("escape format str error")
	}

	sql = Format("in (?)", []int{1, 2, 3})
	t.Logf("sql %s", sql)

	if sql != "in (1,2,3)" {
		t.Fatalf("escape format arr error")

	}

	sql = Format("in (?)", []interface{}{1, 2, 3})
	t.Logf("sql %s", sql)

	if sql != "in (1,2,3)" {
		t.Fatalf("escape format arr error")

	}

	sql = Format("in (?)", []string{"1", "2", "3"})
	t.Logf("sql %s", sql)

	if sql != "in ('1','2','3')" {
		t.Fatalf("escape format arr2 error")

	}

	sql = Format("in (?)", []interface{}{"1", "2", "3"})
	t.Logf("sql %s", sql)

	if sql != "in ('1','2','3')" {
		t.Fatalf("escape format arr2 error")

	}

	sql = Format("in (?)", []interface{}{1, 2, "3"})
	t.Logf("sql %s", sql)

	if sql != "in (1,2,'3')" {
		t.Fatalf("escape format arr error")

	}

	bt, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-01-01 15:00:09", time.Local)

	sql = Format("a=?", bt)
	t.Logf("sql %s", sql)

	if sql != "a='2021-01-01 15:00:09'" {
		t.Fatalf("escape format time error")

	}
}
