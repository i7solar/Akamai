package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/CrimsonAIO/radix"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Bmak BMAK (Akamai) Struct
type Bmak struct {
	frontofSensor     string
	Parameter100      string
	Parameter100Front string
	Parameter100Back  string
	Parameter101      string
	Parameter102      string
	Parameter105      string
	Parameter108      string
	Parameter110      string
	Parameter111      string
	Parameter117      string
	Parameter109      string
	Parameter114      string
	Parameter103      string
	Parameter112      string
	Parameter115      string
	Parameter106      string
	Parameter119      string
	Parameter122      string
	Parameter123      string
	Parameter124      string
	Parameter126      string
	Parameter127      string
	Parameter70       string
	Parameter80       string
	Parameter116      string
	Parameter118      string
	Parameter129      string
	Parameter121      string
	it0Flag           bool

	ver            string
	ke_cnt_lmt     int64
	mme_cnt_lmt    int64
	mduce_cnt_lmt  int64
	pme_cnt_lmt    int64
	pduce_cnt_lmt  int64
	tme_cnt_lmt    int64
	tduce_cnt_lmt  int64
	doe_cnt_lmt    int64
	dme_cnt_lmt    int64
	vc_cnt_lmt     int64
	doa_throttle   int64
	dma_throttle   int64
	js_post        bool
	loc            string
	auth           string
	api_public_key string
	aj_lmt_doact   int64
	aj_lmt_dmact   int64
	aj_lmt_tact    int64
	ce_js_post     int64
	init_time      int64
	informinfo     string
	prevfid        int64
	fidcnt         int64
	sensor_data    string
	xagg           int64
	pen            int64
	wen            int64
	den            int64
	brow           string
	browver        string
	psub           string
	lang           string
	prod           string
	plen           int64
	doadma_en      int64
	d2             int64
	d3             int64
	thr            int64
	cs             string
	hn             string
	z1             int64
	o9             int64
	vc             string
	y1             int64
	ta             int64
	tst            int64
	t_tst          int64
	ckie           string
	n_ck           int64
	ckurl          int64
	bm             bool
	mr             string
	altFonts       bool
	rst            bool
	runFonts       bool
	fsp            bool
	firstLoad      bool
	pstate         bool
	mn_mc_lmt      int64
	mn_state       int64
	mn_mc_indx     int64
	mn_sen         int64
	mn_tout        int64
	mn_stout       int64
	mn_ct          int64
	mn_cc          string
	mn_cd          int64
	mn_lc          []string
	mn_ld          []string
	mn_lcl         int64
	mn_al          []string
	mn_il          []string
	mn_tcl         []string
	mn_r           []string
	mn_rt          int64
	mn_wt          int64
	mn_abck        string
	mn_psn         string
	mn_ts          string
	mn_lg          []string
	loap           int64
	dcs            int64
	wl             int64
	weh            string
	start_ts       int64
	kact           string
	ke_cnt         int64
	ke_vel         int64
	mact           string
	mme_cnt        int64
	mduce_cnt      int64
	me_vel         int64
	pact           string
	pme_cnt        int64
	pduce_cnt      int64
	pe_vel         int64
	tact           string
	tme_cnt        int64
	tduce_cnt      int64
	te_vel         int64
	doact          string
	doe_cnt        int64
	doe_vel        int64
	dmact          string
	dme_cnt        int64
	dme_vel        int64
	vcact          string
	vc_cnt         int64
	aj_indx        int64
	aj_ss          int64
	aj_type        int64
	aj_indx_doact  int64
	aj_indx_dmact  int64
	aj_indx_tact   int64
	me_cnt         int64
	pe_cnt         int64
	te_cnt         int64
	nav_perm       string
	brv            int64
	hbCalc         bool
	fmh            string
	fmz            string
	ssh            string
	wv             string
	wr             string
	fpValStr       string
	rVal           int64
	rCFP           int64
	fpcfTD         int64
	device         AkamaiDevice
}

func (bmak *Bmak) floatToString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

func (bmak *Bmak) GenerateRandomNumber(min int, max int) int64 {
	rand.Seed(time.Now().UTC().UnixNano())
	return int64(min + rand.Intn(max-min))
}

func (bmak *Bmak) initBMAK(device AkamaiDevice) {
	bmak.ver = "1.7"
	bmak.ke_cnt_lmt = 150
	bmak.mme_cnt_lmt = 100
	bmak.mduce_cnt_lmt = 75
	bmak.pme_cnt_lmt = 25
	bmak.pduce_cnt_lmt = 25
	bmak.tme_cnt_lmt = 25
	bmak.tduce_cnt_lmt = 25
	bmak.doe_cnt_lmt = 10
	bmak.dme_cnt_lmt = 10
	bmak.vc_cnt_lmt = 100
	bmak.doa_throttle = 0
	bmak.dma_throttle = 0
	bmak.js_post = false
	bmak.loc = ""
	bmak.auth = ""
	bmak.api_public_key = "afSbep8yjnZUjq3aL010jO15Sawj2VZfdYK8uY90uxq"
	bmak.aj_lmt_doact = 1
	bmak.aj_lmt_dmact = 1
	bmak.aj_lmt_tact = 1
	bmak.ce_js_post = 0
	bmak.init_time = 0
	bmak.informinfo = ""
	bmak.prevfid = -1
	bmak.fidcnt = 0
	bmak.sensor_data = ""
	bmak.xagg = -1
	bmak.pen = -1
	bmak.brow = ""
	bmak.browver = ""
	bmak.psub = "-"
	bmak.lang = "-"
	bmak.prod = "-"
	bmak.plen = 3
	bmak.doadma_en = 0
	bmak.d2 = 0
	bmak.d3 = 0
	bmak.thr = 0
	bmak.cs = "0a46G5m17Vrp4o4c"
	bmak.hn = "unk"
	bmak.z1 = 0
	bmak.o9 = 0
	bmak.vc = ""
	bmak.y1 = 2016
	bmak.ta = 0
	bmak.tst = -1
	bmak.t_tst = 0
	bmak.ckie = "_abck"
	bmak.n_ck = 0
	bmak.ckurl = 0
	bmak.bm = false
	bmak.mr = "-1"
	bmak.altFonts = false
	bmak.rst = false
	bmak.runFonts = false
	bmak.fsp = false
	bmak.firstLoad = true
	bmak.pstate = false
	bmak.mn_mc_lmt = 10
	bmak.mn_state = 0
	bmak.mn_mc_indx = 0
	bmak.mn_sen = 0
	bmak.mn_tout = 100
	bmak.mn_stout = 1e3
	bmak.mn_ct = 1
	bmak.mn_cc = ""
	bmak.mn_cd = 1e4
	bmak.mn_lc = nil
	bmak.mn_ld = nil
	bmak.mn_lcl = 0
	bmak.mn_al = nil
	bmak.mn_il = nil
	bmak.mn_tcl = nil
	bmak.mn_r = nil
	bmak.mn_rt = 0
	bmak.mn_wt = 0
	bmak.mn_abck = ""
	bmak.mn_psn = ""
	bmak.mn_ts = ""
	bmak.mn_lg = nil
	bmak.loap = 1
	bmak.dcs = 0
	bmak.device = device
	bmak.fpValStr = "-1"
	bmak.rVal = -1
	bmak.rCFP = -1
	bmak.fpcfTD = -999999

}

