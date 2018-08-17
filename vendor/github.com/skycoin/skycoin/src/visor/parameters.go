package visor

/*
* CODE GENERATED AUTOMATICALLY WITH FIBER COIN CREATOR
* AVOID EDITING THIS MANUALLY
 */

const (
	// MaxCoinSupply is the maximum supply of coins
	MaxCoinSupply uint64 = 300000000
	// DistributionAddressesTotal is the number of distribution addresses
	DistributionAddressesTotal uint64 = 100
	// DistributionAddressInitialBalance is the initial balance of each distribution address
	DistributionAddressInitialBalance uint64 = MaxCoinSupply / DistributionAddressesTotal
	// InitialUnlockedCount is the initial number of unlocked addresses
	InitialUnlockedCount uint64 = 100
	// UnlockAddressRate is the number of addresses to unlock per unlock time interval
	UnlockAddressRate uint64 = 0
	// UnlockTimeInterval is the distribution address unlock time interval, measured in seconds
	// Once the InitialUnlockedCount is exhausted,
	// UnlockAddressRate addresses will be unlocked per UnlockTimeInterval
	UnlockTimeInterval uint64 = 31536000 // in seconds
	// MaxDropletPrecision represents the decimal precision of droplets
	MaxDropletPrecision uint64 = 3
	//DefaultMaxBlockSize is max block size
	DefaultMaxBlockSize int = 32768 // in bytes
)

