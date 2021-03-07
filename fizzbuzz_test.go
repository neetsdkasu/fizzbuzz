package fizzbuzz

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testdata := []struct {
		Input  uint8
		Output string
	}{
		{0, "FizzBuzz"},
		{1, "1"},
		{10, "Buzz"},
		{100, "Buzz"},
		{101, "101"},
		{102, "Fizz"},
		{103, "103"},
		{104, "104"},
		{105, "FizzBuzz"},
		{106, "106"},
		{107, "107"},
		{108, "Fizz"},
		{109, "109"},
		{11, "11"},
		{110, "Buzz"},
		{111, "Fizz"},
		{112, "112"},
		{113, "113"},
		{114, "Fizz"},
		{115, "Buzz"},
		{116, "116"},
		{117, "Fizz"},
		{118, "118"},
		{119, "119"},
		{12, "Fizz"},
		{120, "FizzBuzz"},
		{121, "121"},
		{122, "122"},
		{123, "Fizz"},
		{124, "124"},
		{125, "Buzz"},
		{126, "Fizz"},
		{127, "127"},
		{128, "128"},
		{129, "Fizz"},
		{13, "13"},
		{130, "Buzz"},
		{131, "131"},
		{132, "Fizz"},
		{133, "133"},
		{134, "134"},
		{135, "FizzBuzz"},
		{136, "136"},
		{137, "137"},
		{138, "Fizz"},
		{139, "139"},
		{14, "14"},
		{140, "Buzz"},
		{141, "Fizz"},
		{142, "142"},
		{143, "143"},
		{144, "Fizz"},
		{145, "Buzz"},
		{146, "146"},
		{147, "Fizz"},
		{148, "148"},
		{149, "149"},
		{15, "FizzBuzz"},
		{150, "FizzBuzz"},
		{151, "151"},
		{152, "152"},
		{153, "Fizz"},
		{154, "154"},
		{155, "Buzz"},
		{156, "Fizz"},
		{157, "157"},
		{158, "158"},
		{159, "Fizz"},
		{16, "16"},
		{160, "Buzz"},
		{161, "161"},
		{162, "Fizz"},
		{163, "163"},
		{164, "164"},
		{165, "FizzBuzz"},
		{166, "166"},
		{167, "167"},
		{168, "Fizz"},
		{169, "169"},
		{17, "17"},
		{170, "Buzz"},
		{171, "Fizz"},
		{172, "172"},
		{173, "173"},
		{174, "Fizz"},
		{175, "Buzz"},
		{176, "176"},
		{177, "Fizz"},
		{178, "178"},
		{179, "179"},
		{18, "Fizz"},
		{180, "FizzBuzz"},
		{181, "181"},
		{182, "182"},
		{183, "Fizz"},
		{184, "184"},
		{185, "Buzz"},
		{186, "Fizz"},
		{187, "187"},
		{188, "188"},
		{189, "Fizz"},
		{19, "19"},
		{190, "Buzz"},
		{191, "191"},
		{192, "Fizz"},
		{193, "193"},
		{194, "194"},
		{195, "FizzBuzz"},
		{196, "196"},
		{197, "197"},
		{198, "Fizz"},
		{199, "199"},
		{2, "2"},
		{20, "Buzz"},
		{200, "Buzz"},
		{201, "Fizz"},
		{202, "202"},
		{203, "203"},
		{204, "Fizz"},
		{205, "Buzz"},
		{206, "206"},
		{207, "Fizz"},
		{208, "208"},
		{209, "209"},
		{21, "Fizz"},
		{210, "FizzBuzz"},
		{211, "211"},
		{212, "212"},
		{213, "Fizz"},
		{214, "214"},
		{215, "Buzz"},
		{216, "Fizz"},
		{217, "217"},
		{218, "218"},
		{219, "Fizz"},
		{22, "22"},
		{220, "Buzz"},
		{221, "221"},
		{222, "Fizz"},
		{223, "223"},
		{224, "224"},
		{225, "FizzBuzz"},
		{226, "226"},
		{227, "227"},
		{228, "Fizz"},
		{229, "229"},
		{23, "23"},
		{230, "Buzz"},
		{231, "Fizz"},
		{232, "232"},
		{233, "233"},
		{234, "Fizz"},
		{235, "Buzz"},
		{236, "236"},
		{237, "Fizz"},
		{238, "238"},
		{239, "239"},
		{24, "Fizz"},
		{240, "FizzBuzz"},
		{241, "241"},
		{242, "242"},
		{243, "Fizz"},
		{244, "244"},
		{245, "Buzz"},
		{246, "Fizz"},
		{247, "247"},
		{248, "248"},
		{249, "Fizz"},
		{25, "Buzz"},
		{250, "Buzz"},
		{251, "251"},
		{252, "Fizz"},
		{253, "253"},
		{254, "254"},
		{255, "FizzBuzz"},
		{26, "26"},
		{27, "Fizz"},
		{28, "28"},
		{29, "29"},
		{3, "Fizz"},
		{30, "FizzBuzz"},
		{31, "31"},
		{32, "32"},
		{33, "Fizz"},
		{34, "34"},
		{35, "Buzz"},
		{36, "Fizz"},
		{37, "37"},
		{38, "38"},
		{39, "Fizz"},
		{4, "4"},
		{40, "Buzz"},
		{41, "41"},
		{42, "Fizz"},
		{43, "43"},
		{44, "44"},
		{45, "FizzBuzz"},
		{46, "46"},
		{47, "47"},
		{48, "Fizz"},
		{49, "49"},
		{5, "Buzz"},
		{50, "Buzz"},
		{51, "Fizz"},
		{52, "52"},
		{53, "53"},
		{54, "Fizz"},
		{55, "Buzz"},
		{56, "56"},
		{57, "Fizz"},
		{58, "58"},
		{59, "59"},
		{6, "Fizz"},
		{60, "FizzBuzz"},
		{61, "61"},
		{62, "62"},
		{63, "Fizz"},
		{64, "64"},
		{65, "Buzz"},
		{66, "Fizz"},
		{67, "67"},
		{68, "68"},
		{69, "Fizz"},
		{7, "7"},
		{70, "Buzz"},
		{71, "71"},
		{72, "Fizz"},
		{73, "73"},
		{74, "74"},
		{75, "FizzBuzz"},
		{76, "76"},
		{77, "77"},
		{78, "Fizz"},
		{79, "79"},
		{8, "8"},
		{80, "Buzz"},
		{81, "Fizz"},
		{82, "82"},
		{83, "83"},
		{84, "Fizz"},
		{85, "Buzz"},
		{86, "86"},
		{87, "Fizz"},
		{88, "88"},
		{89, "89"},
		{9, "Fizz"},
		{90, "FizzBuzz"},
		{91, "91"},
		{92, "92"},
		{93, "Fizz"},
		{94, "94"},
		{95, "Buzz"},
		{96, "Fizz"},
		{97, "97"},
		{98, "98"},
		{99, "Fizz"},
	}

	for _, d := range testdata {
		res := FizzBuzz(d.Input)
		if res != d.Output {
			t.Fatalf(`unmatch %d "%s"!="%s"`, d.Input, d.Output, res)
		}
	}
}

func ExampleFizzBuzz() {
	var i uint8
	for i = 10; i <= 17; i++ {
		fmt.Println(FizzBuzz(i))
	}
	// Output:
	// Buzz
	// 11
	// Fizz
	// 13
	// 14
	// FizzBuzz
	// 16
	// 17
}