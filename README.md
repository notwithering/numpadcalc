# numpadcalc

**x11 application**

this is a program that turns your useless number pad into a fully functioning standard notation calculator, simply press 3 and then + and then 4 and then enter on your numpad and it will place 7 into your clipboard

it also supports reading numbers from your clipboard as well so after you had 7 in your clipboard you can press * and then 3 and then it will place 14 in your clipboard

basically just for simple calculations you need quickly, it doesnt support decimals or anything, mainly because im too lazy to find an alternative to bc (bc -l gives me ugly trailing zeros and python uses IEEE 754 double-precision so the arithmetic is horribly inaccurate)

## prequisities

- bash
- bc
- xbindkeys
- xclip
- xorg-xev for arch, x11-utils for debian