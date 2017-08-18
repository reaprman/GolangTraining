package main

import "fmt"

func main() {
	var x [256]string

	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 0; i < 256; i++ {
		x[i] = string(i)
	}

	for _, v := range x {
		fmt.Printf("%v - %T - %v\n", v, v, []byte(v))
	}
}

/*
Run Results:
256

  - string - [0]
╔ - string - [1]
╗ - string - [2]
╚ - string - [3]
╝ - string - [4]
║ - string - [5]
═ - string - [6]
 - string - [7]
 - string - [8]
         - string - [9]

 - string - [10]

 - string - [11]

 - string - [12]
 - string - [13]
 - string - [14]
 - string - [15]
 - string - [16]
 - string - [17]
 - string - [18]
 - string - [19]
 - string - [20]
 - string - [21]
 - string - [22]
 - string - [23]
 - string - [24]
 - string - [25]
 - string - [26]
? - string - [27]
 - string - [28]
 - string - [29]
 - string - [30]
 - string - [31]
  - string - [32]
! - string - [33]
" - string - [34]
# - string - [35]
$ - string - [36]
% - string - [37]
& - string - [38]
' - string - [39]
( - string - [40]
) - string - [41]
* - string - [42]
+ - string - [43]
, - string - [44]
- - string - [45]
. - string - [46]
/ - string - [47]
0 - string - [48]
1 - string - [49]
2 - string - [50]
3 - string - [51]
4 - string - [52]
5 - string - [53]
6 - string - [54]
7 - string - [55]
8 - string - [56]
9 - string - [57]
: - string - [58]
; - string - [59]
< - string - [60]
= - string - [61]
> - string - [62]
? - string - [63]
@ - string - [64]
A - string - [65]
B - string - [66]
C - string - [67]
D - string - [68]
E - string - [69]
F - string - [70]
G - string - [71]
H - string - [72]
I - string - [73]
J - string - [74]
K - string - [75]
L - string - [76]
M - string - [77]
N - string - [78]
O - string - [79]
P - string - [80]
Q - string - [81]
R - string - [82]
S - string - [83]
T - string - [84]
U - string - [85]
V - string - [86]
W - string - [87]
X - string - [88]
Y - string - [89]
Z - string - [90]
[ - string - [91]
\ - string - [92]
] - string - [93]
^ - string - [94]
_ - string - [95]
` - string - [96]
a - string - [97]
b - string - [98]
c - string - [99]
d - string - [100]
e - string - [101]
f - string - [102]
g - string - [103]
h - string - [104]
i - string - [105]
j - string - [106]
k - string - [107]
l - string - [108]
m - string - [109]
n - string - [110]
o - string - [111]
p - string - [112]
q - string - [113]
r - string - [114]
s - string - [115]
t - string - [116]
u - string - [117]
v - string - [118]
w - string - [119]
x - string - [120]
y - string - [121]
z - string - [122]
{ - string - [123]
| - string - [124]
} - string - [125]
~ - string - [126]

































  - string - [194 160]
¡ - string - [194 161]
¢ - string - [194 162]
£ - string - [194 163]
¤ - string - [194 164]
¥ - string - [194 165]
¦ - string - [194 166]
§ - string - [194 167]
¨ - string - [194 168]
© - string - [194 169]
ª - string - [194 170]
« - string - [194 171]
¬ - string - [194 172]
­ - string - [194 173]
® - string - [194 174]
¯ - string - [194 175]
° - string - [194 176]
± - string - [194 177]
² - string - [194 178]
³ - string - [194 179]
´ - string - [194 180]
µ - string - [194 181]
¶ - string - [194 182]
· - string - [194 183]
¸ - string - [194 184]
¹ - string - [194 185]
º - string - [194 186]
» - string - [194 187]
¼ - string - [194 188]
½ - string - [194 189]
¾ - string - [194 190]
¿ - string - [194 191]
À - string - [195 128]
Á - string - [195 129]
Â - string - [195 130]
Ã - string - [195 131]
Ä - string - [195 132]
Å - string - [195 133]
Æ - string - [195 134]
Ç - string - [195 135]
È - string - [195 136]
É - string - [195 137]
Ê - string - [195 138]
Ë - string - [195 139]
Ì - string - [195 140]
Í - string - [195 141]
Î - string - [195 142]
Ï - string - [195 143]
Ð - string - [195 144]
Ñ - string - [195 145]
Ò - string - [195 146]
Ó - string - [195 147]
Ô - string - [195 148]
Õ - string - [195 149]
Ö - string - [195 150]
× - string - [195 151]
Ø - string - [195 152]
Ù - string - [195 153]
Ú - string - [195 154]
Û - string - [195 155]
Ü - string - [195 156]
Ý - string - [195 157]
Þ - string - [195 158]
ß - string - [195 159]
à - string - [195 160]
á - string - [195 161]
â - string - [195 162]
ã - string - [195 163]
ä - string - [195 164]
å - string - [195 165]
æ - string - [195 166]
ç - string - [195 167]
è - string - [195 168]
é - string - [195 169]
ê - string - [195 170]
ë - string - [195 171]
ì - string - [195 172]
í - string - [195 173]
î - string - [195 174]
ï - string - [195 175]
ð - string - [195 176]
ñ - string - [195 177]
ò - string - [195 178]
ó - string - [195 179]
ô - string - [195 180]
õ - string - [195 181]
ö - string - [195 182]
÷ - string - [195 183]
ø - string - [195 184]
ù - string - [195 185]
ú - string - [195 186]
û - string - [195 187]
ü - string - [195 188]
ý - string - [195 189]
þ - string - [195 190]
ÿ - string - [195 191]
*/
