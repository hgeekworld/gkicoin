package visor

/*
* CODE GENERATED AUTOMATICALLY WITH FIBER COIN CREATOR
* AVOID EDITING THIS MANUALLY
*/

const (
	// MaxCoinSupply is the maximum supply of coins
	MaxCoinSupply uint64 = 100000000
	// DistributionAddressesTotal is the number of distribution addresses
	DistributionAddressesTotal uint64 = 100
	// DistributionAddressInitialBalance is the initial balance of each distribution address
	DistributionAddressInitialBalance uint64 = MaxCoinSupply / DistributionAddressesTotal
	// InitialUnlockedCount is the initial number of unlocked addresses
	InitialUnlockedCount uint64 = 100
	// UnlockAddressRate is the number of addresses to unlock per unlock time interval
	UnlockAddressRate uint64 = 100
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
   "23U9wuB2AVE81P4BKGv9DND4acM6znWLHaY",
   "2XBPKDvv8QFscJRVKJcBtnKMyY4cUhANp4K",
   "WWs81kQmLqeQSv1Znis6j4M7RnzESs6prw",
   "21HWGKPTC46GV7YAuwo6khSs72RGzLV9SS2",
   "2kjjTV1A7ku3kyBEG9KimTaK15kbfRYn7oL",
   "2CCgXAfhSPPnW3FCeY3wtrJBPzGWMAEEdAJ",
   "2HonbQmB7JEnkx9ZeBg7J4fqXfupfPfhRua",
   "8sTH1gjuAHt8Ge6fCsXwKYHTNCZsqYVUm1",
   "21xdgCNVZtF1tCrUiNR7AMUAJHj2waWud7s",
   "49nirfwjpJrBcd73GJYXH7MwhRhw9WQRtw",
   "AT6CYmCP2FsWLGYGQGJb911FptEoPrVB2f",
   "db9Dn5nRBSYSnGLyEgbVDaQjDif6aiFBQD",
   "nTGFZfHvpXGqENBPbN9Mq548z8PJne1re9",
   "26vZLTpz5PBV1QLJ8gxMsSMe7KpTeeue57c",
   "nME6R4KaK6DA5CEogqmcEho6F7TdCaCyL",
   "QvS6BCFbVSAqMfqJf8hbV15Lgh83aGLKjr",
   "qbouktr49iYiHvKouardFfWnWvbWhJZaLr",
   "2VtCsuDTjTR1uhb5JLG2UviDi45rfJbUwZH",
   "2U7Fcy35TjecpggRzk3Y8XdUdobMnarSHCb",
   "2G8xzxSAdvMaEtWZ9buhcpt63GYAGW7oxpd",
   "iyjuBnWJwWd8p3ng31GNzcPwHqj36d8DpN",
   "28TDHYP39CmwFmCXmZhj6vgzvesdNqnnah6",
   "zSGhhXbPRqmKv6MwWUXoSVNTjxeii4A4vq",
   "2PFphMKfrm6GWJSBoSidVvrrTLc8548HhdS",
   "TPwMnMhnVWdsjpHy76GSPrH3t4TBYEHXva",
   "24ZJuPeRdXL7kwmyJZ6E35FJXueRxaNhvue",
   "2ZFyoUb4QLEqP9YJztQBBKQgfKLingrGMh4",
   "2MsSvKQ4S7iuFpYyKRWDBPyeAk6SZSXJDAa",
   "pdTb94AVq9bMSGZQ5v8TRsUk5MF1ooXWE2",
   "2f8dM7Ek8jDJky2mYXBy9Htam98qLsKBnnX",
   "2g7xuLoiaDG44nx9BYtkVfFgg9YLpwSL1vt",
   "8HQ8rCLbBxJkC6grJZFWetcfquMMZcmLHJ",
   "23BnxMSis1Ji34YRNdhWBYTgqU1Wo24gwE2",
   "25yBoo7TsY64yEHvLxQwRYW4qJTAXueXhYV",
   "m9SbsjhQnBwUtSsiBMoLSfUgufMhp3mvdg",
   "2aV7nCB9zUtPCnEt9bVEkHRPEwNZQJ94n7M",
   "2UJibUnV59g1HLFAaSWCLQfjeXoMtBMEpjq",
   "2bNvNQRjXJ2vgZiTTa1186ki9whbnt35NAU",
   "v2P1NzqaJYScDWkgRjSZuRdiKUwpu5pqrR",
   "JK2n1zhgiwjKJvu7o5bmYazWyRZgg1Vgvk",
   "yCxqhvfks9QQqyfzyC1jHVULZpayD5sxwP",
   "PMj3urXxnewRgxBySqxr6S2kvkJ7s7x5ik",
   "BQi9BADefnafRFeijxdWJaf5KwAkrQYKKR",
   "qBMXmL2hRQ2it4jLpWZd8JXUUXic9kCEGR",
   "2fWLKmk92DZjU1SvLKLXrawdDeFy8n57N7a",
   "vwgwk8jsfBRSS1RA4pK5TqgvUXYomX6bGH",
   "2KB4iNDTZJJ5ewXnUu9ZHT3nrCBp9gb2qHn",
   "2fGdgJtpxSnRXWuWA4MQqGYoD1pvwdp5tVz",
   "2B6MLzD3Pe33UqoQq6uuhXH6z9GBZa5Cgu1",
   "q7xKfJiJ1swcYNU1jUP5BfjNASKkfhEcM9",
   "6sj5KC6APyBxQuYNT8tqWUZeoUEo3sv66u",
   "uDBPbrTeQsnryHECW3jbJ4DPEsdiN3NT9o",
   "2HGdHXXKL11n5G9QVp5taBa25Y2MuEt4tLb",
   "Cb6wKn7Pj973S2vrF2sN7zfksFXCqP5Jno",
   "2LagJVfX18XC9HGCR1TGhXqwa9jTa57RDsh",
   "Yc7B4Fg6w3ppxfARiJ2WDV1n8pAk21cN9F",
   "26FDr5aVJdZefgnu7A6sBfA81JGp5kQT92Q",
   "2H5x1QjinowWJXMYokL1RGePKrbAzNiEwDr",
   "WL65ki7TPSiDmv28wmLn6i5Fc7U2eWEugF",
   "AiGDMe98ZY3FuYftQRDFR7BhMvjpPAZzPw",
   "RjiyVYkqHfUQcH7oPancb1nJsQbptAazMB",
   "2UhXJjUyHEByiaPbsnye6AK9u7T7qVLZM7S",
   "NVFiYyMbhnqivUZrTGeL2isEzYwrbh4zR2",
   "2EfHcUJANWfGUptxggfMWPYMUVTSRAaLncP",
   "jWi5in1By8AmFFPUbwJMFLrEA7sb8SFDfc",
   "2hSEwXcu1oEEZsyrRu7M7YVFe8pDETELwvY",
   "73e5B2M9f1rWpbvne838SpS6KPvr9aqSwk",
   "bRYynNzYz3UXLrQZ5SBQyxzPdJQwpqM24g",
   "2SPm8nQKGGhX8ynjbSnxY7THvsbA4ykP65N",
   "24GNjfVkBrvFk28uQVgmt63u39FfpP4RzPq",
   "2iDC1EZHRguTYZwpXGweVXorFYdcjQrPhAE",
   "2awzdx1nke869eXstSKphWqQqAHGXQnJxHn",
   "7SeqmjtFBj3RnP14LEcX6HbTfnsT1cJzdg",
   "LxzbweGomt8kgNayrNUpZi2xoNq4xQAHb9",
   "2EzRF8h6atMAcS2yS4SdVhrg4zMTaUM9WwZ",
   "2DM9sDVtjpSGqZLvgoiopQHTjDAvaEJvzmk",
   "aE2TryCLuqwX6XhLEjqpPaKiRgFyfbXAEM",
   "2RZG5XZcFqheURttevk6erGKcT27AcmZJ9R",
   "2VXzP1c4pYaAmNzZnuFz5KkLtHLAJypRrRs",
   "xHZUHWheTMSRYj3Pajy7kbXtbGehcMpfju",
   "2duNzWBaHYLUBt4PKHCmrqA7Khj674o93ip",
   "2fEuwXoAYPMBPBdtzy8W37i5VXttPkR4BVy",
   "21pwp4bZ1gwr84cT9cZgwMCGpVP2rWYVyx2",
   "2hVHN3rEp9gccg3DYbYZHaELmMbxdytHrg6",
   "z6WA7ASENaJUdctgy46NTfvdr5aFEFpBb",
   "2UGkMmRfeZRqxSnkCftfcE1wMDkCPAZWmU2",
   "2kZsAEPVwgTVe7vDu2vF1uuWXjdnUUinn8W",
   "KaX46XBw1CvkGfSPyqcKbWtZKfxR7zimkz",
   "f2AF1mFzsWUmq17KEsnMe6joVhHSxZa5HM",
   "2EGoPoYyfPD97468w4sgvbqHZZLi3T4wCz7",
   "2467dYfoTrVyG5TXuqXxmHJnpmHvNv58ZqC",
   "2MtGw7riq5FUrYF816xkve4jaxTMFeFWzhK",
   "bSq2SpQCAZuVWRTT6Xs2XtkJwCK1QxXWf2",
   "2QaoTRFRAQY6czPeASBWSLESpD7XAPyu4qk",
   "Uykyw6C9damDLFwC5VLGiUZRkLvr8QSbYA",
   "24KzvmAoYKN9S4Hk3Qx3YjQapXDVGJxfpod",
   "s9aDwV9iHpkWmBB1TdjhAUYxmzP9iyRFni",
   "2LZqr6LdkCogMxrzhAEZnhiEXDWYe7j9uNv",
   "8hMd4wvafxiYMDA8rW6EBRZX9jL9Ujq3fe",
   "2CrWyCHREcdyRKJdCzaboZX9rKnH48RAM39",
}
