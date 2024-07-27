set -e
go run extract.go sort.go story.go progress.go -progress
cd progress
gnuplot plot.gp

convert +append 905_progress.png suvehnux_progress.png temp1.png
convert +append violet_progress.png shade_progress.png temp2.png
convert -append temp1.png temp2.png progress.png
rm temp1.png
rm temp2.png

#convert +append 905_progress.png suvehnux_progress.png shade_progress.png progress.png

#go run extract.go sort.go story.go progress.go -file ../storydump/905.z5_2024-07-22_20-52-10.json
#go run extract.go sort.go story.go progress.go -file ../storydump/suvehnux.z5_2024-06-28_20-46-47.json
#go run extract.go sort.go story.go progress.go -file ../storydump/violet.z8_2024-07-27_09-30-55.json
#go run extract.go sort.go story.go progress.go -file ../storydump/shade.z5_2024-07-26_10-12-18.json