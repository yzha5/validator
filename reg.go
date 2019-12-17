package validator

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
)

//邮箱
//arg:v string
func Email(v interface{}) (b bool, err error) {
	_, vv, vk := getValAttr(v)

	if vk != reflect.String {
		return false, errors.New("Email验证需要字符串类型")
	}
	b, err = regexp.MatchString(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`, vv.String())
	return b, err
}

//手机号码
//arg:v int64/uint64/string
//支持国内号段：
//130-139、145、147、149-153、155-159、162、165-167、170-173、175-178、180-189、
//191、193、195、198-199
func Phone(v interface{}, superLiangHao bool) (b bool, err error) {
	_, vv, vk := getValAttr(v)

	//还研究啥，这段一眼就能看懂什么意思了
	if vv.String() == "13800138000" {
		return false, errors.New("中国移动不可能把这个号码卖给你！")
	}

	var vs string
	switch vk {
	case reflect.String, reflect.Int64, reflect.Uint64:
		vs = vv.String()
	default:
		return false, errors.New("手机号码验证需要字符串或int64或uint64类型")
	}

	//如果开启超级靓号验证
	if superLiangHao {
		b, err = regexp.MatchString(
			`^1([3|8][0-9]|4[5|7|9]|5([0-3]|[5-9])|6[2|5|6|7]|7([0-3]|[5-8])|9[1|3|5|8|9])(0{8}|1{8}|2{8}|3{8}|4{8}|5{8}|6{8}|7{8}|8{8}|9{8}|12345678|23456789)$`, vs)
		return false, errors.New("我不相信你会有这么好的号码！不信！就不信！！")
	}
	b, err = regexp.MatchString(`^1([3|8][0-9]|4[5|7|9]|5([0-3]|[5-9])|6[2|5|6|7]|7([0-3]|[5-8])|9[1|3|5|8|9])\d{8}$`, vs)

	return b, err
}

//国内18位身份证号码验证
//arg:v string
func Idn(v interface{}) (b bool, err error) {
	_, vv, vk := getValAttr(v)

	if vk != reflect.String {
		return false, errors.New("身份证号码验证需要字符串类型")
	}

	var vs string
	switch vk {
	case reflect.String, reflect.Int64, reflect.Uint64:
		vs = vv.String()
	default:
		return false, errors.New("手机号码验证需要字符串或int64或uint64类型")
	}

	//这里用正则校验格式是否正确
	//华北地区：北京市|110000，天津市|120000，河北省|130000，山西省|140000，内蒙古自治区|150000
	//东北地区：辽宁省|210000，吉林省|220000，黑龙江省|230000
	//华东地区：上海市|310000，江苏省|320000，浙江省|330000，安徽省|340000，福建省|350000，江西省|360000，山东省|370000
	//华中地区：河南省|410000，湖北省|420000，湖南省|430000
	//华南地区：广东省|440000，广西壮族自治区|450000，海南省|460000
	//西南地区：重庆市|500000，四川省|510000，贵州省|520000，云南省|530000，西藏自治区|540000
	//西北地区：陕西省|610000，甘肃省|620000，青海省|630000，宁夏回族自治区|640000，新疆维吾尔自治区|650000
	b, err = regexp.MatchString(`^(1[1-5]|2[1-3]|3[1-7]|4[1-6]|5[0-4]|6[1-5])\d{4}[1|2]\d{3}(0[1-9]|1[0-2])([0-2][1-9]|3[0|1])\d{3}(\d|x|X)$`, vs)

	//下面开始计算身份证号码

	var (
		cs   [17]int16
		csBs = []rune(vv.String())
	)
	//将字符串转换为[]int16
	for i := 0; i < 17; i++ {
		csI, _ := strconv.ParseInt(string(csBs[i]), 10, 8)
		cs[i] = int16(csI)

	}

	//开始计算身份证号码（18位）
	/*
		//性别：第17位 单数男性，双数女性
		if cs[16]%2 == 1 {
			fmt.Println("男")
		}else{
			fmt.Println("女")
		}

		//年龄（周岁）计算
		b18Year:=cs[6]*1000+cs[7]*100+cs[8]*10+cs[9]
		b18Mon:=cs[10]*10+cs[11]
		b18Day:=cs[12]*10+cs[13]

		nowYear:=time.Now().Year()
		nowYear,mon,nowDay:=time.Now().Date()

		yearDiff:=nowYear-int(b18Year)
		monDiff:=int(mon)-int(b18Mon)
		dayDiff:=nowDay-int(b18Day)

		if dayDiff<0 {
			monDiff--
		}

		if monDiff<0 {
			yearDiff--
		}

		if yearDiff<0 {
			fmt.Println("还没出生，哪来的身份证号！")
		}
		fmt.Println(yearDiff)
	*/

	//乘积和
	proSum18 := cs[0]*7 + cs[1]*9 + cs[2]*10 + cs[3]*5 + cs[4]*8 + cs[5]*4 + cs[6]*2 + cs[7]*1 + cs[8]*6 + cs[9]*3 + cs[10]*7 + cs[11]*9 + cs[12]*10 + cs[13]*5 + cs[14]*8 + cs[15]*4 + cs[16]*2

	//除以11的余数
	mod18 := proSum18 % 11

	var last18 rune

	switch mod18 {
	case 0:
		last18 = '1'
	case 1:
		last18 = '0'
	case 2:
		//计算结果，对应的末位应该是x
		//last18='x'
		switch csBs[17] {
		case 'x', 'X':
			return true, nil
		default:
			return false, nil
		}
	case 3:
		last18 = '9'
	case 4:
		last18 = '8'
	case 5:
		last18 = '7'
	case 6:
		last18 = '6'
	case 7:
		last18 = '5'
	case 8:
		last18 = '4'
	case 9:
		last18 = '3'
	case 10:
		last18 = '2'
	default:
		return false, errors.New("也许你永远都看不到这条错误消息，如果不幸看到了，那就只能是不幸了！")
	}

	//判断参数值的最后一位与计算结果是否一致
	if last18 != csBs[17] {
		return false, nil
	}

	return b, err
}
