package time

import (
    "database/sql/driver"
    "fmt"
    "time"
)

// Time format json time field by myself
type Time struct {
    time.Time
}

// MarshalJSON on Time format Time field with %Y-%m-%d %H:%M:%S
func (t Time) MarshalJSON() ([]byte, error) {
    formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
    return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t Time) Value() (driver.Value, error) {
    var zeroTime time.Time
    if t.Time.UnixNano() == zeroTime.UnixNano() {
        return nil, nil
    }
    return t.Time, nil
}

// Scan valueof time.Time
func (t *Time) Scan(v interface{}) error {
    value, ok := v.(time.Time)
    if ok {
        *t = Time{Time: value}
        return nil
    }
    return fmt.Errorf("can not convert %v to timestamp", v)
}