func (bmak *Bmak) ir() {

	bmak.start_ts = bmak.get_cf_date() - 1
	bmak.kact = ""
	bmak.ke_cnt = 0
	bmak.ke_vel = 0
	bmak.mact = ""
	bmak.it0Flag = false
	bmak.mme_cnt = 0
	bmak.mduce_cnt = 0
	bmak.me_vel = 0
	bmak.pact = ""
	bmak.pme_cnt = 0
	bmak.pduce_cnt = 0
	bmak.pe_vel = 0
	bmak.tact = ""
	bmak.tme_cnt = 0
	bmak.tduce_cnt = 0
	bmak.te_vel = 0
	bmak.doact = ""
	bmak.doe_cnt = 0
	bmak.doe_vel = 0
	bmak.dmact = ""
	bmak.dme_cnt = 0
	bmak.dme_vel = 0
	bmak.vcact = ""
	bmak.vc_cnt = 0
	bmak.aj_indx = 0
	bmak.aj_ss = 0
	bmak.aj_type = 0
	bmak.aj_indx_doact = 0
	bmak.aj_indx_dmact = 0
	bmak.aj_indx_tact = 0
	bmak.me_cnt = 0
	bmak.pe_cnt = 0
	bmak.te_cnt = 0
	bmak.nav_perm = "8"
	bmak.brv = 0
	bmak.hbCalc = false
	bmak.fmh = ""
	bmak.fmz = ""
	bmak.ssh = ""
	bmak.wv = ""
	bmak.wr = ""
	bmak.weh = ""
	bmak.wl = 0

}

func (bmak *Bmak) get_cf_date() int64 {
	return time.Now().UTC().UnixNano() / 1e6
}

func (bmak *Bmak) uar() string {
	return strings.Replace(bmak.device.UserAgent, "/\\\\|\"/g", "", -1)
}

func (bmak *Bmak) ab(t string) int64 {
	var a int64
	var n rune

	for e := 0; e < len(t); e++ {
		n = rune(t[e])
		if n < 128 {
			a += int64(n)
		}
	}

	return a
}

func (bmak *Bmak) pi(t string) (int64, error) {
	i, err := strconv.Atoi(t)
	if err != nil {
		return 0, err
	}
	return int64(i), nil
}

func (bmak *Bmak) gbrv() int64 {
	return 0 // Brave is always 0.
}

func (bmak *Bmak) get_browser() {
	bmak.psub = bmak.device.productSub
	bmak.lang = "en-US"
	bmak.prod = bmak.device.productName
	bmak.plen = 3 //TODO: Dynamically create different plugin lengths.
}

func (bmak *Bmak) bc() {
	// 11059 firefox, 12147 chrome
	if bmak.device.browserName == "Chrome" || bmak.device.browserName == "Edge" {
		bmak.xagg = 12147
	} else {
		bmak.xagg = 11059
	}
}

func (bmak *Bmak) bmisc() {
	bmak.pen = 0
	bmak.wen = 0
	bmak.den = 0
	if bmak.device.browserName == "Chrome" || bmak.device.browserName == "Edge" {
		bmak.plen = 3
	} else {
		bmak.plen = 0
	}
}

func (bmak *Bmak) x2() float64 {
	/* Converts to (get_cf_date), a useless function on the akamai script.
	t := Bmak.ff
	a := t(98) + t(109) + t(97) + t(107)
	e := t(103) + t(101) + t(116) + t(95) + t(99) + t(100) + t(97) + t(116) + t(101)
	*/
	return float64(bmak.get_cf_date())
}

func (bmak *Bmak) ff(n int) string {
	return string(rune(n))
}

func (bmak *Bmak) cc(t int64) func(t, a int64) int64 {
	var a = t % 4
	if a == 2 {
		a = 3
	}
	var e = 42 + a
	n := func(t, a int64) int64 {
		return 0
	}
	if e == 42 {
		n = func(t, a int64) int64 {
			return t * a
		}
	} else if e == 43 {
		n = func(t, a int64) int64 {
			return t + a
		}
	} else {
		n = func(t, a int64) int64 {
			return t - a
		}
	}

	return n
}

func (bmak *Bmak) to() error {

	t := int64(math.Mod(bmak.x2(), 1e7))
	bmak.d3 = t
	var a = t
	e, err := bmak.pi(bmak.ff(51))
	if err != nil {
		return err
	}

	for n := 0; n < 5; n++ {
		o, err := bmak.pi(strconv.FormatInt(t/int64(math.Pow(10, float64(n)))%10, 10))
		if err != nil {
			return err
		}
		m := o + 1
		op := bmak.cc(o)
		a = op(a, m)
	}

	bmak.o9 = a * e
	return nil
}

func (bmak *Bmak) bd() (string, error) {

	var err error
	bmak.d2, err = bmak.pi(strconv.FormatInt(bmak.z1/23, 10))
	if err != nil {
		return "", err
	}

	if strings.Contains(bmak.device.browserName, "Chrome") || strings.Contains(bmak.device.browserName, "Edge") {
		return ",cpen:0,i1:0,dm:0,cwen:0,non:1,opc:0,fc:0,sc:0,wrc:1,isc:0,vib:1,bat:1,x11:0,x12:1", nil
	} else {
		iscList := [21]string{
			"59.5", "255", "259", "863.3333129882812", "113", "93.5", "518", "73.0999984741211", "68", "64", "64.5999984741211", "51",
			"85", "323.75", "78.19999694824219", "289.2166748046875", "102", "93.5", "198.56666564941406", "647.5", "172.6666717529297"}

		return ",cpen:0,i1:0,dm:0,cwen:0,non:1,opc:0,fc:1,sc:0,wrc:1,isc:" + iscList[bmak.GenerateRandomNumber(0, len(iscList))] + ",vib:1,bat:0,x11:0,x12:1", nil
	}
}

