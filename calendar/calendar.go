package calendar

import (
	"fmt"
	"math"
)

type Month string

func (m Month) NextMonth() Month {
    switch m {
    case Hammer:
        return Alturiak
    case Alturiak:
        return Ches
    case Ches:
        return Tarsakh
    case Tarsakh:
        return Mirtul
    case Mirtul:
        return Kythorn
    case Kythorn:
        return Flamerule
    case Flamerule:
        return Eleasias
    case Eleasias:
        return Eleint
    case Eleint:
        return Marpenoth
    case Marpenoth:
        return Uktar
    case Uktar:
        return Nightal
    case Nightal:
        return Hammer
    }

    panic("unreachable")
}

func (m Month) IsHolidayMonth() bool {
    if m == Hammer || m == Tarsakh || m == Flamerule || m == Eleint || m == Uktar {
        return true
    }

    return false
}

func MonthAndDayFromRest(rest int, shieldmeet bool) (Month, Day) {
	  if rest >= 1 && rest <= 31 {
        return Hammer, Day(rest)
		} else if rest >= 32 && rest <= 61 {
        return Alturiak, Day(rest-31)
		} else if rest >= 62 && rest <= 91 {
        return Ches, Day(rest-61)
		} else if rest >= 92 && rest <= 121 {
        return Tarsakh, Day(rest-91)
		} else if rest >= 123 && rest <= 152 {
        return Mirtul, Day(rest-121)
		} else if rest >= 153 && rest <= 182 {
        return Kythorn, Day(rest-152)
		} else if shieldmeet {
			if rest >= 183 && rest <= 213 {
        return Flamerule, Day(rest-182)
			} else if rest >= 214 && rest <= 243 {
        return Eleasias, Day(rest-213)
			} else if rest >= 244 && rest <= 274 {
        return Eleint, Day(rest-243)
			} else if rest >= 275 && rest <= 304 {
        return Marpenoth, Day(rest-274)
			} else if rest >= 305 && rest <= 335 {
        return Uktar, Day(rest-304)
			} else if rest >= 336 && rest <= 366 {
        return Nightal, Day(rest-235)
			}
		} else {
			if rest >= 182 && rest <= 212 {
        return Flamerule, Day(rest-182)
			} else if rest >= 213 && rest <= 242 {
        return Eleasias, Day(rest-212)
			} else if rest >= 243 && rest <= 273 {
        return Eleint, Day(rest-242)
			} else if rest >= 274 && rest <= 303 {
        return Marpenoth, Day(rest-273)
			} else if rest >= 304 && rest <= 334 {
        return Uktar, Day(rest-303)
			} else if rest >= 335 && rest <= 365 {
        return Nightal, Day(rest-334)
			}
		}

    panic("unreachable")
}

func MonthFromIndex(index int) Month {
    switch index {
    case 1:
        return Hammer
    case 2:
        return Alturiak
    case 3:
        return Ches
    case 4:
        return Tarsakh
    case 5:
        return Mirtul
    case 6:
        return Kythorn
    case 7:
        return Flamerule
    case 8:
        return Eleasias
    case 9:
        return Eleint
    case 10:
        return Marpenoth
    case 11:
        return Uktar
    case 12:
        return Nightal
    }

    panic("unreachable")
}

func (m Month) Index() int {
    switch m {
    case Hammer:
        return 1
    case Alturiak:
        return 2
    case Ches:
        return 3
    case Tarsakh:
        return 4
    case Mirtul:
        return 5
    case Kythorn:
        return 6
    case Flamerule:
        return 7
    case Eleasias:
        return 8
    case Eleint:
        return 9
    case Marpenoth:
        return 10
    case Uktar:
        return 11
    case Nightal:
        return 12
    }

    panic("unreachable")
}

func (m Month) PastHolidays() int {
    if m.Index() < 2 {
        return 0
    }

    if m.Index() < 5 {
        return 1
    }

    if m.Index() < 8 {
        return 2
    }

    if m.Index() < 12 {
        return 3
    }

    return 4
}

func (m Month) daysSinceStartOfYear() int {
    return (m.Index() - 1)*daysPerMonth + m.PastHolidays()
}

type Holiday string
type Day uint
type Year uint

func (y Year) IsShieldmeetYear() bool {
    return y%4 == 0
}

func (y Year) ShieldmeetsCurrentYear() int {
    if int(y)%4 == 0 {
        return 1
    }

    return 0
}

func (y Year) ShieldmeetsSinceDR() int {
    return int(math.Floor(float64(y) / 4))
}

const (
    Hammer    Month = "Hammer"
    Alturiak  Month = "Alturiak"
    Ches      Month = "Ches"
    Tarsakh   Month = "Tarsakh"
    Mirtul    Month = "Mirtul"
    Kythorn   Month = "Kythorn"
    Flamerule Month = "Flamerule"
    Eleasias  Month = "Eleasias"
    Eleint    Month = "Eleint"
    Marpenoth Month = "Marpenoth"
    Uktar     Month = "Uktar"
    Nightal   Month = "Nightal"
)

