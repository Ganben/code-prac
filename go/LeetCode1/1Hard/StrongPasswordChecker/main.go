package main

func strongPasswordChecker(password string) int {
	n := len(password)
	a1, nu, A1 := false, false, false
	m := map[int]byte{}
	if password == nil {
		return 6
	}
	if password[0] >= 0 && password[0] <= 9 {
		nu = true
	} else if password[0] >= 'a' && password[0] <= 'z' {
		a1 = true
	} else if password[0] >= 'A' && password[0] <= 'Z' {
		A1 = true
	}
	t := 1
	for i := 0; i < n; i++ {
		if password[i] >= 0 && password[i] <= 9 {
			nu = true
		} else if password[i] >= 'a' && password[i] <= 'z' {
			a1 = true
		} else if password[i] >= 'A' && password[i] <= 'Z' {
			A1 = true
		}
		if password[i] == password[i-1] {
			t += 1
		} else {
			if t >= 3 {
				m[i-1] = t
			}
			t = 1
		}
	}
	//
	if t >= 3 {
		m[i] = t
	}
	if n < 6 {
		return max(6-n, 3-(int(a1)+int(A1)+int(nu)))
	} else if n <= 20 {
		flag := 0
		for _, i := range m {
			flag += m[i] % 3
		}
		return max(3-(a1+A1+nu), flag)
	}
	x := 3 - (a1 + A1 + nu)
	m0, m1, m2 := map[int]int{}, map[int]int{}, map[int]int{}
	ans := x
	for _, i := range m {
		if m[i]%3 == 2 {
			m2[i] = m[i]
		} else if m[i]%3 == 1 {
			m1[i] = m[i]
		} else {
			m0[i] = m[i]
		}
	}
	//
	if x > 0 {
		for _, i := range m2 {
			for m2[i] >= 3 && x > 0 {
				x -= 1
				m2[i] -= 3
			}
			if x == 0 {
				break
			}
		}
	}
	//
	if x > 0 {
		for _, i := range m1 {
			for m1[i] >= 3 && x > 0 {
				x -= 1
				m1[i] -= 3
			}
			if x == 0 {
				break
			}
		}
	}
	//
	if x > 0 {
		for _, i := range m0 {
			for m0[i] >= 3 && x > 0 {
				x -= 1
				m0[i] -= 3
			}
			if x == 0 {
				break
			}
		}
	}
	//
	if x > 0 {
		return ans + n - 20

	}
	//
	for len(m0) > 0 || len(m1) > 0 || len(m2) > 0 {
		/* 	k0 = list(m0.keys())
			for i in k0:
				a = m0.pop(i)  # 删掉i
				if a >= 3:
					n -= 1
					ans += 1
					m2[i] = a - 1  # 余数变为2
				if n == 20:
					break

			if n > 20:  # n>20说明删m0不够
				k1 = list(m1.keys())
				for i in k1:  # 再删m1
					a = m1.pop(i)
					if a >= 3:
						n -= 2
						ans += 2
						m2[i] = a - 2 #每次删2个，余数变为2
					if n == 20:
						break
					if n == 19:  # 剩19个说明多删了，只需要删1个，省下连续的后面再处理
						n += 1
						ans -= 1
						m0[i] = a - 1 # 删2个变为删1个
						m2.pop(i) #删去刚刚增加的m2的值
						break

			if n > 20:  # 删m1没结束
				k2 = list(m2.keys())
				for i in k2:  # 最后删m2
					a = m2.pop(i)
					if a >= 3:
						n -= 3
						ans += 3
						m2[i] = a - 3
					if n == 20:
						break
					elif n == 19:  # 假如剩下22个，删3个变为删2个
						n += 1
						ans -= 1
						m0[i] = a - 2
						m2.pop(i)
						break
					elif n == 18:  # 剩下21个，删3个变为删1个
						n += 2
						ans -= 2
						m1[i] = a - 1
						m2.pop(i)
						break
			if n <= 20:  # 删除到固定长度
				break
		if n > 20:  # 连续的用完了，删去多余的值即可
			return ans + n - 20
		for i in m2: #还有连续的，则检查连续值，依次替换得结果
			ans+=m2[i]//3
		for i in m1:
			ans+=m1[i]//3
		for i in m0:
			ans+=m0[i]//3
		return ans */

	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
