package lunar

import (
	"time"

	"github.com/hablullah/go-hijri"
)

// Arabic lunar calendar, Hijri, has two calculation mechanisms: arithmetic rules and observation of astronomical rules (Umm Al Qura).
// Countries with significant time difference with the Arab gulf usually follow the consensus as follows:
//     1. Minor holidays are decided based on arithmetic calculation.
//     2. Major holidays that inquire specific obligations are decided based on Umm Al Qura method.
// To author's knowledge, this consensus is common among Malaysia, Indonesia, Singapore, and Brunei.

// Below listed the most common holidays that are celebrated.
// Each var has the Gregorian date of Arabic lunar holidays.

var (
	EidAlFitr    time.Time = hijriToGregorian(10, 1)
	EidAlAdha    time.Time = hijriToGregorian(12, 10)
	HijriNewYear time.Time = hijriToGregorian(1, 1)
	IsraMiraj    time.Time = hijriToGregorian(7, 27)
	Mawlid       time.Time = hijriToGregorian(3, 12)
)

var (
	EidAlFitrUAQ time.Time = ummAlQuraToGregorian(10, 1)
	EidAlAdhaUAQ time.Time = ummAlQuraToGregorian(12, 10)
)

// TODO: Manual calculation
func ummAlQuraToGregorian(month int64, day int64) time.Time {
	uaq, _ := hijri.CreateUmmAlQuraDate(time.Now())
	d := hijri.UmmAlQuraDate{Year: uaq.Year, Month: month, Day: day}
	return d.ToGregorian()
}

// TODO: Manual calculation
func hijriToGregorian(month int64, day int64) time.Time {
	hijria, _ := hijri.CreateHijriDate(time.Now(), hijri.Default)
	d := hijri.UmmAlQuraDate{Year: hijria.Year, Month: month, Day: day}
	return d.ToGregorian()
}
