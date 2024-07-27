set term pngcairo
set yrange [-1:25]
set ylabel "Game Progress (%)"
set xlabel "Steps"

set key top left
#set key horiz
#set key top left maxcolumns 2
#set key top left maxrows 7
set key out vert
set key right

#set style line 1 lw 2
#set style line 2 lw 2
#set style line 3 lw 2
#set style line 4 lw 2
#set style line 5 lw 2
#set style line 6 lw 2
#set style line 7 lw 2
#set style line 8 lw 2
#set style line 9 lw 2

set style line 1 lt 1 lw 2 lc rgb '#1B9E77' # dark teal
set style line 2 lt 1 lw 2 lc rgb '#D95F02' # dark orange
set style line 3 lt 1 lw 2 lc rgb '#7570B3' # dark lilac
set style line 4 lt 1 lw 2 lc rgb '#E7298A' # dark magenta
set style line 5 lt 1 lw 2 lc rgb '#66A61E' # dark lime green
set style line 6 lt 1 lw 2 lc rgb '#E6AB02' # dark banana
set style line 7 lt 1 lw 2 lc rgb '#A6761D' # dark tan
set style line 8 lt 1 lw 2 lc rgb '#666666' # dark gray
set style line 9 lw 2
set style line 10 lw 2
set style line 11 lw 2

set xrange [0:100]

set title "Game Suveh Nux"
set output "suvehnux_progress.png"
load "suvehnux_plotlines.gp"

set yrange [-1:14]

set label "First ending" at 2, 11.5 left
set arrow from 0,11 to 100, 11 nohead lw 2 lc rgb '#000000' dashtype 2

set label "Second ending" at 2, 13.5 left
set arrow from 0,13 to 100, 13 nohead lw 2 lc rgb '#000000' dashtype 2
set ylabel "Game Progress"

set title "Game 9:05"
set output "905_progress.png"
load "905_plotlines.gp"

set yrange [-1:20]
unset arrow
unset label
set title "Game Shade"
set output "shade_progress.png"
load "shade_plotlines.gp"

set yrange [-1:20]
unset arrow
unset label
set title "Game Violet"
set output "violet_progress.png"
load "violet_plotlines.gp"
