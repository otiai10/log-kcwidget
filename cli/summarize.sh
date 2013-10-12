
# Usage : nohup sh cli/summarize.sh > /dev/null &

while true
do
  go run app/proc/summarize.go 1> summarize.log 2> summarize.error.log
  sleep 1h
done