// Parameter 100
func (bmak *Bmak) gd() error {

	d := rand.Float64()
	s, err := bmak.pi(strconv.FormatInt(int64(1e3*d/2), 10))
	if err != nil {
		return err
	}
	k := bmak.floatToString(d) + ""
	k = k[0:11] + strconv.FormatInt(s, 10)

	/*
		Parameter 100 is updated throughout each POST of the script only in the (k) value.
		So if the aj_indx is 0, we'll set the front and back values but each time we'll generate a new k value following the sensor data/script's behavior.
	*/
	if bmak.aj_indx == 0 {

		t := bmak.uar()
		a := bmak.ab(t)
		e := bmak.floatToString(float64(bmak.start_ts) / 2.0)

		bmak.z1, err = bmak.pi(strconv.FormatInt(bmak.start_ts/(bmak.y1*bmak.y1), 10))
		if err != nil {
			return err
		}

		bmak.gbrv()
		bmak.get_browser()
		bmak.bc()
		bmak.bmisc()

		// Bmak.bd() called below.
		bd, err := bmak.bd()
		if err != nil {
			return err
		}
		bmak.Parameter100Front = t + ",uaend," + strconv.FormatInt(bmak.xagg, 10) + "," + bmak.psub + "," + bmak.lang + "," + bmak.prod + "," + strconv.FormatInt(bmak.plen, 10) + "," + strconv.FormatInt(bmak.pen, 10) + "," + strconv.FormatInt(bmak.wen, 10) + "," + strconv.FormatInt(bmak.den, 10) + "," + strconv.FormatInt(bmak.z1, 10) + "," + strconv.FormatInt(bmak.d3, 10) + "," + bmak.device.screenResolution + "," + bd + "," + strconv.FormatInt(a, 10)
		bmak.Parameter100Back = e + "," + strconv.FormatInt(bmak.brv, 10) + ",loc:" + bmak.loc
	}

	// Assign Parameter100 with the front, k value (randomNumber), and the back.
	bmak.Parameter100 = bmak.Parameter100Front + "," + k + "," + bmak.Parameter100Back

	return nil
}

func (bmak *Bmak) getforminfo(getdurl string) {

	if strings.Contains(getdurl, "finishline.com") {
		bmak.Parameter105 = "0,0,0,0,2402,310,0;0,0,0,0,1802,310,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,520,0;0,-1,0,0,-1,520,0;"
		bmak.Parameter102 = "0,0,0,0,2402,310,0;0,0,0,0,1802,310,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,520,0;0,-1,0,0,-1,520,0;"
	} else if strings.Contains(getdurl, "jdsports.com") {
		bmak.Parameter105 = "0,0,0,0,2402,310,0;0,0,0,0,1802,310,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,0;0,-1,0,0,-1,-1,1;0,-1,0,0,1143,520,0;"
		bmak.Parameter102 = "0,0,0,0,2402,310,0;0,0,0,0,1802,310,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,0;0,-1,0,0,-1,-1,1;0,-1,0,0,1143,520,0;"
	} else if strings.Contains(getdurl, "zalando.co.uk") {
		bmak.Parameter105 = "0,0,0,0,-1,113,0;0,-1,0,0,1125,-1,0;"
		bmak.Parameter102 = "0,0,0,0,-1,113,0;0,-1,0,0,1125,-1,0;"
	} else if strings.Contains(getdurl, "nike.com") {
		bmak.Parameter105 = "0,0,0,0,927,630,0;0,-1,0,1,560,-1,0;0,-1,0,1,1412,811,0;"
		bmak.Parameter102 = "0,0,0,0,927,630,0;0,-1,0,1,560,-1,0;"
	} else if strings.Contains(getdurl, "offspring.co.uk") {
		bmak.Parameter105 = "0,0,0,0,927,630,0;"
		bmak.Parameter102 = "0,0,0,0,927,630,0;"
	} else if strings.Contains(getdurl, "jdsports.com") {
		bmak.Parameter105 = "0,0,0,0,2402,310,0;0,0,0,0,1802,310,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,0;0,-1,0,0,-1,-1,1;0,-1,0,0,1136,520,0;1,0,0,0,883,883,0;0,-1,0,0,1143,520,0;0,-1,0,0,2194,706,0;0,-1,0,0,2658,2658,0;0,-1,0,0,2584,2584,0;"
		bmak.Parameter102 = "0,0,0,0,2402,310,0;0,0,0,0,1802,310,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,0;0,-1,0,0,-1,-1,1;0,-1,0,0,-1,-1,0;0,-1,0,0,-1,-1,1;0,-1,0,0,1136,520,0;1,0,0,0,883,883,0;0,-1,0,0,1143,520,0;0,-1,0,0,2194,706,0;0,-1,0,0,2658,2658,0;0,-1,0,0,2584,2584,0;"
	} else if strings.Contains(getdurl, "amd") {
		bmak.Parameter105 = "0,-1,0,0,1240,773,0;0,-1,0,0,1240,773,0;"
		bmak.Parameter102 = "0,-1,0,0,1240,773,0;0,-1,0,0,1240,773,0;"
	} else {
		bmak.Parameter105 = ""
		bmak.Parameter102 = ""
	}
}

func (bmak *Bmak) cal_dis(t []int64) float64 {
	var a = t[0] - t[1]
	var e = t[2] - t[3]
	var n = t[4] - t[5]
	var o = math.Sqrt(float64(a*a + e*e + n*n))
	return math.Floor(o)
}

func (bmak *Bmak) jrs(t int64) (float64, float64) {
	var a float64
	var e string
	var n int
	var m bool
	var o []int64
	a = math.Floor(1e5*rand.Float64() + 1e4)
	e = fmt.Sprintf("%f", float64(t)*a)
	n = 0
	m = len(e) >= 18
	for len(o) < 6 {
		intN, _ := strconv.ParseInt(e[n:n+2], 10, 32)
		o = append(o, intN)
		if m {
			n += 3
		} else {
			n += 2
		}
	}

	return a, bmak.cal_dis(o)
}

func (bmak *Bmak) fas() string {
	/*
		30261693 chrome
		26067385 firefox
	*/
	if bmak.device.browserName == "Chrome" || bmak.device.browserName == "Edge" {
		return "30261693"
	} else {
		return "26067385"
	}

}

func (bmak *Bmak) hbs() string {
	// 0 on all browsers.
	return "0"
}

func (bmak *Bmak) gwd() string {
	// -1 on all browsers.
	return "-1"
}

func (bmak *Bmak) updatet() int64 {
	return bmak.get_cf_date() - bmak.start_ts
}

func (bmak *Bmak) sed() {
	// Parameter122
	bmak.Parameter122 = "0,0,0,0,1,0,0"
}

func (bmak *Bmak) rir(a, b, c, d int) int {
	if a > 47 && a <= 57 {
		a += d % (c - b)
		if a > c {
			a = a - c + b
		}
	}
	return a
}

func (bmak *Bmak) od(a, t string) string {
	var e []string
	n := len(t)
	if n > 0 {
		for o, v := range a {
			r := string(v)
			i := t[o%n]

			vInt := int(v)
			m := bmak.rir(vInt, 47, 57, int(i))
			if m != vInt {
				r = string(rune(m))
			}
			e = append(e, r)
		}
		if len(e) > 0 {
			return strings.Join(e, "")
		}
	}
	return a
}

func (bmak *Bmak) parameter110_check(getdurl string) bool {
	// This function checks the target string in the opeartion. If the conditions match, then we will return true to add the parameter110 values for (istrusted, REF: from akamai script).
	if strings.Contains(getdurl, "finishline") {
		bmak.me_vel = bmak.GenerateRandomNumber(200, 350)
		return true
	} else {
		return false
	}
}

func (bmak *Bmak) generate_dmact(randomNumber int64) {
	t := (bmak.get_cf_date() - bmak.start_ts) + randomNumber
	bmak.dme_vel = t
	bmak.dmact = "0," + strconv.FormatInt(bmak.dme_vel, 10) + ",-1,-1,-1,-1,-1,-1,-1,-1,-1;"
	bmak.Parameter109 = bmak.dmact
}

