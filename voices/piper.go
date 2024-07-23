package voices

// original
// https://huggingface.co/rhasspy/piper-voices/resolve/v1.0.0/ar/ar_JO/kareem/low/ar_JO-kareem-low.onnx

var PiperVoices = map[string]map[string]map[string]string{
	"Arabic": {
		"kareem": {
			"low":    "ar/ar_JO/kareem/low/ar_JO-kareem-low",
			"medium": "ar/ar_JO/kareem/medium/ar_JO-kareem-medium",
		},
	},
	"Catalan": {
		"upc_ona": {
			"x_low":  "ca/ca_ES/upc_ona/x_low/ca_ES-upc_ona-x_low",
			"medium": "ca/ca_ES/upc_ona/medium/ca_ES-upc_ona-medium",
		},
		"upc_pau": {
			"x_low": "ca/ca_ES/upc_pau/x_low/ca_ES-upc_pau-x_low",
		},
	},
	"Czech": {
		"jirka": {
			"low":    "cs/cs_CZ/jirka/low/cs_CZ-jirka-low",
			"medium": "cs/cs_CZ/jirka/medium/cs_CZ-jirka-medium",
		},
	},
	"Welsh": {
		"gwryw_gogleddol": {
			"medium": "cy/cy_GB/gwryw_gogleddol/medium/cy_GB-gwryw_gogleddol-medium",
		},
	},
	"Danish": {
		"talesyntese": {
			"medium": "da/da_DK/talesyntese/medium/da_DK-talesyntese-medium",
		},
	},
	"German": {
		"eva_k": {
			"x_low": "de/de_DE/eva_k/x_low/de_DE-eva_k-x_low",
		},
		"karlsson": {
			"low": "de/de_DE/karlsson/low/de_DE-karlsson-low",
		},
		"kerstin": {
			"low": "de/de_DE/kerstin/low/de_DE-kerstin-low",
		},
		"mls": {
			"medium": "de/de_DE/mls/medium/de_DE-mls-medium",
		},
		"pavoque": {
			"low": "de/de_DE/pavoque/low/de_DE-pavoque-low",
		},
		"ramona": {
			"low": "de/de_DE/ramona/low/de_DE-ramona-low",
		},
		"thorsten": {
			"low":    "de/de_DE/thorsten/low/de_DE-thorsten-low",
			"medium": "de/de_DE/thorsten/medium/de_DE-thorsten-medium",
			"high":   "de/de_DE/thorsten/high/de_DE-thorsten-high",
		},
		"thorsten_emotional": {
			"medium": "de/de_DE/thorsten_emotional/medium/de_DE-thorsten_emotional-medium",
		},
	},
	"Greek": {
		"rapunzelina": {
			"low": "el/el_GR/rapunzelina/low/el_GR-rapunzelina-low",
		},
	},
	"EnglishGB": {
		"alan": {
			"low":    "en/en_GB/alan/low/en_GB-alan-low",
			"medium": "en/en_GB/alan/medium/en_GB-alan-medium",
		},
		"alba": {
			"medium": "en/en_GB/alba/medium/en_GB-alba-medium",
		},
		"aru": {
			"medium": "en/en_GB/aru/medium/en_GB-aru-medium",
		},
		"cori": {
			"medium": "en/en_GB/cori/medium/en_GB-cori-medium",
			"high":   "en/en_GB/cori/high/en_GB-cori-high",
		},
		"jenny_dioco": {
			"medium": "en/en_GB/jenny_dioco/medium/en_GB-jenny_dioco-medium",
		},
		"northern_english_male": {
			"medium": "en/en_GB/northern_english_male/medium/en_GB-northern_english_male-medium",
		},
		"semaine": {
			"medium": "en/en_GB/semaine/medium/en_GB-semaine-medium",
		},
		"southern_english_female": {
			"low": "en/en_GB/southern_english_female/low/en_GB-southern_english_female-low",
		},
		"vctk": {
			"medium": "en/en_GB/vctk/medium/en_GB-vctk-medium",
		},
	},
	"EnglishUS": {
		"amy": {
			"low":    "en/en_US/amy/low/en_US-amy-low",
			"medium": "en/en_US/amy/medium/en_US-amy-medium",
		},
		"arctic": {
			"medium": "en/en_US/arctic/medium/en_US-arctic-medium",
		},
		"bryce": {
			"medium": "en/en_US/bryce/medium/en_US-bryce-medium",
		},
		"danny": {
			"low": "en/en_US/danny/low/en_US-danny-low",
		},
		"hfc_female": {
			"medium": "en/en_US/hfc_female/medium/en_US-hfc_female-medium",
		},
		"hfc_male": {
			"medium": "en/en_US/hfc_male/medium/en_US-hfc_male-medium",
		},
		"joe": {
			"medium": "en/en_US/joe/medium/en_US-joe-medium",
		},
		"john": {
			"medium": "en/en_US/john/medium/en_US-john-medium",
		},
		"kathleen": {
			"low": "en/en_US/kathleen/low/en_US-kathleen-low",
		},
		"kristin": {
			"medium": "en/en_US/kristin/medium/en_US-kristin-medium",
		},
		"kusal": {
			"medium": "en/en_US/kusal/medium/en_US-kusal-medium",
		},
		"l2arctic": {
			"medium": "en/en_US/l2arctic/medium/en_US-l2arctic-medium",
		},
		"lessac": {
			"low":    "en/en_US/lessac/low/en_US-lessac-low",
			"medium": "en/en_US/lessac/medium/en_US-lessac-medium",
			"high":   "en/en_US/lessac/high/en_US-lessac-high",
		},
		"libritts": {
			"high": "en/en_US/libritts/high/en_US-libritts-high",
		},
		"libritts_r": {
			"medium": "en/en_US/libritts_r/medium/en_US-libritts_r-medium",
		},
		"ljspeech": {
			"medium": "en/en_US/ljspeech/medium/en_US-ljspeech-medium",
			"high":   "en/en_US/ljspeech/high/en_US-ljspeech-high",
		},
		"norman": {
			"medium": "en/en_US/norman/medium/en_US-norman-medium",
		},
		"ryan": {
			"low":    "en/en_US/ryan/low/en_US-ryan-low",
			"medium": "en/en_US/ryan/medium/en_US-ryan-medium",
			"high":   "en/en_US/ryan/high/en_US-ryan-high",
		},
	},
	"SpanishES": {
		"carlfm": {
			"x_low": "es/es_ES/carlfm/x_low/es_ES-carlfm-x_low",
		},
		"davefx": {
			"medium": "es/es_ES/davefx/medium/es_ES-davefx-medium",
		},
		"mls_10246": {
			"low": "es/es_ES/mls_10246/low/es_ES-mls_10246-low",
		},
		"mls_9972": {
			"low": "es/es_ES/mls_9972/low/es_ES-mls_9972-low",
		},
		"sharvard": {
			"medium": "es/es_ES/sharvard/medium/es_ES-sharvard-medium",
		},
	},
	"SpanishMX": {
		"ald": {
			"medium": "es/es_MX/ald/medium/es_MX-ald-medium",
		},
		"claude": {
			"high": "es/es_MX/claude/high/es_MX-claude-high",
		},
	},
	"Farsi": {
		"amir": {
			"medium": "fa/fa_IR/amir/medium/fa_IR-amir-medium",
		},
		"gyro": {
			"medium": "fa/fa_IR/gyro/medium/fa_IR-gyro-medium",
		},
	},
	"Finnish": {
		"harri": {
			"low":    "fi/fi_FI/harri/low/fi_FI-harri-low",
			"medium": "fi/fi_FI/harri/medium/fi_FI-harri-medium",
		},
	},
	"French": {
		"gilles": {
			"low": "fr/fr_FR/gilles/low/fr_FR-gilles-low",
		},
		"mls": {
			"medium": "fr/fr_FR/mls/medium/fr_FR-mls-medium",
		},
		"mls_1840": {
			"low": "fr/fr_FR/mls_1840/low/fr_FR-mls_1840-low",
		},
		"siwis": {
			"low":    "fr/fr_FR/siwis/low/fr_FR-siwis-low",
			"medium": "fr/fr_FR/siwis/medium/fr_FR-siwis-medium",
		},
		"tom": {
			"medium": "fr/fr_FR/tom/medium/fr_FR-tom-medium",
		},
		"upmc": {
			"medium": "fr/fr_FR/upmc/medium/fr_FR-upmc-medium",
		},
	},
	"Hungarian": {
		"anna": {
			"medium": "hu/hu_HU/anna/medium/hu_HU-anna-medium",
		},
		"berta": {
			"medium": "hu/hu_HU/berta/medium/hu_HU-berta-medium",
		},
		"imre": {
			"medium": "hu/hu_HU/imre/medium/hu_HU-imre-medium",
		},
	},
	"Icelandic": {
		"bui": {
			"medium": "is/is_IS/bui/medium/is_IS-bui-medium",
		},
		"salka": {
			"medium": "is/is_IS/salka/medium/is_IS-salka-medium",
		},
		"steinn": {
			"medium": "is/is_IS/steinn/medium/is_IS-steinn-medium",
		},
		"ugla": {
			"medium": "is/is_IS/ugla/medium/is_IS-ugla-medium",
		},
	},
	"Italian": {
		"paola": {
			"medium": "it/it_IT/paola/medium/it_IT-paola-medium",
		},
		"riccardo": {
			"x_low": "it/it_IT/riccardo/x_low/it_IT-riccardo-x_low",
		},
	},
	"Georgian": {
		"natia": {
			"medium": "ka/ka_GE/natia/medium/ka_GE-natia-medium",
		},
	},
	"Kazakh": {
		"iseke": {
			"x_low": "kk/kk_KZ/iseke/x_low/kk_KZ-iseke-x_low",
		},
		"issai": {
			"high": "kk/kk_KZ/issai/high/kk_KZ-issai-high",
		},
		"raya": {
			"x_low": "kk/kk_KZ/raya/x_low/kk_KZ-raya-x_low",
		},
	},
	"Luxembourgish": {
		"marylux": {
			"medium": "lb/lb_LU/marylux/medium/lb_LU-marylux-medium",
		},
	},
	"Nepali": {
		"google": {
			"x_low":  "ne/ne_NP/google/x_low/ne_NP-google-x_low",
			"medium": "ne/ne_NP/google/medium/ne_NP-google-medium",
		},
	},
	"DutchBE": {
		"nathalie": {
			"x_low":  "nl/nl_BE/nathalie/x_low/nl_BE-nathalie-x_low",
			"medium": "nl/nl_BE/nathalie/medium/nl_BE-nathalie-medium",
		},
		"rdh": {
			"x_low":  "nl/nl_BE/rdh/x_low/nl_BE-rdh-x_low",
			"medium": "nl/nl_BE/rdh/medium/nl_BE-rdh-medium",
		},
	},
	"DutchNL": {
		"mls": {
			"medium": "nl/nl_NL/mls/medium/nl_NL-mls-medium",
		},
		"mls_5809": {
			"low": "nl/nl_NL/mls_5809/low/nl_NL-mls_5809-low",
		},
		"mls_7432": {
			"low": "nl/nl_NL/mls_7432/low/nl_NL-mls_7432-low",
		},
	},
	"Norwegian": {
		"talesyntese": {
			"medium": "no/no_NO/talesyntese/medium/no_NO-talesyntese-medium",
		},
	},
	"Polish": {
		"darkman": {
			"medium": "pl/pl_PL/darkman/medium/pl_PL-darkman-medium",
		},
		"gosia": {
			"medium": "pl/pl_PL/gosia/medium/pl_PL-gosia-medium",
		},
		"mc_speech": {
			"medium": "pl/pl_PL/mc_speech/medium/pl_PL-mc_speech-medium",
		},
		"mls_6892": {
			"low": "pl/pl_PL/mls_6892/low/pl_PL-mls_6892-low",
		},
	},
	"PortugueseBR": {
		"edresson": {
			"low": "pt/pt_BR/edresson/low/pt_BR-edresson-low",
		},
		"faber": {
			"medium": "pt/pt_BR/faber/medium/pt_BR-faber-medium",
		},
	},
	"PortuguesePT": {
		"tugão": {
			"medium": "pt/pt_PT/tugão/medium/pt_PT-tugão-medium",
		},
	},
	"Romanian": {
		"mihai": {
			"medium": "ro/ro_RO/mihai/medium/ro_RO-mihai-medium",
		},
	},
	"Russian": {
		"denis": {
			"medium": "ru/ru_RU/denis/medium/ru_RU-denis-medium",
		},
		"dmitri": {
			"medium": "ru/ru_RU/dmitri/medium/ru_RU-dmitri-medium",
		},
		"irina": {
			"medium": "ru/ru_RU/irina/medium/ru_RU-irina-medium",
		},
		"ruslan": {
			"medium": "ru/ru_RU/ruslan/medium/ru_RU-ruslan-medium",
		},
	},
	"Slovak": {
		"lili": {
			"medium": "sk/sk_SK/lili/medium/sk_SK-lili-medium",
		},
	},
	"Slovenian": {
		"artur": {
			"medium": "sl/sl_SI/artur/medium/sl_SI-artur-medium",
		},
	},
	"Serbian": {
		"serbski_institut": {
			"medium": "sr/sr_RS/serbski_institut/medium/sr_RS-serbski_institut-medium",
		},
	},
	"Swedish": {
		"nst": {
			"medium": "sv/sv_SE/nst/medium/sv_SE-nst-medium",
		},
	},
	"Swahili": {
		"lanfrica": {
			"medium": "sw/sw_CD/lanfrica/medium/sw_CD-lanfrica-medium",
		},
	},
	"Turkish": {
		"dfki": {
			"medium": "tr/tr_TR/dfki/medium/tr_TR-dfki-medium",
		},
		"fahrettin": {
			"medium": "tr/tr_TR/fahrettin/medium/tr_TR-fahrettin-medium",
		},
		"fettah": {
			"medium": "tr/tr_TR/fettah/medium/tr_TR-fettah-medium",
		},
	},
	"Ukrainian": {
		"lada": {
			"x_low": "uk/uk_UA/lada/x_low/uk_UA-lada-x_low",
		},
		"ukrainian_tts": {
			"medium": "uk/uk_UA/ukrainian_tts/medium/uk_UA-ukrainian_tts-medium",
		},
	},
	"Vietnamese": {
		"25hours_single": {
			"low": "vi/vi_VN/25hours_single/low/vi_VN-25hours_single-low",
		},
		"vais1000": {
			"medium": "vi/vi_VN/vais1000/medium/vi_VN-vais1000-medium",
		},
		"vivos": {
			"x_low": "vi/vi_VN/vivos/x_low/vi_VN-vivos-x_low",
		},
	},
	"Chinese": {
		"huayan": {
			"x_low":  "zh/zh_CN/huayan/x_low/zh_CN-huayan-x_low",
			"medium": "zh/zh_CN/huayan/medium/zh_CN-huayan-medium",
		},
	},
}