var distributionAddresses = [DistributionAddressesTotal]string{
	"2f7gGjLCb3CyV3ogpbxmAtBSUttTGBZckdT",
"2btpmivC5dB67VsUZdC5GqWPQ1dzuuRpb3X",
"ASP2JgBWcNsB5Egts65HbAs9ZnYSunHKzE",
"NfiNNEZLn8SvtHoCdTCcrAnDzNcJQYSFfk",
"2PNQvRbiovsCK9xp1qkykouxMDPGjEVauFS",
"kkTGWV1DjSqsKSdvLUr7UQGKy3BHjVgbBa",
"Pgf8JfXwW5CEmz1k3heV5thmwGBJitJiXe",
"2XAQNci9eLsGZMNs7G5gsGX2kga11pU73LE",
"iSwwzhcRDd5CfB6AcyjH5sMZvoRH721Tcb",
"kQhs9dzFqyd4AHBBShporGY9KcYZcG1qrC",
"Frj82RP6pHfJKC2PS34VrttGpwjpzqZB7L",
"2Z2rFnvBHxRLKHaoYLiSMuj14ZMtE6Cbqy5",
"XUhPVmyogxDSoVKVB2NWuUHQkpEd6h5yvR",
"miTsCdpSEx5YzVLkqrF5Pxc6uZ9XHT1CKh",
"7aBcybKNDK6YD2ZrmMQM6i6UYnGj8TbHBX",
"Hes9VNnq2PWgjkyTSX4mDPiDgn3ECisKyY",
"etnPcrfoYz1kcEJsHuq4aZWijAPi8zJD3f",
"shTjYzrFYPs2Mruf2dZjDiNRj2m7u7kua7",
"2AH1RBfvrB3TxnwJB5whfbvzF8wZbaQdndc",
"wEfQv5vecq8oswYoRYhNUQhNMEBWCEk719",
"2DYB23mQP8sJC4VaKqWKhTYHNw42NzbUEMo",
"GbdSwN2soqgzL5zaWWh5J8dj53ux1YELHV",
"93x27PcUZTsqaFgV9bE8WwCFknbNYQPHKL",
"2PNzJMd1k5NHP1jtPoRoNCsaqgnqRrPoMch",
"DztTRhNnsKtqBJ9Y2aVmW3MiAKHZeTNUuv",
"56WGpo3EV6McA9dtqhfP49M2HpY712nvUA",
"2mouHRxk4S5byNSd5wRPLy56WpSPuMwvGMF",
"2bbtzwrzaesKS4NnWkf9erkQEqFGVm5e1se",
"2QofBCZxHpK9zQGzETmb8FES31VooEv3cAm",
"hyJW9c2DUUFwXcPhMu9LU9vJxtqS9JBJUj",
"2aCZSYELbLLbBFgjTF86owsZuFQbbBAk2dW",
"zRSfnhXw6xWZBmok1M5dVNuM39bLmNbk2z",
"hdU1v2h3J6M7uDDHbyajikPoSLezY9ENAF",
"BPZPZJjAkysrU2xxpMkJpSYMoPPFarbh1b",
"2847b1jw7Cztzr5CXLd3nHQzhgKUfYkKX8L",
"6K2F9Wk1BaKjEadDtFFRxxkBMzgS7M7skW",
"ecRKLXFvMsDkkmRcYLiiJvaNbbTqSS4Kkh",
"AWViDqvD6e8Tk3KG8eCiPpFRvk71sHYZCA",
"23DoWWRw4MhCmvcthcLxUTWFJCDCyBUqEjU",
"R9awCtZcwtQuhXbwBGwj7r3iH8JC9vrwf7",
"2ZvVrw8JxjimFaoss7yA7KUdwyugoyZTPHX",
"2TjrVUzNwoXSZ4LVpHQoN6jbUnZAR3rWp3i",
"39f4qjQtFEBqvbQXLWjtARzX6Rjmxo7Vez",
"26mtewNVE1wXwvgTJyQjEQz9GQH7YyK4TSH",
"yxRy9a9mYUNY4uheh6nXvJub9tZmxL5J89",
"2Bc83jYZ7wadfkFscH4CsMYrmkwjy1JYB3o",
"2PTDowpqRmWrPesZ83PkeHFC7WbmGbpXTgh",
"2j9oHwYYddZpxAXvx5ftmUdcuc1SFZxRzc7",
"gHioiBMiZ4zH98gi1sgMcd9tH8KyPFX4Ug",
"AWChmvwf1XHL75f27x6GYNLx4TRHiYYFH5",
"gQntfDNmNWFWdLTCNq8Y7JGcSxuoVjDmic",
"2U1zn3rDZRAPssd3SEBF4Vfjq8WFgQccWgx",
"2MEKfC9VfaXBjPBjKQhzc5d9pAXsF9aJSaP",
"2ZEkRhvB6mx3DqwieQ3dcTPSLbL5VjdprZz",
"R5eoC8SdzvWgsnkmU1UMgZdHjRs1GbwNV2",
"eNGQxffXwqddPfEWtNwidivcswqXahvxe4",
"XwV4QsUuNm6BrVqs88L58RzjszFN8jqqiQ",
"sr5ETHwvfCjZnQ81WTRe1PhX67QNMpGws1",
"28zmJU4FVhYD8KF8zGsoiWj49BcgYVnnbwk",
"afHCwap5zhrbgbEeBJgjrhK1CgYkTUKkCW",
"daUhNH6aekxQMzUQik1mgQVXm1waZSBxFq",
"W9tgtu9Fh8zLFf9TGFeKBKLhdLLuDR5ckn",
"3mS9kadpuj6kXW4Ty7fH8cDXnnXX7HqgL8",
"NeiD2bwSNVrTHUDaXVsDWo4ikHApVzCxqg",
"uVf6togwTUefF8s5k5RhJmwGdXV9J5eU8g",
"sBPQp96Xp1LKZiu8jvnSnfHP9Ku6YTeHat",
"HhibPgpmmM2nRPW85WDzbSwN2rHHhMYsRv",
"Y79QTve1YvB5ZkkMJL4nj8wfDedaUksMdr",
"PritPwiCtDgmFyQ8kBuQtpPiaMVwZbsWYS",
"2FghaYKjpxbhSS8dEEj2GZog2EQC7DCqf6W",
"uqj1snJvunMzm5f3PoKQPDBnSQjuiGrV9N",
"ine8PW6BdeLkQkcWdAPSie6WBAxKvxRyio",
"Avr7gS9gNyrfNE78Ln85CaZUHcJ3EZFpp7",
"HZZUknJN9EcuD1XS13xJYEEBsz6PoznVLQ",
"2gSgXraKyX8TUPDAwiMm4VYFrkguHobYwJN",
"2TzXWiSRVDzid751ZkPwVwpTpVedb6Sw1zz",
"cpTSQecwfFidbE1ZgxHRGWMV4rVEH6nnsE",
"2gk4P8mN4sbeA8BCcNu6WJePR27tnWZbj2L",
"28shQhYUadUUWhpzfafNgbEJpqQ7U5xi7Xp",
"2Sa4iCzfdzgHnkW3ZGsWNqPDzVfF3FKfPqt",
"22baBznKiyvf8uBSuAbVzxL7N6Cy83CQ5FC",
"ZiNCxaA1AEiirKqJWqcsBx88Y3J5C2HQYX",
"kAnVp2vwHEMk8hma43iUXEWFaqZfWsdLHw",
"cybukMKmbb3fERzTgFN8sM3im9cCCR5DCw",
"cAYiGi8Wtr5MA18spT8xuhcMi64sHUgnzi",
"yQ6in1iHkazpZzwa3XAtQBLixY7RcEnRBt",
"Vi6wo8wUJJ9cSR5oHuP6Le7k88nE98XeeG",
"CSzLBqRTYcfTYJn6rH9P8SbVVFs99sbV1p",
"2ZwUW85emHqY5Co26WhaJkNkaX8Z1ChLA53",
"29KHiRAqmsBMJHrs1tP5eW1J5NVhyyjFJgH",
"dry7SqXZjRnGvLeocZ7wvTuFfZZQkKkJ1B",
"2DGjYmLofrNAqPGk5bVi7Cy8RnvJTsnmxNq",
"TGieXjKXMcG65R9FX5HCjq7ayT8awUJypj",
"2dbBD294rJmkAdvo4R9kzkztUQay34r3Pio",
"ewsUxKfJKs3fvDLb85bWGbqSpAVWD1mKbo",
"phx9vqgXehXo5kZj8GiHv2WS6Eo18JqzCi",
"ztPgcUF42eiU5Fp49mTSmB7GsFnqhkiibg",
"sumL8dhMGwmEETZA971rUrDrsP4HxK49GS",
"N62mT4LYi999ZL5fVvx2ei9Yx5HgN8e2oA",
"2BDAwXRZVMcWEcpPNuA4yvJqWhb8dFkDr9W",}