func (bmak *Bmak) generate_doact(randomNumber int64) {
	t := (bmak.get_cf_date() - bmak.start_ts) + randomNumber + bmak.GenerateRandomNumber(1, 4)
	bmak.doe_vel = t
	bmak.doact = "0," + strconv.FormatInt(bmak.doe_vel, 10) + ",-1,-1,-1;"
	bmak.Parameter111 = bmak.doact
}

func (bmak *Bmak) generate_dmact_doact_timing() {
	n := bmak.GenerateRandomNumber(500, 1000)
	bmak.generate_dmact(n)
	bmak.generate_doact(n)
}

func (bmak *Bmak) generate_fpcf_td() {
	bmak.fpcfTD = bmak.GenerateRandomNumber(5, 25)
}

func (bmak *Bmak) calculate_ta() {
	bmak.ta = bmak.doe_vel + bmak.dme_vel + bmak.me_vel
}

func (bmak *Bmak) check_cookie_for_challenge(cookie string) bool {
	/* Checks if the cookie given back has challenge */
	if strings.Contains(cookie, "||1-") {
		return true
	} else {
		return false
	}
}

// Build Challenge String
func (bmak *Bmak) createChallengeString() {
	challengeString := bmak.mn_al[0] + "," + bmak.mn_al[1] + "," + bmak.mn_al[2] + "," + bmak.mn_al[3] + "," + bmak.mn_al[4] + "," + bmak.mn_al[5] + "," + bmak.mn_al[6] + "," + bmak.mn_al[7] + "," + bmak.mn_al[8] + "," + bmak.mn_al[9] + ";"
	challengeString = challengeString + bmak.mn_tcl[0] + "," + bmak.mn_tcl[1] + "," + bmak.mn_tcl[2] + "," + bmak.mn_tcl[3] + "," + bmak.mn_tcl[4] + "," + bmak.mn_tcl[5] + "," + bmak.mn_tcl[6] + "," + bmak.mn_tcl[7] + "," + bmak.mn_tcl[8] + "," + bmak.mn_tcl[9] + ";"
	challengeString = challengeString + bmak.mn_il[0] + "," + bmak.mn_il[1] + "," + bmak.mn_il[2] + "," + bmak.mn_il[3] + "," + bmak.mn_il[4] + "," + bmak.mn_il[5] + "," + bmak.mn_il[6] + "," + bmak.mn_il[7] + "," + bmak.mn_il[8] + "," + bmak.mn_il[9] + ";"
	challengeString = challengeString + bmak.mn_lg[0] + "," + bmak.mn_lg[1] + "," + bmak.mn_lg[2] + "," + bmak.mn_lg[3] + "," + bmak.mn_lg[4] + "," + bmak.mn_lg[5] + "," + bmak.mn_lg[6] + "," + bmak.mn_lg[7] + "," + bmak.mn_lg[8] + "," + bmak.mn_lg[9]
	challengeString = challengeString + "," + strconv.FormatInt(bmak.get_cf_date(), 10) + ";"

	bmak.Parameter124 = challengeString
}

func (bmak *Bmak) ParseABCKForFront(cookieString string) string {

	cookieStringArray := strings.Split(cookieString, "~")

	return cookieStringArray[0]

}

func (bmak *Bmak) mnS(t string) []int {
	bytesdt := []byte(t)
	hasher := sha256.New()
	hasher.Write(bytesdt)
	return ByteArrayToInt(hasher.Sum(nil))
}

// Takes the cookieString and splits it into an array of 7 values but also appends mn_ts
func (bmak *Bmak) mnUpdateChallengeDetails(cookieString string) error {

	/* mn_update_challenge_details takes the cookieString and splits it into an array of 7 values but also appends mn_ts.*/
	/////////////////////////////////
	// mn_poll calls get_mn_params_from_abck which grabs values from the cookies/sensorData

	/*
		mn_get_new_challenge_params uses that value as a param and is returned an array of (7) values
		these 7 values are from the challenge string/front cookie.
			ex. 1-abcd-1-10-1000-2
			ex(cont). (1), (front_cookie_value), (abcd), (1), (10), (1000), (2)

			mn_sen = 1, mn_abck = front cookie value, mn_psn = abc, mn_cd = 1, mn_tout = 10, mn_stout = 1000, mn_ct = 2
	*/
	// mn_ts = start timestamp
	// mn_cc will contain the ParsedFront of the cookie during sensor v2, the start timestamp and the challenge string.

	var (
		mn_sen, mn_cd, mn_tout, mn_stout, mn_ct int
		err                                     error
	)

	// We'll split it by the | entities.
	cookieValueArray := strings.Split(cookieString, "|")
	challengeValueArray := strings.Split(cookieValueArray[2], "-")

	mn_sen, err = strconv.Atoi(challengeValueArray[0])
	if err != nil {
		return err
	}
	mn_cd, err = strconv.Atoi(challengeValueArray[2])
	if err != nil {
		return err
	}
	mn_tout, err = strconv.Atoi(challengeValueArray[3])
	if err != nil {
		return err
	}
	mn_stout, err = strconv.Atoi(challengeValueArray[4])
	if err != nil {
		return err
	}
	mn_ct, err = strconv.Atoi(challengeValueArray[5])
	if err != nil {
		return err
	}

	bmak.mn_sen = int64(mn_sen)
	bmak.mn_abck = bmak.ParseABCKForFront(cookieString)
	bmak.mn_psn = challengeValueArray[1]
	bmak.mn_cd = int64(mn_cd)
	bmak.mn_tout = int64(mn_tout)
	bmak.mn_stout = int64(mn_stout)
	bmak.mn_ct = int64(mn_ct)
	bmak.mn_ts = strconv.FormatInt(bmak.start_ts, 10)
	bmak.mn_cc = bmak.mn_abck + strconv.FormatInt(bmak.start_ts, 10) + bmak.mn_psn

	return nil
}

func (bmak *Bmak) bdm(t []int, a int) int {

	var e int

	for index /* n */ := 0; index < len(t); index++ {
		e = int(uint32(e<<8|t[index]) >> 0)
		e %= a
	}

	return e
}

