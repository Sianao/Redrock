package api

import "testing"

func TestRoom_Bing(t *testing.T) {
	m := ChessTable()
	table := NewLogic(m)
	var r Room
	r.Info = m
	r.Table = table
	if err := r.Bing([4]int{0, 3, 0, 4}, 1); err != nil {
		t.Errorf("err %#v", err.Error())
	}
}
func TestRoom_Che(t *testing.T) {
	m := ChessTable()
	table := NewLogic(m)
	var r Room
	r.Info = m
	r.Table = table
	if err := r.Che([4]int{0, 0, 1, 0}); err != nil {
		t.Errorf("err %#v", err.Error())

	}
}
func TestRoom_Pao(t *testing.T) {
	test := []struct {
		step   [4]int
		expect string
	}{{[4]int{9, 2, 8, 2}, "你这炮不是意大利炮 不行"},
		{[4]int{9, 2, 8, 0}, "你这炮不是意大利炮 不行"},
		{[4]int{9, 2, 9, 0}, ""},
		{[4]int{5, 2, 5, 1}, ""}}
	m := ChessTable()
	table := NewLogic(m)
	var r Room
	r.Info = m
	r.Table = table
	for _, v := range test {
		if _, err := r.Pao(v.step); err != nil && err.Error() != v.expect {
			t.Errorf("err %#v", err.Error())
		}
	}
}
func TestRoom_Ma(t *testing.T) {

	m := ChessTable()
	table := NewLogic(m)
	var r Room
	r.Info = m
	r.Table = table
	if err := r.Ma([4]int{0, 1, 3, 2}); err == nil {
		t.Errorf("err %#v", err.Error())
	}
	if err := r.Ma([4]int{0, 1, 2, 2}); err != nil {
		t.Errorf("err %#v", err.Error())
	}
}
func TestRoom_Xian(t *testing.T) {
	test := []struct {
		step   [4]int
		param  string
		expect string
	}{{[4]int{9, 2, 7, 0}, "", ""},
		{[4]int{9, 2, 8, 0}, "", "瞎几把下"},
		{[4]int{9, 2, 8, 0}, "", "瞎几把下"},
		{[4]int{5, 2, 3, 1}, "", "不能过去"}}
	m := ChessTable()
	table := NewLogic(m)
	var r Room
	r.Info = m
	r.Table = table
	for _, v := range test {
		if err := r.Xian(v.step); err != nil {

		}
	}
	if err := r.Xian([4]int{9, 2, 7, 0}); err != nil {
		t.Errorf("err %#v", err.Error())
	}
}
func TestRoom_Jiang(t *testing.T) {
	test := []struct {
		step   [4]int
		param  string
		expect string
	}{{[4]int{0, 4, 1, 4}, "", ""},
		{[4]int{0, 4, 1, 5}, "", "瞎几把走"},
		{[4]int{9, 4, 8, 4}, "", "瞎几把走"}}
	m := ChessTable()
	table := NewLogic(m)
	var r Room
	r.Info = m
	r.Table = table
	for _, v := range test {
		if err := r.Jiang(v.step); err != nil {
			if err.Error() != v.expect {
				t.Errorf("err%#v", err.Error())
			}
		}
	}

}
func TestRoom_Shi(t *testing.T) {
	test := []struct {
		step   [4]int
		param  string
		expect string
	}{{[4]int{0, 3, 1, 4}, "", ""},
		{[4]int{2, 4, 2, 5}, "", "瞎几把走"},
		{[4]int{9, 3, 8, 4}, "", ""}}
	m := ChessTable()
	table := NewLogic(m)
	var r Room
	r.Info = m
	r.Table = table
	for _, v := range test {
		if err := r.Shi(v.step); err != nil {
			if err.Error() != v.expect {
				t.Errorf("err%#v", err.Error())
			}
		}
	}

}
