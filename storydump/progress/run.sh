set -e
gnuplot plot.gp
convert -append 905_progress.png suvehnux_progress.png progress.png
