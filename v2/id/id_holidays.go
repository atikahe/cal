// (c) Rick Arnold. Licensed under the BSD licence (see LICENSE).

// Package id provides holiday definitions for the Republic of Indonesia.
package id

import (
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/aa"
	"github.com/rickar/cal/v2/lunar"
)

var (
	// Indonesia does not enforce alt day for when a holiday overlapped with another.
	// Hence Alt Day is not needed.

	// NewYear represents Gregorian New Year's Day on 1 Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{Name: "Tahun Baru", Type: cal.ObservancePublic})

	// GoodFriday represents Good Friday - two days before Easter
	GoodFriday = aa.NewYear.Clone(&cal.Holiday{Name: "Wafat Yesus Kristus", Type: cal.ObservanceReligious})

	// WorkersDay represents International Worker's Day on the first day of May
	WorkersDay = aa.WorkersDay.Clone(&cal.Holiday{Name: "Hari Buruh Internasional", Type: cal.ObservancePublic})

	// AscensionDay represents Ascension Day on the 39th day after Easter
	AscensionDay = aa.AscensionDay.Clone(&cal.Holiday{Name: "Kenaikan Yesus Kristus", Type: cal.ObservanceReligious})

	// PancasilaDay represents Pancasila Day on 1-Jun
	PancasilaDay = &cal.Holiday{
		Name:  "Hari Lahir Pancasila",
		Type:  cal.ObservancePublic,
		Month: time.June,
		Day:   1,
		Func:  cal.CalcDayOfMonth,
	}

	// IndependenceDay represents Independence Day on 17-Aug
	IndependenceDay = &cal.Holiday{
		Name:  "Hari Kemerdekaan Republik Indonesia",
		Type:  cal.ObservancePublic,
		Month: time.August,
		Day:   17,
		Func:  cal.CalcDayOfMonth,
	}

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{Name: "Hari Raya Natal", Type: cal.ObservanceReligious})

	// EidAlFitr represents Eid Al-Fitr celebration on 1 Syawal
	EidAlFitr = &cal.Holiday{
		Name:  "Hari Raya Idulfitri",
		Type:  cal.ObservanceReligious,
		Month: lunar.EidAlFitrUAQ.Month(),
		Day:   lunar.EidAlFitrUAQ.Day(),
		Func:  cal.CalcDayOfMonth,
	}

	// EidAlAdha represents Eid Al-Adha celebration on 10 Dzulhijjah
	EidAlAdha = &cal.Holiday{
		Name:  "Hari Raya Iduladha",
		Type:  cal.ObservanceReligious,
		Month: lunar.EidAlAdhaUAQ.Month(),
		Day:   lunar.EidAlAdhaUAQ.Day(),
		Func:  cal.CalcDayOfMonth,
	}

	// EidAlAdha represents Arabic Lunar New Year every 1 Muharram
	IslamicNewYear = &cal.Holiday{
		Name:  "Tahun Baru Hijriah",
		Type:  cal.ObservanceReligious,
		Month: lunar.HijriNewYear.Month(),
		Day:   lunar.HijriNewYear.Day(),
		Func:  cal.CalcDayOfMonth,
	}

	// IsraMiraj represents Isra Miraj holiday on 27 Rajab
	IsraMiraj = &cal.Holiday{
		Name:  "Isra' Mi'raj",
		Type:  cal.ObservanceReligious,
		Month: lunar.IsraMiraj.Month(),
		Day:   lunar.IsraMiraj.Day(),
		Func:  cal.CalcDayOfMonth,
	}

	// Mawlid represents Mawlid Nabi on 12 Rabi'ul Awal
	Mawlid = &cal.Holiday{
		Name:  "Maulid Nabi",
		Type:  cal.ObservanceReligious,
		Month: lunar.Mawlid.Month(),
		Day:   lunar.Mawlid.Day(),
		Func:  cal.CalcDayOfMonth,
	}

	// TODO: Add Lunar Chinese Holiday
	// TODO: Add Lunar Saka Holiday
	// TODO: Add Waisak Holiday
)