const (
    Midwinter          Holiday = "Midwinter"
    Greengrass         Holiday = "Greengrass"
    Midsummer          Holiday = "Midsummer"
    Shieldmeet         Holiday = "Shieldmeet"
    Highharvestide     Holiday = "Highharvestide"
    TheFeastOfTheMooon Holiday = "The Feast of the Mooon"
    None               Holiday = "None"
)

type Date struct {
    Year    Year
    Month   Month
    Day     Day
    Holiday Holiday
}

func DaysBetween(from Date, to Date) int {
    fromDays := int(from.DaysSinceDR())
    toDays := int(to.DaysSinceDR())

    if fromDays > toDays {
        return fromDays - toDays
    }

    return toDays - fromDays
}

func (d Date) getHoliday() Holiday {
    if int(d.Day) <= 30 {
        return None
    }

    if d.Year.IsShieldmeetYear() && d.Holiday == Midsummer {
        return Shieldmeet
    }

    switch d.Month {
    case Nightal:
    case Hammer:
        return Midwinter
    case Alturiak:
    case Ches:
    case Tarsakh:
        return Greengrass
    case Mirtul:
    case Kythorn:
    case Flamerule:
        return Midwinter
    case Eleasias:
    case Eleint:
        return Highharvestide
    case Marpenoth:
    case Uktar:
        return TheFeastOfTheMooon
    }

    panic("unreachable")
}

const monthsPerYear = 12
const daysPerMonth = 30
const holidaysPerYear = 5
const daysPerYear = (monthsPerYear * daysPerMonth) + holidaysPerYear
const daysBetweenTwoShieldmeets = daysPerYear * 4

// DSD = Days since DR
type DSD int

func (d Date) DaysSinceDR() DSD {

    sumDaysPreviousYears := (int(d.Year)) * daysPerYear
    sumDaysCurrentYear := int(d.Day) + d.Month.daysSinceStartOfYear()
		fmt.Println(sumDaysCurrentYear)
    shieldmeetCountThisYear := 0

    if d.Year%4 == 0 {
        shieldmeetCountThisYear = 1
    }

    return DSD(sumDaysPreviousYears + d.Year.ShieldmeetsSinceDR() + sumDaysCurrentYear + shieldmeetCountThisYear)
}

func (d DSD) ToDate() Date {
    current := int(d)
    date := Date{}
    shieldmeets := current / daysBetweenTwoShieldmeets

    days := math.Floor(float64(current-shieldmeets) / daysPerYear)
    date.Year = Year(days)

    rest := current - (int(date.Year)*daysPerYear + date.Year.ShieldmeetsSinceDR())
		fmt.Println(current, date.Year, "*", daysPerYear, "+", date.Year.ShieldmeetsSinceDR(),  rest)
    //date.Month = MonthFromIndex(int(math.Ceil(float64(rest) / daysPerMonth)))
		date.Month, date.Day = MonthAndDayFromRest(rest, date.Year.IsShieldmeetYear())
    //holidaysCount := Month.PastHolidays(date.Month)
    //date.Day = Day(rest - (holidaysCount + date.Month.Index()*daysPerMonth))
    //date.Day = Day(rest - (holidaysCount + (date.Month.Index()-1)*daysPerMonth))

    date.Holiday = date.getHoliday()

    return date
}

type Calendar struct {
    current Date
}

func NewCalendar() *Calendar {
    return &Calendar{}
}

func (c *Calendar) Current() Date {
    return c.current
}

func (c *Calendar) SetCurrent(d Date) {
    c.current = d
}

func (c *Calendar) Year() Year {
    return c.current.Year
}

func (c *Calendar) Month() Month {
    return c.current.Month
}

func (c *Calendar) Day() Day {
    return c.current.Day
}

func (c *Calendar) Holiday() Holiday {
    return c.current.Holiday
}

func (c *Calendar) SetDate(year Year, month Month, day Day) {
    c.current = Date{Year: year, Month: month, Day: day}
    c.current.Holiday = c.current.getHoliday()
}

func (c *Calendar) Next() {
    date := &c.current

    if date.Day < 30 {
        date.Day += 1
        return
    }

    if date.Month.IsHolidayMonth() {
        if date.Day == 30 {
            date.Day = 31
        } else if date.Year.IsShieldmeetYear() && date.Holiday == Midsummer {
            date.Day = 32
        } else {
						date.Day = 1
						date.Month = date.Month.NextMonth()
				}
    } else {
        date.Day = 1
        date.Month = date.Month.NextMonth()
    }

    date.Holiday = date.getHoliday()
}
