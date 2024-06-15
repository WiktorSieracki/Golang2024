set datafile separator ","
set title "Forest Density vs. Percentage of Trees Burned"
set xlabel "Forest Density"
set ylabel "Percentage of Trees Burned"
set xrange [0:1.0]
set yrange [0:100]
set terminal png size 800,600
set output "density_plot.png"
plot "density_results.csv" using 2:3 with linespoints title 'Burned Trees'