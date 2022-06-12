package api

import (
	"Redrock/models"
	"errors"
	"math"
)

//判断能不能走到

func (r *Room) Bing(s [4]int, che int) error {
	// 世俗总要英雄无畏
	if (s[3]-s[1])*(s[3]-s[1])+(s[2]-s[0])*(s[2]-s[0]) != 1 {
		return errors.New("瞎几把走")
	}
	if che == 0 {
		if s[2] < s[0] {
			return errors.New("你怎么能回头呢")
		}
		if s[0] <= 4 && s[1] != s[3] {
			return errors.New("没有过河 不能走")
		}
	} else {
		if s[2] > s[0] {
			return errors.New("你怎么能回头呢")
		}
		if s[0] >= 5 && s[1] != s[3] {
			return errors.New("没有过河 不能走")
		}
	}

	return nil
}
func (r *Room) Pao(s [4]int) (int, error) {
	//判断直线
	//判断 中间是否只有一个
	var target = 0
	if s[0] == s[2] && s[1] != s[3] {
		m := s[1]
		l := s[3]
		if s[1] > s[3] {
			l = s[1]
			m = s[3]
		}
		for i := m + 1; i <= l; i++ {
			//判断有咩有阻挡
			if r.Info[s[0]][i] != models.Kong {
				target++
			}
		}
	}
	if s[0] != s[2] && s[1] == s[3] {
		m := s[1]
		l := s[3]
		if s[1] > s[3] {
			l = s[1]
			m = s[3]
		}
		for i := m + 1; i < l; i++ {
			//判断有咩有阻挡
			if r.Info[i][s[0]] != models.Kong {
				target++
			}
		}
	}
	if target == 0 || target == 2 {
		//inform[s[2]][s[3]] = models.Pao
		//inform[s[0]][s[1]] = models.Kong
		return target, nil
	} else {
		return -1, errors.New("你这炮不是意大利炮 不行")
	}
}
func (r *Room) Jiang(s [4]int) error {
	// 距离问题
	if (s[3]-s[1])*(s[3]-s[1])+(s[2]-s[0])*(s[2]-s[0]) != 1 {
		return errors.New("瞎几把走")
	}
	if s[3] > 5 || s[3] < 3 {
		return errors.New("出格了")
	}
	// 越界问题
	if math.Abs(4.5-float64(s[2])) < 2 {
		return errors.New("错误")
	}
	if s[2] > 6 {
		for i := s[3]; i <= 9; i++ {
			//将军喝酒
			if r.Info[s[2]][i] == models.Kong {
				continue
			} else if r.Info[s[2]][i] != models.Jiang {
				break
			}
			return errors.New("不行哦")

		}
	} else {
		for i := 9; i <= s[3]; i-- {
			//将军喝酒
			if r.Info[s[2]][i] == models.Kong {
				continue
			} else if r.Info[s[2]][i] != models.Jiang {
				break
			}
			return errors.New("不行哦")

		}
	}
	if s[0] > 5 {
		r.Jia[0][0] = s[2]
		r.Jia[0][1] = s[3]

	} else {
		r.Jia[1][0] = s[2]
		r.Jia[1][1] = s[3]
	}

	//r.Info[s[2]][s[3]] = models.Jiang
	//r.Info[s[0]][s[1]] = models.Kong
	return nil
}
func (r *Room) Shi(s [4]int) error {
	//是否符合规格
	if (s[3]-s[1])*(s[3]-s[1])+(s[2]-s[0])*(s[2]-s[0]) != 2 {
		return errors.New("瞎几把走")
	}
	//越界问题
	if s[3] > 5 || s[3] < 3 {
		return errors.New("出格了")
	}
	// 越界问题
	if math.Abs(4.5-float64(s[2])) < 2 {
		return errors.New("错误")
	}
	//R.Info[s[2]][s[3]] = models.Shi
	//R.Info[s[0]][s[1]] = models.Kong
	return nil
}
func (r *Room) Xian(s [4]int) error {
	// 敌方象 与我方象
	if r.Info[(s[0]+s[2])/2][(s[1]+s[3])/2] != models.Kong {
		return errors.New("蹩脚")
	}
	if s[0] <= 4 {
		if s[2] > 4 {
			return errors.New("不能过去")
		}
		if math.Abs(float64(s[0]-s[2])) != 2 && math.Abs(float64(s[1]-s[3])) != 2 {
			return errors.New("瞎几把下")
		}
	} else {
		if s[2] < 5 {
			return errors.New("过不去")
		}
		if math.Abs(float64(s[0]-s[2])) != 2 && math.Abs(float64(s[1]-s[3])) != 2 {
			return errors.New("瞎几把下")
		}

	}
	//r.Info[s[2]][s[3]] = models.Xian
	//r.Info[s[0]][s[1]] = models.Kong
	return nil
}
func (r *Room) Ma(s [4]int) error {
	// 判断方向 满足走的方向
	if (s[3]-s[1])*(s[3]-s[1])+(s[2]-s[0])*(s[2]-s[0]) == 5 {
		x := math.Abs(float64(s[2] - s[0]))
		y := math.Abs(float64(s[3] - s[1]))
		if x > y {
			if s[2]-s[0] > 0 {
				if r.Info[s[0]+1][s[1]] != models.Kong {
					return errors.New("bie")
				}
			} else {
				if r.Info[s[0]-1][s[1]] != models.Kong {
					return errors.New("bie")
				}
			}
		} else {
			if s[3]-s[1] > 0 {
				if r.Info[s[0]][s[1]+1] != models.Kong {
					return errors.New("bie")
				}
			} else {
				if r.Info[s[0]][s[1]-1] != models.Kong {
					return errors.New("bie")
				}
			}
		}
		//r.Info[s[2]][s[3]] = models.Ma
		//r.Info[s[0]][s[1]] = models.Kong
		return nil
	} else {
		return errors.New("走位错误")
	}

}
func (r *Room) Che(s [4]int) error {
	// 判断是不是直线

	if s[0] == s[2] && s[1] != s[3] {
		m := s[1]
		l := s[3]
		if s[1] > s[3] {
			l = s[1]
			m = s[3]
		}
		for i := m + 1; i < l; i++ {
			//判断有咩有阻挡
			if r.Info[s[0]][i] != models.Kong {
				err := errors.New("中间过不去")
				return err
			}
		}
	}
	if s[0] != s[2] && s[1] == s[3] {
		m := s[1]
		l := s[3]
		if s[1] > s[3] {
			l = s[1]
			m = s[3]
		}
		for i := m + 1; i < l; i++ {
			//判断有咩有阻挡
			if r.Info[i][s[0]] != models.Kong {
				err := errors.New("中间过不去")
				return err
			}
		}
	}

	return nil
}
