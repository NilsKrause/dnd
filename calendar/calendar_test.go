package calendar

import (
	"fmt"
	"testing"
)

func TestNext(t *testing.T) {

	calendar := NewCalendar();
	calendar.SetDate(1493, Hammer, 30)
	calendar.Next()
	fmt.Println(calendar.Current(), " ", calendar.Current().DaysSinceDR())
	calendar.Next()
	fmt.Println(calendar.Current(), " ", calendar.Current().DaysSinceDR())
	calendar.Next()
	fmt.Println(calendar.Current(), " ", calendar.Current().DaysSinceDR())

	dr := calendar.Current().DaysSinceDR()
	nCalendar := NewCalendar()
	nCalendar.SetCurrent(dr.ToDate())
	fmt.Println(nCalendar.Current(), " ", nCalendar.Current().DaysSinceDR())
}

func TestNextMonth(t *testing.T) {
	month := Hammer
	month = month.NextMonth()

	if month != Alturiak {
		t.Errorf("got %q, want Alturiak", month)
	}

	month = month.NextMonth()
	if month != Ches {
		t.Errorf("got %q, want Ches", month)
	}

	month = month.NextMonth()
	if month != Tarsakh {
		t.Errorf("got %q, want Tarsakh", month)
	}

	month = month.NextMonth()
	if month != Mirtul {
		t.Errorf("got %q, want Mirtul", month)
	}

	month = month.NextMonth()
	if month != Kythorn {
		t.Errorf("got %q, want Kythorn", month)
	}

	month = month.NextMonth()
	if month != Flamerule {
		t.Errorf("got %q, want Flamerule", month)
	}

	month = month.NextMonth()
	if month != Eleasias {
		t.Errorf("got %q, want Eleasias", month)
	}

	month = month.NextMonth()
	if month != Eleint {
		t.Errorf("got %q, want Eleint", month)
	}

	month = month.NextMonth()
	if month != Marpenoth {
		t.Errorf("got %q, want Marpenoth", month)
	}

	month = month.NextMonth()
	if month != Uktar {
		t.Errorf("got %q, want Uktar", month)
	}

	month = month.NextMonth()
	if month != Nightal {
		t.Errorf("got %q, want Nightal", month)
	}

	month = month.NextMonth()
	if month != Hammer {
		t.Errorf("got %q, want Hammer", month)
	}
}

func TestIsHolidayMonth(t *testing.T) {
	month := Hammer
	for {
		month = month.NextMonth()

		if month.IsHolidayMonth() {
			if month != Hammer && month != Tarsakh && month != Flamerule && month != Eleint && month != Uktar {
				t.Errorf("got %q as a holiday month, while it ISN'T", month)
			}
		} else {
			if month != Alturiak && month != Ches && month != Mirtul && month != Kythorn && month != Eleasias && month != Marpenoth && month != Nightal {
				t.Errorf("got %q as NOT a holiday month, while it is", month)
			}
		}

		if month == Nightal {
			break
		}
	}
}
