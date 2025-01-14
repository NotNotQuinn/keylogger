package keylogger

// keyCodeMap connects the code with human readable key
var keyCodeMap = map[uint16]string{
	1:   "ESC",
	2:   "1",
	3:   "2",
	4:   "3",
	5:   "4",
	6:   "5",
	7:   "6",
	8:   "7",
	9:   "8",
	10:  "9",
	11:  "0",
	12:  "-",
	13:  "=",
	14:  "BS",
	15:  "TAB",
	16:  "q",
	17:  "w",
	18:  "e",
	19:  "r",
	20:  "t",
	21:  "y",
	22:  "u",
	23:  "i",
	24:  "o",
	25:  "p",
	26:  "[",
	27:  "]",
	28:  "ENTER",
	29:  "L_CTRL",
	30:  "a",
	31:  "s",
	32:  "d",
	33:  "f",
	34:  "g",
	35:  "h",
	36:  "j",
	37:  "k",
	38:  "l",
	39:  ";",
	40:  "'",
	41:  "`",
	42:  "L_SHIFT",
	43:  "\\",
	44:  "z",
	45:  "x",
	46:  "c",
	47:  "v",
	48:  "b",
	49:  "n",
	50:  "m",
	51:  ",",
	52:  ".",
	53:  "/",
	54:  "R_SHIFT",
	55:  "*",
	56:  "L_ALT",
	57:  "SPACE",
	58:  "CAPS_LOCK",
	59:  "F1",
	60:  "F2",
	61:  "F3",
	62:  "F4",
	63:  "F5",
	64:  "F6",
	65:  "F7",
	66:  "F8",
	67:  "F9",
	68:  "F10",
	69:  "NUM_LOCK",
	70:  "SCROLL_LOCK",
	71:  "HOME",
	72:  "UP_8",
	73:  "PGUP_9",
	74:  "-",
	75:  "LEFT_4",
	76:  "5",
	77:  "RT_ARROW_6",
	78:  "+",
	79:  "END_1",
	80:  "DOWN",
	81:  "PGDN_3",
	82:  "INS",
	83:  "DEL",
	84:  "",
	85:  "",
	86:  "",
	87:  "F11",
	88:  "F12",
	89:  "",
	90:  "",
	91:  "",
	92:  "",
	93:  "",
	94:  "",
	95:  "",
	96:  "R_ENTER",
	97:  "R_CTRL",
	98:  "/",
	99:  "PRT_SCR",
	100: "R_ALT",
	101: "",
	102: "Home",
	103: "Up",
	104: "PgUp",
	105: "Left",
	106: "Right",
	107: "End",
	108: "Down",
	109: "PgDn",
	110: "Insert",
	111: "Del",
	112: "",
	113: "",
	114: "",
	115: "",
	116: "",
	117: "",
	118: "",
	119: "Pause",
	201: "ESC",
	202: "!",
	203: "@",
	204: "#",
	205: "$",
	206: "%",
	207: "^",
	208: "&",
	209: "*",
	210: "(",
	211: ")",
	212: "_",
	213: "+",
	214: "BS",
	215: "TAB",
	216: "Q",
	217: "W",
	218: "E",
	219: "R",
	220: "T",
	221: "Y",
	222: "U",
	223: "I",
	224: "O",
	225: "P",
	226: "{",
	227: "}",
	228: "ENTER",
	229: "L_CTRL",
	230: "A",
	231: "S",
	232: "D",
	233: "F",
	234: "G",
	235: "H",
	236: "J",
	237: "K",
	238: "L",
	239: ":",
	240: "\"",
	241: "~",
	242: "L_SHIFT",
	243: "|",
	244: "Z",
	245: "X",
	246: "C",
	247: "V",
	248: "B",
	249: "N",
	250: "M",
	251: "<",
	252: ">",
	253: "?",
	254: "R_SHIFT",
	255: "*",
	256: "L_ALT",
	257: "SPACE",
	258: "CAPS_LOCK",
	//@todo below are special keys like numpad that depend on the state of numlock
	259: "F1",
	260: "F2",
	261: "F3",
	262: "F4",
	263: "F5",
	264: "F6",
	265: "F7",
	266: "F8",
	267: "F9",
	268: "F10",
	269: "NUM_LOCK",
	270: "SCROLL_LOCK",
	271: "HOME",
	272: "UP_8",
	273: "PGUP_9",
	274: "-",
	275: "LEFT_4",
	276: "5",
	277: "RT_ARROW_6",
	278: "+",
	279: "END_1",
	280: "DOWN",
	281: "PGDN_3",
	282: "INS",
	283: "DEL",
	284: "",
	285: "",
	286: "",
	287: "F11",
	288: "F12",
	289: "",
	290: "",
	291: "",
	292: "",
	293: "",
	294: "",
	295: "",
	296: "R_ENTER",
	297: "R_CTRL",
	298: "/",
	299: "PRT_SCR",
	300: "R_ALT",
	301: "",
	302: "Home",
	303: "Up",
	304: "PgUp",
	305: "Left",
	306: "Right",
	307: "End",
	308: "Down",
	309: "PgDn",
	310: "Insert",
	311: "Del",
	312: "",
	313: "",
	314: "",
	315: "",
	316: "",
	317: "",
	318: "",
	319: "Pause",
}
