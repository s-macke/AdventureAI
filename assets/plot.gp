set encoding utf8

#set term pngcairo enhanced color dashed font "Alegreya, 24" rounded size 1024,512
set term pngcairo enhanced color dashed font "Arial, 24" rounded size 1024,512
set output "prices.png"

set style line 1 lc rgb '#E41A1C' pt 1 ps 1 lt 1 lw 4 # red
set style line 2 lc rgb '#377EB8' pt 6 ps 1 lt 1 lw 4 # blue

set grid

set xlabel "Number of input commands"
set ylabel "Total Price in US $"

set border lw 2

set label "Token limit" at 77,11.7 center textcolor rgb "#377EB8" front

#set ytics nomirror

set format y "$%g"

plot [0:90] [0:13] \
"price.data" u 1:4 axis x1y1 w lp ls 1 notitle, 10.7 ls 2 notitle

#title "Total Price in US$"

#"storydump/price.data" u 1:3 axis x1y2 w l lw 2 title "Prompt Tokens",		\
#"storydump/price.data" u 1:2 axis x1y2 w l lw 2 title "Completion Tokens"
