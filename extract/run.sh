set -e
go run extract.go sort.go story.go progress.go -progress
cd progress
gnuplot plot.gp
convert -append 905_progress.png suvehnux_progress.png progress.png
