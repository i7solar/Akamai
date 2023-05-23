package main

import (
	"math/rand"
	"strconv"
)

// BuildProcessor Creates a processor array that contains the process's homeURL, akamaiURL, and device UA.
func BuildProcessor(targetSite string, userAgent string) []string {
	switch targetSite {
	case "finishline":
		/* Status: Working */
		return ConstructProcessor("https://www.finishline.com", "https://www.finishline.com/staticweb/962aac4d841ti2110a2a4ca61844715d6", userAgent)
	case "jdsports":
		/* Status: Working */
		return ConstructProcessor("https://www.jdsports.com", "https://www.jdsports.com/staticweb/398497171f6ti267f2eea19daba6cf7d6", userAgent)
	case "zalando":
		/* Status: Working */
		return ConstructProcessor("https://accounts.zalando.com", "https://accounts.zalando.com/Vywp7n/Ba3/LCV/2BB7xw/uOrucpm6Gu/BCxDFUhs/Y0/QSOXgHYw8B", userAgent)
	case "ys":
		/* Status: Working */
		return ConstructProcessor("https://www.yeezysupply.com", "https://www.yeezysupply.com/5vfqfG7UQ5/NT/lgSu_Ezx/rEEzQkhkaN/MngDAQ/flh0/BBd4M20", userAgent)
	case "nike":
		/* Status: Working */
		return ConstructProcessor("https://www.nike.com/de", "https://www.nike.com/yhd7iBGwUuLr3maBnw0iQ5fZ/aO1GrDbtNY/dDswFj8m/EBQb/Wyc2RlQB", userAgent)
	case "offspring":
		/* Currently not supported. TODO: Find a fix for cookie length values. */
		return ConstructProcessor("https://www.offspring.co.uk", "https://www.offspring.co.uk/XBcBbN_6Pemo/Gg/nFlrBgfw1B/EkGiNwrb/XQ5VPXkABA/dUFoVlU/_bQQ", userAgent)
	case "mpv":
		/* Status: Working */
		return ConstructProcessor("https://www.mpv.tickets.com", "https://mpv.tickets.com/1_-iWB/VKAyjN/TQvuA4/4NTQes8U/0/i75pVJGf/EzQmBQE/PBI/VX3o_VlQ", userAgent)
	case "bestbuy_us":
		/* Status: Working */
		return ConstructProcessor("https://www.bestbuy.com", "https://www.bestbuy.com/-cJDnfjvk/Et/H6CugMg/QY1EprDt/RUpYAQ/aUlSW/n0jBDM", userAgent)
	case "asos":
		/* Status: Working */
		return ConstructProcessor("https://www.asos.com", "https://www.asos.com/lZBu/mDb2/5/ZhA/7Ou4gg/7k3SmpGJumOc/CiR9aQ1H/KC/wKFUlFKzc", userAgent)
	case "dsg":
		/* Status: Working */
		return ConstructProcessor("https://www.dickssportinggoods.com", "https://www.dickssportinggoods.com/YNotrs/O/k/tT17U-5zEg/Ez1bJbNE/LRJ3AEFcVA/DR4EAmZb/JzE", userAgent)
	case "qvc":
		/* Status: Working */
		return ConstructProcessor("https://www.qvc.com", "https://www.qvc.com/tkW6_0/taZa/632Le/Vhsx/Tqf-s/zaf9XSfGa5/TiMYOnwD/ViQEb/SQBQlw", userAgent)
	case "kickz":
		/* Status: Working */
		return ConstructProcessor("https://www.kickz.com", "https://www.kickz.com/SzDa/AQDi/WAi8/Xm/uIBg/YS5a0bJQS7/O0lncRp8PRA/e0J/CfFEmZHk", userAgent)
	case "notebooksbilliger":
		/* Status: Working */
		return ConstructProcessor("https://www.notebooksbilliger.de", "https://www.notebooksbilliger.de/5SVTajjSLk6C/ctAy1d/BWnXCC/3OXikQ8XLYr7/M3MLaQE/cFQrf/F5FcCY", userAgent)
	case "evga":
		/* Status: Working */
		return ConstructProcessor("https://www.evga.com", "https://www.evga.com/L7pQWVHi-/jne45/9nAw/7wLEVVb2D7/OWQXGX0/ZgF/rDR4wXnM", userAgent)
	case "converse":
		/* Currently untested. TODO: Test. */
		return ConstructProcessor("https://www.converse.com", "https://www.converse.com/EKmutX/fPtxH5/yUA/5dq5L/14tFP6U/pEuXbVmw/IntzETIaAwk/TFl/HNTJNCx8", userAgent)
	case "aldi":
		/* Status: Working */
		return ConstructProcessor("https://www.aldi.co.uk", "https://www.aldi.co.uk/mJ5lqqzQx/kmFJw4j/tQ/t5f9Xz8kub/OE1GanRlDQE/MWE/tCTcNZic", userAgent)
	case "argos":
		/* Status: Working */
		return ConstructProcessor("https://www.argos.co.uk", "https://www.argos.co.uk/sgDAPh/2soc/F_s/lHB/QG1A5GYY21k/iu7Xkprt/dyIVCmg/TA4LAF9/XTA4", userAgent)
	case "disney":
		/* Status: Working */
		return ConstructProcessor("https://www.shopdisney.com", "https://www.shopdisney.com/fFqIl/c1Wl/W3yd/I1/7gP8S/amuJhkfG/LRJ3AEFcVA/AFJw/JhlnFWA", userAgent)
	case "gamestop":
		/* Status: Working */
		return ConstructProcessor("https://www.gamestop.com", "https://www.gamestop.com/m_g7TZFAJOiVwho3ZwGI-mCR/L35QV63i/S20SQQE/FF4R/BFBUP3U", userAgent)
	case "pacsun":
		/* Status: Working */
		return ConstructProcessor("https://www.pacsun.com", "https://www.pacsun.com/nExe_u/p/Z/8jgPN-NBCQ/L9cOzGzX/NGVhcQ/RVBm/J2wgaSQ", userAgent)
	case "net-a-porter":
		/* Status: Working */
		return ConstructProcessor("https://www.net-a-porter.com", "https://www.net-a-porter.com/xZAgfY9BFoxruJOEJi31IwpdUv0/arumXw3b1GaD/VGB1VA/fS94XXd/kPFA", userAgent)
	case "mrporter":
		/* Status: Working */
		return ConstructProcessor("https://www.mrporter.com", "https://www.mrporter.com/aoBiSwUOs7yup/d6GV2cVa/ix4Bzk/YhVumS9fQE/FEF2AQ/QiQ/1Jj4kFAk", userAgent)
	case "prodirectbasketball":
		/* Status: Working */
		return ConstructProcessor("https://www.prodirectbasketball.com", "https://www.prodirectbasketball.com/lCMbO5yJHk1Mp1mulq0i/5Gc1cQpNam/C0QIeWs8Bw/KDh4G/hlhCzY", userAgent)
	case "prodirectsoccer":
		/* Status: Working */
		return ConstructProcessor("https://www.prodirectsoccer.com", "https://www.prodirectsoccer.com/RAXXvI/pyWvJ/B_Giv/WA/ai9Yt8bw7fYN/dj0xMno/dXFIBw/w0OgUB", userAgent)
	case "fanatics":
		/* Status: Working */
		return ConstructProcessor("https://www.fanatics.com", "https://www.fanatics.com/k1cKF6yqgtpU_e-_FEln/mYLO2tkmOD/Bi1t/BDYdXXJ/efhE", userAgent)
	case "fansedge":
		/* Status: Working */
		return ConstructProcessor("https://www.fansedge.com", "https://www.fansedge.com/2NDXJvDVHMzEP/PpFX/Xy-HMqxcQo/QaYD2QzS1b/QysCRnJjBw/bWwEB0B/TCQo", userAgent)
	case "sportsmemorabilia":
		/* Status: Working */
		return ConstructProcessor("https://www.sportsmemorabilia.com", "https://www.sportsmemorabilia.com/MhTQ07/Vb4Y/K-2y/n9xm/aPJ8xvg7k/YQurLkNkib/U2MDMBYB/H0QqY3A/INFI", userAgent)
	case "target_au":
		/* Status: Working */
		return ConstructProcessor("https://www.target.com.au", "https://www.target.com.au/shOC/NJJW/KCbQ/WoV7/-A/h13YrJXDwa/AB5YVy8/HXE/OBi0GB2g", userAgent)
	case "footasylum_debug":
		/* Status: TODO: Add KACT */
		return ConstructProcessor("https://www.footasylum.com", "https://www.footasylum.com/lrI7iKHV6TFkggpcNNB_/3LGYwbzm/RSlQbENwKAI/dC/A2Eik8JmY", userAgent)
	case "nbastore":
		/* Stauts: Working */
		return ConstructProcessor("https://store.nba.com", "https://store.nba.com/e30DTr/81/Yi/z_iW/9HTpmUtYrRE3k/YmikrtLzS99h/ACR4OywxBA/Vk/J5FgVXYXA", userAgent)
	case "zozo":
		/* Stauts: Working */
		return ConstructProcessor("https://zozo.jp", "https://zozo.jp/mkdbZ8F5ST9q58gMSg/uaE5rchpc7/XmZRLEFQPQE/WTF/0fTYUDlk", userAgent)
	case "ebuyer":
		/* Stauts: Working */
		return ConstructProcessor("https://www.ebuyer.com", "https://www.ebuyer.com/6cc378/vs/q2/p7bt/8b9lqN4F0vuo8/9mm3hz2k3GEh/VElEcw/IU/9XdlwwOkk", userAgent)
	case "allstate":
		/* Stauts: Working */
		return ConstructProcessor("https://myaccountrwd.allstate.com", "https://myaccountrwd.allstate.com/DgrWrJbcwX-cE/cWs2nKYA/1YdYFo/Viua2rrhL5/b0Rn/dUw_b/XcyZ04", userAgent)
	case "stoneisland":
		/* Stauts: Working */
		return ConstructProcessor("https://www.stoneisland.com", "https://www.stoneisland.com/rmhLVH9lQ/NadirP3yd/V/1F-oxmls/imi5SbOpG7/NSMwXA/Aw93X0oD/CkQ", userAgent)
	case "bigw":
		/* Stauts: Working */
		return ConstructProcessor("https://www.bigw.com.au", "https://www.bigw.com.au/YcW5DUkqqd5Ef/wKb11AnImOJzL/Y/5OEhQ4Nb/eQUCVyttMg/IlY0cio/ySQUB", userAgent)
	case "sizeer":
		/* Stauts: Working */
		return ConstructProcessor("https://sklep.sizeer.com", "https://sklep.sizeer.com/SmAp/69k9/TiC/u2e/xYKw/1ic1ftpz/TXUGYQ/bxB7E/FZJcUM", userAgent)
	case "sizeer_de":
		return ConstructProcessor("https://sizeer.de", "https://sizeer.de/ZHqk1UH88GP8A4G-o6O3/9V3fNStQ/OUxrAQ/JzF/FQiJKeVEB", userAgent)
	case "sizeer_ro":
		/* Stauts: Working */
		return ConstructProcessor("https://sizeer.ro", "https://sizeer.ro/shzx6w01bTQGT/JH-D/-6pVCtz0_M/1kiVJcSkL5/YxU3A31SAw/eF9VM/UMKMT8", userAgent)
	case "bestbuy_ca":
		/* Stauts: Working */
		return ConstructProcessor("https://www.bestbuy.ca", "https://www.bestbuy.ca/vv6AFS7TFNvw/YTcIHA3zx2/fJ/iXEQSrp2uE/YG9R/cUdsN2/IyD1A", userAgent)
	case "adidas_au":
		return ConstructProcessor("https://www.adidas.com.au", "https://www.adidas.com.au/IEnrK7LPSzz8GN_Y7xsl/k1afJVkk/bXBh/ITUhRCw/TZ2g", userAgent)
	case "zalando_ch":
		return ConstructProcessor("https://www.zalando.ch", "https://www.zalando.ch/m-iNF6ykeCqWfcJaAeNMIPPoK6k/OLQa4pXV/TDx9OQ/L1ova2I/KIl0", userAgent)
	case "luisaviaroma":
		return ConstructProcessor("https://www.luisaviaroma.com", "https://www.luisaviaroma.com/rcRy75/Q4LD/Vu_/hLj/o6TTkW-CDvA/YYk7zGGQmub9/VElEcw/TFsAcR/1nJjE", userAgent)
	case "homedepot":
		return ConstructProcessor("https://www.homedepot.com", "https://www.homedepot.com/SF_wF_AiadQqc/t_i9/SUBjojZJJE/caaacmt4V3/HBEQC1o/PWoyJ/DVYZwk", userAgent)
	case "panera":
		return ConstructProcessor("https://www.panerabread.com", "https://www.panerabread.com/ng_Y_0HpQTjz/W3SPJ1vPsV/D_/OtkarGzi/An5oVQE/DUpI/B3p1AQkC", userAgent)
	case "woolworthsrewards":
		return ConstructProcessor("https://www.woolworthsrewards.com.au", "https://www.woolworthsrewards.com.au/TI7n3cukh/TD5TS/w-AQ/1Y7Grc4r5z/ejN9VGIBBQ/Kng/SB3sDLXEB", userAgent)
	case "kohls":
		return ConstructProcessor("https://www.kohls.com", "https://www.kohls.com/TxVrCOGXm_LQw_G9UDcwkZtUkrQ/X33zpwbG/OTJhcFIKBg/PSl5/RG4QOAkC", userAgent)
	case "macys":
		return ConstructProcessor("https://www.macys.com", "https://www.macys.com/82XOFWM2QOPis2RUWBMJP5mL/Y5afLkDh/TgxyDCYfAw/Wk5zDj/BYVUIB", userAgent)
	case "cabelas":
		return ConstructProcessor("https://www.cabelas.com/shop/en", "https://www.cabelas.com/xxLoqv1z_o/JsIEUxfg/0f/auuVk2EpEuiu/TiMYOnwD/KAI-Fn/Y4KWoB", userAgent)
	case "scottycameron":
		return ConstructProcessor("https://www.scottycameron.com", "https://www.scottycameron.com/k1-oRj/uECV-v/VlQGVI/_RiY/Ezixw/aLG5ccfVm9/aRc-Ag/Slg3S/kNuJBU", userAgent)
	case "sephora":
		return ConstructProcessor("https://www.sephora.com", "https://www.sephora.com/D9MCqv/G/G/B05pOkh29yQn/E9ciQpui/LFtobS0ANwE/VTZ4K/j1tcBc", userAgent)
	case "autozone":
		return ConstructProcessor("https://www.autozone.com", "https://www.autozone.com/RdfGbuBQQfk4/OBARSNpNDq/Mq/i3kYNww9/QBdNBBIUJAo/bnxF/Yx5CcyA", userAgent)
	case "swatch":
		return ConstructProcessor("https://www.swatch.com", "https://www.swatch.com/JO2GmN/kXMW_u/rwC9F/26wlU/cn/uaSabDrJ7Qw1/QWkwVg/BTdX/FWUsPGA", userAgent)
	case "ae":
		return ConstructProcessor("https://www.ae.com", "https://www.ae.com/Ep5OCDYLR/4gj/fpt/iSxO_-e8tGb0/h5OXG42hEO/ejN9VGIBBQ/S1U-G/ixVdAQB", userAgent)
	}

	return nil
}

// ConstructProcessor Constructs a processor array.
func ConstructProcessor(requestURL string, akamaiURL string, userAgent string) []string {

	var ProcessorArray []string

	ProcessorArray = append(ProcessorArray, requestURL)
	ProcessorArray = append(ProcessorArray, akamaiURL)
	ProcessorArray = append(ProcessorArray, userAgent)

	return ProcessorArray
}

// ByteArrayToString Converts the byteArray given to a string array.
func ByteArrayToString(arr []byte) []string {
	result := make([]string, len(arr))
	intArray := ByteArrayToInt(arr)
	for i := 0; i < len(result); i++ {
		result[i] = strconv.Itoa(intArray[i])
	}
	return result
}

// ByteArrayToInt Converts the byteArray given to a int array.
func ByteArrayToInt(arr []byte) []int {
	result := make([]int, len(arr))
	for i := 0; i < len(result); i++ {
		result[i] = int(arr[i])
	}
	return result
}

/* Creates random numbers. */
func randNumGen(mod int) int {
	b := make([]byte, 1)
	// safe to ignore error here as rand.Read documentation states the returned error is always nil
	_, _ = rand.Read(b)
	return int(b[0]) % mod
}