// Builds all of the Challenge Values
func (bmak *Bmak) buildChallenge(cookieString string) func() {
	// This is the mn_w function that is represented in the Akamai script.
	//alEntites make up the random strings we need to create, ex: 0.2d78cd1964972e

	err := bmak.mnUpdateChallengeDetails(cookieString)
	if err != nil {
		return nil
	}
	bmak.mn_lc = make([]string, 1)
	bmak.mn_ld = make([]string, 1)

	var t = 0
	var a = 0
	var e = 0
	var n string
	var o = int(bmak.get_cf_date()) - int(bmak.GenerateRandomNumber(0, 20))
	var m = bmak.mn_cd + bmak.mn_mc_indx

	for t == 0 {
		n = radix.ToString(rand.Float64(), 16)
		var r = bmak.mn_cc + strconv.Itoa(int(m)) + n
		var i = bmak.mnS(r)
		if 0 == bmak.bdm(i, int(m)) {
			t = 1
			e = int(bmak.get_cf_date()) - o
			bmak.mn_al = append(bmak.mn_al, n)

			if bmak.mn_mc_indx == 0 {
				bmak.mn_lg = append(bmak.mn_lg, bmak.mn_abck)
				bmak.mn_lg = append(bmak.mn_lg, bmak.mn_ts)
				bmak.mn_lg = append(bmak.mn_lg, bmak.mn_psn)
				bmak.mn_lg = append(bmak.mn_lg, bmak.mn_cc)
				bmak.mn_lg = append(bmak.mn_lg, strconv.Itoa(int(bmak.mn_cd)))
				bmak.mn_lg = append(bmak.mn_lg, strconv.Itoa(int(m)))
				bmak.mn_lg = append(bmak.mn_lg, n)
				bmak.mn_lg = append(bmak.mn_lg, r)
				bmak.mn_lg = append(bmak.mn_lg, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(i)), ","), "[]"))
				bmak.mn_lg = append(bmak.mn_lg, strconv.Itoa(int(bmak.mn_rt)))
			}

		} else if (a)%1e3 == 0 && int(bmak.get_cf_date())-o > int(bmak.mn_stout) {
			bmak.mn_wt += int64(e)
			//time.Sleep(time.Duration(bmak.mn_stout))
			return bmak.buildChallenge(cookieString)

		}

	}

	bmak.mn_mc_indx += 1

	if bmak.mn_mc_indx < bmak.mn_mc_lmt {
		//time.AfterFunc(time.Duration(e) * time.Second, bmak.buildChallenge(cookieString))
		//time.Sleep(time.Duration(e) * time.Second)
		bmak.buildChallenge(cookieString)
	} else {

		bmak.mn_mc_indx = 0
		//bmak.mn_lc[bmak.mn_lcl] = bmak.mn_cc
		bmak.mn_lc[0] = bmak.mn_cc
		bmak.mn_ld[bmak.mn_lcl] = strconv.Itoa(int(bmak.mn_cd))
		bmak.mn_lcl = bmak.mn_lcl + 1
		bmak.mn_state = 0
		bmak.mn_lg = append(bmak.mn_lg, strconv.Itoa(int(bmak.mn_wt)))
		bmak.mn_lg = append(bmak.mn_lg, strconv.Itoa(int(bmak.get_cf_date())))

		b := make([]byte, 10)

		_, _ = rand.Read(b)
		bmak.mn_tcl = ByteArrayToString(b)

		_, _ = rand.Read(b)
		bmak.mn_il = ByteArrayToString(b)

	}

	return func() {
	}
}

// Master for Challenge
func (bmak *Bmak) solveChallenge(cookie string) {
	err := bmak.mnUpdateChallengeDetails(cookie)
	if err != nil {
		return
	}
	bmak.buildChallenge(cookie)
	bmak.createChallengeString()
}

