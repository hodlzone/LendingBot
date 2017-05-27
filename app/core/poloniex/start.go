package poloniex

import ()

// StartPoloniex returns a Poloniex object to begin making API calls
func StartPoloniex() *Poloniex {
	// Poloniex client
	var exch Exchanges
	exch.Enabled = true
	exch.Verbose = false
	exch.Websocket = false
	exch.RESTPollingDelay = 10
	exch.AvailablePairs = "BTC_XUSD,BTC_FCT,BTC_MMNXT,BTC_NMC,BTC_BITUSD,BTC_RDD,BTC_XMR,BTC_XST,BTC_DSH,BTC_MAID,BTC_DGB,BTC_NEOS,BTC_BLK,BTC_NAUT,BTC_NBT,BTC_XCP,BTC_STR,BTC_BTCD,BTC_GRC,BTC_HUC,BTC_BBR,BTC_XDN,BTC_INDEX,BTC_IOC,BTC_SWARM,BTC_EMC2,BTC_MCN,BTC_NOXT,BTC_MINT,BTC_PTS,BTC_SC,BTC_GEO,BTC_XRP,BTC_FLO,BTC_BITS,BTC_HYP,BTC_XCR,BTC_LTBC,BTC_SYS,BTC_GMC,BTC_ETH,BTC_SYNC,BTC_GAP,BTC_BCN,BTC_C2,BTC_PINK,BTC_FIBRE,BTC_POT,BTC_QTL,BTC_SDC,BTC_XC,BTC_DASH,BTC_SILK,BTC_CLAM,BTC_NAV,BTC_PIGGY,BTC_BCY,BTC_MIL,BTC_XCN,BTC_YACC,BTC_BTS,BTC_QBK,BTC_SJCX,BTC_LQD,BTC_BURST,BTC_RIC,BTC_VRC,BTC_LTC,BTC_XPB,BTC_GRS,BTC_XCH,BTC_ARCH,BTC_QORA,BTC_HZ,BTC_NSR,BTC_XPM,BTC_BITCNY,BTC_EXE,BTC_XMG,BTC_BTC,BTC_BTM,BTC_NOBL,BTC_NXT,BTC_DOGE,BTC_CURE,BTC_MNTA,BTC_ADN,BTC_EXP,BTC_VTC,BTC_FLDC,BTC_MRS,BTC_MYR,BTC_OMNI,BTC_VNL,BTC_USDT,BTC_NOTE,BTC_WDC,BTC_BELA,BTC_VIA,BTC_CGA,BTC_DIEM,BTC_IFC,BTC_XDP,BTC_BLOCK,BTC_MMC,BTC_1CR,BTC_UNITY,BTC_XBC,BTC_GEMZ,BTC_FLT,BTC_PPC,BTC_XEM,BTC_RBY,BTC_CNMT,BTC_ABY,XMR_XDN,XMR_IFC,XMR_DIEM,XMR_BBR,XMR_DSH,XMR_BCN,XMR_LTC,XMR_MAID,XMR_DASH,XMR_BTCD,XMR_HYP,XMR_BLK,XMR_QORA,XMR_MNTA,XMR_NXT,USDT_BTC,USDT_ETH,USDT_XRP,USDT_DASH,USDT_LTC,USDT_NXT,USDT_XMR,USDT_STR"
	exch.EnabledPairs = "BTC_LTC,BTC_ETH,BTC_DOGE,BTC_DASH,BTC_XRP,BTC_FCT"
	exch.BaseCurrencies = "USD"

	polo := new(Poloniex)
	polo.Setup(exch)

	return polo
}
