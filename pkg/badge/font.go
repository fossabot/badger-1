package badge

const chartWidthTableLength = 127
const defaultCharWidth = 120
const scale = 10

// fontFamily=Verdana, fontSize=11, scale=10
var charWidthTable = [chartWidthTableLength]int{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 39,
	43, 50, 90, 70, 120, 80, 30, 50, 50, 70, 90,
	40, 50, 40, 50, 70, 70, 70, 70, 70, 70, 70,
	70, 70, 70, 50, 50, 90, 90, 90, 60, 110, 75,
	75, 77, 85, 70, 63, 85, 83, 46, 50, 76, 61,
	93, 82, 87, 66, 87, 76, 75, 68, 81, 75, 110,
	75, 68, 75, 50, 50, 50, 90, 70, 70, 66, 69,
	57, 69, 66, 39, 69, 70, 30, 38, 65, 30, 110,
	70, 67, 69, 69, 47, 57, 43, 70, 65, 90, 65,
	65, 58, 70, 50, 70, 90,
}

func computeTextWidth(text string) int {
	totalWidth := 0
	for _, c := range text {
		if charCode := int(c); charCode < chartWidthTableLength {
			totalWidth += charWidthTable[charCode]
		} else {
			totalWidth += defaultCharWidth
		}
	}

	return totalWidth / scale
}
