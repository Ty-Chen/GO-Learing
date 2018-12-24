package main

func max(a int, b int) int {
	if a > b {	
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func sign(a int) int {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	} else {
		return 0
	}
}

func prize1(score int) string {
	switch score / 10 {
	case 0, 1, 2, 3, 4, 5:
		return "不及格"
	case 6, 7:
		return "及格"
	case 8:
		return "良"
	case 9, 10:
		return "优秀"
	}
	return ""
}

// 表达式匹配
func prize2(score int) string {
    // 注意 switch 后面什么也没有
    switch {
        case score < 60:
            return "差"
        case score < 80:
            return "及格"
        case score < 90:
            return "良"
        default:
            return "优"
    }
}