func (bmak *Bmak) getSensor(device AkamaiDevice, cookie string, getdurl string, fpCompute bool, mactCompute bool) (string, error) {

	// FIRST ROUND
	// If the fingerprint and MACT computation flags are both false, we're going through the first round.
	if fpCompute == false && mactCompute == false {
		bmak.t_tst = bmak.get_cf_date() // Assign t_tst value for it's timestamp value.
		//t := Bmak.get_cf_date() // Give t a timestamp value for it's start-time (which mimics the previous script).
		bmak.initBMAK(device) // Assign all values of BMAK init values to their variables.
		bmak.ir()             // Called at the start of Akamai script.
		err := bmak.to()      // Gives Bmak.o9 it's value and Bmak.d3.
		if err != nil {
			return "", err
		}

		var a = bmak.updatet() // Perfroms math from a get_cf_date and the start-time.
		e := cookie            // Cookie from the inital GET request as this is the init phase.
		o := "do_en"
		m := "dm_en"
		r := "t_en"
		bmak.Parameter101 = o + "," + m + "," + r // do_en,dm_en,t_en
		bmak.getforminfo(getdurl)                 //c = Bmak.getforminfo()
		bmak.Parameter112 = getdurl               // Website location of where the sensor is being generated.
		if err := bmak.gd(); err != nil {         // Bmak.Parameter100 and z1.
			return "", err
		}
		d := strconv.FormatInt(bmak.aj_type, 10) + "," + strconv.FormatInt(bmak.aj_indx, 10) // Type is 0 for first computation (init), aj_indx for the first POST will be 0. aj_indx increments for each post.
		bmak.Parameter106 = d

		/* During the first iteration, all of the values are 0, except for 1: TBD */
		var s = bmak.ke_vel + bmak.me_vel + bmak.doe_vel + bmak.dme_vel + bmak.te_vel + bmak.pe_vel
		//        0 (-110)   0 (set to 32?)  0(-111)         0(-109)         0(765)         0 (609)

		k := bmak.ff                                 // Assings the Bmak.ff function to the variable k.
		l := k(80) + k(105) + k(90) + k(116) + k(69) // Converts to PiZtE
		u0, u1 := bmak.jrs(bmak.start_ts)            // Returns two values back from jrs that are used in the (g) array which do change upon each increment of post.
		__ := bmak.get_cf_date() - bmak.start_ts + (bmak.GenerateRandomNumber(1, 3))
		f, err := bmak.pi(strconv.FormatInt(bmak.d2/6, 10))
		if err != nil {
			return "", err
		}
		p := bmak.fas() // Depends on the browser that's picked.
		v := bmak.hbs() // Always 0.
		h := bmak.gwd() // Always -1.
		g := []string{
			strconv.FormatInt(bmak.ke_vel+1, 10), strconv.FormatInt(bmak.me_vel+32, 10), strconv.FormatInt(bmak.te_vel+32, 10),
			strconv.FormatInt(bmak.doe_vel, 10), strconv.FormatInt(bmak.dme_vel, 10), strconv.FormatInt(bmak.pe_vel, 10), strconv.FormatInt(s, 10),
			strconv.FormatInt(a, 10), strconv.FormatInt(bmak.init_time, 10), strconv.FormatInt(bmak.start_ts, 10), strconv.FormatInt(bmak.fpcfTD, 10), strconv.FormatInt(bmak.d2, 10),
			strconv.FormatInt(bmak.ke_cnt, 10), strconv.FormatInt(bmak.me_cnt, 10), strconv.FormatInt(f, 10), strconv.FormatInt(bmak.pe_cnt, 10), strconv.FormatInt(bmak.te_cnt, 10),
			strconv.FormatInt(__, 10), strconv.FormatInt(bmak.ta, 10), strconv.FormatInt(bmak.n_ck, 10), e, strconv.FormatInt(bmak.ab(e), 10), strconv.FormatInt(bmak.rVal, 10),
			strconv.FormatInt(bmak.rCFP, 10), p, l, bmak.floatToString(u0), bmak.floatToString(u1), v, h,
		}
		w := strings.Join(g, ",")
		bmak.Parameter115 = w
		y := "" + strconv.FormatInt(bmak.ab(bmak.fpValStr), 10) // Bmak.Parameter80
		bmak.sed()                                              // Bmak.Parameter122

		// C = Bmak.mn_get_current_challenges() // Currently we'll end up supporting this inside of fingerprint computation from the Bmak.go file that was made previously.

		/* These values are usually always blank, just copied it for similarity */
		B := ""
		x := ""
		M := ""
		bmak.Parameter123 = B
		bmak.Parameter124 = x
		bmak.Parameter126 = M

		/*
			Upon the first stage, the following parameters are blank:
			-1,2,-94,-108, -1,2,-94,-110, -1,2,-94,-117, -1,2,-94,-111, -1,2,-94,-109, -1,2,-94,-114,
			-1,2,-94,-103, -1,2,-94,-123, -1,2,-94,-124, -1,2,-94,-126, -1,2,-94,-129,
		*/

		bmak.Parameter119 = bmak.mr // MR value during the first computation is "-1".
		bmak.Parameter127 = bmak.nav_perm
		bmak.Parameter70 = bmak.fpValStr
		bmak.Parameter80 = y
		bmak.Parameter116 = strconv.FormatInt(bmak.o9, 10)

		bmak.sensor_data = bmak.ver + "-1,2,-94,-100," + bmak.Parameter100 + "-1,2,-94,-101," + bmak.Parameter101 + "-1,2,-94,-105," + bmak.Parameter105 + "-1,2,-94,-102," + bmak.Parameter102 + "-1,2,-94,-108," + bmak.Parameter108 + "-1,2,-94,-110,"
		bmak.sensor_data = bmak.sensor_data + bmak.Parameter110 + "-1,2,-94,-117," + bmak.Parameter117 + "-1,2,-94,-111," + bmak.Parameter111 + "-1,2,-94,-109," + bmak.Parameter109 + "-1,2,-94,-114," + bmak.Parameter114 + "-1,2,-94,-103,"
		bmak.sensor_data = bmak.sensor_data + "-1,2,-94,-112," + bmak.Parameter112 + "-1,2,-94,-115," + bmak.Parameter115 + "-1,2,-94,-106," + bmak.Parameter106 + "-1,2,-94,-119," + bmak.Parameter119 + "-1,2,-94,-122," + bmak.Parameter122
		bmak.sensor_data = bmak.sensor_data + "-1,2,-94,-123," + bmak.Parameter123 + "-1,2,-94,-124," + bmak.Parameter124 + "-1,2,-94,-126," + bmak.Parameter126 + "-1,2,-94,-127," + bmak.Parameter127
		P := 24 ^ bmak.ab(bmak.sensor_data)
		bmak.Parameter118 = strconv.FormatInt(P, 10)
		bmak.Parameter129 = ""
		bmak.sensor_data = bmak.sensor_data + "-1,2,-94,-70," + bmak.Parameter70 + "-1,2,-94,-80," + bmak.Parameter80 + "-1,2,-94,-116," + bmak.Parameter116 + "-1,2,-94,-118," + bmak.Parameter118 + "-1,2,-94,-129," + bmak.Parameter129 + "-1,2,-94,-121,"

		F := bmak.od(bmak.cs, bmak.api_public_key)[0:16]
		D := math.Floor(float64(bmak.get_cf_date() / 36e5))
		R := bmak.get_cf_date()
		N := F + bmak.od(bmak.floatToString(D), F) // Create front sensor value that does not change after it is set.
		bmak.frontofSensor = N                     // Assign it to the frontofSensor value in struct property.
		N = bmak.frontofSensor + bmak.sensor_data  // Add the front value to the sensor data.
		//Bmak.sensor_data = N + ";" + strconv.FormatInt(Bmak.get_cf_date() - t, 10) + ";" + strconv.FormatInt(Bmak.tst, 10) + ";" + strconv.FormatInt(Bmak.get_cf_date() - R, 10)      // Assign final timings to the sensor data.
		bmak.sensor_data = N + ";" + strconv.Itoa(int(bmak.GenerateRandomNumber(3, 5))) + ";" + strconv.FormatInt(bmak.tst, 10) + ";" + strconv.FormatInt(bmak.get_cf_date()-R, 10) // Assign final timings to the sensor data.

		return bmak.sensor_data, nil
	}

	// SECOND ROUND
	// If the fpCompute is true, we'll go through fingerprint computation is the which is the second part of the sensor data POST.
	if fpCompute == true {
		/*

				AJ Type is always '9" for fingerprint computation stage and the index will always be "1" as it is the second increment in the phase of genning the sensor data.

				These values remain empty during the fpCompute stage:
				-1,2,-94,-108, -1,2,-94,-117, -1,2,-94,-114,
				-1,2,-94,-103, -1,2,-94,-123, -1,2,-94,-124, -1,2,-94,-126,

				We will fill/edit the following parameters during this stage:
				-1,2,-94,-100 (edit), -1,2,-94,-110 (add ONLY for certain sites) , -1,2,-94,-111 (add), -1,2,-94,-109 (add)
				-1,2,-94,-106 (edit), -1,2,-94,-127 (edit), -1,2,-94,-70 (edit), -1,2,-94,-80 (edit),
			    -1,2,-94,-118 (edit), -1,2,-94,-129 (add), -1,2,-94,-121 (edit)
		*/

		/* Append rcfpValue and rVal onto the sensor data durng the second computation. (fp) */
		rValInt, rvalError := strconv.Atoi(bmak.device.rValValue)
		if rvalError != nil {
			return "", rvalError
		}
		rcfpValue, rcfpError := strconv.Atoi(bmak.device.rCFPValue)
		if rcfpError != nil {
			return "", rcfpError
		}
		bmak.rVal = int64(rValInt)
		bmak.rCFP = int64(rcfpValue)

		bmak.tst = 6
		time_helper := bmak.GenerateRandomNumber(325, 375)
		bmak.aj_type = 9 // aj_type is set to 9 for fingerprint computation.
		//Bmak.aj_type = 8
		bmak.aj_indx = 1                                                                                    // index is changed to 1 for post interval.
		bmak.Parameter106 = strconv.FormatInt(bmak.aj_type, 10) + "," + strconv.FormatInt(bmak.aj_indx, 10) // Update parameter 106 to reflect (9,1).

		var a = bmak.updatet() + time_helper // Performs math from a get_cf_date and the start-time while adding a random number between 325-375 to make it look real.
		var e = cookie                       // Grab the new cookie from the arguments the function provides. (cookie :: type: string)

		/* Check if the cookie has a challenge on it. */
		if bmak.check_cookie_for_challenge(cookie) == true {
			bmak.solveChallenge(cookie)
		}

		err := bmak.gd()
		if err != nil { // Call gd function to update Parameter100 with a new k value (random number sliced), all other values will remain the same.
			return "", err
		}
		bmak.generate_dmact_doact_timing() // Fills in the dmact and doact values. (Parameter 111 & 109)

		// If the URL returns true from our include parameter 110 checker add the 0,2,(randomValue),-1,-1,-1,it0 value to our mact.

		if bmak.parameter110_check(getdurl) {
			bmak.mact = "0,2," + strconv.FormatInt(bmak.me_vel, 10) + ",-1,-1,-1,it0;"
			bmak.Parameter110 = bmak.mact
			bmak.it0Flag = true
		}

		/* Assign the Bmak.mr struct the device's mr value and assign it in parameter 119. */
		bmak.mr = bmak.device.mrValue
		bmak.Parameter119 = bmak.mr

		/* Assign the Bmak.nav_perm struct the device's np value and assign it in parameter 127. */
		bmak.nav_perm = bmak.device.np
		bmak.Parameter127 = bmak.nav_perm

		/* Assign the Bmak.fpValString struct the device's fpValString value and assign it in parameter 70. */
		bmak.fpValStr = bmak.device.fpValString
		bmak.Parameter70 = bmak.fpValStr

		/* Assign y value to the value computed by passing fpvalstr to ab(), assign to parameter 80 */
		y := "" + strconv.FormatInt(bmak.ab(bmak.fpValStr), 10)
		bmak.Parameter80 = y

		/* Calculate the fpcf.TD value by creating a random number from 5-25. */
		bmak.generate_fpcf_td()

		/* Calculate the TA which is based on the: Bmak.doe_vel + Bmak.dme_vel + Bmak.me_vel :: Called upon the fpCompute and MACT parts only! */
		bmak.calculate_ta()

		/*  Add all values together to create (s) known as the sum. */
		var s = bmak.ke_vel + bmak.me_vel + bmak.doe_vel + bmak.dme_vel + bmak.te_vel + bmak.pe_vel

		/*  Change me_cnt to 1 to indicate that a small movement was made prior to the script being sent off during fingerprint computation.. */
		bmak.me_cnt = 1

		k := bmak.ff                                 // Assings the Bmak.ff function to the variable k.
		l := k(80) + k(105) + k(90) + k(116) + k(69) // Converts to PiZtE
		u0, u1 := bmak.jrs(bmak.start_ts)            // Returns two values back from jrs that are used in the (g) array which do change upon each increment of post.
		__ := bmak.get_cf_date() - bmak.start_ts + (time_helper) - bmak.GenerateRandomNumber(4, 9)
		f, err := bmak.pi(strconv.FormatInt(bmak.d2/6, 10))
		if err != nil {
			return "", err
		}
		p := bmak.fas() // Depends on the browser that's picked.
		v := bmak.hbs() // Always 0.
		h := bmak.gwd() // Always -1.
		g := []string{
			strconv.FormatInt(bmak.ke_vel+1, 10), strconv.FormatInt(bmak.me_vel+32, 10), strconv.FormatInt(bmak.te_vel+32, 10),
			strconv.FormatInt(bmak.doe_vel, 10), strconv.FormatInt(bmak.dme_vel, 10), strconv.FormatInt(bmak.pe_vel, 10), strconv.FormatInt(s, 10),
			strconv.FormatInt(a, 10), strconv.FormatInt(bmak.init_time, 10), strconv.FormatInt(bmak.start_ts, 10), strconv.FormatInt(bmak.fpcfTD, 10), strconv.FormatInt(bmak.d2, 10),
			strconv.FormatInt(bmak.ke_cnt, 10), strconv.FormatInt(bmak.me_cnt, 10), strconv.FormatInt(f, 10), strconv.FormatInt(bmak.pe_cnt, 10), strconv.FormatInt(bmak.te_cnt, 10),
			strconv.FormatInt(__, 10), strconv.FormatInt(bmak.ta, 10), strconv.FormatInt(bmak.n_ck, 10), e, strconv.FormatInt(bmak.ab(e), 10), strconv.FormatInt(bmak.rVal, 10),
			strconv.FormatInt(bmak.rCFP, 10), p, l, bmak.floatToString(u0), bmak.floatToString(u1), v, h,
		}
		w := strings.Join(g, ",")
		bmak.Parameter115 = w

		bmak.sensor_data = bmak.ver + "-1,2,-94,-100," + bmak.Parameter100 + "-1,2,-94,-101," + bmak.Parameter101 + "-1,2,-94,-105," + bmak.Parameter105 + "-1,2,-94,-102," + bmak.Parameter102 + "-1,2,-94,-108," + bmak.Parameter108 + "-1,2,-94,-110,"
		bmak.sensor_data = bmak.sensor_data + bmak.Parameter110 + "-1,2,-94,-117," + bmak.Parameter117 + "-1,2,-94,-111," + bmak.Parameter111 + "-1,2,-94,-109," + bmak.Parameter109 + "-1,2,-94,-114," + bmak.Parameter114 + "-1,2,-94,-103,"
		bmak.sensor_data = bmak.sensor_data + "-1,2,-94,-112," + bmak.Parameter112 + "-1,2,-94,-115," + bmak.Parameter115 + "-1,2,-94,-106," + bmak.Parameter106 + "-1,2,-94,-119," + bmak.Parameter119 + "-1,2,-94,-122," + bmak.Parameter122
		bmak.sensor_data = bmak.sensor_data + "-1,2,-94,-123," + bmak.Parameter123 + "-1,2,-94,-124," + bmak.Parameter124 + "-1,2,-94,-126," + bmak.Parameter126 + "-1,2,-94,-127," + bmak.Parameter127
		P := 24 ^ bmak.ab(bmak.sensor_data)
		bmak.Parameter118 = strconv.FormatInt(P, 10)
		bmak.Parameter129 = bmak.device.fmh + ",1," + bmak.device.RawDevice.SSH + ",n,n,n,0"
		bmak.sensor_data = bmak.sensor_data + "-1,2,-94,-70," + bmak.Parameter70 + "-1,2,-94,-80," + bmak.Parameter80 + "-1,2,-94,-116," + bmak.Parameter116 + "-1,2,-94,-118," + bmak.Parameter118 + "-1,2,-94,-129," + bmak.Parameter129 + "-1,2,-94,-121,"
		R := bmak.get_cf_date()
		bmak.sensor_data = bmak.frontofSensor + bmak.sensor_data
		bmak.sensor_data = bmak.sensor_data + ";" + strconv.FormatInt(bmak.GenerateRandomNumber(12, 16), 10) + ";" + strconv.FormatInt(bmak.tst, 10) + ";" + strconv.FormatInt(bmak.get_cf_date()-R, 10) // Assign final timings to the sensor data.

		return bmak.sensor_data, nil

	}

	// THIRD ROUND
	// If the mactCompute is true, we'll go through mact computation is the which is the third and final part of the sensor data POST.
	if mactCompute == true {
		/*

			AJ Type is either 7 or 1 for mactCompute computation stage and the index will always be "2" as it is the third increment in the phase of genning the sensor data.

			These values remain empty during the mactCompute stage:
			-1,2,-94,-108, -1,2,-94,-117, -1,2,-94,-114,
			-1,2,-94,-123, -1,2,-94,-124, -1,2,-94,-126,

			We will fill/edit the following parameters during this stage:
			-1,2,-94,-100 (edit), -1,2,-94,-103 (add), -1,2,-94,-118 (edit),
			-1,2,-94,-106 (edit), -1,2,-94,-115 (edit), , -1,2,-94,-121 (edit)
		*/

		bmak.Parameter124 = ""
		time_helper := bmak.GenerateRandomNumber(7000, 9000)

		var a = bmak.updatet() + time_helper // Performs math from a get_cf_date and the start-time while adding a random number between 7000-9000 to make it look real.
		var e = cookie                       // Grab the new cookie from the arguments the function provides. (cookie :: type: string)

		/* Check if the cookie has a challenge on it. */
		if bmak.check_cookie_for_challenge(cookie) == true {
			bmak.solveChallenge(cookie)
		}

		/* The aj_type can be different based on either a  7 for DeviceMovement event or it could be a 1 for mouseclick event. */
		if bmak.GenerateRandomNumber(0, 1000)%2 == 0 {
			bmak.aj_type = 7
		} else {
			bmak.aj_type = 1
		}

		bmak.aj_type = 6
		bmak.aj_indx = 3                                                                                    // index is changed to 2 for post interval.
		bmak.Parameter106 = strconv.FormatInt(bmak.aj_type, 10) + "," + strconv.FormatInt(bmak.aj_indx, 10) // Update parameter 106 to reflect (7 OR 1, 2).

		if err := bmak.gd(); err != nil { // Call gd function to update Parameter100 with a new k value (random number sliced), all other values will remain the same.
			return "", err
		}

		/* genMact returns the mactString used for 110, the timeStamp total for me_vel, and the Bmak.ta). */
		if !strings.Contains(getdurl, "nike") {
			mactString, mactTimeStampTotal, mactCompleteTotal := simplexNoiseGen(bmak.it0Flag)
			bmak.me_vel += mactTimeStampTotal /* Append the time of all of the mact timestamps to me_vel including the one from fingerprint computation (step 2). */
			bmak.Parameter110 += mactString   /* Append the string of the mact string created from genMact */
			bmak.ta += mactCompleteTotal      /* Append the complete total of the (index, event_type, timestamp, x, y) of all MACT values into TA which includes doe_vel, dme_vel, and me_vel */
		}

		/*  Add all values together to create (s) known as the sum. */
		var s = bmak.ke_vel + bmak.me_vel + bmak.doe_vel + bmak.dme_vel + bmak.te_vel + bmak.pe_vel

		/*  Change me_cnt to a random number from 109-289 to indicate that alot of movement events happened. */
		bmak.me_cnt = bmak.GenerateRandomNumber(109, 289)

		k := bmak.ff                                 // Assings the Bmak.ff function to the variable k.
		l := k(80) + k(105) + k(90) + k(116) + k(69) // Converts to PiZtE
		u0, u1 := bmak.jrs(bmak.start_ts)            // Returns two values back from jrs that are used in the (g) array which do change upon each increment of post.
		__ := bmak.get_cf_date() - bmak.start_ts + (time_helper) - bmak.GenerateRandomNumber(4, 9)
		f, err := bmak.pi(strconv.FormatInt(bmak.d2/6, 10))
		if err != nil {
			return "", err
		}
		p := bmak.fas() // Depends on the browser that's picked.
		v := bmak.hbs() // Always 0.
		h := bmak.gwd() // Always -1.
		g := []string{
			strconv.FormatInt(bmak.ke_vel+1, 10), strconv.FormatInt(bmak.me_vel+32, 10), strconv.FormatInt(bmak.te_vel+32, 10),
			strconv.FormatInt(bmak.doe_vel, 10), strconv.FormatInt(bmak.dme_vel, 10), strconv.FormatInt(bmak.pe_vel, 10), strconv.FormatInt(s, 10),
			strconv.FormatInt(a, 10), strconv.FormatInt(bmak.init_time, 10), strconv.FormatInt(bmak.start_ts, 10), strconv.FormatInt(bmak.fpcfTD, 10), strconv.FormatInt(bmak.d2, 10),
			strconv.FormatInt(bmak.ke_cnt, 10), strconv.FormatInt(bmak.me_cnt, 10), strconv.FormatInt(f, 10), strconv.FormatInt(bmak.pe_cnt, 10), strconv.FormatInt(bmak.te_cnt, 10),
			strconv.FormatInt(__, 10), strconv.FormatInt(bmak.ta, 10), strconv.FormatInt(bmak.n_ck, 10), e, strconv.FormatInt(bmak.ab(e), 10), strconv.FormatInt(bmak.rVal, 10),
			strconv.FormatInt(bmak.rCFP, 10), p, l, bmak.floatToString(u0), bmak.floatToString(u1), v, h,
		}

		w := strings.Join(g, ",")
		bmak.Parameter115 = w

		bmak.sensor_data = bmak.ver + "-1,2,-94,-100," + bmak.Parameter100 + "-1,2,-94,-101," + bmak.Parameter101 + "-1,2,-94,-105," + bmak.Parameter105 + "-1,2,-94,-102," + bmak.Parameter102 + "-1,2,-94,-108," + bmak.Parameter108 + "-1,2,-94,-110,"
		bmak.sensor_data = bmak.sensor_data + bmak.Parameter110 + "-1,2,-94,-117," + bmak.Parameter117 + "-1,2,-94,-111," + bmak.Parameter111 + "-1,2,-94,-109," + bmak.Parameter109 + "-1,2,-94,-114," + bmak.Parameter114 + "-1,2,-94,-103,"
		bmak.sensor_data = bmak.sensor_data + "-1,2,-94,-112," + bmak.Parameter112 + "-1,2,-94,-115," + bmak.Parameter115 + "-1,2,-94,-106," + bmak.Parameter106 + "-1,2,-94,-119," + bmak.Parameter119 + "-1,2,-94,-122," + bmak.Parameter122
		bmak.sensor_data = bmak.sensor_data + "-1,2,-94,-123," + bmak.Parameter123 + "-1,2,-94,-124," + bmak.Parameter124 + "-1,2,-94,-126," + bmak.Parameter126 + "-1,2,-94,-127," + bmak.Parameter127
		P := 24 ^ bmak.ab(bmak.sensor_data)
		bmak.Parameter118 = strconv.FormatInt(P, 10)
		bmak.sensor_data = bmak.sensor_data + "-1,2,-94,-70," + bmak.Parameter70 + "-1,2,-94,-80," + bmak.Parameter80 + "-1,2,-94,-116," + bmak.Parameter116 + "-1,2,-94,-118," + bmak.Parameter118 + "-1,2,-94,-129," + bmak.Parameter129 + "-1,2,-94,-121,"
		R := bmak.get_cf_date()
		bmak.sensor_data = bmak.frontofSensor + bmak.sensor_data
		bmak.sensor_data = bmak.sensor_data + ";" + strconv.FormatInt(bmak.GenerateRandomNumber(9, 18), 10) + ";" + strconv.FormatInt(bmak.tst, 10) + ";" + strconv.FormatInt(bmak.get_cf_date()-R, 10) // Assign final timings to the sensor data.

		return bmak.sensor_data, nil

	}

	return "", nil
}
