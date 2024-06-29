set term pngcairo
set output "progress.png"
set yrange [-1:41]
set ylabel "Game Progress (%)"
set xlabel "Steps"
set key top left


set style line 1 lw 2
set style line 2 lw 2
set style line 3 lw 2
set style line 4 lw 2
set style line 5 lw 2
set style line 6 lw 2
set style line 7 lw 2

load "plotlines.gp"


